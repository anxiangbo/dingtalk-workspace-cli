# 鉴权与凭证（Auth & Credentials）— 前置必读

调用前先把鉴权与凭证准备好；凭证视为敏感，绝不外泄。

## 1. 登录与 token 状态

```bash
dws auth status --format json      # 查看当前登录态与身份
dws auth login                     # 扫码 / Device Flow 登录，自动刷新 token
dws auth logout                    # 清除认证信息
dws auth reset                     # 重置认证（清本地 token，触发重新授权）
```

| 子命令 | 用途 |
| --- | --- |
| `status` | 调用前确认是否已登录、token 是否有效 |
| `login` | 未登录 / `token is illegal` / 401 时执行 |
| `export` / `import` | 迁移认证包（跨机器） |

- token 过期或 `token is illegal`：执行 `dws auth login`，**不要**把 token 贴进对话或日志。

## 1b. access_token 端点（HTTP 原始换取）

```http
# 新版（推荐）：token 用于业务请求 Header x-acs-dingtalk-access-token
POST https://api.dingtalk.com/v1.0/oauth2/accessToken
Content-Type: application/json

{"appKey":"<APP_KEY>","appSecret":"<APP_SECRET>"}
```
```http
# 旧版：token 用于业务请求 query ?access_token=
GET https://oapi.dingtalk.com/gettoken?appkey=<APP_KEY>&appsecret=<APP_SECRET>
```

- token 必须缓存复用，禁止每请求重换（高频换取会触发限流）。
- 文档：[获取企业内部应用 accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-an-internal-app) · [获取第三方应用授权企业 accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-the-authorized-enterprise-1)。
- 用 dws 时一般无需手动换取，`dws api` 用自有应用凭证登录后自动注入；凭证读取见 §3。

## 2. token 类型

| 类型 | 主体 | 典型用途 |
| --- | --- | --- |
| `app` / `tenant` | 企业自建应用 | 大多数服务端 OpenAPI |
| `user` | 用户授权 | 读当前用户数据、网页授权 |
| `isv` | 第三方应用 | ISV 套件场景 |

> 应用类型（企业自建 vs 第三方）查询 API 缺失是高频凭证混淆来源（gap #4），不确定时先 `dws devapp get` 看应用属性。

## 3. 应用凭证读取（敏感）

```bash
dws devapp credentials get --unified-app-id UNIFIED_APP_ID --format json
```

- 返回含 `clientId` / `clientSecret`(=`appSecret`)，**按敏感凭证处理**：仅用于构造 `dws api` 的 `--client-id/--client-secret`，不回显、不写文件。
- 不能用 `dws devapp get` 代替（`get` 不返回完整 secret）。
- `unifiedAppId` 来自 `dws devapp list --name X --format json`，多条命中先让用户指定。

## 4. raw API 调用的凭证约束

- `dws api` 仅限**自有应用凭证**（`--client-id/--client-secret` 登录后）使用。
- 通过 MCP 默认凭证登录获取的加密 token **不支持** raw API 调用 → 改用对应产品糖命令或先用自有应用登录。

## 常见错误

| 现象 | 应对 |
| --- | --- |
| `token is illegal` / 401 | `dws auth status` → `dws auth login` |
| raw api 报 token 不支持 | 用 `--client-id/--client-secret` 自有应用凭证 |
| 凭证泄露风险 | 只走 `credentials get`，输出不回显 secret |
