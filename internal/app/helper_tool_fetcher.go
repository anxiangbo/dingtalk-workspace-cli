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
	"context"
	"sync"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cli"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/transport"
)

// newDevAppToolFetcher returns a cli.HelperToolFetcher that loads the op-app
// (devapp) MCP tools/list LIVE and projects each tool into a cli.HelperToolSchema
// (name, description, inputSchema properties/required). It is injected into the
// schema command so the cli package can render `dws schema dev.*` from real
// server schema without importing app/transport.
//
// The fetch is memoized per process (sync.Once on success) so repeated
// `dws schema dev.*` in one invocation hit the network at most once. A failed
// fetch is not cached, allowing a later retry within the same process.
func newDevAppToolFetcher() cli.HelperToolFetcher {
	var (
		once   sync.Once
		cached map[string]cli.HelperToolSchema
	)
	return func(ctx context.Context) (map[string]cli.HelperToolSchema, error) {
		if cached != nil {
			return cached, nil
		}
		schemas, err := fetchDevAppToolSchemas(ctx)
		if err != nil {
			return nil, err
		}
		once.Do(func() { cached = schemas })
		return cached, nil
	}
}

// fetchDevAppToolSchemas performs the live tools/list call against the pinned
// op-app endpoint and converts the descriptors. Auth and identity headers are
// resolved the same way the runner does for direct-runtime invocations.
func fetchDevAppToolSchemas(ctx context.Context) (map[string]cli.HelperToolSchema, error) {
	token := resolveRuntimeAuthToken(ctx, "")
	headers := resolveIdentityHeaders()
	client := transport.NewClient(nil).WithAuth(token, headers)

	result, err := client.ListTools(ctx, devappEndpoint)
	if err != nil {
		return nil, err
	}

	out := make(map[string]cli.HelperToolSchema, len(result.Tools))
	for _, td := range result.Tools {
		out[td.Name] = cli.HelperToolSchema{
			Name:        td.Name,
			Description: td.Description,
			Properties:  inputSchemaProperties(td.InputSchema),
			Required:    inputSchemaRequired(td.InputSchema),
		}
	}
	return out, nil
}

// inputSchemaProperties pulls the "properties" object out of a deserialized
// MCP inputSchema map. Returns an empty (non-nil) map when absent.
func inputSchemaProperties(schema map[string]any) map[string]any {
	if schema == nil {
		return map[string]any{}
	}
	props, _ := schema["properties"].(map[string]any)
	if props == nil {
		return map[string]any{}
	}
	return props
}

// inputSchemaRequired pulls the "required" string list out of a deserialized
// MCP inputSchema map.
func inputSchemaRequired(schema map[string]any) []string {
	if schema == nil {
		return nil
	}
	raw, ok := schema["required"].([]any)
	if !ok {
		return nil
	}
	out := make([]string, 0, len(raw))
	for _, v := range raw {
		if s, ok := v.(string); ok && s != "" {
			out = append(out, s)
		}
	}
	return out
}
