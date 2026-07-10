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

package app

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cache"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/compat"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/market"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/plugin"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/transport"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/pkg/config"
	"github.com/spf13/cobra"
)

// resolveStdioOverlay resolves the CLIOverlay for a stdio plugin server
// from its manifest. It supports two forms:
//
//  1. inline JSON object in manifest.MCPServers[key].CLI
//  2. a relative file path (JSON string) pointing to an external overlay
//     file anchored at the plugin root (e.g. "overlay.json")
//
// When no CLI metadata is present, a minimal overlay keyed by the server
// name is returned so callers can still build an identity descriptor.
func resolveStdioOverlay(p *plugin.Plugin, sc plugin.StdioServerClient) market.CLIOverlay {
	serverID := sc.Key
	overlay := market.CLIOverlay{
		ID:      serverID,
		Command: serverID,
	}
	srv, ok := p.Manifest.MCPServers[sc.Key]
	if !ok || len(srv.CLI) == 0 {
		return overlay
	}

	cliData := srv.CLI
	// A JSON string is interpreted as a relative path to an external
	// overlay file (e.g. "overlay.json") anchored at the plugin root.
	if len(cliData) > 0 && cliData[0] == '"' {
		var cliPath string
		if err := json.Unmarshal(cliData, &cliPath); err == nil && cliPath != "" {
			absPath := filepath.Join(p.Root, cliPath)
			if fileData, readErr := os.ReadFile(absPath); readErr == nil {
				cliData = fileData
			} else {
				slog.Warn("plugin: failed to read CLI overlay file",
					"plugin", p.Manifest.Name, "path", absPath, "error", readErr)
			}
		}
	}
	if err := json.Unmarshal(cliData, &overlay); err != nil {
		slog.Warn("plugin: failed to parse CLI overlay for stdio server",
			"plugin", p.Manifest.Name, "server", sc.Key, "error", err)
	}
	if overlay.ID == "" {
		overlay.ID = serverID
	}
	if overlay.Command == "" {
		overlay.Command = serverID
	}
	return overlay
}

// toolsToDetails converts discovered ToolDescriptors to the DetailTool map
// shape expected by compat.BuildDynamicCommands (keyed by overlay ID).
// Returns nil if tools is empty.
func toolsToDetails(tools []transport.ToolDescriptor, overlayID string) map[string][]market.DetailTool {
	if len(tools) == 0 {
		return nil
	}
	detailTools := make([]market.DetailTool, 0, len(tools))
	for _, tool := range tools {
		schemaJSON := ""
		if tool.InputSchema != nil {
			if data, marshalErr := json.Marshal(tool.InputSchema); marshalErr == nil {
				schemaJSON = string(data)
			}
		}
		detailTools = append(detailTools, market.DetailTool{
			ToolName:    tool.Name,
			ToolTitle:   tool.Title,
			ToolDesc:    tool.Description,
			IsSensitive: tool.Sensitive,
			ToolRequest: schemaJSON,
		})
	}
	return map[string][]market.DetailTool{overlayID: detailTools}
}

// registerStdioServerFromOverlay builds cobra commands for a stdio plugin
// server using only its manifest + overlay.json — no subprocess required.
//
// Returns (cmds, descriptor, true) when the overlay carries toolOverrides,
// otherwise (nil, zero, false). Startup skips such plugins because runtime
// discovery is disabled.
//
// When a warm tools cache exists for this server, its DetailTools are passed
// to BuildDynamicCommands so flag types are enriched from the last successful
// discovery. Fresh installs (or evicted caches) get overlay-declared flags
// only.
func registerStdioServerFromOverlay(
	p *plugin.Plugin,
	sc plugin.StdioServerClient,
	runner executor.Runner,
	store *cache.Store,
) ([]*cobra.Command, market.ServerDescriptor, bool) {
	overlay := resolveStdioOverlay(p, sc)
	if len(overlay.ToolOverrides) == 0 {
		return nil, market.ServerDescriptor{}, false
	}

	descriptor := market.ServerDescriptor{
		Key:         sc.Key,
		DisplayName: p.Manifest.Name + "/" + sc.Key,
		Description: p.Manifest.Description,
		Endpoint:    StdioEndpoint(p.Manifest.Name, sc.Key),
		Source:      "plugin",
		CLI:         overlay,
		HasCLIMeta:  true,
	}

	AppendDynamicServer(descriptor)
	RegisterStdioClient(p.Manifest.Name+"/"+sc.Key, sc.Client)

	// Migration enrichment: an existing snapshot from an older version may
	// still enrich flag types, but cache absence never triggers discovery.
	var detailsByID map[string][]market.DetailTool
	if store != nil {
		cacheKey := pluginCacheKey(p.Manifest.Name, sc.Key)
		if snapshot, _, err := store.LoadTools(config.DefaultPartition, cacheKey); err == nil && len(snapshot.Tools) > 0 {
			detailsByID = toolsToDetails(snapshot.Tools, overlay.ID)
		}
	}

	cmds := compat.BuildDynamicCommands(
		[]market.ServerDescriptor{descriptor}, runner, detailsByID)

	slog.Debug("plugin: stdio server registered from overlay",
		"plugin", p.Manifest.Name, "server", sc.Key,
		"toolOverrides", len(overlay.ToolOverrides),
		"commands", len(cmds),
		"enriched", detailsByID != nil)

	return cmds, descriptor, true
}
