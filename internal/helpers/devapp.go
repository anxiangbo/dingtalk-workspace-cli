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

package helpers

import (
	"strings"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cobracmd"
	apperrors "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/errors"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
	"github.com/spf13/cobra"
)

const (
	devAppProduct = "devapp"

	devAppMemberListTool     = "list_open_dev_app_members"
	devAppMemberAddTool      = "add_open_dev_app_members"
	devAppMemberRemoveTool   = "remove_open_dev_app_members"
	devAppSecurityConfigTool = "update_app_security_config"
)

func init() {
	RegisterPublic(func() Handler {
		return devAppHandler{}
	})
}

type devAppHandler struct{}

func (devAppHandler) Name() string {
	return "devapp"
}

func (devAppHandler) Command(runner executor.Runner) *cobra.Command {
	return newDevAppCommand(runner)
}

func newDevAppCommand(runner executor.Runner) *cobra.Command {
	root := &cobra.Command{
		Use:               "devapp",
		Short:             "开放平台应用",
		Long:              "管理开放平台开发者应用：成员查询、成员增删、安全配置等。",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	member := &cobra.Command{
		Use:               "member",
		Short:             "开放平台应用成员管理",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	member.AddCommand(
		newDevAppMemberListCommand(runner),
		newDevAppMemberAddCommand(runner),
		newDevAppMemberRemoveCommand(runner),
	)

	security := &cobra.Command{
		Use:               "security",
		Short:             "开放平台应用安全设置",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	security.AddCommand(newDevAppSecurityConfigCommand(runner))

	root.AddCommand(member, security)
	return root
}

func newDevAppMemberListCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list",
		Short:             "查询开放平台应用成员",
		Example:           "  dws devapp member list --app-id <unifiedAppId>",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			appID, err := requiredDevAppID(cmd)
			if err != nil {
				return err
			}
			return runDevAppTool(runner, cmd, devAppMemberListTool, map[string]any{
				"unifiedAppId": appID,
			})
		},
	}
	cmd.Flags().String("app-id", "", "开放平台统一应用 ID (必填)")
	preferLegacyLeaf(cmd)
	return cmd
}

func newDevAppMemberAddCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "add",
		Short:             "添加开放平台应用成员",
		Example:           "  dws devapp member add --app-id <unifiedAppId> --users userId1,userId2 --member-type DEVELOPER",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDevAppMemberMutation(runner, cmd, devAppMemberAddTool)
		},
	}
	registerDevAppMemberMutationFlags(cmd)
	preferLegacyLeaf(cmd)
	return cmd
}

func newDevAppMemberRemoveCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "remove",
		Short:             "移除开放平台应用成员",
		Example:           "  dws devapp member remove --app-id <unifiedAppId> --users userId1,userId2 --member-type DEVELOPER",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDevAppMemberMutation(runner, cmd, devAppMemberRemoveTool)
		},
	}
	registerDevAppMemberMutationFlags(cmd)
	preferLegacyLeaf(cmd)
	return cmd
}

func newDevAppSecurityConfigCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "更新开放平台应用安全配置",
		Example: "  dws devapp security config --app-id <unifiedAppId> " +
			"--ip-whitelist 103.211.230.150 --redirect-url https://example.com/callback --sso-url https://example.com/sso",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			appID, err := requiredDevAppID(cmd)
			if err != nil {
				return err
			}

			params := map[string]any{"unifiedAppId": appID}
			if values := parseDevAppListFlag(cmd, "ip-whitelist"); len(values) > 0 {
				params["ipWhiteList"] = values
			}
			if values := parseDevAppListFlag(cmd, "redirect-url"); len(values) > 0 {
				params["redirectUrls"] = values
			}
			if values := parseDevAppListFlag(cmd, "sso-url"); len(values) > 0 {
				params["otherAuthUrls"] = values
			}
			if len(params) == 1 {
				return apperrors.NewValidation("one of --ip-whitelist, --redirect-url, or --sso-url is required")
			}
			return runDevAppTool(runner, cmd, devAppSecurityConfigTool, params)
		},
	}
	cmd.Flags().String("app-id", "", "开放平台统一应用 ID (必填)")
	cmd.Flags().String("ip-whitelist", "", "出口 IP 白名单，多个用逗号或分号分隔")
	cmd.Flags().String("redirect-url", "", "登录重定向 URL，多个用逗号或分号分隔")
	cmd.Flags().String("sso-url", "", "端内免登地址，多个用逗号或分号分隔")
	preferLegacyLeaf(cmd)
	return cmd
}

func registerDevAppMemberMutationFlags(cmd *cobra.Command) {
	cmd.Flags().String("app-id", "", "开放平台统一应用 ID (必填)")
	cmd.Flags().String("users", "", "成员 userId 列表，多个用逗号分隔 (必填)")
	cmd.Flags().String("member-type", "", "成员类型，如 DEVELOPER (必填)")
}

func runDevAppMemberMutation(runner executor.Runner, cmd *cobra.Command, tool string) error {
	appID, err := requiredDevAppID(cmd)
	if err != nil {
		return err
	}
	users, err := requiredDevAppUsers(cmd)
	if err != nil {
		return err
	}
	memberType, err := requiredDevAppMemberType(cmd)
	if err != nil {
		return err
	}

	params := map[string]any{
		"unifiedAppId":  appID,
		"memberUserIds": users,
		"memberType":    memberType,
	}
	return runDevAppTool(runner, cmd, tool, params)
}

func runDevAppTool(runner executor.Runner, cmd *cobra.Command, tool string, params map[string]any) error {
	invocation := executor.NewHelperInvocation(
		cobracmd.LegacyCommandPath(cmd),
		devAppProduct,
		tool,
		params,
	)
	invocation.DryRun = commandDryRun(cmd)
	result, err := runner.Run(cmd.Context(), invocation)
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func requiredDevAppID(cmd *cobra.Command) (string, error) {
	appID, _ := cmd.Flags().GetString("app-id")
	appID = strings.TrimSpace(appID)
	if appID == "" {
		return "", apperrors.NewValidation("--app-id is required")
	}
	return appID, nil
}

func requiredDevAppUsers(cmd *cobra.Command) ([]string, error) {
	usersRaw, _ := cmd.Flags().GetString("users")
	if strings.TrimSpace(usersRaw) == "" {
		return nil, apperrors.NewValidation("--users is required")
	}
	users := splitDevAppList(usersRaw)
	if len(users) == 0 {
		return nil, apperrors.NewValidation("--users must contain at least one userId")
	}
	return users, nil
}

func requiredDevAppMemberType(cmd *cobra.Command) (string, error) {
	memberType, _ := cmd.Flags().GetString("member-type")
	memberType = strings.TrimSpace(memberType)
	if memberType == "" {
		return "", apperrors.NewValidation("--member-type is required")
	}
	return memberType, nil
}

func parseDevAppListFlag(cmd *cobra.Command, name string) []string {
	raw, _ := cmd.Flags().GetString(name)
	return splitDevAppList(raw)
}

func splitDevAppList(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	raw = strings.ReplaceAll(raw, ";", ",")
	parts := strings.Split(raw, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		if value := strings.TrimSpace(part); value != "" {
			values = append(values, value)
		}
	}
	return values
}
