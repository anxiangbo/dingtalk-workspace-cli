package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSheetBatchOperationTranslationCoversEveryMapping(t *testing.T) {
	input := map[string]any{
		"sheet-id": "Sheet1", "range": "A1:B2", "type": "all", "values": []any{"value"},
		"merge-type": "mergeRows", "source-range": "A1", "target-range": "B2",
		"target-sheet-id": "Sheet2", "paste-type": "values", "dimension": "row",
		"length": float64(2), "position": "3", "start-index": 1, "end-index": "2",
		"destination-index": float64(4), "options": []string{"a", "b"}, "multi-select": true,
		"csv": "a,b\r\n1,2", "start-cell": "A1", "allow-overwrite": true,
		"float-image-id": "image-id", "pixel-size": "24", "hidden": true,
		"group-state": "fold",
	}
	for name, mapping := range batchOpDispatch {
		got, err := translateBatchOp(map[string]any{"toolName": name, "input": input})
		if err != nil {
			t.Errorf("translateBatchOp(%q): %v", name, err)
			continue
		}
		if got["toolName"] != mapping.mcpTool {
			t.Errorf("translateBatchOp(%q) tool = %v, want %q", name, got["toolName"], mapping.mcpTool)
		}
		if _, ok := got["input"].(map[string]any); !ok {
			t.Errorf("translateBatchOp(%q) input = %#v", name, got["input"])
		}
	}
	if _, err := translateBatchOp(map[string]any{"toolName": "unknown"}); err == nil {
		t.Fatal("unknown batch operation should fail")
	}
	if _, err := translateBatchOp(map[string]any{"toolName": "range clear"}); err != nil {
		t.Fatalf("nil input should be accepted: %v", err)
	}
}

func TestSheetBatchValueConversionsAndDefaults(t *testing.T) {
	if got := batchStr(map[string]any{"second": 42}, "first", "second"); got != "42" {
		t.Fatalf("batchStr() = %q", got)
	}
	if got := batchStr(nil, "missing"); got != "" {
		t.Fatalf("missing batchStr() = %q", got)
	}
	for _, tc := range []struct {
		input map[string]any
		want  int
	}{
		{map[string]any{"n": float64(3)}, 3},
		{map[string]any{"n": 4}, 4},
		{map[string]any{"n": "5"}, 5},
		{nil, 0},
	} {
		if got := batchInt(tc.input, "missing", "n"); got != tc.want {
			t.Errorf("batchInt(%v) = %d, want %d", tc.input, got, tc.want)
		}
	}
	for _, input := range []map[string]any{nil, {"type": nil}, {"type": ""}} {
		if got := batchStrOr(input, "type", "content"); got != "content" {
			t.Errorf("batchStrOr(%v) = %q", input, got)
		}
	}
	if got := batchStrOr(map[string]any{"type": "all"}, "type", "content"); got != "all" {
		t.Fatalf("explicit batchStrOr() = %q", got)
	}

	if got := BuildMergeCellsArgs(nil)["mergeType"]; got != "mergeAll" {
		t.Fatalf("default merge type = %v", got)
	}
	if _, ok := BuildFillRangeArgs(nil)["fillType"]; ok {
		t.Fatal("empty fill type should be omitted")
	}
	if got := BuildGroupDimensionArgs(nil)["groupState"]; got != "expand" {
		t.Fatalf("default group state = %v", got)
	}
	if _, ok := BuildUpdateDimensionArgs(nil)["pixelSize"]; ok {
		t.Fatal("zero pixel size should be omitted")
	}
	if _, ok := BuildSetDropdownArgs(nil)["enableMultiSelect"]; ok {
		t.Fatal("unset multi-select should be omitted")
	}
}

func TestResolveCSVContent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "input.csv")
	if err := os.WriteFile(path, []byte("\xef\xbb\xbfhead\r\nvalue"), 0o600); err != nil {
		t.Fatalf("write csv: %v", err)
	}
	if got := resolveCsvContent("@" + path); got != "head\nvalue" {
		t.Fatalf("file csv = %q", got)
	}
	missing := "@" + filepath.Join(dir, "missing.csv")
	if got := resolveCsvContent(missing); got != missing {
		t.Fatalf("missing file csv = %q", got)
	}

	oldStdin := os.Stdin
	pipeRead, pipeWrite, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	if _, err := pipeWrite.WriteString("a,b\r\n1,2"); err != nil {
		t.Fatalf("write pipe: %v", err)
	}
	_ = pipeWrite.Close()
	os.Stdin = pipeRead
	t.Cleanup(func() {
		os.Stdin = oldStdin
		_ = pipeRead.Close()
	})
	if got := resolveCsvContent("-"); got != "a,b\n1,2" {
		t.Fatalf("stdin csv = %q", got)
	}
	if got := resolveCsvContent("plain\r\n"); got != strings.ReplaceAll("plain\r\n", "\r", "") {
		t.Fatalf("plain csv = %q", got)
	}
}
