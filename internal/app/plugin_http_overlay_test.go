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
	"testing"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/market"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/plugin"
)

func TestRegisterHTTPServerFromOverlayBuildsWithoutDiscovery(t *testing.T) {
	p := &plugin.Plugin{Manifest: plugin.Manifest{Name: "http-overlay"}}
	srv := market.ServerDescriptor{
		Key:      "offline-http",
		Endpoint: "http://127.0.0.1:1/mcp",
		CLI: market.CLIOverlay{
			ID:      "http-overlay",
			Command: "http-overlay",
			ToolOverrides: map[string]market.CLIToolOverride{
				"ping": {CLIName: "ping", Description: "Ping the service"},
			},
		},
		HasCLIMeta: true,
	}

	cmds := registerHTTPServerFromOverlay(p, srv, executor.EchoRunner{}, nil)
	if len(cmds) != 1 || cmds[0].Name() != "http-overlay" {
		t.Fatalf("commands = %#v, want one http-overlay command", cmds)
	}
	if child, _, err := cmds[0].Find([]string{"ping"}); err != nil || child == nil || child.Name() != "ping" {
		t.Fatalf("overlay leaf not registered: child=%v err=%v", child, err)
	}
}

func TestRegisterHTTPServerWithoutOverridesSkipsCommand(t *testing.T) {
	p := &plugin.Plugin{Manifest: plugin.Manifest{Name: "http-no-overlay"}}
	srv := market.ServerDescriptor{
		Key:        "offline-http",
		Endpoint:   "http://127.0.0.1:1/mcp",
		CLI:        market.CLIOverlay{ID: "http-no-overlay", Command: "http-no-overlay"},
		HasCLIMeta: true,
	}
	if cmds := registerHTTPServerFromOverlay(p, srv, executor.EchoRunner{}, nil); len(cmds) != 0 {
		t.Fatalf("commands = %#v, want none without toolOverrides", cmds)
	}
}
