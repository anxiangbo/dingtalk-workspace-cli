# 应用基础操作

应用列表查询、详情、创建、修改、生命周期启停和删除。

## 应用定位

写操作前必须定位到唯一应用：

| 优先级 | 标识 | 处理 |
|--------|------|------|
| 1 | `--unified-app-id` | 直接使用 |
| 2 | `--agent-id` / `--app-id` | 直接使用 |
| 3 | `--app-key` / `--custom-key` | 先查询，唯一命中才继续 |
| 4 | `--name` | 模糊搜索，写操作必须唯一命中 |

## 应用列表

```bash
dws devapp list --format json
dws devapp list --name DemoApp --format json
dws devapp list --agent-id 123456 --format json
dws devapp list --creator 张三 --sort gmt_modified --order desc --format json
```

MCP tool: `list_open_dev_apps_by_condition`

| CLI | MCP | 说明 |
|-----|-----|------|
| `--page` | `currentPage` | 1-based，默认 1 |
| `--page-size` | `pageSize` | 默认 20 |
| `--name` / `--keyword` | `appName` | 应用名搜索 |
| `--agent-id` | `agentId` | 精确定位 |
| `--app-key` | `appKey` | appKey/clientId |
| `--creator` | `creator` | 创建人关键词 |
| `--sort` | `sortType` | 如 `gmt_modified` |
| `--order` | `sortOrder` | `asc` / `desc` |

## 应用详情

```bash
dws devapp get --unified-app-id UNIFIED_APP_ID --format json
dws devapp get --agent-id 123456 --format json
```

MCP tool: `get_open_dev_app_detail`

详情展示 `agentId/clientId/appKey`，但不能用来读 `clientSecret/appSecret`。

## 创建应用

```bash
dws devapp create --name DemoApp --desc "内部应用" --type internal --dry-run --format json
dws devapp create --name DemoApp --desc "内部应用" --type internal --yes --format json
```

MCP tool: `create_inner_app`

| CLI | MCP | 必填 |
|-----|-----|------|
| `--name` | `appName` | 是 |
| `--desc` | `appDesc` | 否 |
| `--icon` | `appIcon` | 否 |

`--type` 只做 CLI 校验（当前仅支持 `internal`），不下发 MCP。

## 修改应用

```bash
dws devapp update --unified-app-id ID --name DemoApp2 --desc "新描述" --dry-run --format json
dws devapp update --unified-app-id ID --name DemoApp2 --desc "新描述" --yes --format json
```

MCP tool: `update_inner_app`

至少提供一个更新字段：`--name` / `--desc` / `--icon`。

## 停用 / 启用应用

```bash
dws devapp inactive --unified-app-id ID --dry-run --format json
dws devapp inactive --unified-app-id ID --yes --format json

dws devapp active --unified-app-id ID --dry-run --format json
dws devapp active --unified-app-id ID --yes --format json
```

MCP tools: `inactive_inner_app` / `active_inner_app`

停用保留数据但应用不可用，可通过 `active` 恢复。

## 删除应用

```bash
dws devapp delete --unified-app-id ID --dry-run --format json
dws devapp delete --unified-app-id ID --yes --format json
```

MCP tool: `delete_inner_app`

删除前必须展示应用摘要。删除为异步操作，成功后应用延迟从列表消失。

## 错误处理

| 情况 | 处理 |
|------|------|
| `unknown command` | CLI 构建不含 devapp helper |
| `endpoint_not_resolved` | 检查 edition endpoint 注入 |
| 多应用命中 | 展示候选，停止写操作 |
| `ServiceResult.success=false` | 透传 `errorCode/errorMsg` |
