// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/app"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/generator/agentmetadata"
)

func main() {
	var root string
	var skillPath string
	var productsDir string
	var intentGuidePath string
	var outputPath string
	var maxExamples int
	var validateSurface bool
	flag.StringVar(&root, "root", ".", "Repository root")
	flag.StringVar(&skillPath, "skill", "skills/mono/SKILL.md", "Main DWS SKILL.md path")
	flag.StringVar(&productsDir, "products", "skills/mono/references/products", "Product skill reference directory")
	flag.StringVar(&intentGuidePath, "intent-guide", "skills/mono/references/intent-guide.md", "Cross-product intent guide path")
	flag.StringVar(&outputPath, "output", "internal/cli/schema_agent_metadata.json", "Output embedded Agent metadata JSON")
	flag.IntVar(&maxExamples, "max-examples", 2, "Maximum examples retained per command")
	flag.BoolVar(&validateSurface, "validate-surface", true, "Keep only paths present in the current runtime command schema")
	flag.Parse()
	var surface commandSurface
	if validateSurface {
		var err error
		surface, err = loadCommandSurface()
		if err != nil {
			fail(fmt.Errorf("load runtime command surface: %w", err))
		}
	}

	metadata, stats, err := agentmetadata.Generate(agentmetadata.Options{
		Root:             root,
		SkillPath:        skillPath,
		ProductsDir:      productsDir,
		IntentGuidePath:  intentGuidePath,
		MaxExamples:      maxExamples,
		ToolPaths:        surface.ToolPaths,
		ProductIDs:       surface.ProductIDs,
		SurfaceHash:      surface.Hash,
		SurfaceToolCount: surface.ToolCount,
	})
	if err != nil {
		fail(err)
	}
	encoded, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		fail(fmt.Errorf("encode metadata: %w", err))
	}
	encoded = append(encoded, '\n')
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		fail(fmt.Errorf("create output directory: %w", err))
	}
	if err := os.WriteFile(outputPath, encoded, 0o644); err != nil {
		fail(fmt.Errorf("write output: %w", err))
	}
	_, _ = fmt.Fprintf(
		os.Stderr,
		"generated schema Agent metadata: output=%s sources=%d products=%d tools=%d intents=%d examples=%d risk_rules=%d unmatched=%d surface_tools=%d\n",
		outputPath,
		stats.SourceFiles,
		stats.Products,
		stats.Tools,
		stats.ToolIntents,
		stats.Examples,
		stats.RiskRules,
		stats.UnmatchedTools,
		surface.ToolCount,
	)
}

type commandSurface struct {
	ToolPaths  map[string]string
	ProductIDs map[string]bool
	Hash       string
	ToolCount  int
}

func loadCommandSurface() (commandSurface, error) {
	root := app.NewRootCommand()
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	root.SetArgs([]string{"schema", "--all", "--format", "json"})
	if err := root.Execute(); err != nil {
		return commandSurface{}, fmt.Errorf("execute dws schema --all: %w: %s", err, strings.TrimSpace(stderr.String()))
	}
	var payload struct {
		Products []struct {
			ID    string `json:"id"`
			Tools []struct {
				CLIPath string   `json:"cli_path"`
				Aliases []string `json:"aliases"`
			} `json:"tools"`
		} `json:"products"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		return commandSurface{}, fmt.Errorf("decode schema catalog: %w", err)
	}
	surface := commandSurface{
		ToolPaths:  map[string]string{},
		ProductIDs: map[string]bool{},
	}
	rows := make([]string, 0)
	for _, product := range payload.Products {
		productID := strings.TrimSpace(product.ID)
		if productID != "" {
			surface.ProductIDs[productID] = true
		}
		for _, tool := range product.Tools {
			primary := strings.TrimSpace(tool.CLIPath)
			if primary == "" {
				continue
			}
			surface.ToolPaths[primary] = primary
			surface.ToolCount++
			aliases := append([]string(nil), tool.Aliases...)
			sort.Strings(aliases)
			for _, alias := range aliases {
				alias = strings.TrimSpace(alias)
				if alias != "" {
					surface.ToolPaths[alias] = primary
				}
			}
			rows = append(rows, productID+"\x00"+primary+"\x00"+strings.Join(aliases, "\x00"))
		}
	}
	sort.Strings(rows)
	sum := sha256.Sum256([]byte(strings.Join(rows, "\n")))
	surface.Hash = "sha256:" + hex.EncodeToString(sum[:])
	return surface, nil
}

func fail(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "generate-schema-agent-metadata: %v\n", err)
	os.Exit(1)
}
