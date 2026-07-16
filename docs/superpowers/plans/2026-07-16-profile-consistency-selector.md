# Profile Consistency and Friendly Selector Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make identity Token state authoritative, add deterministic organization/account aliases, and preserve historical profile data and commands.

**Architecture:** `profiles.json` owns identity selection through exact `corpId:userId` pointers and a new `orgCurrentProfiles` map. Identity-scoped Keychain entries remain the only authentication facts; organization and global Token slots remain compatibility mirrors. All user-friendly selectors resolve to an exact identity before storage or execution.

**Tech Stack:** Go, Cobra, platform Keychain/file-backed test Keychain, Bash E2E, embedded mono/multi Skills.

## Global Constraints

- Preserve `profiles.json` version 1 and accept every historical field.
- Preserve historical `--profile <corpId>` and `--profile <corpId:userId>` commands.
- Never select an account by array order or login/use timestamps.
- Never use profile Token summaries as authentication facts.
- Keep `profile switch -` through `previousProfile`.
- Keep `primaryProfile` and `isPrimary` in JSON for compatibility, but do not use them for selection.
- Use `gofmt` for every modified Go file.
- Write each behavior test first and verify the expected failure before production changes.

---

### Task 1: Exact selection state and deterministic migration

**Files:**
- Modify: `internal/auth/profiles.go`
- Test: `internal/auth/token_test.go`

**Interfaces:**
- Produces: `ProfilesConfig.OrgCurrentProfiles map[string]string`
- Produces: `ResolveProfileSelector(configDir, selector string) (*Profile, error)`
- Produces: exact `currentProfile`, `previousProfile`, and organization-current selectors for all new writes.

- [ ] **Step 1: Write failing tests**

Add tests proving:

```go
func TestMigrationBuildsOrgCurrentFromMatchingCorpMirror(t *testing.T)
func TestMigrationDoesNotGuessOrgCurrentForMultipleAccounts(t *testing.T)
func TestLoginStoresExactCurrentAndPreviousSelectors(t *testing.T)
func TestOrganizationSelectionWithoutExplicitDefaultFails(t *testing.T)
```

The test fixtures must include one same-organization pair with no timestamps and one
legacy `currentProfile=corpId`.

- [ ] **Step 2: Verify RED**

Run:

```bash
go test ./internal/auth -run 'Test(MigrationBuildsOrgCurrent|MigrationDoesNotGuessOrgCurrent|LoginStoresExactCurrent|OrganizationSelectionWithoutExplicitDefault)' -count=1
```

Expected: failures because `orgCurrentProfiles` does not exist and organization selection still falls back to timestamps.

- [ ] **Step 3: Implement exact selection state**

Add:

```go
OrgCurrentProfiles map[string]string `json:"orgCurrentProfiles,omitempty"`
```

Normalize only valid exact selectors. Migrate organization defaults from a matching
organization Token mirror, an exact historical pointer, or a sole account. Stop using
`profileRecency`, `mostRecentProfile`, `primaryProfile`, or `previousProfile` as
selection fallbacks.

- [ ] **Step 4: Verify GREEN**

Run the focused command from Step 2 and:

```bash
go test ./internal/auth -count=1
```

Expected: PASS.

### Task 2: Friendly organization and account aliases

**Files:**
- Modify: `internal/auth/profiles.go`
- Modify: `internal/app/profile_command.go`
- Test: `internal/auth/token_test.go`
- Test: `internal/app/profile_command_test.go`

**Interfaces:**
- Consumes: `ResolveProfileSelector`
- Produces: support for `corpId:userId`, `corpId:userName`, `corpName:userId`, and `corpName:userName`.
- Produces: typed ambiguity errors containing stable `corpId:userId` candidates.

- [ ] **Step 1: Write failing tests**

Add table-driven resolution tests for:

```text
corp-a:user-1
corp-a:张三
组织A:user-1
组织A:张三
```

Add separate tests proving duplicate organization names and duplicate usernames fail
and include every stable candidate.

- [ ] **Step 2: Verify RED**

Run:

```bash
go test ./internal/auth ./internal/app -run 'Test.*FriendlyIdentitySelector|Test.*Ambiguous.*Selector' -count=1
```

Expected: name-based compound selectors fail as not found.

- [ ] **Step 3: Implement alias resolution**

Split the input on the last colon, resolve the organization segment first, then resolve
the account segment only inside that organization. Match stable IDs before names.
Return the exact stored profile and never persist the input alias.

- [ ] **Step 4: Verify GREEN**

Run the focused tests and:

```bash
go test ./internal/auth ./internal/app -run 'Profile|IdentitySelector' -count=1
```

Expected: PASS.

### Task 3: Current, previous, and organization-current transitions

**Files:**
- Modify: `internal/auth/profiles.go`
- Modify: `internal/auth/token.go`
- Test: `internal/auth/token_test.go`

**Interfaces:**
- Consumes: exact selector resolution and `OrgCurrentProfiles`
- Produces: deterministic login, switch, previous-switch, and logout transitions.

- [ ] **Step 1: Write failing tests**

Cover:

```go
func TestExactSwitchUpdatesCurrentPreviousAndOrgCurrent(t *testing.T)
func TestUsePreviousUpdatesDestinationOrgCurrent(t *testing.T)
func TestOneShotExactProfileDoesNotChangeSelectionState(t *testing.T)
func TestDeleteOrgCurrentPromotesOnlyRemainingAccount(t *testing.T)
func TestDeleteOrgCurrentLeavesMultipleAccountsUnselected(t *testing.T)
```

- [ ] **Step 2: Verify RED**

Run:

```bash
go test ./internal/auth -run 'Test(ExactSwitchUpdates|UsePreviousUpdates|OneShotExact|DeleteOrgCurrent)' -count=1
```

Expected: state assertions fail under organization-mirror/timestamp fallback behavior.

- [ ] **Step 3: Implement transitions**

Every persistent login or switch writes exact `currentProfile` and updates
`orgCurrentProfiles`. `profile switch -` swaps exact pointers and updates the
destination organization. Exact one-shot execution only reads. Logout promotes a sole
remaining account and otherwise clears the organization default.

- [ ] **Step 4: Verify GREEN**

Run:

```bash
go test ./internal/auth -count=1
```

Expected: PASS.

### Task 4: Profile list from real Token state

**Files:**
- Modify: `internal/app/profile_command.go`
- Test: `internal/app/profile_command_test.go`

**Interfaces:**
- Consumes: identity Token storage and exact selection state.
- Produces: `profile list` views whose status and expiry values come from the real Token.

- [ ] **Step 1: Write failing tests**

Add tests where `profiles.json` says `22:29`, the identity Token says `17:38`, and the
JSON view must return `17:38`. Add tests for exactly one `isCurrent`, at most one
`isOrgCurrent` per organization, deprecated `isPrimary`, and a table without `PRI`.

- [ ] **Step 2: Verify RED**

Run:

```bash
go test ./internal/app -run 'TestProfileList.*(RealToken|CurrentFlags|DeprecatedPrimary|Table)' -count=1
```

Expected: the view still returns cached profile timestamps and the table still contains `PRI`.

- [ ] **Step 3: Implement live views**

Load each exact identity Token without refreshing it. Populate `status`, `expiresAt`,
and `refreshExpAt` from that Token. Compute `isCurrent` and `isOrgCurrent` from exact
selection state. Keep `isPrimary` and top-level `primaryProfile` only for wire
compatibility.

- [ ] **Step 4: Verify GREEN**

Run:

```bash
go test ./internal/app -run 'ProfileList|ProfileSwitch' -count=1
```

Expected: PASS.

### Task 5: Refresh diagnostics

**Files:**
- Modify: `internal/app/auth_command.go`
- Modify: `internal/app/access_token_resolve.go`
- Modify: `internal/auth/oauth_provider.go`
- Test: `internal/app/auth_command_test.go`
- Test: `internal/app/access_token_resolve_test.go`

**Interfaces:**
- Produces: `auth status` with `authenticated=false` after refresh failure.
- Produces: business commands that preserve the refresh failure instead of returning an empty Token.

- [ ] **Step 1: Write failing tests**

Inject a refresh endpoint failure while the local Refresh Token timestamp remains in
the future. Assert that status is unauthenticated with a diagnostic and token
resolution returns the refresh failure.

- [ ] **Step 2: Verify RED**

Run:

```bash
go test ./internal/app -run 'TestAuthStatus.*RefreshFailure|TestResolveAccessToken.*RefreshFailure' -count=1
```

Expected: status reports authenticated and resolver returns an empty Token without the original error.

- [ ] **Step 3: Preserve refresh failures**

Use the refresh error as the status diagnostic, calculate authenticated state only
after a successful refresh or a valid Access Token, and return the OAuth provider
error when no legacy Token succeeds.

- [ ] **Step 4: Verify GREEN**

Run:

```bash
go test ./internal/app -count=1
```

Expected: PASS.

### Task 6: Skills, help, and compatibility documentation

**Files:**
- Modify: `skills/mono/SKILL.md`
- Modify: `skills/multi/dws-shared/SKILL.md`
- Modify: `skills/multi/dingtalk-profile/SKILL.md`
- Modify: `README.md`
- Modify: `internal/app/profile_command.go`
- Modify: `docs/plans/2026-07-16-multi-account-profiles-design.md`
- Test: `internal/app/skill_setup_embed_test.go`
- Test: `internal/app/profile_command_test.go`

**Interfaces:**
- Produces: identical selector and ambiguity guidance in mono and multi Skill modes.

- [ ] **Step 1: Write failing content tests**

Assert the embedded mono and multi Skills contain:

```text
corpId:userName
组织名:userId
组织名:userName
重名时报错
```

Assert profile command Help documents the same selector forms.

- [ ] **Step 2: Verify RED**

Run:

```bash
go test ./internal/app -run 'Test.*Skill.*ProfileSelector|TestProfile.*Help.*Selector' -count=1
```

Expected: mono content and Help assertions fail.

- [ ] **Step 3: Update documentation**

Document canonical `corpId:userId`, friendly aliases, ambiguity errors, exact current
state, deprecated primary state, and the rule that Skills should prefer the stable
`profile` field returned by `profile list`.

- [ ] **Step 4: Verify GREEN**

Run:

```bash
go test ./internal/app -run 'Skill|Profile' -count=1
```

Expected: PASS.

### Task 7: Isolated end-to-end and full verification

**Files:**
- Modify: `scripts/dev/test-multi-profile-e2e.sh`

**Interfaces:**
- Consumes: all previous tasks.
- Produces: production CLI coverage for profile storage, selection, execution, logout, migration, and installed Skills.

- [ ] **Step 1: Extend E2E assertions**

Seed same-organization accounts and verify all four compound selector forms, duplicate
name rejection, exact flags, real Token expiry output, previous switching, logout
transitions, historical migration, and installed mono/multi Skill contents.

- [ ] **Step 2: Run focused E2E**

Run:

```bash
bash scripts/dev/test-multi-profile-e2e.sh --skip-go-tests
```

Expected: PASS.

- [ ] **Step 3: Run full repository gates**

Run:

```bash
gofmt -w <all modified Go files>
go build ./cmd
DWS_PACKAGE_VERSION=0.0.0-test go test ./...
go generate ./internal/cli
./scripts/policy/check-generated-drift.sh
./scripts/policy/check-schema-catalog.sh
```

Expected: every command exits 0 and generated files have no unexplained drift.

- [ ] **Step 4: Build and install beta.3**

Run:

```bash
DWS_PACKAGE_VERSION=v1.0.53-beta.3 go build -o /Users/qinze/.local/bin/dws ./cmd
/Users/qinze/.local/bin/dws version
```

Expected: version output contains `v1.0.53-beta.3`.

- [ ] **Step 5: Run local acceptance**

Run `profile list`, exact `auth status`, and non-mutating selector checks against the
installed binary. Do not send messages or modify remote DingTalk data.
