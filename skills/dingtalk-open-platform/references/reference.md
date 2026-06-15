# 一等公民：错误码 / 状态码 / 限流 / 版本化

把 Status codes / Request limits / Versioning / Webhooks 作为一等公民单列，调用前明确契约边界。

## 1. 错误诊断：devdoc error diagnose

```bash
dws devdoc error diagnose --error-code "40014" --query "获取用户信息 access_token" --format json
dws devdoc error diagnose --request-id "<REQUEST_ID>" --query "接口调用失败" --format json
dws devdoc error diagnose --error-message "token is illegal" --api "获取用户信息" --format json
```

| flag | 说明 |
| --- | --- |
| `--query` | 原始排查问题（必填语义）|
| `--error-code` | 错误码 |
| `--error-message` | 错误描述，合并进问题 |
| `--api` | API 名称，合并进问题 |
| `--request-id` | 开放平台 requestId（不是所有接口都返回）|
| `--context` | 额外排查上下文 |

- 后端返回 `PARAM_ERROR - 未找到指定工具` / `unknown tool` → 记 `needs_gateway_tool_registration`，降级 `devdoc article search` 或本 skill 的 [llm.md](../llm.md) 目录。

## 2. 错误码三元组（自恢复）

每个错误码应能回答：**触发原因 / 排查路径 / next_action**；未知错误码写 `源文档未提供`，不要编造。

| 错误码 | 触发原因 | next_action |
| --- | --- | --- |
| `40014` | access_token 无效/过期 | `dws auth login`，自有应用重取凭证 |
| `token is illegal` | 未登录 / token 失效 | `dws auth status` → `dws auth login` |
| `PARAM_ERROR - 未找到指定工具` | 网关工具未注册 | 记 `needs_gateway_tool_registration`，降级本地索引 |
| `ServiceResult.success=false` | 业务校验失败 | 透传 `errorCode/errorMsg`，按总表排查 |

> 未知错误码写 `源文档未提供`，不要编造原因或 next_action（gap #3：错误码总表数据仍在增量建设）。

## 3. 状态码 / HTTP 语义

| HTTP | 含义 | 处理 |
| --- | --- | --- |
| 2xx | 成功 | 校验业务成功字段后继续 |
| 401 / 403 | 鉴权 / 权限不足 | `auth login` 或 `devapp permission add` |
| 429 | 限流 | 退避重试（见限流）|
| 5xx | 平台错误 | 附 requestId 回流，不反复变通 |

## 4. 限流（Request limits）

- 每接口 `maxQps` 写在 `doc_open_api.extension`（如 `{"maxQps":100}`）；批量调用前先看 QPS。
- 429 / 限流：指数退避重试，单命令重试 ≤ 3 次；超限停止并报告。
- 幂等键无平台级统一约定（gap #10）→ 写接口自动重试需谨慎，优先 readback 确认而非盲目重发。

## 5. 版本化（Versioning）

- 路径版本：`/v1.0/` 与 `/v2.0/` 并存；旧 `dingtalk.oapi.*` 与 `/topapi/*` 为老接口面。
- `dws upgrade`：能力缺失 / 新接口未注册时，先升级 dws 再试（新能力先在新版本上线）。
- 应用版本发布走 `dws devapp version ...`（命令树以 `dws devapp version --help` 为准）。

## 6. 事件 / 回调（Webhooks）

- 事件订阅 `dws devapp event ...` 为**待实现能力**（后端待发布）→ 标 `TODO(待验证)`，不要当可用。
- 企业回调相关接口见本 skill [llm.md](../llm.md) 的相应能力域。
