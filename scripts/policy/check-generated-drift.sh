#!/bin/sh
set -eu

# Check committed version metadata against its deterministic Skill sources.

ROOT="$(CDPATH= cd -- "$(dirname -- "$0")/../.." && pwd)"
cd "$ROOT"

tmp="$(mktemp)"
trap 'rm -f "$tmp"' EXIT HUP INT TERM

go run ./internal/generator/cmd_schema_agent_metadata \
  -root . \
  -validate-surface=false \
  -output "$tmp"

if ! cmp -s internal/cli/schema_agent_metadata.json "$tmp"; then
  printf '%s\n' 'generated drift: internal/cli/schema_agent_metadata.json is stale' >&2
  printf '%s\n' 'run: make generate-schema-agent-metadata' >&2
  diff -u internal/cli/schema_agent_metadata.json "$tmp" || true
  exit 1
fi

printf 'generated drift check: ok\n'
