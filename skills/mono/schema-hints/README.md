# DWS Agent Schema Hints

This directory contains versioned, structured Agent metadata. Files are build inputs for `internal/cli/schema_agent_metadata/`; generated runtime metadata must not be edited directly.

## Source kinds

- `explicit`: reviewed DWS hints. Scalar fields override imported baselines.
- `imported`: sanitized metadata from a fixed external revision. It fills missing Agent semantics but cannot redefine command paths or parameter contracts.

The Agent metadata generator also reads the committed `internal/cli/schema_mcp_metadata.json` after Skill and Hint parsing. A sanitized MCP description can fill an otherwise empty `agent_summary`; it is marked `reviewed: false`, retains revision provenance, and cannot infer or override risk/effect fields.

Tool keys should use stable `canonical_path` values from `internal/cli/schema_command_surface.json`. CLI paths and aliases are also accepted and are reconciled to the canonical public tool during generation.

`selection-review.json` fixes the reviewed selection contract for every public
tool: `use_when`, `avoid_when`, safe examples, `interface_mode`, availability,
and the reason for local, composite, or unavailable implementations. These
values are build inputs; the generator must not derive them from the previous
Catalog.

`reference-review.json` classifies every Skill command reference that is not a
current public leaf. `alias` entries bind an old or cross-product path to an
explicit current target. `group`, `stale`, and `out_of_surface` entries remain
visible in the audit but are never fuzzy-matched to a leaf.

`interface_ref` is a separate interface binding. Use it when a public helper/canonical tool calls a differently named MCP RPC or another source product:

```json
{
  "version": 1,
  "source": {"kind": "explicit", "name": "reviewed-interface-map"},
  "tools": {
    "chat.bot_search": {
      "interface_ref": {
        "product_id": "bot",
        "rpc_name": "search_my_robots"
      }
    }
  }
}
```

An entry containing only `interface_ref` participates in interface projection but does not count as Agent semantic coverage. It cannot add a command, change a flag, or expose a Wukong-only tool.

`interface_mode` has four reviewed values:

- `mcp`: exactly one fixed `interface_ref` implements the command.
- `composite`: multiple RPC/local steps implement the command; a singular ref would be misleading.
- `local`: the command only changes local process or policy state.
- `unavailable`: the compatibility command is retained but no reviewed backend is shipped.

The missing `notify` MCP service is separately dispositioned in
`internal/cli/schema_mcp_service_review.json`; it is outside the public command
surface and must not trigger runtime discovery.

Interface metadata may enrich type and description, but MCP `required` never promotes an optional Cobra flag. CLI required/one-of/conditional rules must be represented by current Cobra markers, typed runtime constraints, or reviewed parameter hints.

```json
{
  "version": 1,
  "source": {
    "kind": "explicit",
    "name": "calendar-schema-review"
  },
  "products": {
    "calendar": {
      "agent_summary": "管理日程、参与人、会议室和闲忙信息"
    }
  },
  "tools": {
    "calendar.get_calendar_detail": {
      "agent_summary": "读取一个日程的完整详情",
      "use_when": ["已经取得 eventId，需要查看详情"],
      "effect": "read",
      "reviewed": true
    }
  }
}
```

Run `make generate-schema` after changing Hint or Skill sources. External Wukong metadata must be refreshed by the controlled offline import pipeline with an immutable revision, then committed together with its audit before regenerating the Catalog; runtime refresh is forbidden.
