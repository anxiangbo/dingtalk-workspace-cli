# 权限管理

查询、申请、取消开放平台应用的 APP 应用权限和 SNS 个人权限。

## 权限列表

```bash
dws devapp permission list --unified-app-id ID --format json
dws devapp permission list --unified-app-id ID --keyword "机器人发消息" --status UNAUTHED --format json
dws devapp permission list --unified-app-id ID --scope-type SNS --format json
dws devapp permission list --unified-app-id ID --scope qyapi_robot_sendmsg --format json
```

MCP tool: `list_open_dev_app_permissions`

| CLI | MCP | 说明 |
|-----|-----|------|
| `--unified-app-id` | `unifiedAppId` | 应用定位 |
| `--agent-id` | `agentId` | 应用定位 |
| `--keyword` | `keyword` | 权限名/API 名关键词 |
| `--status` | `authStatus` | `ALL` / `AUTHED` / `UNAUTHED` |
| `--scope-type` | `firstLevelType` | `APP` / `SNS`，为空返回两者 |
| `--scope` | `scopeValue` | 单权限详情模式 |
| `--limit` | `limit` | 默认 20 |

**规则：**
- `permission search` 和 `permission detail` 是 `list` 的 CLI 别名，不是独立 MCP tool。
- 默认同时返回 APP 和 SNS 权限，无服务端分页。
- 列表模式只返回 `apiPreview`；`--scope` 详情模式返回完整 `apiList`。

**scopeValue 选择顺序：**

1. 用户给了 `scopeValue` → 精确匹配
2. 用户给了 API 名 → `keyword` 搜索，匹配 `apiPreview.name`
3. 用户给了权限名 → 匹配 `scopeName/scopeDesc`
4. 多个候选 → 展示列表让用户选择，不自动取第一条

## 申请权限

```bash
dws devapp permission add --unified-app-id ID --permissions qyapi_robot_sendmsg --dry-run --format json
dws devapp permission add --unified-app-id ID --permissions Contact.User.mobile,qyapi_robot_sendmsg --yes --format json
```

MCP tool: `apply_open_dev_app_permissions`

**规则：**
- `--permissions` 传 `scopeValue`，多个逗号分隔，必须来自 `permission list` 的返回。
- 已开通跳过，不可编辑拒绝。
- `requiredApproval=true` 允许申请——写入版本变更，审批在版本发布时处理。
- 不在此处选审批人。

## 取消权限

```bash
dws devapp permission remove --unified-app-id ID --permission qyapi_robot_sendmsg --dry-run --format json
dws devapp permission remove --unified-app-id ID --permission qyapi_robot_sendmsg --yes --format json
```

MCP tool: `remove_open_dev_app_permission`

一次只取消一个权限点。未开通返回 `NOT_AUTHED`；不可编辑返回 no-edit 原因。
