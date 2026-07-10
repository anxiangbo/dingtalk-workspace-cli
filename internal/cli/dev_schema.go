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
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// helperSchemaRoots are top-level command names whose subtrees are helper-only
// hard-coded cobra commands. Their schema is projected from the executable
// Cobra flags plus embedded metadata. The mapping from a leaf command to its
// canonical tool comes from the `mcp-tool` annotation set by the helper.
var helperSchemaRoots = map[string]bool{"dev": true}

// renderHelperSchema builds the `dws schema` payload for helper-only command
// subtrees. Returns (payload, true) when the path targets a helper subtree;
// (nil, false) otherwise so the caller falls back to runtime command schema.
//
// Leaf commands render the same gws-flat object as runtime schema leaves.
// Group/root paths render a browse listing {path, commands:[...]} from the
// Cobra tree. This path performs no network I/O.
func renderHelperSchema(root *cobra.Command, rawPath string) (map[string]any, bool, error) {
	if root == nil {
		return nil, false, nil
	}
	tokens := splitSchemaPathTokens(rawPath)
	if len(tokens) == 0 || !helperSchemaRoots[tokens[0]] {
		return nil, false, nil
	}
	helperRoot, _, err := root.Find([]string{tokens[0]})
	if err != nil || helperRoot == nil || !helperRoot.HasParent() {
		return nil, false, nil
	}
	if len(tokens) == 1 {
		product, ok := helperProductSummary(helperRoot)
		if !ok {
			return nil, false, nil
		}
		return map[string]any{
			"kind":    "schema",
			"level":   "product",
			"count":   schemaProductToolCount(product),
			"product": product,
			"source":  "helper-command",
		}, true, nil
	}
	if strings.Contains(rawPath, ".") && len(tokens) == 2 {
		if leaf, ok := helperLeafByToolName(helperRoot, tokens[1]); ok {
			payload, err := helperLeafSchema(leaf)
			return payload, true, err
		}
		return nil, false, nil
	}

	target, rest, err := root.Find(tokens)
	if err != nil || target == nil {
		target = root
		rest = tokens[1:]
	}
	// Find resolves to the deepest matching command and returns trailing tokens
	// it couldn't match as (sub)commands. Any non-flag leftover means a typo'd
	// or unknown subcommand — surface it with the closest group's children.
	if unknown := firstNonFlag(rest); unknown != "" {
		return map[string]any{
			"path":      rawPath,
			"error":     "unknown subcommand \"" + unknown + "\" under \"" + helperCommandPath(target) + "\"",
			"available": helperSubcommands(target),
		}, true, nil
	}

	// A runnable leaf → emit its schema in gws-flat shape from Cobra and
	// versioned metadata.
	// A group → browse its subcommands.
	if target.Runnable() && !target.HasAvailableSubCommands() {
		payload, err := helperLeafSchema(target)
		return payload, true, err
	}

	return map[string]any{
		"path":     helperCommandPath(target),
		"commands": helperSubcommands(target),
	}, true, nil
}

// helperLeafSchema renders a helper leaf from its executable Cobra surface.
func helperLeafSchema(cmd *cobra.Command) (map[string]any, error) {
	return helperLocalLeafSchema(cmd), nil
}

func helperLocalLeafSchema(cmd *cobra.Command) map[string]any {
	meta := helperSchemaLeafForCommand(cmd)
	productID := meta.ProductID
	if productID == "" {
		productID = helperProductID(cmd)
	}
	toolName := meta.ToolName
	if toolName == "" {
		toolName = helperLocalToolName(cmd)
	}
	canonicalPath := meta.CanonicalPath
	if canonicalPath == "" {
		canonicalPath = productID + "." + toolName
	}
	hint := schemaHintForCanonicalPath(canonicalPath)
	constraints := runtimeCommandConstraints(cmd)
	parameters := runtimeCommandParameters(cmd, hint.Parameters, nil, constraints)
	if parameters == nil {
		parameters = map[string]any{}
	}
	primaryCLIPath := meta.PrimaryCLIPath
	if primaryCLIPath == "" {
		primaryCLIPath = helperCommandPath(cmd)
	}
	payload := map[string]any{
		"name":             toolName,
		"cli_name":         cmd.Name(),
		"canonical_path":   canonicalPath,
		"path":             canonicalPath,
		"cli_path":         helperCommandPath(cmd),
		"primary_cli_path": primaryCLIPath,
		"is_alias":         meta.IsAlias,
		"source":           "hardcoded:" + productID,
		"product_id":       productID,
		"display":          helperProductDisplay(cmd),
		"title":            strings.TrimSpace(cmd.Short),
		"description":      runtimeCommandDescription(cmd),
		"parameters":       parameters,
		"has_parameters":   len(parameters) > 0,
		"parameter_count":  len(parameters),
	}
	if rendered := runtimeConstraintsPayload(constraints); len(rendered) > 0 {
		payload["constraints"] = rendered
	}
	if positionals := runtimeCommandPositionals(cmd); len(positionals) > 0 {
		payload["positionals"] = positionals
	}
	if len(meta.Aliases) > 0 {
		payload["aliases"] = meta.Aliases
	}
	paths := []string{primaryCLIPath, helperCommandPath(cmd), canonicalPath}
	paths = append(paths, meta.Aliases...)
	applyAgentToolMetadata(payload, true, paths...)
	return payload
}

type helperSchemaLeaf struct {
	Command        *cobra.Command
	ProductID      string
	ToolName       string
	Source         string
	CanonicalPath  string
	CLIPath        string
	PrimaryCLIPath string
	Aliases        []string
	IsAlias        bool
}

func collectHelperSchemaLeaves(top *cobra.Command) []helperSchemaLeaf {
	if top == nil {
		return nil
	}
	productID := helperProductID(top)
	leaves := []helperSchemaLeaf{}
	walkLeafCommands(top, func(leaf *cobra.Command) {
		toolName, _ := helperLeafToolBinding(leaf)
		if toolName == "" {
			toolName = helperLocalToolName(leaf)
		}
		leaves = append(leaves, helperSchemaLeaf{
			Command:       leaf,
			ProductID:     productID,
			ToolName:      toolName,
			Source:        "hardcoded:" + productID,
			CanonicalPath: productID + "." + toolName,
			CLIPath:       helperCommandPath(leaf),
		})
	})
	sort.Slice(leaves, func(i, j int) bool {
		if leaves[i].CanonicalPath != leaves[j].CanonicalPath {
			return leaves[i].CanonicalPath < leaves[j].CanonicalPath
		}
		return leaves[i].CLIPath < leaves[j].CLIPath
	})
	annotateHelperSchemaAliases(leaves)
	return leaves
}

func annotateHelperSchemaAliases(leaves []helperSchemaLeaf) {
	groups := map[string][]int{}
	for idx, leaf := range leaves {
		groups[leaf.CanonicalPath] = append(groups[leaf.CanonicalPath], idx)
	}
	for _, indexes := range groups {
		primary := choosePrimaryHelperLeaf(leaves, indexes)
		aliases := make([]string, 0, len(indexes)-1)
		for _, idx := range indexes {
			if idx == primary {
				continue
			}
			aliases = append(aliases, leaves[idx].CLIPath)
		}
		sort.Strings(aliases)
		primaryPath := leaves[primary].CLIPath
		for _, idx := range indexes {
			leaves[idx].PrimaryCLIPath = primaryPath
			leaves[idx].Aliases = append([]string(nil), aliases...)
			leaves[idx].IsAlias = idx != primary
		}
	}
}

func choosePrimaryHelperLeaf(leaves []helperSchemaLeaf, indexes []int) int {
	if len(indexes) == 0 {
		return 0
	}
	if primaryHint := schemaPrimaryCLIPath(leaves[indexes[0]].ProductID, leaves[indexes[0]].ToolName); primaryHint != "" {
		for _, idx := range indexes {
			if leaves[idx].CLIPath == primaryHint {
				return idx
			}
		}
	}
	toolParts := map[string]bool{}
	for _, part := range strings.FieldsFunc(leaves[indexes[0]].ToolName, func(r rune) bool { return r == '_' || r == '-' }) {
		if part != "" {
			toolParts[part] = true
		}
	}
	for _, idx := range indexes {
		if toolParts[leaves[idx].Command.Name()] {
			return idx
		}
	}
	return indexes[0]
}

func helperSchemaLeafForCommand(cmd *cobra.Command) helperSchemaLeaf {
	top := topLevelCommand(cmd)
	for _, leaf := range collectHelperSchemaLeaves(top) {
		if leaf.Command == cmd {
			return leaf
		}
	}
	return helperSchemaLeaf{}
}

func helperLeafByToolName(top *cobra.Command, toolName string) (*cobra.Command, bool) {
	canonicalPath := helperProductID(top) + "." + strings.TrimSpace(toolName)
	for _, leaf := range collectHelperSchemaLeaves(top) {
		if leaf.CanonicalPath == canonicalPath && !leaf.IsAlias {
			return leaf.Command, true
		}
	}
	return nil, false
}

// helperProductSummaries returns light product entries for every helper-only
// subtree, appended to the no-arg `dws schema` product listing so agents
// browsing all products also see helper commands. Tools are listed by path +
// summary only; drill in with `dws schema "<path>"` for full parameter schema.
func helperProductSummaries(root *cobra.Command) []map[string]any {
	if root == nil {
		return nil
	}
	out := []map[string]any{}
	for name := range helperSchemaRoots {
		top, _, err := root.Find([]string{name})
		if err != nil || top == nil || !top.HasParent() {
			continue
		}
		if product, ok := helperProductSummary(top); ok {
			out = append(out, product)
		}
	}
	return out
}

func helperProductSummary(top *cobra.Command) (map[string]any, bool) {
	if top == nil {
		return nil, false
	}
	leafEntries := collectHelperSchemaLeaves(top)
	tools := []map[string]any{}
	for _, leaf := range leafEntries {
		if leaf.IsAlias {
			continue
		}
		tool := map[string]any{
			"name":             leaf.ToolName,
			"cli_name":         leaf.Command.Name(),
			"canonical_path":   leaf.CanonicalPath,
			"cli_path":         leaf.CLIPath,
			"primary_cli_path": leaf.PrimaryCLIPath,
			"source":           leaf.Source,
			"description":      strings.TrimSpace(leaf.Command.Short),
		}
		if len(leaf.Aliases) > 0 {
			tool["aliases"] = leaf.Aliases
		}
		paths := []string{leaf.PrimaryCLIPath, leaf.CLIPath, leaf.CanonicalPath}
		paths = append(paths, leaf.Aliases...)
		applyAgentToolMetadata(tool, false, paths...)
		tools = append(tools, tool)
	}
	if len(tools) == 0 {
		return nil, false
	}
	product := map[string]any{
		"id":          helperProductID(top),
		"name":        strings.TrimSpace(top.Short),
		"description": "helper 命令组；schema 来自可执行命令和版本内元数据",
		"helper":      true,
		"tool_count":  len(tools),
		"tools":       tools,
	}
	applyAgentProductMetadata(product, helperProductID(top))
	return product, true
}

// walkLeafCommands invokes fn for every runnable leaf under cmd (depth-first).
func walkLeafCommands(cmd *cobra.Command, fn func(*cobra.Command)) {
	if cmd.Runnable() && !cmd.HasAvailableSubCommands() {
		fn(cmd)
		return
	}
	for _, sub := range cmd.Commands() {
		if !sub.IsAvailableCommand() || sub.Name() == "help" {
			continue
		}
		walkLeafCommands(sub, fn)
	}
}

// helperSubcommands lists a group's runnable children for browse mode, sorted
// by name for deterministic output.
func helperSubcommands(cmd *cobra.Command) []map[string]any {
	out := []map[string]any{}
	for _, sub := range cmd.Commands() {
		if !sub.IsAvailableCommand() || sub.Name() == "help" {
			continue
		}
		if !helperCommandHasSchemaLeaf(sub) {
			continue
		}
		out = append(out, map[string]any{
			"cli_path":    helperCommandPath(sub),
			"description": strings.TrimSpace(sub.Short),
		})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i]["cli_path"].(string) < out[j]["cli_path"].(string)
	})
	return out
}

func helperLeafToolBinding(cmd *cobra.Command) (toolName, source string) {
	if cmd == nil || cmd.Annotations == nil {
		return "", ""
	}
	return strings.TrimSpace(cmd.Annotations["mcp-tool"]), strings.TrimSpace(cmd.Annotations["mcp-source"])
}

func helperCommandHasSchemaLeaf(cmd *cobra.Command) bool {
	if cmd == nil {
		return false
	}
	if cmd.Runnable() && !cmd.HasAvailableSubCommands() {
		return true
	}
	for _, sub := range cmd.Commands() {
		if !sub.IsAvailableCommand() || sub.Name() == "help" {
			continue
		}
		if helperCommandHasSchemaLeaf(sub) {
			return true
		}
	}
	return false
}

func helperLocalToolName(cmd *cobra.Command) string {
	parts := splitSchemaPathTokens(helperCommandPath(cmd))
	if len(parts) <= 1 {
		return "command"
	}
	return strings.ReplaceAll(strings.Join(parts[1:], "_"), "-", "_")
}

func helperProductID(cmd *cobra.Command) string {
	parts := splitSchemaPathTokens(helperCommandPath(cmd))
	if len(parts) > 0 {
		return parts[0]
	}
	return "helper"
}

func helperProductDisplay(cmd *cobra.Command) string {
	for c := cmd; c != nil && c.HasParent(); c = c.Parent() {
		if !c.Parent().HasParent() {
			return strings.TrimSpace(c.Short)
		}
	}
	return ""
}

// helperCommandPath returns the space-joined path from root to cmd, e.g.
// "dev app robot config".
func helperCommandPath(cmd *cobra.Command) string {
	parts := []string{}
	for c := cmd; c != nil && c.HasParent(); c = c.Parent() {
		parts = append([]string{c.Name()}, parts...)
	}
	return strings.Join(parts, " ")
}

// firstNonFlag returns the first token that is not a flag (does not start with
// "-"), or "" if there is none.
func firstNonFlag(tokens []string) string {
	for _, t := range tokens {
		if t != "" && !strings.HasPrefix(t, "-") {
			return t
		}
	}
	return ""
}
