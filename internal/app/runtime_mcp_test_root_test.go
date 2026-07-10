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
	"os"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cli"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/pipeline"
	"github.com/spf13/cobra"
)

// newRuntimeMCPTestRoot exposes the catalog-driven command builder only to
// runner tests. The production root deliberately does not register `dws mcp`.
func newRuntimeMCPTestRoot(ctx context.Context, engine *pipeline.Engine) *cobra.Command {
	if ctx == nil {
		ctx = context.Background()
	}
	flags := &GlobalFlags{}
	loader := cli.EnvironmentLoader{
		LookupEnv:              os.LookupEnv,
		CatalogBaseURLOverride: DiscoveryBaseURL(),
		AuthTokenFunc: func(ctx context.Context) string {
			return resolveRuntimeAuthToken(ctx, "")
		},
		LoggerFunc: FileLoggerInstance,
	}
	runner := newCommandRunnerWithFlags(loader, flags)
	root := &cobra.Command{
		Use:           "dws",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	bindPersistentFlags(root, flags)
	root.AddCommand(cli.NewMCPCommand(ctx, loader, runner, engine))
	root.SetFlagErrorFunc(flagErrorWithSuggestions)
	root.SetContext(ctx)
	return root
}
