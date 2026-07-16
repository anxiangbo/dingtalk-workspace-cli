# Multi-Account Profiles Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 支持同一组织多个账号并保持历史 profile 数据、Token 镜像和命令语义兼容。

**Architecture:** `profiles.json` 保持 v1 和扁平数组，唯一性从 `corpId` 改为 `corpId + userId`。账号 Token 使用哈希身份槽位作为事实源，现有组织槽位和全局槽位继续作为兼容镜像；所有选择器先解析成实际身份后再执行。

**Tech Stack:** Go、Cobra、平台 Keychain/加密文件、Go testing。

## Global Constraints

- 不改变 `profiles.json` 的版本、已有字段名和顶层结构。
- 历史 `--profile <corpId|组织名>` 行为不变。
- `profile list` 默认展示全部账号。
- 所有修改的 Go 文件必须运行 `gofmt`。
- 使用测试驱动：每个行为先写失败测试并确认预期失败。

---

### Task 1: 身份选择器和 Profile 唯一性

**Files:**
- Modify: `internal/auth/profiles.go`
- Test: `internal/auth/token_test.go`

**Interfaces:**
- Produces: `ProfileSelector(Profile) string`
- Produces: `ParseIdentitySelector(string) (corpID, userID string, ok bool)`
- Produces: `ResolveProfile(configDir, selector string) (*Profile, error)` 支持组织和精确身份。

- [ ] 写失败测试：相同 `corpId`、不同 `userId` 保存后保留两条记录。
- [ ] 写失败测试：相同 `corpId + userId` 重登只更新对应记录。
- [ ] 写失败测试：normalize 只删除相同身份，不删除同组织其他账号。
- [ ] 运行 `go test ./internal/auth -run 'Test.*SameCorp|Test.*IdentitySelector' -count=1`，确认因当前按 `corpId` 去重而失败。
- [ ] 实现身份选择器、身份查找、组织账号集合和确定性组织当前账号解析。
- [ ] 运行上述测试确认通过。

### Task 2: 账号 Token 事实源和兼容镜像

**Files:**
- Modify: `internal/auth/keychain_store.go`
- Modify: `internal/auth/token.go`
- Modify: `internal/auth/profiles.go`
- Test: `internal/auth/token_test.go`
- Test: `internal/auth/token_preflight_test.go`

**Interfaces:**
- Produces: `TokenAccountForIdentity(corpID, userID string) string`
- Produces: `LoadTokenDataKeychainForIdentity(corpID, userID string) (*TokenData, error)`
- Produces: `SaveTokenDataKeychainForIdentity(corpID, userID string, data *TokenData) error`
- Produces: `DeleteTokenDataKeychainForIdentity(corpID, userID string) error`

- [ ] 写失败测试：同组织两个账号的 Token 均可精确读取。
- [ ] 写失败测试：组织 Token 镜像指向最后登录或显式切换账号。
- [ ] 写失败测试：历史组织 Token 自动复制到身份槽位且保持原槽位。
- [ ] 运行专项测试确认失败原因是只有组织 Token 槽位。
- [ ] 实现 SHA-256 身份槽位和兼容镜像同步。
- [ ] 修改预检、全量删除和迁移逻辑覆盖身份槽位。
- [ ] 运行 `go test ./internal/auth -count=1`。

### Task 3: 无覆盖窗口的登录身份补全

**Files:**
- Modify: `internal/app/tool_caller_adapter.go`
- Modify: `internal/app/auth_command.go`
- Modify: `internal/auth/oauth_provider.go`
- Modify: `internal/auth/device_flow.go`
- Test: `internal/app/auth_command_test.go`

**Interfaces:**
- Produces: internal optional interface `CallToolWithToken(ctx, token, productID, toolName, args)`.
- Produces: `enrichAuthLoginTokenBeforePersist(...)`，在持久化前补齐身份。

- [ ] 写失败测试：已有同组织账号时，新登录 Token 在身份识别完成前不会覆盖旧 Token。
- [ ] 写失败测试：身份识别失败保持原 profiles 和 Token 不变。
- [ ] 写失败测试：识别成功后同组织新增账号并成为组织当前账号。
- [ ] 运行 `go test ./internal/app -run 'TestAuthLogin.*Identity|TestAuthLogin.*SameCorp' -count=1` 确认失败。
- [ ] 给 ToolCaller adapter 增加进程内临时 Token 调用能力。
- [ ] 把 OAuth、设备流和 exchange 的最终持久化移动到身份补全之后。
- [ ] 运行专项测试确认通过。

### Task 4: 切换、指针和退出

**Files:**
- Modify: `internal/auth/profiles.go`
- Modify: `internal/auth/token.go`
- Modify: `internal/app/profile_command.go`
- Modify: `internal/app/auth_command.go`
- Test: `internal/auth/token_test.go`
- Test: `internal/app/profile_command_test.go`
- Test: `internal/app/auth_command_test.go`

**Interfaces:**
- `SetCurrentProfile` 接受组织或身份选择器。
- `UsePreviousProfile` 原样交换选择器。
- `DeleteTokenDataForProfile` 对组织选择器删除全部账号，对身份选择器只删除一个账号。

- [ ] 写失败测试：三个指针分别保存和恢复 `corpId:userId`。
- [ ] 写失败测试：`switch corpId` 使用组织当前账号，`switch corpId:userId` 更新组织当前账号。
- [ ] 写失败测试：身份级退出保留同组织其他账号；组织级退出删除全部账号。
- [ ] 运行专项测试确认失败。
- [ ] 实现指针、镜像和退出状态转换。
- [ ] 运行 `go test ./internal/auth ./internal/app -run 'Profile|AuthLogout' -count=1`。

### Task 5: 列表、批量执行和长生命周期进程

**Files:**
- Modify: `internal/app/profile_command.go`
- Modify: `internal/app/runner.go`
- Modify: `internal/app/event_personal_command.go`
- Test: `internal/app/profile_command_test.go`
- Test: `internal/app/multi_profile_runner_test.go`
- Test: `internal/app/event_stdin_gate_test.go`

**Interfaces:**
- Profile JSON 增加 `profile` 和 `isOrgCurrent`，保留所有旧字段。
- 多 profile 结果增加 `userId/userName/profile`，保留旧字段。

- [ ] 写失败测试：profile list 默认返回同组织全部账号。
- [ ] 写失败测试：CSV 解析后按实际身份去重。
- [ ] 写失败测试：Event 子进程转发 `corpId:userId`。
- [ ] 运行专项测试确认失败。
- [ ] 实现列表标记、身份去重和子进程选择器透传。
- [ ] 运行 `go test ./internal/app -run 'ProfileList|MultiProfile|PersonalBusSpawn' -count=1`。

### Task 6: 导入导出、E2E 和文档

**Files:**
- Modify: `internal/auth/portable_store_test.go`
- Modify: `scripts/dev/test-multi-profile-e2e.sh`
- Modify: `README.md`
- Modify: `skills/mono/SKILL.md`
- Modify: `skills/multi/dingtalk-profile/SKILL.md`
- Modify: `skills/multi/dws-shared/SKILL.md`

- [ ] 写失败测试：认证包往返后同组织两个账号均可读取。
- [ ] 扩展 E2E：种入同组织双账号，验证列表、切换、一次性执行和退出。
- [ ] 更新 README 和 Skill，说明组织选择器与身份选择器。
- [ ] 运行 `go test ./internal/auth -run Portable -count=1`。
- [ ] 运行 `bash scripts/dev/test-multi-profile-e2e.sh --skip-go-tests`。

### Task 7: 完整验证

**Files:**
- Verify only.

- [ ] 运行 `gofmt` 格式化所有修改的 Go 文件。
- [ ] 运行 `go test ./internal/auth ./internal/app -count=1`。
- [ ] 运行 `DWS_PACKAGE_VERSION=0.0.0-test go test ./...`。
- [ ] 运行 `go build ./cmd`。
- [ ] 运行 `git diff --check`。
- [ ] 检查 `git status --short` 和完整 diff，确认没有生成文件漂移或无关改动。
