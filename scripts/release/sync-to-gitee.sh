#!/usr/bin/env bash
# Copyright 2026 Alibaba Group
# Licensed under the Apache License, Version 2.0
#
# Mirror release artifacts to a Gitee release so China users can install without
# hitting GitHub. The repo code itself is kept in sync by Gitee's repo-mirror
# feature; this script handles what the mirror does NOT carry — the GitHub
# Release *attachments* (binaries, checksums, skills zip) — by uploading them to
# the matching Gitee release via the Gitee OpenAPI v5.
#
# Consumed by install.sh when DWS_GITEE_REPO is set (it resolves each asset's
# real download_url from the Gitee API, since Gitee attachment URLs carry an
# unstable numeric id).
#
# Required environment (CI secrets):
#   GITEE_TOKEN   Gitee private access token (scopes: projects)
#   GITEE_REPO    "owner/repo" on Gitee, e.g. DingTalk-Real-AI/dingtalk-workspace-cli
# Optional:
#   VERSION       release tag (default: git describe)
#   DIST_DIR      artifacts dir (default: ./dist)
#   GITEE_API     API base (default: https://gitee.com/api/v5)
#
# Gating: if GITEE_TOKEN / GITEE_REPO are unset, exit 0 with a notice so the
# step can live in release.yml without breaking forks that lack the secret.

set -eu

DIST_DIR="${DIST_DIR:-dist}"
GITEE_API="${GITEE_API:-https://gitee.com/api/v5}"

missing=""
[ -z "${GITEE_TOKEN:-}" ] && missing="$missing GITEE_TOKEN"
[ -z "${GITEE_REPO:-}" ]  && missing="$missing GITEE_REPO"
if [ -n "$missing" ]; then
  echo "ℹ️  Gitee mirror sync skipped — missing:${missing}"
  echo "   Set these as repo secrets to auto-mirror releases to Gitee for China users."
  exit 0
fi

VERSION="${VERSION:-$(git describe --tags --always 2>/dev/null || echo dev)}"
OWNER="${GITEE_REPO%%/*}"
NAME="${GITEE_REPO##*/}"
base="${GITEE_API}/repos/${OWNER}/${NAME}"

echo "📦 Mirroring release ${VERSION} → Gitee ${GITEE_REPO}"

# ── Resolve or create the Gitee release for this tag ──────────────────────────
# Gitee mirror sync brings the git tag over, so the tag should already exist.
rel_json="$(curl -fsSL "${base}/releases/tags/${VERSION}?access_token=${GITEE_TOKEN}" 2>/dev/null || true)"
release_id="$(printf '%s' "$rel_json" | grep -o '"id":[ ]*[0-9]*' | head -1 | grep -o '[0-9]*' || true)"

if [ -z "$release_id" ]; then
  echo "   No Gitee release for ${VERSION} yet — creating it."
  rel_json="$(curl -fsSL -X POST "${base}/releases" \
    -F "access_token=${GITEE_TOKEN}" \
    -F "tag_name=${VERSION}" \
    -F "name=${VERSION}" \
    -F "body=Mirror of GitHub release ${VERSION} for China users." \
    -F "target_commitish=main" 2>/dev/null || true)"
  release_id="$(printf '%s' "$rel_json" | grep -o '"id":[ ]*[0-9]*' | head -1 | grep -o '[0-9]*' || true)"
fi
[ -n "$release_id" ] || { echo "❌ Could not get/create Gitee release for ${VERSION}. Response: ${rel_json}" >&2; exit 1; }
echo "   Gitee release id = ${release_id}"

# ── Mirror each artifact, verifying content so re-runs self-heal ─────────────
# Pull the current asset list (name + attach id + download url) so we can, per
# file:
#   • skip it when it is already on Gitee AND byte-identical,
#   • REPLACE it when present but stale (different bytes) — e.g. the darwin
#     binaries are re-signed and so differ between GitHub and a prior mirror
#     run, which breaks install.sh's checksums.txt verification on macOS,
#   • upload it when missing.
# This brings the Gitee release into byte-for-byte agreement with $DIST_DIR,
# which the caller fills from the GitHub release whose checksums.txt is what
# install.sh verifies against.
assets_map="$(curl -fsSL "${base}/releases/${release_id}?access_token=${GITEE_TOKEN}" 2>/dev/null \
  | python3 -c 'import json,sys
try:
    for a in json.load(sys.stdin).get("assets",[]):
        n=a.get("name",""); i=a.get("id",""); u=a.get("browser_download_url","")
        if n:
            print("%s\t%s\t%s" % (n, i, u))
except Exception:
    pass' 2>/dev/null || true)"

sha256_of() {  # sha256 of a file ($1) or, with no arg, of stdin
  if command -v sha256sum >/dev/null 2>&1; then sha256sum ${1:+"$1"} | awk '{print $1}';
  else shasum -a 256 ${1:+"$1"} | awk '{print $1}'; fi
}

gitee_attach() {  # upload file $1; success when the response carries a download url
  printf '%s' "$(curl -fsSL -X POST "${base}/releases/${release_id}/attach_files" \
    -F "access_token=${GITEE_TOKEN}" -F "file=@${1}" 2>/dev/null || true)" \
    | grep -q '"browser_download_url"'
}

uploaded=0
replaced=0
skipped=0
for f in "$DIST_DIR"/dws-*.tar.gz "$DIST_DIR"/dws-*.zip "$DIST_DIR"/checksums.txt; do
  [ -f "$f" ] || continue
  fn="$(basename "$f")"
  local_sha="$(sha256_of "$f")"
  row="$(printf '%s\n' "$assets_map" | awk -F'\t' -v n="$fn" '$1==n {print; exit}')"
  if [ -n "$row" ]; then
    aid="$(printf '%s' "$row" | cut -f2)"
    aurl="$(printf '%s' "$row" | cut -f3)"
    gitee_sha="$(curl -fsSL "$aurl" 2>/dev/null | sha256_of || true)"
    if [ "$gitee_sha" = "$local_sha" ]; then
      echo "   ✓ ${fn} already correct on Gitee — skip"
      skipped=$((skipped + 1))
      continue
    fi
    echo "   ↻ ${fn} differs on Gitee (stale) — deleting + re-uploading"
    curl -fsSL -X DELETE "${base}/releases/${release_id}/attach_files/${aid}?access_token=${GITEE_TOKEN}" >/dev/null 2>&1 || true
    if gitee_attach "$f"; then replaced=$((replaced + 1)); else echo "   ⚠ re-upload may have failed for ${fn}" >&2; fi
    continue
  fi
  echo "   ⬆ ${fn} (new)"
  if gitee_attach "$f"; then uploaded=$((uploaded + 1)); else echo "   ⚠ upload may have failed for ${fn}" >&2; fi
done

if [ "$uploaded" -eq 0 ] && [ "$replaced" -eq 0 ] && [ "$skipped" -eq 0 ]; then
  echo "❌ No artifacts found to mirror. Did the build (goreleaser) run / were assets downloaded into ${DIST_DIR}?" >&2
  exit 1
fi
echo "✅ Gitee release ${VERSION}: uploaded ${uploaded}, replaced ${replaced}, skipped ${skipped} (already correct)."
echo "   China install:  DWS_GITEE_REPO=${GITEE_REPO} \\"
echo "     curl -fsSL https://gitee.com/${GITEE_REPO}/raw/main/scripts/install.sh | sh"
