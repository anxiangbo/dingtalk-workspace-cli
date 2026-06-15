# 发现层（Discovery）— 找到接口与文档

先从文档目录定位能力域与接口，再下钻到单接口官方文档页。

## 发现优先级

1. **文档目录（首选，离线可用）**：[llm.md](../llm.md) ＝ [api-catalog.md](api-catalog.md) 同源 — 全量**有效**官方文档链接（经内容长度校验，按 8 能力域 + 功能子类归类），URL 模板 `https://open.dingtalk.com/document/development/<slug>`。
2. **运行态搜索**：`dws devdoc article search`（命中即返回标题/摘要/文档链接）。
3. **官方文档站**：`https://open.dingtalk.com/document`。

## devdoc article search

```bash
dws devdoc article search "MCP" --format json
dws devdoc article search --query "OAuth2 接入" --format json
dws devdoc article search --query "消息卡片" --page 2 --size 5 --format json
```

| flag | 说明 |
| --- | --- |
| 位置参数 / `--query` / `--keyword` | 搜索关键词（必填，建议传用户原话里的 API 名/能力名/错误码）|
| `--page` | 页码，从 1 开始，默认 1 |
| `--size` | 每页条数，默认 10 |

- 只做**搜索**，不做读取；命中条目返回标题、摘要、`doc_link`，由 Agent 引用链接或进一步浏览。
- devdoc 查的是**开放平台开发者文档**（面向研发），不要拿它查业务数据（那是 doc / aitable / report）。

## 运行态风险与降级

- 真实调用若返回 `PARAM_ERROR - 未找到指定工具` / Forbidden（pre-mcp-gw）：这是**网关/工具注册未闭环**，不是用户参数问题。
- 处理：标记 `needs_gateway_tool_registration`，降级到 [llm.md](../llm.md) 本地目录或已挂载的逐页 Markdown。
- `--help` / `--dry-run` 只证明 CLI 命令面与映射存在，**不**证明后端工具已注册。

## 上线/注入验收（六步门禁）

```bash
dws devdoc --help
dws devdoc article search --query "OAuth2 接入" --dry-run --format json
dws devdoc error diagnose --error-code "40014" --query "access_token" --dry-run --format json
dws schema devdoc.search_open_platform_docs --format json || true
dws schema devdoc.search_open_error_code_rag --format json || true
# real call：未注册则记录 needs_gateway_tool_registration
```
