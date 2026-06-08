package helpers

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestDevAppMemberCommandsBuildToolParams(t *testing.T) {
	cases := []struct {
		name       string
		cmd        string
		args       []string
		wantTool   string
		wantParams map[string]any
	}{
		{
			name:     "list",
			cmd:      "list",
			args:     []string{"--app-id", "app-001"},
			wantTool: "list_open_dev_app_members",
			wantParams: map[string]any{
				"unifiedAppId": "app-001",
			},
		},
		{
			name:     "add multiple users",
			cmd:      "add",
			args:     []string{"--app-id", "app-001", "--users", "userId1,userId2,userId3,userId4", "--member-type", "DEVELOPER"},
			wantTool: "add_open_dev_app_members",
			wantParams: map[string]any{
				"unifiedAppId":  "app-001",
				"memberUserIds": []string{"userId1", "userId2", "userId3", "userId4"},
				"memberType":    "DEVELOPER",
			},
		},
		{
			name:     "remove trims users",
			cmd:      "remove",
			args:     []string{"--app-id", "app-001", "--users", " userId1 , userId2 ", "--member-type", "DEVELOPER"},
			wantTool: "remove_open_dev_app_members",
			wantParams: map[string]any{
				"unifiedAppId":  "app-001",
				"memberUserIds": []string{"userId1", "userId2"},
				"memberType":    "DEVELOPER",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runner := &captureRunner{}
			root := newDevAppCommand(runner)
			var out bytes.Buffer
			root.SetOut(&out)
			root.SetErr(&out)
			root.SetArgs(append([]string{"member", tc.cmd}, tc.args...))

			if err := root.Execute(); err != nil {
				t.Fatalf("Execute() error = %v\noutput:\n%s", err, out.String())
			}

			if got := runner.last.CanonicalProduct; got != "devapp" {
				t.Fatalf("CanonicalProduct = %q, want devapp", got)
			}
			if got := runner.last.Tool; got != tc.wantTool {
				t.Fatalf("Tool = %q, want %q", got, tc.wantTool)
			}
			if !reflect.DeepEqual(runner.last.Params, tc.wantParams) {
				t.Fatalf("Params = %#v, want %#v", runner.last.Params, tc.wantParams)
			}
		})
	}
}

func TestDevAppMemberCommandsValidateRequiredFlags(t *testing.T) {
	cases := []struct {
		name    string
		cmd     string
		args    []string
		wantErr string
	}{
		{name: "list requires app", cmd: "list", args: nil, wantErr: "--app-id is required"},
		{name: "add requires users", cmd: "add", args: []string{"--app-id", "app-001", "--member-type", "DEVELOPER"}, wantErr: "--users is required"},
		{name: "add rejects empty users", cmd: "add", args: []string{"--app-id", "app-001", "--users", " , ", "--member-type", "DEVELOPER"}, wantErr: "--users must contain at least one userId"},
		{name: "remove requires member type", cmd: "remove", args: []string{"--app-id", "app-001", "--users", "userId1"}, wantErr: "--member-type is required"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runner := &captureRunner{}
			root := newDevAppCommand(runner)
			var out bytes.Buffer
			root.SetOut(&out)
			root.SetErr(&out)
			root.SetArgs(append([]string{"member", tc.cmd}, tc.args...))

			err := root.Execute()
			if err == nil {
				t.Fatalf("Execute() error = nil, want %q", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error = %q, want to contain %q", err.Error(), tc.wantErr)
			}
			if runner.last.Tool != "" {
				t.Fatalf("tool = %q, want no invocation", runner.last.Tool)
			}
		})
	}
}

func TestDevAppSecurityConfigBuildsOnlyProvidedLists(t *testing.T) {
	runner := &captureRunner{}
	root := newDevAppCommand(runner)
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{
		"security", "config",
		"--app-id", "app-001",
		"--ip-whitelist", "103.211.230.150,103.211.230.151",
		"--redirect-url", "https://example.com/callback",
		"--sso-url", "https://example.com/sso",
	})

	if err := root.Execute(); err != nil {
		t.Fatalf("Execute() error = %v\noutput:\n%s", err, out.String())
	}

	if got := runner.last.CanonicalProduct; got != "devapp" {
		t.Fatalf("CanonicalProduct = %q, want devapp", got)
	}
	if got := runner.last.Tool; got != "update_app_security_config" {
		t.Fatalf("Tool = %q, want update_app_security_config", got)
	}
	want := map[string]any{
		"unifiedAppId":  "app-001",
		"ipWhiteList":   []string{"103.211.230.150", "103.211.230.151"},
		"redirectUrls":  []string{"https://example.com/callback"},
		"otherAuthUrls": []string{"https://example.com/sso"},
	}
	if !reflect.DeepEqual(runner.last.Params, want) {
		t.Fatalf("Params = %#v, want %#v", runner.last.Params, want)
	}
}

func TestDevAppSecurityConfigOmitsAbsentOptionalLists(t *testing.T) {
	runner := &captureRunner{}
	root := newDevAppCommand(runner)
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"security", "config", "--app-id", "app-001", "--redirect-url", "https://example.com/callback"})

	if err := root.Execute(); err != nil {
		t.Fatalf("Execute() error = %v\noutput:\n%s", err, out.String())
	}

	want := map[string]any{
		"unifiedAppId": "app-001",
		"redirectUrls": []string{"https://example.com/callback"},
	}
	if !reflect.DeepEqual(runner.last.Params, want) {
		t.Fatalf("Params = %#v, want %#v", runner.last.Params, want)
	}
}

func TestDevAppSecurityConfigRequiresAtLeastOneConfig(t *testing.T) {
	runner := &captureRunner{}
	root := newDevAppCommand(runner)
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"security", "config", "--app-id", "app-001"})

	err := root.Execute()
	if err == nil {
		t.Fatal("Execute() error = nil, want validation error")
	}
	if !strings.Contains(err.Error(), "one of --ip-whitelist, --redirect-url, or --sso-url is required") {
		t.Fatalf("error = %q", err.Error())
	}
	if runner.last.Tool != "" {
		t.Fatalf("tool = %q, want no invocation", runner.last.Tool)
	}
}
