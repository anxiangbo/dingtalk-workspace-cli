---
name: dingtalk-open-platform
description: >-
  指导 AI Agent 发现、鉴权、调用、验证并排错钉钉开放平台 OpenAPI / CLI / 应用能力的执行手册。
  Use when 用户要 调钉钉开放平台接口 / OpenAPI、建企业内部或第三方应用、拿 access_token、申请权限 scope、
  订阅事件、查开发文档或错误码、用 dws 调接口。给出 base URL、Header、应用类型选型、分步工作流与校验清单。
  Distinct from 单纯的应用增删改管理 与 钉钉云文档内容读写。
  Do NOT use for 钉钉客户端 UI 设计、后台人工运营、未开放的内部能力。
  口语同义词：开放平台、开发者后台、open api、调接口、接口怎么调。
cli_version: ">=1.0.15"
metadata:
  category: developer-platform
  api_base: "https://api.dingtalk.com"
  legacy_api_base: "https://oapi.dingtalk.com"
  doc_base: "https://open.dingtalk.com/document"
---

# 钉钉开放平台 Skill（AI Agent 执行手册）

每一步都回答「你下一步敲什么」。不得凭记忆猜 endpoint / 参数 / scope / ID；先查 [llm.md](llm.md) 文档大目录或官方文档，再执行。

## 1. Product summary

钉钉开放平台是 REST OpenAPI，覆盖通讯录 / 消息 / 机器人 / 文档 / 钉盘 / 表格 / AI 表格 / 日历 / 审批 / 应用管理 / 事件订阅 / JSAPI / 小程序 / 三方应用。

| 项 | 值 |
| --- | --- |
| 新版 Base URL | `https://api.dingtalk.com`（路径 `/v1.0`、`/v2.0`；token 走 Header）|
| 旧版 Base URL | `https://oapi.dingtalk.com`（路径 `/topapi/...`；token 走 query `access_token`）|
| 鉴权 Header（新版）| `x-acs-dingtalk-access-token: <ACCESS_TOKEN>` + `Content-Type: application/json` |
| 鉴权类型 | 应用身份（appKey/appSecret 换 token）/ 用户身份（OAuth 授权码）/ JSAPI 免登 |
| 文档站 | `https://open.dingtalk.com/document`；全量有效链接见 [llm.md](llm.md) |
| 访问方式 | `dws` CLI（本仓库）、MCP、原始 HTTPS |

## 2. 约束 (NEVER DO)

- 不要编造 endpoint / 参数 / scopeValue / ID / URL / 错误码；ID 从返回值提取，链接来自 [llm.md](llm.md) 或官方文档。
- 不要把 `appSecret` / `clientSecret` / `accessToken` / `refreshToken` / JSAPI ticket / Cookie 写进对话、日志或 prompt。
- 不要混用新版与旧版鉴权；不要把「权限名称」当 `scopeValue`；不要默认拥有全员数据范围。
- **[危险]** 写 / 删 / 撤回 / 发消息 / 发布 / 加权限 等操作：先 `--dry-run` 展示摘要，**获用户当前轮确认**后才加 `--yes`。
- **Harmness**：不输出诱导有害内容；不教唆越权 / 绕过审批 / 绕过安全边界 / 泄露凭证；越权请求一律拒绝并说明原因。

## 3. When to use

调用 OpenAPI；读写协作资源；建/管应用、权限、机器人、版本；拿 access_token；订阅事件；查开发文档 / 字段 / 错误码 / requestId；安装与使用 `dws`。

不要用：钉钉客户端 UI 设计、后台人工运营、未公开开放的内部能力 → 直接告知不支持。

## 4. Quick reference

- **应用类型选型** → [references/quick-reference.md](references/quick-reference.md#应用类型)（企业内部 / 第三方企业 / 第三方个人 / 移动接入）。
- **核心能力域 + 代表端点**（链接 `https://open.dingtalk.com/document/development/<slug>`；全量 8 域 1513 条见 [llm.md](llm.md) / [references/api-catalog.md](references/api-catalog.md)）：
  - 通讯录：[新增或修改限制查看通讯录设置](https://open.dingtalk.com/document/development/add-or-modify-visibility-settings-for-address-book-restrictions)
  - 消息/IM：[添加互通群成员](https://open.dingtalk.com/document/development/add-a-group-member-1) · 机器人：[批量撤回机器人消息](https://open.dingtalk.com/document/development/batch-message-recall-chat)
  - 日历：[预定会议室](https://open.dingtalk.com/document/development/add-a-meeting-room) · 待办：[创建钉钉待办任务](https://open.dingtalk.com/document/development/add-dingtalk-to-do-task)
  - 审批：[归档审批实例](https://open.dingtalk.com/document/development/api-archiveprocessinstance) · 文档：[添加知识库成员](https://open.dingtalk.com/document/development/add-permissions-for-team-space-members)
  - 考勤：[添加假期规则](https://open.dingtalk.com/document/development/add-holiday-rules) · 互动卡片：[更新卡片场域信息](https://open.dingtalk.com/document/development/add-field-interface)
  - 鉴权：[获取企业内部应用 accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-an-internal-app) · 调用规范：[服务端 API 怎么调](https://open.dingtalk.com/document/development/how-to-call-apis)
- **必需 Header / 限流 / 分页 / 大小** → [references/quick-reference.md](references/quick-reference.md#header限流分页)。
- **dws CLI 速查**（可选工具，等价于直接发 HTTP 的一种调用/发现方式；命令均 `--help` 核验）：
```bash
dws auth status --format json                  # 看登录态
dws auth login --device                        # 无头环境扫码登录
dws devdoc article search "<关键词>" --format json   # 查开发文档
dws api GET /v1.0/contact/users/me --format json     # 原始调用（自有应用凭证）
dws devapp list --name <应用名> --format json        # 开放平台应用 CRUD / 凭证 / 权限
```

## 5. Decision guidance

| 选择题 | 用哪个 | 依据 |
| --- | --- | --- |
| 应用类型 | 单企业自建→企业内部应用；多企业 ISV→第三方企业应用；面向个人→第三方个人应用；端内免登→移动接入 | [quick-reference.md](references/quick-reference.md#应用类型) |
| 新版 vs 旧版 | 优先新版 `api.dingtalk.com`（Header token）；仅旧接口用 `oapi.dingtalk.com`（query token） | [api-invocation.md](references/api-invocation.md) |
| token 类型 | 组织服务端能力→应用 accessToken；个人数据/免登→用户身份；端内→JSAPI | [auth.md](references/auth.md) |
| 轮询 vs 事件订阅 | 一次性结果→同步调用；持续变更→事件订阅（`configure-event-subcription`）| [reference.md](references/reference.md) |
| 访问方式 | 有产品糖命令→`dws <product>`；只有 HTTP 契约→`dws api`；CLI 缺失→HTTPS | [api-invocation.md](references/api-invocation.md) |

## 6. Workflow

发现 → 前置 → 调用 → 验证 → 恢复，每步真实命令见 [references/workflow.md](references/workflow.md)：

1. **发现**：[llm.md](llm.md) 定位能力域→接口；运行态 `dws devdoc article search "<关键词>" --format json`。
2. **建应用 / 拿凭证**：`dws devapp create ...`（[危险]，先 dry-run）→ `dws devapp credentials get --unified-app-id <ID> --format json`（凭证敏感，不回显）。
3. **申请权限 scope**：`dws devapp permission add --unified-app-id <ID> --permissions <scopeValue> --dry-run --format json` → 确认 → `--yes`。
4. **拿 access_token**：`POST https://api.dingtalk.com/v1.0/oauth2/accessToken` body `{"appKey","appSecret"}`；token 必须缓存，禁止每请求重换。
5. **首次调用**：`dws api {METHOD} {PATH} --data '{...}' --format json`（Header token 由 dws 注入）。
6. **readback 校验**：成功(2xx+业务字段) / `needs_permission` / `needs_resource`；写接口走 setup→target→verify→cleanup。
7. **错误排查**：`dws devdoc error diagnose --error-code <code> | --request-id <id> --query "..." --format json`，保留 `requestId`。
8. **清理**：测试资源用完回收（删除/取消属 [危险]，需确认）。

## 7. Common gotchas

- MCP 默认凭证 ≠ raw API 凭证：`dws api` 只认自有应用 `--client-id/--client-secret`，MCP 加密 token 不支持 raw 调用。
- scope 未开通就调用 → 403/无权限：先 `dws devapp permission add` 申请并发版。
- 新旧 path / token 混用（Header token 配 `oapi` query 接口）→ 鉴权失败。
- token 类型用错（用应用 token 调个人数据接口）→ 无权限。
- devdoc 运行态网关 Forbidden / `PARAM_ERROR - 未找到指定工具` → 记 `needs_gateway_tool_registration`，降级 [llm.md](llm.md)。
- 多候选默认选第一个 / 写操作未 dry-run / 丢失 requestId / 高频刷新 token → 见 [references/gotchas.md](references/gotchas.md)。

## 8. Verification checklist

上线前逐项核对（详见 [references/checklist.md](references/checklist.md)）：已查 llm.md 或官方文档；已解析稳定 ID；已确认 token 类型 + scope + 数据范围；Header 正确；限流 / 分页已处理；写操作已 dry-run + 确认；错误处理覆盖鉴权/权限/限流/参数；能回读则已回读；保留 operation/requestId。

## 9. Resources

- [llm.md](llm.md)：开放平台文档大目录（`development/<slug>` 全量有效链接，按 8 能力域归类）。
- [references/api-catalog.md](references/api-catalog.md)：同源全量有效链接目录（skill 内离线）。
- `https://open.dingtalk.com/document`：官方文档站。
- 细节参考：[auth.md](references/auth.md) · [quick-reference.md](references/quick-reference.md) · [api-invocation.md](references/api-invocation.md) · [workflow.md](references/workflow.md) · [reference.md](references/reference.md) · [gotchas.md](references/gotchas.md) · [checklist.md](references/checklist.md) · [discovery.md](references/discovery.md)
