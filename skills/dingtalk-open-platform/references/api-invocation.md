# API 调用层（Invocation）— HTTP 契约 1:1 映射

文档里的 HTTP 主契约（Method + Path + Body）可直接落成 `dws` 命令或原始 HTTPS 请求，无需猜 SDK。`dws` 是可选 CLI，与直接发 HTTP 等价。

## 1. 原始直通：dws api

```bash
# api.dingtalk.com：token 走 Header (x-acs-dingtalk-access-token)，路径 /v1.0|/v2.0
dws api GET /v1.0/contact/users/me --format json
dws api POST /v1.0/contact/users/search \
  --data '{"queryWord":"张三","offset":0,"size":10}' --format json

# oapi.dingtalk.com：token 走 URL 参数 (access_token)，路径 /topapi/...
dws api POST /topapi/v2/user/get --data '{"userid":"<USER_ID>"}' --format json
```

| 域名 | token 传递 | 路径形态 |
| --- | --- | --- |
| `api.dingtalk.com` | Header `x-acs-dingtalk-access-token` | `/v1.0/xxx`、`/v2.0/xxx` |
| `oapi.dingtalk.com` | URL 参数 `access_token` | `/topapi/...` 或完整 URL |

- raw api 仅限自有应用凭证（见 [auth.md](auth.md)）。

## 2. 从文档契约到命令（1:1）

| 文档字段 | 命令位置 |
| --- | --- |
| Method | `dws api {METHOD}` |
| URL Path | `dws api {METHOD} {PATH}` |
| Body (JSON) | `--data '{...}'`（保持嵌套，不要展平）|
| Header token | 由 `dws` 注入，不要手填 |
| 输出解析 | `--format json` + 可选 `--jq` |

## 3. 产品糖命令优先

高频能力可用 `dws <product> ...` 糖命令（自带翻页/轮询/校验），而非手搓 `dws api`，命令树与 flag 以 `dws <product> --help` 为准：

- 应用管理 → `dws devapp ...`（开放平台应用 CRUD / 凭证 / 权限 / 机器人 / 版本）
- 通讯录 / 日历 / 审批 / 消息等 → `dws contact|calendar|oa|chat ...`

`dws api` 用于：糖命令未覆盖、或文档明确给 HTTP 契约要 1:1 复现时。

## 4. 命令发现（事实源）

```bash
dws <command-path> --help                 # 人读：Usage / Example / Flags
dws schema                                # 机读：所有产品与工具
dws schema <product>.<canonical_name>     # 单工具 JSON Schema
dws schema <path> --jq '.tool.required'   # 只看必填字段
```

> 参考文档里的 flag 列表是便于理解，不是契约。冲突时**以 `--help` / `dws schema` 为准**。

## 5. 写操作铁律

- 写/删/发/审批类先 `--dry-run --format json` 展示摘要 → 用户确认 → 加 `--yes`（详见 [workflow.md](workflow.md)）。
- 单次批量 ≤ 30 条。
