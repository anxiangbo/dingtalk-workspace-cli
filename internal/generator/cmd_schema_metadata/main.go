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

package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cache"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/market"
)

type catalog struct {
	Products []product `json:"products"`
}

type product struct {
	ID    string `json:"id"`
	Tools []tool `json:"tools"`
}

type tool struct {
	RPCName     string         `json:"rpc_name"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"input_schema"`
}

type metadataFile struct {
	Version    int                     `json:"version"`
	Source     string                  `json:"source"`
	SourceHash string                  `json:"source_hash"`
	Tools      map[string]toolMetadata `json:"tools"`
}

type toolMetadata struct {
	Title       string                   `json:"title,omitempty"`
	Description string                   `json:"description,omitempty"`
	Parameters  map[string]paramMetadata `json:"parameters,omitempty"`
}

type paramMetadata struct {
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	Default     string   `json:"default,omitempty"`
	Format      string   `json:"format,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Required    *bool    `json:"required,omitempty"`
}

func main() {
	var catalogPath string
	var registryPath string
	var outputPath string
	flag.StringVar(&catalogPath, "catalog", "docs/generated/schema/catalog.json", "Input catalog snapshot path")
	flag.StringVar(&registryPath, "registry", "", "Input cached CLI registry snapshot (takes precedence over --catalog)")
	flag.StringVar(&outputPath, "output", "internal/cli/schema_mcp_metadata.json", "Output embedded MCP metadata JSON")
	flag.Parse()

	inputPath := catalogPath
	source := "mcp-catalog"
	if strings.TrimSpace(registryPath) != "" {
		inputPath = registryPath
		source = "cli-registry"
	}
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fail(fmt.Errorf("read %s: %w", source, err))
	}

	var out metadataFile
	if source == "cli-registry" {
		var snapshot cache.RegistrySnapshot
		if err := json.Unmarshal(data, &snapshot); err != nil {
			fail(fmt.Errorf("decode CLI registry: %w", err))
		}
		out = metadataFromRegistry(snapshot)
	} else {
		var cat catalog
		if err := json.Unmarshal(data, &cat); err != nil {
			fail(fmt.Errorf("decode MCP catalog: %w", err))
		}
		out = metadataFromCatalog(cat)
	}
	out.Version = 1
	out.Source = source
	out.SourceHash = metadataHash(out.Tools)

	encoded, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fail(fmt.Errorf("encode metadata: %w", err))
	}
	encoded = append(encoded, '\n')
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		fail(fmt.Errorf("create output directory: %w", err))
	}
	if err := os.WriteFile(outputPath, encoded, 0o644); err != nil {
		fail(fmt.Errorf("write metadata: %w", err))
	}
	_, _ = fmt.Fprintf(os.Stderr, "generated schema interface metadata: output=%s source=%s tools=%d hash=%s\n", outputPath, source, len(out.Tools), out.SourceHash)
}

func metadataFromCatalog(cat catalog) metadataFile {
	out := metadataFile{Tools: map[string]toolMetadata{}}
	for _, product := range cat.Products {
		productID := strings.TrimSpace(product.ID)
		if productID == "" {
			continue
		}
		for _, tool := range product.Tools {
			toolName := strings.TrimSpace(tool.RPCName)
			if toolName == "" {
				continue
			}
			meta := toolMetadata{
				Title:       strings.TrimSpace(tool.Title),
				Description: strings.TrimSpace(tool.Description),
				Parameters:  parameterMetadata(tool.InputSchema),
			}
			if meta.Title == "" && meta.Description == "" && len(meta.Parameters) == 0 {
				continue
			}
			out.Tools[productID+"."+toolName] = meta
		}
	}
	return out
}

func metadataFromRegistry(snapshot cache.RegistrySnapshot) metadataFile {
	out := metadataFile{Tools: map[string]toolMetadata{}}
	for _, server := range snapshot.Servers {
		productID := firstNonEmpty(server.CLI.ID, server.CLI.Command)
		for _, tool := range server.CLI.Tools {
			toolName := strings.TrimSpace(tool.Name)
			if productID == "" || toolName == "" || tool.Hidden {
				continue
			}
			mergeMetadata(out.Tools, productID+"."+toolName, toolMetadata{
				Title:       strings.TrimSpace(tool.Title),
				Description: strings.TrimSpace(tool.Description),
			})
		}
		for toolName, override := range server.CLI.ToolOverrides {
			toolName = strings.TrimSpace(toolName)
			canonicalProduct := firstNonEmpty(override.ServerOverride, productID)
			if canonicalProduct == "" || toolName == "" || override.Hidden || strings.TrimSpace(override.RedirectTo) != "" {
				continue
			}
			mergeMetadata(out.Tools, canonicalProduct+"."+toolName, toolMetadata{
				Description: strings.TrimSpace(override.Description),
				Parameters:  registryParameterMetadata(override.Flags),
			})
		}
	}
	return out
}

func registryParameterMetadata(flags map[string]market.CLIFlagOverride) map[string]paramMetadata {
	out := map[string]paramMetadata{}
	for property, flag := range flags {
		property = strings.TrimSpace(property)
		if property == "" || flag.Hidden || flag.PipelineLocal {
			continue
		}
		meta := paramMetadata{
			Type:        registryFlagType(flag.Type),
			Description: strings.TrimSpace(flag.Description),
			Default:     strings.TrimSpace(flag.Default),
		}
		if flag.Required {
			required := true
			meta.Required = &required
		}
		if meta.Type == "" && meta.Description == "" && meta.Default == "" && meta.Required == nil {
			continue
		}
		out[property] = meta
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func registryFlagType(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "int", "int8", "int16", "int32", "int64", "integer":
		return "integer"
	case "bool", "boolean":
		return "boolean"
	case "stringslice", "stringarray", "array":
		return "array"
	case "string":
		return "string"
	default:
		return ""
	}
}

func mergeMetadata(tools map[string]toolMetadata, key string, incoming toolMetadata) {
	if incoming.Title == "" && incoming.Description == "" && len(incoming.Parameters) == 0 {
		return
	}
	existing := tools[key]
	if incoming.Title != "" {
		existing.Title = incoming.Title
	}
	if incoming.Description != "" {
		existing.Description = incoming.Description
	}
	if len(incoming.Parameters) > 0 {
		if existing.Parameters == nil {
			existing.Parameters = map[string]paramMetadata{}
		}
		for property, metadata := range incoming.Parameters {
			existing.Parameters[property] = metadata
		}
	}
	tools[key] = existing
}

func metadataHash(tools map[string]toolMetadata) string {
	data, _ := json.Marshal(tools)
	sum := sha256.Sum256(data)
	return fmt.Sprintf("sha256:%x", sum[:])
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			return value
		}
	}
	return ""
}

func parameterMetadata(schema map[string]any) map[string]paramMetadata {
	properties, _ := schema["properties"].(map[string]any)
	if len(properties) == 0 {
		return nil
	}
	required := map[string]bool{}
	for _, raw := range anySlice(schema["required"]) {
		if name, ok := raw.(string); ok && strings.TrimSpace(name) != "" {
			required[strings.TrimSpace(name)] = true
		}
	}

	keys := make([]string, 0, len(properties))
	for key := range properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	out := map[string]paramMetadata{}
	for _, key := range keys {
		prop, _ := properties[key].(map[string]any)
		if len(prop) == 0 {
			continue
		}
		meta := paramMetadata{
			Type:        stringField(prop, "type"),
			Description: firstStringField(prop, "description", "title"),
			Default:     defaultString(prop["default"]),
			Format:      stringField(prop, "format"),
			Enum:        stringEnum(prop["enum"]),
		}
		if required[key] {
			v := true
			meta.Required = &v
		}
		out[key] = meta
	}
	return out
}

func anySlice(raw any) []any {
	values, _ := raw.([]any)
	return values
}

func stringField(values map[string]any, key string) string {
	value, _ := values[key].(string)
	return strings.TrimSpace(value)
}

func firstStringField(values map[string]any, keys ...string) string {
	for _, key := range keys {
		if value := stringField(values, key); value != "" {
			return value
		}
	}
	return ""
}

func defaultString(raw any) string {
	if raw == nil {
		return ""
	}
	switch value := raw.(type) {
	case string:
		return strings.TrimSpace(value)
	default:
		return fmt.Sprint(value)
	}
}

func stringEnum(raw any) []string {
	values := anySlice(raw)
	if len(values) == 0 {
		return nil
	}
	out := make([]string, 0, len(values))
	for _, value := range values {
		text := strings.TrimSpace(fmt.Sprint(value))
		if text != "" {
			out = append(out, text)
		}
	}
	return out
}

func fail(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "generate-schema-mcp-metadata: %v\n", err)
	os.Exit(1)
}
