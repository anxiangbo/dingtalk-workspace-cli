# 端到端工作流（发现 → 前置 → 调用 → 验证 → 恢复）

把运行闭环固化为可复制步骤。所有 dws 命令均 `--help` 核验；dws 仅为可选 CLI，与直接发 HTTP 等价。

## 0. 拿 access_token（调用前置）

```http
# 新版（推荐）：token 走业务请求 Header
POST https://api.dingtalk.com/v1.0/oauth2/accessToken
Content-Type: application/json

{"appKey":"<APP_KEY>","appSecret":"<APP_SECRET>"}
```
```http
# 旧版：token 走 query access_token
GET https://oapi.dingtalk.com/gettoken?appkey=<APP_KEY>&appsecret=<APP_SECRET>
```
- token 必须缓存，禁止每请求重换；日志只记 token hash / operation / path / errcode / requestId。
- 用 dws 时凭证读取走 `dws devapp credentials get --unified-app-id <ID> --format json`（敏感，不回显）。

## 1. 只读接口（最常见）

```bash
# 发现：先查 llm.md 定位文档，运行态可用时：
dws devdoc article search --query "获取用户信息" --format json
# 前置：未登录则 dws auth login（无头环境 dws auth login --device）
dws auth status --format json
# 调用：
dws api GET /v1.0/contact/users/me --format json
# 验证：HTTP 2xx + 业务字段；失败转 §4
```

## 2. 写接口（setup → target → verify → cleanup）

```bash
dws devapp permission list --unified-app-id ID --status UNAUTHED --format json
dws devapp permission add --unified-app-id ID --permissions qyapi_robot_sendmsg --dry-run --format json
#   → 展示摘要 → 用户当前轮确认 → 加 --yes
dws devapp permission add --unified-app-id ID --permissions qyapi_robot_sendmsg --yes --format json
dws devapp permission list --unified-app-id ID --scope qyapi_robot_sendmsg --format json   # verify
# cleanup：测试资源用完回收（取消权限 / 删除测试应用属 [危险]，需确认）
```

## 3. readback 三态

| 态 | 判据 | 下一步 |
| --- | --- | --- |
| 成功 | HTTP 2xx + 业务成功字段 | 继续后续动作 |
| `needs_permission` | 401/403 / 权限未授 | `dws devapp permission add`（→ §2）|
| `needs_resource` | 缺前置 ID / 资源 | 按文档 `resource_requirements` 用 `dws ...` 获取 |

## 4. 恢复与回流

1. 失败先加 `--verbose` 重试**一次**。
2. 有 requestId → `dws devdoc error diagnose --request-id <ID> --query "..." --format json`（[reference.md](reference.md)）。
3. 按错误码三元组取 next_action 自恢复；网关未注册记 `needs_gateway_tool_registration`。
4. 同一命令重试 **≤ 3 次**；仍失败 → 停止，附完整错误与 requestId 报告。

## 5. 危险操作确认清单（[危险]，必须用户确认）

| 产品域 | 命令 | 影响 |
| --- | --- | --- |
| devapp | `delete` | 删除应用（异步生效，不可逆）|
| devapp | `inactive` | 停用应用 |
| devapp | `permission remove` | 取消已授权限 |
| devapp | `version publish` | 发布版本（含高敏权限需 `--confirm-sensitive`）|
| devapp | `member remove` | 移除应用成员 |
| devapp | `robot offline` | 下线机器人 |
| api | 任意 `POST/PUT/PATCH/DELETE` 写接口 | 按接口 side_effect 评估，先 `--dry-run` |

确认流程：`Step 1 展示摘要（操作+目标+影响）→ Step 2 用户明确确认 → Step 3 加 --yes 执行`。

## 6. 完整示例（创建机器人应用）

```text
devapp create --dry-run → 确认 → create --yes
  → devapp get 确认 unifiedAppId
  → credentials get（敏感，不回显）
  → permission add（机器人发消息，dry-run → 确认 → --yes）
  → robot create（dry-run → 确认 → --yes）
  → version create → check-approval → publish（确认）→ status 验证
```

> 应用管理每步 flag 细节以 `dws devapp <子命令> --help` 为准。
