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

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cli"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/pkg/edition"
)

func TestRootDoesNotRegisterCanonicalMCPCommand(t *testing.T) {
	withCleanDynamicRegistry(t)
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("DWS_CONFIG_DIR", home)
	t.Setenv(cli.CacheDirEnv, t.TempDir())

	previous := edition.Get()
	edition.Override(&edition.Hooks{
		Name: "no-discovery-test",
		StaticServers: func() []edition.ServerInfo {
			return nil
		},
	})
	t.Cleanup(func() { edition.Override(previous) })

	root := NewRootCommand()
	for _, command := range root.Commands() {
		if command.Name() == "mcp" {
			t.Fatal("root must not register the deprecated hidden 'dws mcp' discovery tree")
		}
	}
}
