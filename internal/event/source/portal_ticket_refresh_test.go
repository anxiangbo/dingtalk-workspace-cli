package source

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	dwsevent "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/event"
	"github.com/gorilla/websocket"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
)

// TestPortalStart401RefreshRetryEndToEnd drives the full production chain
// DingtalkSource.Start → startPortalTicket → requestPortalTicket: the first
// ticket request is rejected with 401, ForceRefreshToken rotates the token,
// the in-chain retry succeeds with the fresh token and a WebSocket event is
// delivered to emit.
func TestPortalStart401RefreshRetryEndToEnd(t *testing.T) {
	var ticketCalls, refreshCalls atomic.Int64
	var rejectedSeen atomic.Value

	upgrader := websocket.Upgrader{}
	var wsURL string
	mux := http.NewServeMux()
	mux.HandleFunc("/ticket", func(w http.ResponseWriter, r *http.Request) {
		ticketCalls.Add(1)
		switch r.Header.Get("x-user-access-token") {
		case "fresh-token":
			_ = json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"result":  map[string]string{"endpoint": wsURL, "ticket": "ticket-1"},
			})
		default:
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = io.WriteString(w, "token expired")
		}
	})
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		df := payload.DataFrame{Type: "event", Headers: payload.DataFrameHeader{payload.DataFrameHeaderKMessageId: "msg-1"}, Data: `{}`}
		_ = conn.WriteJSON(df)
		_, _, _ = conn.ReadMessage()
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	wsURL = "ws" + strings.TrimPrefix(srv.URL, "http") + "/ws"

	s, err := New(Config{PortalTicket: &PortalTicketConfig{
		TicketURL: srv.URL + "/ticket",
		AccessTokenProvider: func(context.Context) (string, error) {
			return "stale-token", nil
		},
		ForceRefreshToken: func(_ context.Context, rejectedToken string) (string, error) {
			refreshCalls.Add(1)
			rejectedSeen.Store(rejectedToken)
			return "fresh-token", nil
		},
		SourceID:   "source",
		HTTPClient: srv.Client(),
	}})
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	emitted := make(chan struct{}, 1)
	done := make(chan error, 1)
	go func() { done <- s.Start(ctx, func(*dwsevent.RawEvent) { emitted <- struct{}{} }) }()
	select {
	case <-emitted:
	case <-time.After(2 * time.Second):
		t.Fatal("portal event timeout after 401 refresh retry")
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("portal stop = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("portal stop timeout")
	}

	if got := ticketCalls.Load(); got != 2 {
		t.Fatalf("ticket calls = %d, want 2", got)
	}
	if got := refreshCalls.Load(); got != 1 {
		t.Fatalf("refresh calls = %d, want 1", got)
	}
	if got, _ := rejectedSeen.Load().(string); got != "stale-token" {
		t.Fatalf("rejected token = %q, want %q", got, "stale-token")
	}
}

// TestRequestPortalTicketRetryUsesRotatedTokenDirectly asserts the in-chain
// retry sends the token returned by ForceRefreshToken instead of re-invoking
// the provider (which could still serve the stale token).
func TestRequestPortalTicketRetryUsesRotatedTokenDirectly(t *testing.T) {
	providerCalls := 0
	var attemptTokens []string
	client := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		token := req.Header.Get("x-user-access-token")
		attemptTokens = append(attemptTokens, token)
		if token != "rotated" {
			return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"endpoint":"wss://x","ticket":"t"}`)), Header: make(http.Header)}, nil
	})}
	ticket, err := requestPortalTicket(context.Background(), &PortalTicketConfig{
		TicketURL: "https://x",
		AccessTokenProvider: func(context.Context) (string, error) {
			providerCalls++
			return "stale", nil
		},
		ForceRefreshToken: func(_ context.Context, rejectedToken string) (string, error) {
			if rejectedToken != "stale" {
				t.Fatalf("rejected token = %q, want %q", rejectedToken, "stale")
			}
			return "rotated", nil
		},
		SourceID:   "s",
		HTTPClient: client,
	})
	if err != nil {
		t.Fatalf("requestPortalTicket = %v", err)
	}
	if ticket.Endpoint != "wss://x" || ticket.Ticket != "t" {
		t.Fatalf("ticket = %#v", ticket)
	}
	if providerCalls != 1 {
		t.Fatalf("provider calls = %d, want 1", providerCalls)
	}
	if len(attemptTokens) != 2 || attemptTokens[0] != "stale" || attemptTokens[1] != "rotated" {
		t.Fatalf("attempt tokens = %v", attemptTokens)
	}
}

// TestRequestPortalTicketRefreshFailureKeepsBothErrors asserts a failing
// refresh neither retries nor drops the refresh error or the original 401.
func TestRequestPortalTicketRefreshFailureKeepsBothErrors(t *testing.T) {
	refreshErr := errors.New("refresh_token exchange failed")
	attempts := 0
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
	})}
	_, err := requestPortalTicket(context.Background(), &PortalTicketConfig{
		TicketURL:   "https://x",
		AccessToken: "stale",
		ForceRefreshToken: func(context.Context, string) (string, error) {
			return "", refreshErr
		},
		SourceID:   "s",
		HTTPClient: client,
	})
	if !errors.Is(err, refreshErr) {
		t.Fatalf("error should wrap refresh error, got %v", err)
	}
	if err == nil || !strings.Contains(err.Error(), "HTTP 401") {
		t.Fatalf("error should keep original 401, got %v", err)
	}
	if attempts != 1 {
		t.Fatalf("attempts = %d, want 1 (no retry after failed refresh)", attempts)
	}
}

// TestRequestPortalTicketWithoutRefreshCallback401StaysFatal covers backward
// compatibility: nil ForceRefreshToken keeps the single-attempt fatal 401.
func TestRequestPortalTicketWithoutRefreshCallback401StaysFatal(t *testing.T) {
	attempts := 0
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
	})}
	_, err := requestPortalTicket(context.Background(), &PortalTicketConfig{
		TicketURL: "https://x", AccessToken: "stale", SourceID: "s", HTTPClient: client,
	})
	if err == nil || !strings.Contains(err.Error(), "HTTP 401") {
		t.Fatalf("fatal 401 expected, got %v", err)
	}
	if attempts != 1 {
		t.Fatalf("attempts = %d, want 1", attempts)
	}
}

// TestRequestPortalTicketSecond401IsFatal guards against refresh loops: the
// controlled retry happens exactly once even if the rotated token is also
// rejected.
func TestRequestPortalTicketSecond401IsFatal(t *testing.T) {
	attempts := 0
	refreshCalls := 0
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
	})}
	_, err := requestPortalTicket(context.Background(), &PortalTicketConfig{
		TicketURL:   "https://x",
		AccessToken: "stale",
		ForceRefreshToken: func(context.Context, string) (string, error) {
			refreshCalls++
			return "rotated-but-still-rejected", nil
		},
		SourceID:   "s",
		HTTPClient: client,
	})
	if err == nil || !strings.Contains(err.Error(), "HTTP 401") {
		t.Fatalf("fatal 401 expected after single retry, got %v", err)
	}
	if attempts != 2 || refreshCalls != 1 {
		t.Fatalf("attempts = %d refreshCalls = %d, want 2/1", attempts, refreshCalls)
	}
}

// TestRequestPortalTicketRefreshEmptyTokenIsFatal asserts an empty rotated
// token is rejected instead of being sent to the server.
func TestRequestPortalTicketRefreshEmptyTokenIsFatal(t *testing.T) {
	attempts := 0
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
	})}
	_, err := requestPortalTicket(context.Background(), &PortalTicketConfig{
		TicketURL:   "https://x",
		AccessToken: "stale",
		ForceRefreshToken: func(context.Context, string) (string, error) {
			return "  ", nil
		},
		SourceID:   "s",
		HTTPClient: client,
	})
	if err == nil || !strings.Contains(err.Error(), "empty token") {
		t.Fatalf("empty rotated token error expected, got %v", err)
	}
	if attempts != 1 {
		t.Fatalf("attempts = %d, want 1", attempts)
	}
}

// TestPersonalFetchTicket401RefreshRetry mirrors the portal behavior for the
// personal stream ticket path.
func TestPersonalFetchTicket401RefreshRetry(t *testing.T) {
	var attemptTokens []string
	src, err := NewPersonal(PersonalConfig{
		AccessTokenProvider: func(context.Context) (string, error) { return "stale", nil },
		ForceRefreshToken: func(_ context.Context, rejectedToken string) (string, error) {
			if rejectedToken != "stale" {
				t.Fatalf("rejected token = %q, want %q", rejectedToken, "stale")
			}
			return "rotated", nil
		},
		ClientID:  "client",
		SourceID:  "source",
		TicketURL: "https://ticket.test",
		HTTPClient: &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			token := req.Header.Get("x-user-access-token")
			attemptTokens = append(attemptTokens, token)
			if token != "rotated" {
				return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
			}
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"endpoint":"wss://stream.test","ticket":"ticket"}`)), Header: make(http.Header)}, nil
		})},
	})
	if err != nil {
		t.Fatal(err)
	}
	ticket, err := src.fetchTicket(context.Background())
	if err != nil {
		t.Fatalf("fetchTicket = %v", err)
	}
	if ticket.Endpoint != "wss://stream.test" || ticket.Ticket != "ticket" {
		t.Fatalf("ticket = %#v", ticket)
	}
	if len(attemptTokens) != 2 || attemptTokens[0] != "stale" || attemptTokens[1] != "rotated" {
		t.Fatalf("attempt tokens = %v", attemptTokens)
	}
}

// TestPersonalFetchTicket401RefreshFailureStaysFatal asserts a failed refresh
// keeps the 401 fatal (not retryable) and wraps the refresh error.
func TestPersonalFetchTicket401RefreshFailureStaysFatal(t *testing.T) {
	refreshErr := errors.New("refresh_token exchange failed")
	src, err := NewPersonal(PersonalConfig{
		AccessToken: "stale",
		ForceRefreshToken: func(context.Context, string) (string, error) {
			return "", refreshErr
		},
		ClientID:  "client",
		SourceID:  "source",
		TicketURL: "https://ticket.test",
		HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(strings.NewReader("expired")), Header: make(http.Header)}, nil
		})},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = src.fetchTicket(context.Background())
	if !errors.Is(err, refreshErr) {
		t.Fatalf("error should wrap refresh error, got %v", err)
	}
	if isRetryablePersonalError(err) {
		t.Fatalf("failed refresh should stay fatal, got retryable %v", err)
	}
}
