# Verification Checklist — 调用 / 上线前勾选

每次开放平台调用或交付前逐项核对，任一不过则停止并补齐。

## 发现 / 契约
- [ ] 已在 [llm.md](../llm.md) 或官方文档定位到目标接口（method / path / 文档链接）。
- [ ] 已解析稳定资源 ID（unifiedAppId / userId / unionId / nodeId 等），来自返回值，非编造。

## 鉴权 / 权限
- [ ] 已确认 token 类型（应用 / 用户 / JSAPI）与 base URL（新版 Header / 旧版 query）匹配。
- [ ] 已确认所需 `scopeValue` 已授权（`dws devapp permission list`）；缺则已申请并发版。
- [ ] 已确认数据范围（data_range），未默认全员。
- [ ] 凭证未出现在对话 / 日志 / prompt。

## 调用 / Header / 限流
- [ ] 必需 Header 正确（`x-acs-dingtalk-access-token` + `Content-Type`）。
- [ ] 分页 / 批量大小符合接口约定；限流退避已考虑。
- [ ] token 已缓存复用，未每请求重换。

## 写操作 / 安全
- [ ] 写 / 删 / 发 / 发布 / 撤回 已先 `--dry-run` 展示摘要并经用户当前轮确认（[危险]）。
- [ ] 不可逆操作目标精确（无批量误伤）。

## 验证 / 错误处理
- [ ] readback 三态明确：成功 / needs_permission / needs_resource。
- [ ] 错误处理覆盖鉴权 / 权限 / 限流 / 参数 / 平台错误；保留 `requestId`。
- [ ] 能回读则已回读校验；不一致字段已报告。
