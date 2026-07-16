# 同组织多账号 Profile 设计

## 目标

通过兼容扩展 `profiles.json`，支持同一组织在一台机器上同时保存多个账号：

- `profiles` 允许相同 `corpId`，以 `corpId + userId` 唯一。
- `--profile <corpId>` 使用 `orgCurrentProfiles[corpId]` 明确记录的账号。
- 支持 `corpId/corpName + userId/userName` 四种组合，最终都解析成 `corpId:userId`。
- `currentProfile`、`previousProfile` 和组织默认账号统一写精确身份。
- `primaryProfile` 只兼容读取和输出，不再参与选择。
- 历史数据、历史 Token、历史导入包和历史命令无需人工迁移。
- `profile list` 默认展示全部账号。

## 兼容边界

必须保证：

- 现有 `version: 1`、字段名和历史摘要继续可读；选择关系迁移后写为 `version: 2`。
- 历史 `primaryProfile/currentProfile/previousProfile = corpId` 能明确解析时自动迁移成精确身份；不能解析时保留并在执行时报错。
- `profile switch <corpId>`、全局 `--profile <corpId>`、`auth status/logout --profile <corpId>` 保持原语义。
- `auth-token:<corpId>` 和 `auth-token` 继续存在，供历史逻辑和宿主程序读取。
- 新版本能自动读取并补全历史单账号数据。

不承诺：

- 新版本产生同组织多账号后，再由不认识重复 `corpId` 的旧版本长期交替写入。旧版本仍能使用组织兼容 Token，但可能在保存 `profiles.json` 时删除其他账号元数据。

## 数据模型

`profiles.json` 使用兼容扩展结构：

```json
{
  "version": 2,
  "primaryProfile": "corp-a",
  "currentProfile": "corp-a:user-2",
  "previousProfile": "corp-a:user-1",
  "orgCurrentProfiles": {
    "corp-a": "corp-a:user-2"
  },
  "profiles": [
    {
      "name": "组织A-user-1",
      "corpId": "corp-a",
      "corpName": "组织A",
      "userId": "user-1",
      "userName": "账号一"
    },
    {
      "name": "组织A-user-2",
      "corpId": "corp-a",
      "corpName": "组织A",
      "userId": "user-2",
      "userName": "账号二"
    }
  ]
}
```

身份选择器定义：

```text
organization selector = corpId
identity selector     = corpId:userId
```

`userId` 已知时按身份选择器去重；历史 `userId` 为空的记录按 `corpId` 保留，直到安全补全。
`version: 2` 还用于保留“该组织明确没有默认账号”的状态，防止后续迁移重新猜测。

## Token 存储

三层 Token：

```text
auth-token:id:<identity-hash>  账号真实 Token，事实源
auth-token:<corpId>            组织当前账号兼容镜像
auth-token                     全局当前账号兼容镜像
```

`identity-hash = sha256(corpId + "\x00" + userId)`，避免直接拼接后的文件名转义碰撞。

写入规则：

- 登录、刷新首先写账号真实 Token。
- 登录或持久切换账号时更新组织兼容镜像。
- 当前默认身份变化时更新全局兼容镜像。
- 一次性 `--profile corpId:userId` 不修改任何兼容镜像。
- 刷新非组织当前账号时只更新账号真实 Token。

## 选择规则

选择器解析优先级：

1. `corpId:userId`：精确身份。
2. `corpId`：读取 `orgCurrentProfiles[corpId]`。
3. 唯一组织名：解析 corpId 后读取组织默认账号。
4. 唯一本地 profile 名：精确身份。

账号段先匹配 userId，再匹配该组织内唯一 userName。组织名或用户名重名时
报错并列出稳定 `corpId:userId` 候选。

组织默认账号缺失时，只有该组织恰好一个账号才可直接使用；多个账号必须报错。
禁止依赖时间、数组顺序、`primaryProfile`、`previousProfile` 或 Token 可读性猜测。

## 指针行为

- 历史 `corpId` 指针迁移后保存精确身份。
- `profile switch <corpId>` 解析组织默认账号后保存精确身份。
- `profile switch <组织:账号>` 保存精确身份，并更新 `orgCurrentProfiles` 和兼容镜像。
- `profile switch -` 在两个精确身份间切换，并更新目标组织默认账号。
- 登录保持现有“登录结果成为当前组织”的行为：
  - 新账号写入 `currentProfile` 和 `orgCurrentProfiles[corpId]`。
  - 旧全局当前身份写入 `previousProfile`。
- 新登录不写 `primaryProfile`。

## 登录流程

OAuth/设备流初始响应可能没有 `userId`，不得先写 `auth-token:<corpId>`。

流程：

1. 在内存中持有新 Token。
2. 通过内部 ToolCaller 的临时 Token 覆盖调用 `contact.get_current_user_profile`。
3. 补全 `userId/userName/corpName`。
4. 在同一认证锁内写账号 Token、profile 元数据和兼容镜像。
5. 身份识别失败时不覆盖任何已有组织数据。

兼容的隐藏 `auth exchange --uid` 可直接使用传入的 `uid`；没有 `uid` 时也走身份补全。

## 列表与命令

`profile list` 默认输出全部账号，每条包含原有字段，并增加：

```json
{
  "profile": "corp-a:user-1",
  "isOrgCurrent": true
}
```

原有 `primaryProfile/currentProfile/previousProfile` 字段继续输出；新版本的当前和上一账号为精确选择器。`primaryProfile/isPrimary` 仅兼容保留。

`status/expiresAt/refreshExpAt` 每次从真实身份 Token 计算，列表不触发刷新；历史摘要只兼容读取，不参与显示判断。表格移除 `PRI` 列，也不按时间排序。

命令语义：

- `auth status --profile <corpId>`：组织当前账号。
- `auth status --profile <corpId>:<userId>`：指定账号。
- `auth logout --profile <corpId>`：退出该组织全部账号。
- `auth logout --profile <corpId>:<userId>`：只退出指定账号。
- 精确退出组织当前账号时：只剩一个账号则设为组织默认；仍有多个则清空默认关系，后续只传组织必须报错。

## 批量执行

CSV `--profile` 先解析为实际身份，再按 `corpId + userId` 去重：

- `corpA,corpA`：一次。
- `corpA:user1,corpA:user2`：两次。
- `corpA,corpA:user1`：若组织当前账号是 user1，则一次；否则两次。

结果保留 `corpId/corpName`，新增 `userId/userName/profile`，不删除历史字段。

## 子进程与缓存

- Runtime profile 保存原始选择器，Token 缓存键继续使用选择器。
- Event 子进程必须转发精确身份选择器，不能只传 `corpId`。
- Connect 守护进程原样保存和重放 `--profile`，天然支持精确身份。
- 审计按实际加载到的 Token 写入 `corpId + userId`。

## 迁移、导入与重置

- 加载 `version: 1` 历史 profile 时按明确证据迁移，完成后写 `version: 2`。
- 若 `auth-token:<corpId>` 中有 `userId`，自动复制到身份 Token 槽位。
- 若 profile 的 `userId` 为空但兼容 Token 有 `userId`，补全该记录。
- 历史组织默认账号只接受：合法 `orgCurrentProfiles`、匹配的组织镜像、精确指针或单账号组织。
- `version: 2` 中缺失的组织默认关系视为明确未选择，不再从 `previousProfile` 等字段补回。
- 迁移幂等，身份 Token 写成功前不删除或改写兼容 Token。
- 导出继续打包全部 `auth-token*` 条目和 `profiles.json`。
- 导入旧包后首次加载自动补全身份 Token。
- `auth reset` 删除所有 profile 身份 Token、组织镜像和全局镜像。

## 测试

必须覆盖：

- 同组织两个账号不会覆盖。
- 同账号重新登录只更新自身。
- 历史单账号数据自动补全且指令不变。
- 历史三种指针继续可读，新写入的当前/上一账号为精确身份。
- `--profile corpId` 只使用明确组织默认账号。
- 四种组织/账号 ID/名称组合和重名报错。
- 一次性精确执行和刷新不改变组织当前账号。
- 组织级与身份级退出。
- 默认列表显示重复 `corpId` 的全部账号，并使用真实身份 Token 状态。
- CSV 按实际身份去重。
- Event 子进程转发精确身份。
- 导入导出保留全部账号。
- 全部旧多组织回归继续通过。
