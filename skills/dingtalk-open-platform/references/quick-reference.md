# Quick Reference — 应用类型 / Header / 限流 / CLI

## 应用类型

来自各接口「权限」表（企业内部应用 / 第三方企业应用 / 第三方个人应用）与端内能力实测。

| 应用类型 | 适用 | 鉴权 / 身份 | 典型场景 |
| --- | --- | --- | --- |
| 企业内部应用 | 单企业自建 | 企业内部应用 `accessToken`（appKey/appSecret 换取） | 服务端调组织数据、机器人、文档 |
| 第三方企业应用 | ISV 多企业 | 授权企业 `accessToken`（suiteTicket / 授权码换取） | 上架应用市场、跨企业服务 |
| 第三方个人应用 | 面向个人用户 | 个人授权（OAuth） | 个人维度能力（部分接口标「暂不支持」）|
| 移动接入应用 | 钉钉端内 H5 / 小程序 / 酷应用 | JSAPI 免登 ticket / 免登 code | 端内免登、容器 JSAPI 能力 |

> 选型见 [SKILL.md](../SKILL.md) §5 Decision guidance；应用增删改管理用 `dws devapp ...`（以 `dws devapp --help` 为准）。

## Header / 限流 / 分页 {#header限流分页}

### 必需 Header

| 域名 | token 传递 | 其它 |
| --- | --- | --- |
| `api.dingtalk.com`（新版）| Header `x-acs-dingtalk-access-token: <ACCESS_TOKEN>` | `Content-Type: application/json` |
| `oapi.dingtalk.com`（旧版）| query `?access_token=<ACCESS_TOKEN>` | 路径 `/topapi/...` |

### 限流

- 每接口有 `maxQps`（官方文档 / API Explorer 标注，部分接口 `doc_open_api.extension.maxQps` 可查）。
- 触发限流返回 429 / 限流码：遵守 `Retry-After`，指数退避，单命令重试 ≤ 3 次。
- 幂等键无平台级统一约定 → 写接口自动重试前优先 readback 确认，避免重复写。

### 分页 / 大小

- 分页两类：`offset`/`size` 偏移式；`nextToken`/`maxResults` 游标式（以接口文档为准）。
- 批量大小遵循各接口文档上限；dws 侧批量操作建议单次 ≤ 30 条。

## dws CLI 速查（均 `--help` 核验）

```bash
dws --help                                     # 顶层产品与 utility 命令
dws auth status --format json                  # 登录态 / 当前身份
dws auth login [--device]                      # 扫码 / 设备流登录
dws schema [<product>.<tool>]                  # MCP 工具 JSON Schema（必填字段 / flag 别名）
dws devdoc article search "<关键词>" --format json          # 查开发文档
dws devdoc error diagnose --error-code <code> --query "..." --format json  # 错误排查
dws api <METHOD> <PATH> --data '<JSON>' --format json       # 原始 OpenAPI 直通
dws devapp list --name <应用名> --format json               # 开放平台应用 CRUD / 凭证 / 权限
dws doc read --node <URL_OR_NODE_ID> --content-format markdown --format json  # 读钉钉文档
```

- `dws api` 仅认自有应用凭证（`--client-id/--client-secret` 登录），MCP 默认加密 token 不支持 raw 调用。
- 命令事实源以 `dws <path> --help` / `dws schema` 为准；参考文档与之冲突时以 `--help` 为准。
