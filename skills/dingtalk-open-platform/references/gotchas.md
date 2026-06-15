# Common Gotchas — 钉钉开放平台真实坑

钉钉开放平台调用高频踩坑，来自线上失败归因与排错沉淀。格式：错 → 对 → 后果。

| 坑 | 错误做法 | 正确做法 | 后果 |
| --- | --- | --- | --- |
| 凭证混用 | 用 MCP 默认（加密）token 调 `dws api` raw 接口 | `dws api` 用自有应用 `--client-id/--client-secret` 登录 | raw 调用被拒 |
| scope 未开通 | 直接调用未授权接口 | 先 `dws devapp permission add` 申请并发版 | 403 / 无权限 |
| 新旧 path 混用 | Header token 配 `oapi` query 接口（或反之）| 新版走 `api.dingtalk.com` Header；旧版走 `oapi` query | 鉴权失败 |
| token 类型用错 | 用应用 token 调个人数据 / 免登接口 | 个人数据用用户身份，端内用 JSAPI | 无权限 |
| 权限名当 scope | `--permissions 机器人发消息` | 只传 `scopeValue`（如 `qyapi_robot_sendmsg`）| 申请失败 |
| 数据范围默认全员 | 假设拥有全员可见范围 | 显式确认 data_range，按授权范围调用 | 越权 / 空数据 |
| 多候选选第一个 | 名称命中多条直接取第一条 | 展示候选，停止写操作，等用户指定 | 操作错对象 |
| 写操作未 dry-run | 直接 `--yes` 执行写/删/发 | 先 `--dry-run` 展示摘要 → 确认 → `--yes` | 不可逆误操作 |
| 丢失 requestId | 报错后不留痕 | 保留 `requestId` 交 `devdoc error diagnose` | 无法链路定位 |
| 高频刷新 token | 每请求重新换 token | token 必须缓存复用 | 触发限流 |
| devdoc 网关未注册 | 把 `PARAM_ERROR - 未找到指定工具` 当用户参数错 | 记 `needs_gateway_tool_registration`，降级 [llm.md](../llm.md) | 误判阻塞原因 |
