// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"testing"

	"github.com/spf13/cobra"
)

// buildHelperTestTree mirrors the shape of the real `dws dev` subtree closely
// enough to exercise the local schema renderer: a group and leaves carrying the
// `mcp-tool` annotation that provides a stable canonical tool name.
func buildHelperTestTree() *cobra.Command {
	root := &cobra.Command{Use: "dws"}

	create := &cobra.Command{
		Use:         "create",
		Short:       "创建应用",
		Annotations: map[string]string{"mcp-tool": "create_dev_app"},
		Run:         func(*cobra.Command, []string) {},
	}
	create.Flags().String("name", "", "应用名称")
	create.Flags().String("desc", "", "应用描述")

	config := &cobra.Command{
		Use:         "config",
		Short:       "配置机器人",
		Annotations: map[string]string{"mcp-tool": "set_extension_robot_config"},
		Run:         func(*cobra.Command, []string) {},
	}
	config.Flags().String("unified-app-id", "", "统一应用 ID")
	config.Flags().String("event-callback-url", "", "事件回调地址")
	config.Flags().StringSlice("skills", nil, "技能列表")
	config.Flags().String("mode", "", "机器人模式")
	_ = config.MarkFlagRequired("unified-app-id")
	_ = config.Flags().SetAnnotation("mode", "x-cli-enum", []string{"HTTPS", "STREAM", "AISKILL"})

	// A local leaf without an mcp-tool annotation (e.g. dev connect status/stop).
	noTool := &cobra.Command{Use: "connect", Short: "无 MCP 工具", Run: func(*cobra.Command, []string) {}}
	noTool.Flags().String("robot-client-id", "", "机器人 clientId")

	robot := &cobra.Command{Use: "robot", Short: "机器人能力"}
	robot.AddCommand(config)

	app := &cobra.Command{Use: "app", Short: "应用"}
	app.AddCommand(create, robot)

	dev := &cobra.Command{Use: "dev", Short: "开放平台开发者命令"}
	dev.AddCommand(app, noTool)

	root.AddCommand(dev)
	return root
}

func TestRenderHelperSchema_LeafGwsFlat(t *testing.T) {
	root := buildHelperTestTree()

	payload, ok, err := renderHelperSchema(root, "dev app robot config")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected helper renderer to claim the path")
	}

	// Flat top-level: runtime-aligned leaf identity + parameters; no wrapper.
	if payload["name"] != "set_extension_robot_config" {
		t.Fatalf("name = %v", payload["name"])
	}
	if payload["product_id"] != "dev" {
		t.Fatalf("product_id = %v", payload["product_id"])
	}
	if payload["canonical_path"] != "dev.set_extension_robot_config" {
		t.Fatalf("canonical_path = %v", payload["canonical_path"])
	}
	if payload["description"] != "配置机器人" {
		t.Fatalf("description = %v", payload["description"])
	}
	if payload["path"] != "dev.set_extension_robot_config" {
		t.Fatalf("path = %v", payload["path"])
	}
	if payload["cli_path"] != "dev app robot config" {
		t.Fatalf("cli_path = %v", payload["cli_path"])
	}
	if payload["primary_cli_path"] != "dev app robot config" || payload["is_alias"] != false {
		t.Fatalf("alias markers = primary:%v is_alias:%v", payload["primary_cli_path"], payload["is_alias"])
	}
	if payload["source"] != "hardcoded:dev" {
		t.Fatalf("source = %v", payload["source"])
	}
	for _, leaked := range []string{"kind", "tool", "product", "helper"} {
		if _, present := payload[leaked]; present {
			t.Fatalf("gws-flat output must not carry %q wrapper key", leaked)
		}
	}

	params, _ := payload["parameters"].(map[string]any)
	if params == nil {
		t.Fatalf("no parameters: %#v", payload)
	}

	// Keys and parameter metadata come from the executable Cobra flags.
	uid, _ := params["unified-app-id"].(map[string]any)
	if uid == nil {
		t.Fatalf("missing unified-app-id param: %#v", params)
	}
	if uid["type"] != "string" || uid["required"] != true {
		t.Fatalf("unified-app-id = %#v, want string+required", uid)
	}
	if uid["property"] != "unifiedAppId" {
		t.Fatalf("unified-app-id property = %#v", uid["property"])
	}
	if _, hasDefault := uid["default"]; hasDefault {
		t.Fatal("unified-app-id must not carry an empty default")
	}

	cb, _ := params["event-callback-url"].(map[string]any)
	if cb == nil || cb["required"] != false {
		t.Fatalf("event-callback-url = %#v, want required=false", cb)
	}

	skills, _ := params["skills"].(map[string]any)
	if skills == nil || skills["type"] != "array" {
		t.Fatalf("skills = %#v, want array", skills)
	}

	mode, _ := params["mode"].(map[string]any)
	if mode == nil || mode["type"] != "string" {
		t.Fatalf("mode = %#v, want string", mode)
	}
	if _, hasDefault := mode["default"]; hasDefault {
		t.Fatalf("mode default = %v, want none", mode["default"])
	}
	if mode["required"] != false {
		t.Fatalf("mode required = %v, want false", mode["required"])
	}
	if enum, _ := mode["enum"].([]string); len(enum) != 3 {
		t.Fatalf("mode enum = %#v, want three values", mode["enum"])
	}
}

func TestRenderHelperSchema_CanonicalPath(t *testing.T) {
	root := buildHelperTestTree()

	payload, ok, err := renderHelperSchema(root, "dev.set_extension_robot_config")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected helper renderer to claim canonical helper path")
	}
	if payload["path"] != "dev.set_extension_robot_config" || payload["cli_path"] != "dev app robot config" {
		t.Fatalf("payload paths = %#v", payload)
	}
}

func TestRenderHelperSchema_Group(t *testing.T) {
	root := buildHelperTestTree()
	payload, ok, err := renderHelperSchema(root, "dev app")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected claim")
	}
	if payload["path"] != "dev app" {
		t.Fatalf("path = %v", payload["path"])
	}
	cmds, _ := payload["commands"].([]map[string]any)
	if len(cmds) != 2 { // create + robot
		t.Fatalf("commands count = %d, want 2", len(cmds))
	}
}

func TestRenderHelperSchema_Product(t *testing.T) {
	root := buildHelperTestTree()
	payload, ok, err := renderHelperSchema(root, "dev")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected claim")
	}
	if payload["level"] != "product" || payload["count"] != 3 {
		t.Fatalf("payload = %#v", payload)
	}
	product, _ := payload["product"].(map[string]any)
	tools, _ := product["tools"].([]map[string]any)
	if product["id"] != "dev" || len(tools) != 3 {
		t.Fatalf("product = %#v, want dev with three tools", product)
	}
}

func TestRenderHelperSchema_LocalLeaf(t *testing.T) {
	root := buildHelperTestTree()
	payload, ok, err := renderHelperSchema(root, "dev connect")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected local helper leaf to be claimed")
	}
	if payload["source"] != "hardcoded:dev" {
		t.Fatalf("source = %v", payload["source"])
	}
	if payload["canonical_path"] != "dev.connect" || payload["cli_path"] != "dev connect" {
		t.Fatalf("payload paths = %#v", payload)
	}
	params, _ := payload["parameters"].(map[string]any)
	if _, ok := params["robot-client-id"]; !ok {
		t.Fatalf("local helper flag missing from parameters: %#v", params)
	}
}

func TestHelperProductSummariesIncludeLocalCommands(t *testing.T) {
	root := buildHelperTestTree()
	summaries := helperProductSummaries(root)
	if len(summaries) != 1 {
		t.Fatalf("summaries len = %d, want 1", len(summaries))
	}
	tools, _ := summaries[0]["tools"].([]map[string]any)
	found := false
	for _, tool := range tools {
		if tool["cli_path"] == "dev connect" {
			found = true
			if tool["source"] != "hardcoded:dev" {
				t.Fatalf("local helper source = %v", tool["source"])
			}
		}
	}
	if !found {
		t.Fatalf("local helper command missing from schema list: %#v", tools)
	}
}

func TestRenderHelperSchema_UnknownSubcommand(t *testing.T) {
	root := buildHelperTestTree()
	payload, ok, err := renderHelperSchema(root, "dev app nope")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected claim")
	}
	if payload["error"] == nil {
		t.Fatalf("expected error for unknown subcommand, got %#v", payload)
	}
	if avail, _ := payload["available"].([]map[string]any); len(avail) == 0 {
		t.Fatal("expected available subcommands listed")
	}
}

func TestRenderHelperSchema_AnnotatedLeafNeedsNoDiscovery(t *testing.T) {
	root := buildHelperTestTree()
	payload, ok, err := renderHelperSchema(root, "dev app create")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatal("expected claim")
	}
	if payload["canonical_path"] != "dev.create_dev_app" {
		t.Fatalf("canonical_path = %v", payload["canonical_path"])
	}
	if payload["source"] != "hardcoded:dev" {
		t.Fatalf("source = %v", payload["source"])
	}
}

func TestRenderHelperSchema_NonHelperPathDeclined(t *testing.T) {
	root := buildHelperTestTree()
	if _, ok, _ := renderHelperSchema(root, "ding.message.send"); ok {
		t.Fatal("non-helper path must not be claimed by the helper renderer")
	}
}

func TestKebabCase(t *testing.T) {
	cases := map[string]string{
		"eventCallbackUrl": "event-callback-url",
		"unifiedAppId":     "unified-app-id",
		"disableSSLVerify": "disable-ssl-verify",
		"mode":             "mode",
		"skills":           "skills",
		"i18nName":         "i18n-name",
	}
	for in, want := range cases {
		if got := kebabCase(in); got != want {
			t.Errorf("kebabCase(%q) = %q, want %q", in, got, want)
		}
	}
}
