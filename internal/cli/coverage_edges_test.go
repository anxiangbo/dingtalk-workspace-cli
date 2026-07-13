// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/pkg/edition"
)

func TestCanonicalCommandsAndFlags(t *testing.T) {
	mcp := NewMCPCommand(t.Context(), nil, nil, nil)
	mcp.SetOut(&bytes.Buffer{})
	if err := mcp.Execute(); err != nil {
		t.Fatal(err)
	}

	schema := map[string]any{"properties": map[string]any{
		"str":     map[string]any{"type": "string", "description": " text "},
		"enum":    map[string]any{"enum": []any{"a"}},
		"int":     map[string]any{"type": "integer"},
		"num":     map[string]any{"type": "number"},
		"bool":    map[string]any{"type": "boolean"},
		"object":  map[string]any{"type": "object"},
		"strings": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
		"enums":   map[string]any{"type": "array", "items": map[string]any{"enum": []any{"x"}}},
		"ints":    map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
		"nums":    map[string]any{"type": "array", "items": map[string]any{"type": "number"}},
		"bools":   map[string]any{"type": "array", "items": map[string]any{"type": "boolean"}},
		"objects": map[string]any{"type": "array", "items": map[string]any{"type": "object"}},
		"unknown": map[string]any{"type": "future"},
		"array":   map[string]any{"type": "array"},
		"invalid": "not schema",
	}}
	specs := BuildFlagSpecs(schema, map[string]CLIFlagHint{
		"str":  {Alias: "string-alias", Shorthand: "s"},
		"int":  {Alias: "int-alias", Shorthand: "i"},
		"num":  {Alias: "num-alias", Shorthand: "n"},
		"bool": {Alias: "bool-alias", Shorthand: "b"},
	})
	if len(specs) != 13 {
		t.Fatalf("flag specs = %d, want 13", len(specs))
	}
	if BuildFlagSpecs(nil, nil) != nil || BuildFlagSpecs(map[string]any{}, nil) != nil {
		t.Fatal("missing properties produced flag specs")
	}
	cmd := &cobra.Command{Use: "flags"}
	cmd.Flags().String("occupied", "", "")
	cmd.Flags().StringP("short", "s", "", "")
	specs = append(specs,
		FlagSpec{PropertyName: "skip", FlagName: "json", Kind: flagString},
		FlagSpec{PropertyName: "same", FlagName: "same", Alias: "same", Kind: flagString},
		FlagSpec{PropertyName: "occupied", FlagName: "occupied", Kind: flagString},
		FlagSpec{PropertyName: "default", FlagName: "default", Kind: flagString},
	)
	applyFlagSpecs(cmd, specs)
	applyFlagSpecs(cmd, []FlagSpec{{PropertyName: "array-alias", FlagName: "array-alias", Alias: "array-hidden", Kind: flagStringArray}})
	for _, name := range []string{"str", "string-alias", "int", "int-alias", "num", "num-alias", "bool", "bool-alias", "strings", "ints", "nums", "bools", "objects", "array", "default"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Fatalf("flag %q not registered", name)
		}
	}
	if nested, ok := nestedMap(map[string]any{"x": map[string]any{"y": true}}, "x"); !ok || !nested["y"].(bool) {
		t.Fatal("nested map lookup failed")
	}
	if _, ok := nestedMap(nil, "x"); ok {
		t.Fatal("nil nested map matched")
	}
	if _, ok := nestedMap(map[string]any{}, "x"); ok {
		t.Fatal("missing nested map matched")
	}
	if _, ok := nestedMap(map[string]any{"x": 1}, "x"); ok {
		t.Fatal("non-map nested value matched")
	}
	if schemaDescription(map[string]any{"description": 1}) != "" {
		t.Fatal("non-string description rendered")
	}
}

func TestSchemaCommandModes(t *testing.T) {
	root := &cobra.Command{Use: "dws"}
	dev := &cobra.Command{Use: "dev"}
	group := &cobra.Command{Use: "app"}
	leaf := &cobra.Command{
		Use:   "create",
		Short: "create app",
		Run:   func(*cobra.Command, []string) {},
		Annotations: map[string]string{
			"mcp-tool":   "app.create",
			"mcp-source": "source",
		},
	}
	group.AddCommand(leaf)
	dev.AddCommand(group)
	root.AddCommand(dev)
	fetch := func(context.Context, string) (map[string]HelperToolSchema, error) {
		return map[string]HelperToolSchema{
			"app.create": {
				Description: "create",
				Properties:  map[string]any{"appName": map[string]any{"type": "string", "default": "demo"}},
				Required:    []string{"appName"},
			},
		}, nil
	}
	schemaCmd := NewSchemaCommand(nil, fetch)
	var out bytes.Buffer
	schemaCmd.SetOut(&out)
	root.AddCommand(schemaCmd)
	root.SetArgs([]string{"schema", "dev.app.create"})
	if err := root.Execute(); err != nil || !strings.Contains(out.String(), "app-name") {
		t.Fatalf("helper schema output = %q, %v", out.String(), err)
	}

	root = &cobra.Command{Use: "dws"}
	schemaCmd = NewSchemaCommand(nil, nil)
	out.Reset()
	schemaCmd.SetOut(&out)
	root.AddCommand(schemaCmd)
	root.SetArgs([]string{"schema"})
	if err := root.Execute(); err != nil || !strings.Contains(out.String(), "static endpoint mode") {
		t.Fatalf("static schema output = %q, %v", out.String(), err)
	}
	root = &cobra.Command{Use: "dws"}
	schemaCmd = NewSchemaCommand(nil, nil)
	root.AddCommand(schemaCmd)
	root.SetArgs([]string{"schema", "one", "--cli-path", "two"})
	if err := root.Execute(); err == nil {
		t.Fatal("conflicting schema paths accepted")
	}
	root = &cobra.Command{Use: "dws"}
	schemaCmd = NewSchemaCommand(nil, nil)
	out.Reset()
	schemaCmd.SetOut(&out)
	root.AddCommand(schemaCmd)
	root.SetArgs([]string{"schema", "--cli-path", "other"})
	if err := root.Execute(); err != nil || !strings.Contains(out.String(), "static endpoint mode") {
		t.Fatalf("cli-path schema = %q, %v", out.String(), err)
	}
	root = &cobra.Command{Use: "dws"}
	dev = &cobra.Command{Use: "dev"}
	leaf = &cobra.Command{Use: "leaf", Run: func(*cobra.Command, []string) {}, Annotations: map[string]string{"mcp-tool": "tool"}}
	dev.AddCommand(leaf)
	root.AddCommand(dev)
	schemaCmd = NewSchemaCommand(nil, func(context.Context, string) (map[string]HelperToolSchema, error) {
		return nil, errors.New("fetch failed")
	})
	root.AddCommand(schemaCmd)
	root.SetArgs([]string{"schema", "dev.leaf"})
	if err := root.Execute(); err == nil {
		t.Fatal("schema fetch error did not propagate")
	}
}

func TestHelperSchemaEdgeCases(t *testing.T) {
	if payload, ok, err := renderHelperSchema(t.Context(), nil, "dev", nil); err != nil || ok || payload != nil {
		t.Fatalf("nil helper root = %#v, %v, %v", payload, ok, err)
	}
	root := &cobra.Command{Use: "dws"}
	dev := &cobra.Command{Use: "dev", Short: "dev"}
	leafNoTool := &cobra.Command{Use: "connect", Run: func(*cobra.Command, []string) {}}
	leaf := &cobra.Command{Use: "leaf", Run: func(*cobra.Command, []string) {}, Annotations: map[string]string{"mcp-tool": "tool"}}
	hidden := &cobra.Command{Use: "hidden", Hidden: true, Run: func(*cobra.Command, []string) {}}
	dev.AddCommand(leafNoTool, leaf, hidden)
	root.AddCommand(dev)
	for _, path := range []string{"", "other", "dev.unknown", "dev", "dev --flag"} {
		if _, _, err := renderHelperSchema(t.Context(), root, path, nil); err != nil {
			t.Fatalf("render %q: %v", path, err)
		}
	}
	if payload, err := helperLeafSchema(t.Context(), leafNoTool, nil); err != nil || payload["error"] == nil {
		t.Fatalf("unbound leaf = %#v, %v", payload, err)
	}
	if payload, err := helperLeafSchema(t.Context(), leaf, nil); err != nil || payload["error"] == nil {
		t.Fatalf("nil fetcher leaf = %#v, %v", payload, err)
	}
	failure := errors.New("fetch failed")
	if _, err := helperLeafSchema(t.Context(), leaf, func(context.Context, string) (map[string]HelperToolSchema, error) {
		return nil, failure
	}); !errors.Is(err, failure) {
		t.Fatalf("fetch error = %v", err)
	}
	if payload, err := helperLeafSchema(t.Context(), leaf, func(context.Context, string) (map[string]HelperToolSchema, error) {
		return map[string]HelperToolSchema{}, nil
	}); err != nil || payload["error"] == nil {
		t.Fatalf("missing tool = %#v, %v", payload, err)
	}
	leaf.Annotations["mcp-source"] = "custom"
	payload, err := helperLeafSchema(t.Context(), leaf, func(_ context.Context, source string) (map[string]HelperToolSchema, error) {
		if source != "custom" {
			t.Fatalf("source = %q", source)
		}
		return map[string]HelperToolSchema{"tool": {Description: "desc"}}, nil
	})
	if err != nil || payload["source"] != "mcp:custom" {
		t.Fatalf("leaf payload = %#v, %v", payload, err)
	}
}

func TestHelperSchemaPrimitiveEdges(t *testing.T) {
	tool := HelperToolSchema{
		Properties: map[string]any{
			"nil":      nil,
			"string":   map[string]any{"type": "string", "default": "x"},
			"integer":  map[string]any{"type": "integer", "default": float64(2)},
			"fraction": map[string]any{"type": "number", "default": 2.5},
			"bool":     map[string]any{"type": "boolean", "default": true},
			"array":    map[string]any{"type": "array"},
			"object":   map[string]any{"type": "object"},
			"unknown":  map[string]any{"type": "unknown"},
			"missing":  map[string]any{},
		},
		Required: []string{"string"},
	}
	if params := helperFlatParameters(tool); len(params) != len(tool.Properties) {
		t.Fatalf("parameters = %#v", params)
	}
	for _, typ := range []string{"string", "integer", "number", "boolean", "array", "object", "unknown"} {
		_ = mcpJSONType(map[string]any{"type": typ})
	}
	if mcpString(nil, "x") != "" || mcpString(map[string]any{"x": 1}, "x") != "" {
		t.Fatal("invalid MCP string rendered")
	}
	for _, prop := range []map[string]any{nil, {}, {"default": nil}, {"default": "x"}, {"default": float64(1)}, {"default": 1.25}, {"default": true}} {
		_, _ = mcpDefault(prop)
	}
	for input, want := range map[string]string{"SSLVerify": "ssl-verify", "__a--b__": "a-b", "aB2C": "a-b2-c"} {
		if got := kebabCase(input); got != want {
			t.Fatalf("kebabCase(%q) = %q, want %q", input, got, want)
		}
	}
	if got := helperSubcommands(&cobra.Command{Use: "empty"}); len(got) != 0 {
		t.Fatalf("empty subcommands = %#v", got)
	}
}

func TestLoaderAndPriorityEdges(t *testing.T) {
	for _, reason := range []CatalogDegradedReason{DegradedUnauthenticated, DegradedMarketUnreachable, DegradedRuntimeAllFailed, "other"} {
		degraded := newCatalogDegraded(reason, 2)
		if degraded.Error() == "" || degradedHint(reason, 2) == "" || degraded.Hint == "" {
			t.Fatalf("degraded %q = %#v", reason, degraded)
		}
	}
	catalog := Catalog{Products: []CanonicalProduct{{ID: "product", Tools: []ToolDescriptor{{RPCName: "tool"}}}}}
	product, ok := catalog.FindProduct("product")
	if !ok {
		t.Fatal("product not found")
	}
	if _, ok := catalog.FindProduct("missing"); ok {
		t.Fatal("missing product found")
	}
	if _, ok := product.FindTool("tool"); !ok {
		t.Fatal("tool not found")
	}
	if _, ok := product.FindTool("missing"); ok {
		t.Fatal("missing tool found")
	}
	if got, err := (StaticLoader{Catalog: catalog}).Load(t.Context()); err != nil || len(got.Products) != 1 {
		t.Fatalf("static load = %#v, %v", got, err)
	}
	failure := errors.New("load failed")
	if got, err := CatalogLoaderFrom(catalog, failure).Load(t.Context()); !errors.Is(err, failure) || len(got.Products) != 1 {
		t.Fatalf("preloaded load = %#v, %v", got, err)
	}
	loader := NewEnvironmentLoader()
	if got, err := loader.Load(t.Context()); err != nil || len(got.Products) != 0 {
		t.Fatalf("environment load = %#v, %v", got, err)
	}
	cmd := &cobra.Command{Use: "priority"}
	SetOverridePriority(cmd, 42)
	if got := OverridePriority(cmd); got != 42 {
		t.Fatalf("priority = %d", got)
	}
	original := edition.Get()
	edition.Override(&edition.Hooks{IsEmbedded: true})
	t.Cleanup(func() { edition.Override(original) })
	for _, reason := range []CatalogDegradedReason{DegradedUnauthenticated, DegradedMarketUnreachable, DegradedRuntimeAllFailed} {
		if degradedHint(reason, 2) == "" {
			t.Fatalf("embedded degraded hint %q is empty", reason)
		}
	}
}

func TestSchemaValidationCompleteMatrix(t *testing.T) {
	if err := ValidateInputSchema(nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := ValidateInputSchema(nil, map[string]any{"type": "object", "required": []string{"x"}}); err == nil {
		t.Fatal("missing required property accepted")
	}
	schema := map[string]any{
		"type":     "object",
		"required": []any{"known", 1, ""},
		"properties": map[string]any{
			"known": map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
			"skip":  "invalid",
		},
		"additionalProperties": false,
	}
	if err := ValidateInputSchema(map[string]any{"known": []any{1, 2.5}}, schema); err == nil {
		t.Fatal("invalid array item accepted")
	}
	if err := ValidateInputSchema(map[string]any{"known": []any{1}, "extra": true}, schema); err == nil {
		t.Fatal("unknown property accepted")
	}
	additional := map[string]any{
		"type":                 "object",
		"properties":           map[string]any{"known": map[string]any{"type": "string"}},
		"additionalProperties": map[string]any{"type": "number"},
	}
	if err := ValidateInputSchema(map[string]any{"known": "x", "extra": "bad"}, additional); err == nil {
		t.Fatal("invalid additional property accepted")
	}
	if err := validateSchemaValue("$", "x", nil); err != nil {
		t.Fatal(err)
	}
	if err := validateSchemaValue("$", "x", map[string]any{"enum": []string{"y"}}); err == nil {
		t.Fatal("invalid enum accepted")
	}

	typeCases := []struct {
		value any
		typ   string
		want  bool
	}{
		{map[string]any{}, "object", true}, {[]any{}, "array", true}, {"x", "string", true},
		{true, "boolean", true}, {1.5, "number", true}, {"x", "number", false},
		{2, "integer", true}, {2.5, "integer", false}, {"x", "integer", false},
		{nil, "null", true}, {"x", "future", true},
	}
	for _, tc := range typeCases {
		if got := matchesType(tc.value, tc.typ); got != tc.want {
			t.Fatalf("matchesType(%#v,%q) = %v, want %v", tc.value, tc.typ, got, tc.want)
		}
	}
	for _, typed := range []any{"", "string", []string{"", "number"}, []any{"", 1, "boolean"}, 1} {
		_ = schemaTypes(map[string]any{"type": typed})
	}
	if hasType([]string{"a", "b"}, "c") || !hasType([]string{"a", "b"}, "b") {
		t.Fatal("hasType failed")
	}
	if matchesAnyType("x", []string{"number", "boolean"}) || !matchesAnyType("x", []string{"number", "string"}) {
		t.Fatal("matchesAnyType failed")
	}
	for _, enum := range []any{[]any{"x"}, []string{"x"}, "invalid"} {
		_ = schemaEnum(map[string]any{"enum": enum})
	}
	if valuesEqual("1", 1) || !valuesEqual(json.Number("1"), float64(1)) {
		t.Fatal("enum value equality failed")
	}
	for _, number := range []any{float64(1), float32(1), int(1), int8(1), int16(1), int32(1), int64(1), uint(1), uint8(1), uint16(1), uint32(1), uint64(1), json.Number("1"), json.Number("x"), "x"} {
		_, _ = numberValue(number)
	}
	if got := schemaProperties(map[string]any{}); len(got) != 0 {
		t.Fatal("missing properties were nonempty")
	}
	for _, required := range []any{[]string{"x"}, []any{"x", 1, ""}, "invalid"} {
		_ = schemaRequired(map[string]any{"required": required})
	}
	for _, additional := range []map[string]any{{}, {"additionalProperties": true}, {"additionalProperties": false}, {"additionalProperties": map[string]any{"type": "string"}}, {"additionalProperties": "invalid"}} {
		_, _, _ = additionalProperties(additional)
	}
}

func TestStdinCompleteMatrix(t *testing.T) {
	original := os.Stdin
	t.Cleanup(func() { os.Stdin = original })
	closed, err := os.CreateTemp(t.TempDir(), "closed")
	if err != nil {
		t.Fatal(err)
	}
	_ = closed.Close()
	os.Stdin = closed
	if StdinIsPipe() {
		t.Fatal("closed stdin reported as pipe")
	}
	if _, err := ReadStdin(); err == nil {
		t.Fatal("closed stdin read succeeded")
	}

	stdin, err := os.CreateTemp(t.TempDir(), "stdin")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := stdin.WriteString("piped"); err != nil {
		t.Fatal(err)
	}
	_, _ = stdin.Seek(0, 0)
	os.Stdin = stdin
	if !StdinIsPipe() {
		t.Fatal("regular file stdin did not report pipe")
	}
	if got, err := ReadStdinIfPiped(); err != nil || got != "piped" {
		t.Fatalf("piped stdin = %q, %v", got, err)
	}
	_, _ = stdin.Seek(0, 0)
	guard := NewStdinGuard()
	if got, err := ResolveInputSource("@-", "text", guard); err != nil || got != "piped" {
		t.Fatalf("resolved stdin = %q, %v", got, err)
	}
	if _, err := ResolveInputSource("@-", "other", guard); err == nil {
		t.Fatal("stdin reused")
	}
	if _, err := ResolveInputSource("@-", "text", nil); err == nil {
		t.Fatal("nil stdin guard accepted")
	}

	dir := t.TempDir()
	if _, err := readFileBounded(dir); err == nil {
		t.Fatal("directory read as file")
	}
	large := filepath.Join(t.TempDir(), "large")
	if err := os.WriteFile(large, bytes.Repeat([]byte("x"), maxStdinSize+1), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := readFileBounded(large); err == nil {
		t.Fatal("oversized file accepted")
	}
	largeStdin, err := os.Open(large)
	if err != nil {
		t.Fatal(err)
	}
	defer largeStdin.Close()
	os.Stdin = largeStdin
	if _, err := readStdinBounded(); err == nil {
		t.Fatal("oversized stdin accepted")
	}
	for _, value := range []string{"@A", "@a", "@1", "@.", "@/", "@~", "@_", "@-", "@!", "@所有人", "plain"} {
		_ = looksLikeFilePath(value)
	}
	if _, _, err := ReadFileArg("@"); err == nil {
		t.Fatal("bare file argument accepted")
	}
	if _, err := ResolveInputSource("@", "text", guard); err == nil {
		t.Fatal("bare input source accepted")
	}
}
