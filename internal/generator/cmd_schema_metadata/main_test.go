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
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cache"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/market"
)

func TestMetadataFromRegistrySanitizesAndMapsOverrides(t *testing.T) {
	snapshot := cache.RegistrySnapshot{
		SavedAt: time.Now(),
		Servers: []market.ServerDescriptor{{
			Endpoint: "https://secret.example/server/calendar",
			CLI: market.CLIOverlay{
				ID: "calendar",
				ToolOverrides: map[string]market.CLIToolOverride{
					"list_calendar_events": {
						Description: "查询日程",
						Flags: map[string]market.CLIFlagOverride{
							"startTime": {Type: "string", Description: "开始时间"},
							"limit":     {Type: "int", Default: "20", Required: true},
						},
					},
					"search_my_robots": {
						ServerOverride: "bot",
						Description:    "查询机器人",
					},
					"legacy": {RedirectTo: "calendar event list", Description: "旧入口"},
				},
			},
		}},
	}

	metadata := metadataFromRegistry(snapshot)
	list := metadata.Tools["calendar.list_calendar_events"]
	if list.Description != "查询日程" || list.Parameters["startTime"].Type != "string" {
		t.Fatalf("calendar metadata = %#v", list)
	}
	limit := list.Parameters["limit"]
	if limit.Type != "integer" || limit.Default != "20" || limit.Required == nil || !*limit.Required {
		t.Fatalf("limit metadata = %#v", limit)
	}
	if got := metadata.Tools["bot.search_my_robots"].Description; got != "查询机器人" {
		t.Fatalf("serverOverride metadata = %q", got)
	}
	if _, exists := metadata.Tools["calendar.legacy"]; exists {
		t.Fatal("redirect metadata must be omitted")
	}

	encoded, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	if strings.Contains(string(encoded), "secret.example") || strings.Contains(string(encoded), "saved_at") {
		t.Fatalf("generated metadata leaked registry transport data: %s", encoded)
	}
}
