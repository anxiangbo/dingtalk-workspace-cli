# Profile 一致性与友好选择器设计

## 目标

在保留历史 `profiles.json`、历史 Token 槽位和历史命令的前提下，解决两类问题：

- 同一账号的 profile 摘要与真实 Token 状态不一致。
- 精确账号只能使用 `corpId:userId`，不能使用组织名或用户名。

## 数据事实

每类数据只保留一个事实源：

```text
账号身份与选择关系     profiles.json
账号真实凭证           auth-token:id:<identity-hash>
旧版本组织兼容凭证     auth-token:<corpId>
旧版本全局兼容凭证     auth-token
```

`profiles.json` 新增可选字段：

```json
{
  "version": 2,
  "currentProfile": "corp-a:user-2",
  "previousProfile": "corp-a:user-1",
  "orgCurrentProfiles": {
    "corp-a": "corp-a:user-2"
  }
}
```

`version: 1` 历史数据继续读取；完成一次选择关系迁移后写为 `version: 2`。
版本号用于区分“历史数据尚未推导组织默认账号”和“新版本明确保持未选择”，
避免删除默认账号后又被 `previousProfile` 补回。

- `currentProfile`：不传 `--profile` 时使用的账号。
- `previousProfile`：仅用于兼容 `profile switch -`。
- `orgCurrentProfiles[corpId]`：只传组织时使用的账号。
- `primaryProfile`：兼容读取和输出，不再参与选择。
- `expiresAt`、`refreshExpAt`、`status`、`lastLoginAt`、`lastUsedAt`、`updatedAt`：兼容读取，但不再用于选择、排序或认证判断。

新版本写入的 `currentProfile`、`previousProfile` 和
`orgCurrentProfiles` 值统一使用稳定的 `corpId:userId`。

## 默认账号

选择规则：

1. 不传 `--profile`：读取 `currentProfile`。
2. 传组织：读取 `orgCurrentProfiles[corpId]`。
3. 传组织和账号：精确读取该账号。

不再按时间、数组顺序、`primaryProfile` 或 `previousProfile` 猜测组织默认账号。

状态变化：

- 登录成功：账号成为该组织默认账号和全局当前账号；旧全局当前账号写入 `previousProfile`。
- `profile switch <组织>`：切换到该组织已经明确记录的默认账号。
- `profile switch <组织:账号>`：该账号成为组织默认账号和全局当前账号。
- 一次性 `--profile`：不修改任何选择关系。
- `profile switch -`：交换 `currentProfile` 和 `previousProfile`，并将切回的账号设为其组织默认账号。

删除组织默认账号时：

- 没有剩余账号：删除该组织默认关系。
- 只剩一个账号：将剩余账号设为组织默认账号。
- 仍有多个账号：不猜测，清空该组织默认关系；以后只传组织时明确报错。

## 历史数据迁移

迁移只接受明确证据：

1. `orgCurrentProfiles` 中已有合法精确账号。
2. `auth-token:<corpId>` 中的 `userId` 能匹配本地账号。
3. `currentProfile`、`previousProfile` 或 `primaryProfile` 已经是精确账号。
4. 该组织只有一个账号。

多账号组织仍无法确定时保持未设置，不使用时间推断。

历史 `currentProfile=corpId`：

- 能确定组织默认账号时，迁移成精确账号。
- 不能确定时，保留历史值；真正执行时返回明确错误，要求指定账号。

## 选择器

单段选择器继续支持：

```text
corpId
唯一组织名
唯一本地 profile 名
```

双段精确选择器支持：

```text
corpId:userId
corpId:userName
组织名:userId
组织名:userName
```

名称只作为输入别名。解析成功后统一得到 `corpId:userId`，不把名称写入
`currentProfile`、`previousProfile`、`orgCurrentProfiles` 或 Token 键。

匹配优先级：

```text
组织段：corpId > 唯一 corpName
账号段：userId > 该组织内唯一 userName
```

组织名或用户名匹配多个候选时直接报错，并列出候选
`corpId:userId`；不得选择第一项或最近使用账号。

## Profile 列表

保留：

- `isCurrent`：账号是否为全局当前账号。
- `isOrgCurrent`：账号是否为所在组织默认账号。

兼容保留但废弃：

- 顶层 `primaryProfile`
- 每项 `isPrimary`

表格不再显示 `PRI` 列，也不按时间排序。

`profile list` 对每个账号读取真实身份 Token，现场计算：

- `status`
- `expiresAt`
- `refreshExpAt`

列表不触发 Token 刷新。Access Token 当前不可用时显示 `expired`；不能继续展示
`profiles.json` 中的历史到期摘要。

## 刷新错误

- `auth status` 刷新失败时返回 `authenticated=false`，并保留刷新失败原因。
- 业务命令不能吞掉刷新错误后统一返回“未登录”。
- 本地 `refreshExpAt` 未到期只能说明“本地记录未到期”，不能证明服务端仍接受该 Refresh Token。

## Skill 与文档

同步更新：

- `skills/mono/SKILL.md`
- `skills/multi/dws-shared/SKILL.md`
- `skills/multi/dingtalk-profile/SKILL.md`
- `README.md`
- CLI Help
- 原同组织多账号设计文档

Skill 必须说明友好名称只用于输入、重名时报错，以及推荐从
`profile list` 的 `profile` 字段取得稳定 `corpId:userId`。

## 验证

- 单元测试覆盖四种精确选择器、重名错误、默认账号状态、历史迁移和真实 Token 列表。
- 隔离 E2E 覆盖双账号登录、列表、切换、状态、业务执行、退出、旧数据迁移和 Skill 安装。
- 运行完整 Go 测试、构建、Schema 生成漂移检查和策略检查。
- 构建 `v1.0.53-beta.3` 覆盖本地 dws 后，再运行本地命令验收。
