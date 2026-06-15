# 钉钉开放平台文档大目录

> 钉钉开放平台官方文档大目录，供 AI Agent 发现接口：按能力域组织，每条为官方文档有效链接（`https://open.dingtalk.com/document/development/<slug>`，经内容长度校验）。有效 / 无效 / 总数统计见下方区块首行。发现接口后按文档调用（可用 dws CLI 或直接 HTTP）。

## 手册

- [执行手册 SKILL.md](SKILL.md): 发现 → 鉴权 → 调用 → 验证 → 排错 的分步执行手册。
- [API 文档目录 api-catalog.md](references/api-catalog.md): 同源全量有效链接目录。
- 官方文档站：`https://open.dingtalk.com/document`

<!-- BEGIN GENERATED: capability_domains -->
<!-- 生成: python scripts/build_llm_md.py --inject <file>；源 /Users/xuan/Skills/docs/all；模板 https://open.dingtalk.com/document/development/<slug> -->

> 全量 slug **1520** 条；有效(≥50000B) **1513** 条，无效(空壳) **7** 条；本目录收录 **1513** 条真实文档链接。

## 应用开发（642 条）

二级 Tab：开发指南 / 服务端API / 客户端JSAPI / 事件订阅 / 钉钉CLI / 开发工具 / 平台服务（公开 URL 待 `doc_url_mapping` 灌入）。

### 钉钉文档（71 条）

- [修改知识库文档成员权限](https://open.dingtalk.com/document/development/update-team-space-document-user-permissions): `update-team-space-document-user-permissions`
- [创建工作表](https://open.dingtalk.com/document/development/create-a-worksheet): `create-a-worksheet`
- [创建快捷方式](https://open.dingtalk.com/document/development/api-createshortcut): `api-createshortcut`
- [创建条件格式规则](https://open.dingtalk.com/document/development/create-conditional-formatting-rules): `create-conditional-formatting-rules`
- [创建浮动图片](https://open.dingtalk.com/document/development/api-createfloatimage): `api-createfloatimage`
- [创建知识库文档](https://open.dingtalk.com/document/development/create-team-space-document): `create-team-space-document`
- [创建筛选](https://open.dingtalk.com/document/development/api-createfilter): `api-createfilter`
- [创建筛选视图](https://open.dingtalk.com/document/development/api-createfilterview): `api-createfilterview`
- [删除下拉列表](https://open.dingtalk.com/document/development/delete-drop-down-list): `delete-drop-down-list`
- [删除列](https://open.dingtalk.com/document/development/delete-column): `delete-column`
- [删除块元素](https://open.dingtalk.com/document/development/api-docdeleteblock): `api-docdeleteblock`
- [删除工作表](https://open.dingtalk.com/document/development/delete-classic-workbooks): `delete-classic-workbooks`
- [删除浮动图片](https://open.dingtalk.com/document/development/api-deletefloatimage): `api-deletefloatimage`
- [删除知识库成员](https://open.dingtalk.com/document/development/delete-team-space-user-permissions): `delete-team-space-user-permissions`
- [删除知识库文档](https://open.dingtalk.com/document/development/delete-team-space-documents): `delete-team-space-documents`
- [删除知识库文档成员](https://open.dingtalk.com/document/development/delete-team-space-document-permissions): `delete-team-space-document-permissions`
- [删除筛选](https://open.dingtalk.com/document/development/api-deletefilter): `api-deletefilter`
- [删除筛选条件](https://open.dingtalk.com/document/development/api-clearfiltercriteria): `api-clearfiltercriteria`
- [删除筛选视图](https://open.dingtalk.com/document/development/api-deletefilterview): `api-deletefilterview`
- [删除筛选视图条件](https://open.dingtalk.com/document/development/api-clearfilterviewcriteria): `api-clearfilterviewcriteria`
- [删除行](https://open.dingtalk.com/document/development/delete-row): `delete-row`
- [合并单元格](https://open.dingtalk.com/document/development/merge-cells): `merge-cells`
- [在段落末尾追加文本](https://open.dingtalk.com/document/development/api-docappendtext): `api-docappendtext`
- [在段落末尾追加行内元素](https://open.dingtalk.com/document/development/api-docappendparagraph): `api-docappendparagraph`
- [复制文档](https://open.dingtalk.com/document/development/api-copydoc): `api-copydoc`
- [工作表中追加行](https://open.dingtalk.com/document/development/append-line): `append-line`
- [批量设置列宽](https://open.dingtalk.com/document/development/api-setcolumnswidth): `api-setcolumnswidth`
- [批量设置行高](https://open.dingtalk.com/document/development/api-setrowsheight): `api-setrowsheight`
- [指定列左侧插入若干列](https://open.dingtalk.com/document/development/insert-column-before-column): `insert-column-before-column`
- [指定行上方插入若干行](https://open.dingtalk.com/document/development/insert-rows-before-rows): `insert-rows-before-rows`
- [插入下拉列表](https://open.dingtalk.com/document/development/insert-drop-down-list): `insert-drop-down-list`
- [插入内容](https://open.dingtalk.com/document/development/api-insertcontent): `api-insertcontent`
- [插入块元素](https://open.dingtalk.com/document/development/api-docinsertblocks): `api-docinsertblocks`
- [新建知识库](https://open.dingtalk.com/document/development/create-a-team-space): `create-a-team-space`
- [更新单元格区域](https://open.dingtalk.com/document/development/update-cell-properties): `update-cell-properties`
- [更新块元素](https://open.dingtalk.com/document/development/api-docblocksmodify): `api-docblocksmodify`
- [更新工作表](https://open.dingtalk.com/document/development/update-worksheet): `update-worksheet`
- [更新浮动图片](https://open.dingtalk.com/document/development/api-updatefloatimage): `api-updatefloatimage`
- [更新知识库成员权限](https://open.dingtalk.com/document/development/update-team-space-user-permissions): `update-team-space-user-permissions`
- [更新筛选](https://open.dingtalk.com/document/development/api-updatefilter): `api-updatefilter`
- [更新筛选视图](https://open.dingtalk.com/document/development/api-updatefilterview): `api-updatefilterview`
- [查找工作表中的单元格](https://open.dingtalk.com/document/development/find-the-next-eligible-cell): `find-the-next-eligible-cell`
- [查找所有符合条件的单元格](https://open.dingtalk.com/document/development/find-all-matching-cells): `find-all-matching-cells`
- [查询块元素](https://open.dingtalk.com/document/development/api-docblocksquery): `api-docblocksquery`
- [查询用户有权限的知识库列表](https://open.dingtalk.com/document/development/querying-the-list-of-user-team-spaces): `querying-the-list-of-user-team-spaces`
- [查询知识库下的目录结构](https://open.dingtalk.com/document/development/query-the-directory-tree-in-the-knowledge-base): `query-the-directory-tree-in-the-knowledge-base`
- [查询知识库信息](https://open.dingtalk.com/document/development/query-team-space): `query-team-space`
- [查询知识库节点信息](https://open.dingtalk.com/document/development/query-knowledge-base-node-information): `query-knowledge-base-node-information`
- [根据 dentryUuid 获取 spaceId](https://open.dingtalk.com/document/development/api-getdentryidbyuuid): `api-getdentryidbyuuid`
- [添加知识库成员](https://open.dingtalk.com/document/development/add-permissions-for-team-space-members): `add-permissions-for-team-space-members`
- [添加知识库文档成员](https://open.dingtalk.com/document/development/add-workspace-document-user-permissions): `add-workspace-document-user-permissions`
- [清除单元格区域内所有内容](https://open.dingtalk.com/document/development/clear-all): `clear-all`
- [清除单元格区域内数据](https://open.dingtalk.com/document/development/clear-cell-data): `clear-cell-data`
- [知识库转交所有者](https://open.dingtalk.com/document/development/api-handoveryworkspace): `api-handoveryworkspace`
- [筛选排序](https://open.dingtalk.com/document/development/api-sortfilter): `api-sortfilter`
- [置顶知识库](https://open.dingtalk.com/document/development/api-pinspace): `api-pinspace`
- [获取 dentryUuid 信息](https://open.dingtalk.com/document/development/api-getuuidbydentryid): `api-getuuidbydentryid`
- [获取任务状态](https://open.dingtalk.com/document/development/api-gettaskinfo): `api-gettaskinfo`
- [获取单元格区域](https://open.dingtalk.com/document/development/get-cell-properties): `get-cell-properties`
- [获取工作表](https://open.dingtalk.com/document/development/obtain-worksheet-properties): `obtain-worksheet-properties`
- [获取所有工作表](https://open.dingtalk.com/document/development/obtain-all-worksheets): `obtain-all-worksheets`
- [获取筛选](https://open.dingtalk.com/document/development/api-getfilter): `api-getfilter`
- [获取筛选视图列表](https://open.dingtalk.com/document/development/api-getfilterviews): `api-getfilterviews`
- [获取资源上传信息](https://open.dingtalk.com/document/development/api-getresourceuploadinfo): `api-getresourceuploadinfo`
- [覆写文档（个人授权）](https://open.dingtalk.com/document/development/api-docupdatecontent): `api-docupdatecontent`
- [覆写文档（应用授权）](https://open.dingtalk.com/document/development/api-doc-updatecontent): `api-doc-updatecontent`
- [设置列隐藏或显示](https://open.dingtalk.com/document/development/set-column-visibility): `set-column-visibility`
- [设置筛选条件](https://open.dingtalk.com/document/development/api-setfiltercriteria): `api-setfiltercriteria`
- [设置筛选视图条件](https://open.dingtalk.com/document/development/api-setfilterviewcriteria): `api-setfilterviewcriteria`
- [设置自动行高](https://open.dingtalk.com/document/development/set-row-height-automatically): `set-row-height-automatically`
- [设置行隐藏或显示](https://open.dingtalk.com/document/development/set-row-visibility): `set-row-visibility`

### 存储（53 条）

- [修改权限](https://open.dingtalk.com/document/development/modify-permissions-file): `modify-permissions-file`
- [修改权限](https://open.dingtalk.com/document/development/modify-storage-permissions): `modify-storage-permissions`
- [初始化文件分片上传](https://open.dingtalk.com/document/development/initialize-a-multipart-upload-object): `initialize-a-multipart-upload-object`
- [删除回收项](https://open.dingtalk.com/document/development/delete-recycle-item): `delete-recycle-item`
- [删除文件或文件夹](https://open.dingtalk.com/document/development/delete-a-file-or-folder): `delete-a-file-or-folder`
- [删除文件或文件夹的应用属性](https://open.dingtalk.com/document/development/delete-file-app-attribute): `delete-file-app-attribute`
- [删除权限](https://open.dingtalk.com/document/development/delete-permissions-file): `delete-permissions-file`
- [删除权限](https://open.dingtalk.com/document/development/delete-storage-permissions): `delete-storage-permissions`
- [取消订阅文件变更事件](https://open.dingtalk.com/document/development/unsubscribe-from-file-change-events): `unsubscribe-from-file-change-events`
- [复制文件或文件夹](https://open.dingtalk.com/document/development/copy-an-object): `copy-an-object`
- [恢复文件历史版本](https://open.dingtalk.com/document/development/restore-previous-versions-of-files): `restore-previous-versions-of-files`
- [批量删除回收项](https://open.dingtalk.com/document/development/batch-delete-recycle-items): `batch-delete-recycle-items`
- [批量删除文件或文件夹](https://open.dingtalk.com/document/development/delete-files-or-folders-in-bulk): `delete-files-or-folders-in-bulk`
- [批量复制文件或文件夹](https://open.dingtalk.com/document/development/copy-files-or-folders-in-bulk): `copy-files-or-folders-in-bulk`
- [批量移动文件或文件夹](https://open.dingtalk.com/document/development/bulk-move-files-or-folders): `bulk-move-files-or-folders`
- [批量获取文件或文件夹信息](https://open.dingtalk.com/document/development/get-file-or-folder-information-in-bulk): `get-file-or-folder-information-in-bulk`
- [批量获取文件缩略图](https://open.dingtalk.com/document/development/get-file-thumbnails-in-bulk): `get-file-thumbnails-in-bulk`
- [批量还原回收项](https://open.dingtalk.com/document/development/batch-restore-recycled-items): `batch-restore-recycled-items`
- [提交文件](https://open.dingtalk.com/document/development/submit-documents): `submit-documents`
- [提交文件](https://open.dingtalk.com/document/development/submittal-file): `submittal-file`
- [搜索文件](https://open.dingtalk.com/document/development/search-for-files): `search-for-files`
- [搜索知识库](https://open.dingtalk.com/document/development/search-knowledge-base): `search-knowledge-base`
- [更新文件或文件夹的应用属性](https://open.dingtalk.com/document/development/update-file-application-properties): `update-file-application-properties`
- [添加文件夹](https://open.dingtalk.com/document/development/add-folder): `add-folder`
- [添加权限](https://open.dingtalk.com/document/development/add-permissions-file): `add-permissions-file`
- [添加权限](https://open.dingtalk.com/document/development/add-storage-permissions): `add-storage-permissions`
- [添加空间](https://open.dingtalk.com/document/development/add-space): `add-space`
- [清空回收站](https://open.dingtalk.com/document/development/empty-the-recycle-bin): `empty-the-recycle-bin`
- [知识库下载文件](https://open.dingtalk.com/document/development/knowledge-base-download-file): `knowledge-base-download-file`
- [移动文件或文件夹](https://open.dingtalk.com/document/development/move-a-file-or-folder): `move-a-file-or-folder`
- [获取企业信息](https://open.dingtalk.com/document/development/obtain-enterprise-storage-related-information): `obtain-enterprise-storage-related-information`
- [获取回收站信息](https://open.dingtalk.com/document/development/obtain-information-about-the-recycle-bin): `obtain-information-about-the-recycle-bin`
- [获取回收项信息](https://open.dingtalk.com/document/development/obtain-recycling-item-information): `obtain-recycling-item-information`
- [获取回收项列表](https://open.dingtalk.com/document/development/gets-the-list-of-recycle-items): `gets-the-list-of-recycle-items`
- [获取存储中异步任务信息](https://open.dingtalk.com/document/development/get-the-asynchronous-task-information-in-storage): `get-the-asynchronous-task-information-in-storage`
- [获取应用信息](https://open.dingtalk.com/document/development/obtains-the-information-about-the-current-application): `obtains-the-information-about-the-current-application`
- [获取文件上传信息](https://open.dingtalk.com/document/development/obtain-file-upload-informations): `obtain-file-upload-informations`
- [获取文件上传信息](https://open.dingtalk.com/document/development/obtain-storage-upload-information): `obtain-storage-upload-information`
- [获取文件下载信息](https://open.dingtalk.com/document/development/obtains-the-download-information-about-a-file): `obtains-the-download-information-about-a-file`
- [获取文件分片上传信息](https://open.dingtalk.com/document/development/obtains-the-information-about-multipart-uploads-of-an-object): `obtains-the-information-about-multipart-uploads-of-an-object`
- [获取文件或文件夹信息](https://open.dingtalk.com/document/development/obtain-file-or-folder-information): `obtain-file-or-folder-information`
- [获取文件或文件夹列表](https://open.dingtalk.com/document/development/get-a-list-of-files-or-folders): `get-a-list-of-files-or-folders`
- [获取文件版本列表](https://open.dingtalk.com/document/development/obtains-a-list-of-file-versions): `obtains-a-list-of-file-versions`
- [获取文件预览或编辑信息](https://open.dingtalk.com/document/development/obtains-the-object-preview-or-editing-information): `obtains-the-object-preview-or-editing-information`
- [获取权限列表](https://open.dingtalk.com/document/development/get-permission-list): `get-permission-list`
- [获取权限列表](https://open.dingtalk.com/document/development/get-the-storage-permission-list): `get-the-storage-permission-list`
- [获取权限继承模式](https://open.dingtalk.com/document/development/get-permission-inheritance-mode): `get-permission-inheritance-mode`
- [获取空间下所有文件或文件夹列表](https://open.dingtalk.com/document/development/get-a-list-of-all-files-or-folders-under-a): `get-a-list-of-all-files-or-folders-under-a`
- [获取空间信息](https://open.dingtalk.com/document/development/get-space-information): `get-space-information`
- [订阅文件变更事件](https://open.dingtalk.com/document/development/subscribe-to-file-change-events): `subscribe-to-file-change-events`
- [设置权限继承模式](https://open.dingtalk.com/document/development/set-permission-inheritance-mode): `set-permission-inheritance-mode`
- [还原回收项](https://open.dingtalk.com/document/development/restore-recycle-items): `restore-recycle-items`
- [重命名文件或文件夹](https://open.dingtalk.com/document/development/rename-a-file-or-folder): `rename-a-file-or-folder`

### 项目管理（50 条）

- [任务迁移至回收站](https://open.dingtalk.com/document/development/archive-tasks): `archive-tasks`
- [创建实际工时](https://open.dingtalk.com/document/development/create-actual-work): `create-actual-work`
- [创建或更新项目概览中自定义字段值](https://open.dingtalk.com/document/development/create-or-update-field-values-project-overview): `create-or-update-field-values-project-overview`
- [创建自由任务](https://open.dingtalk.com/document/development/create-a-free-task): `create-a-free-task`
- [创建计划工时](https://open.dingtalk.com/document/development/create-planned-work): `create-planned-work`
- [创建项目](https://open.dingtalk.com/document/development/create-project): `create-project`
- [创建项目任务](https://open.dingtalk.com/document/development/create-a-project-task): `create-a-project-task`
- [删除任务](https://open.dingtalk.com/document/development/delete-task): `delete-task`
- [删除项目成员](https://open.dingtalk.com/document/development/delete-project-members): `delete-project-members`
- [增加或删除自由任务的参与者](https://open.dingtalk.com/document/development/change-task-participant): `change-task-participant`
- [归档项目](https://open.dingtalk.com/document/development/archiving-project): `archiving-project`
- [恢复项目归档](https://open.dingtalk.com/document/development/cancel-project-archiving): `cancel-project-archiving`
- [批量获取自由任务详情](https://open.dingtalk.com/document/development/obtains-details-about-multiple-free-tasks): `obtains-details-about-multiple-free-tasks`
- [搜索任务工作流状态](https://open.dingtalk.com/document/development/search-task-workflow-status): `search-task-workflow-status`
- [搜索企业项目模板](https://open.dingtalk.com/document/development/search-for-enterprise-custom-templates-by-project-template-name): `search-for-enterprise-custom-templates-by-project-template-name`
- [更新任务优先级](https://open.dingtalk.com/document/development/update-task-priority): `update-task-priority`
- [更新任务参与者](https://open.dingtalk.com/document/development/update-task-participants): `update-task-participants`
- [更新任务备注](https://open.dingtalk.com/document/development/update-task-notes): `update-task-notes`
- [更新任务工作流状态](https://open.dingtalk.com/document/development/update-task-workflow-status): `update-task-workflow-status`
- [更新任务开始时间](https://open.dingtalk.com/document/development/update-task-start-time): `update-task-start-time`
- [更新任务截止时间](https://open.dingtalk.com/document/development/update-task-deadline): `update-task-deadline`
- [更新任务执行者](https://open.dingtalk.com/document/development/update-task-performer): `update-task-performer`
- [更新任务标题](https://open.dingtalk.com/document/development/update-task-content): `update-task-content`
- [更新自由任务备注](https://open.dingtalk.com/document/development/update-free-task-notes): `update-free-task-notes`
- [更新自由任务截止时间](https://open.dingtalk.com/document/development/change-free-task-deadline): `change-free-task-deadline`
- [更新自由任务执行者](https://open.dingtalk.com/document/development/change-free-task-executor): `change-free-task-executor`
- [更新自由任务标题](https://open.dingtalk.com/document/development/change-free-task-title): `change-free-task-title`
- [更新自由任务状态](https://open.dingtalk.com/document/development/change-free-task-status): `change-free-task-status`
- [更新自由任务的优先级](https://open.dingtalk.com/document/development/change-free-task-priority): `change-free-task-priority`
- [更新项目任务的自定义字段值](https://open.dingtalk.com/document/development/update-task-custom-field-value): `update-task-custom-field-value`
- [更新项目所在的分组](https://open.dingtalk.com/document/development/update-project-grouping): `update-project-grouping`
- [查询任务分组](https://open.dingtalk.com/document/development/query-task-grouping): `query-task-grouping`
- [查询任务工作流](https://open.dingtalk.com/document/development/query-task-workflow): `query-task-workflow`
- [查询优先级列表](https://open.dingtalk.com/document/development/query-a-priority-list): `query-a-priority-list`
- [查询员工可见的项目分组](https://open.dingtalk.com/document/development/query-available-project-groups): `query-available-project-groups`
- [查询用户任务信息列表](https://open.dingtalk.com/document/development/querying-user-tasks): `querying-user-tasks`
- [查询项目](https://open.dingtalk.com/document/development/query-enterprise-all-projects): `query-enterprise-all-projects`
- [查询项目中的任务](https://open.dingtalk.com/document/development/query-tasks-in-a-project): `query-tasks-in-a-project`
- [查询项目状态](https://open.dingtalk.com/document/development/query-project-status): `query-project-status`
- [根据userId获取Teambition项目用户ID](https://open.dingtalk.com/document/development/obtain-dingtalk-teambition-user-id-based-on-userid): `obtain-dingtalk-teambition-user-id-based-on-userid`
- [根据项目模板创建项目](https://open.dingtalk.com/document/development/create-a-project-from-a-project-template): `create-a-project-from-a-project-template`
- [添加任务的关联内容](https://open.dingtalk.com/document/development/create-a-linked-object-associated-with-a-task): `create-a-linked-object-associated-with-a-task`
- [添加项目成员](https://open.dingtalk.com/document/development/add-project-members): `add-project-members`
- [获取Teambition项目企业ID](https://open.dingtalk.com/document/development/obtain-the-teambition-enterprise-id): `obtain-the-teambition-enterprise-id`
- [获取任务列表](https://open.dingtalk.com/document/development/get-task-list): `get-task-list`
- [获取任务详情](https://open.dingtalk.com/document/development/get-task-details): `get-task-details`
- [获取用户加入的项目](https://open.dingtalk.com/document/development/get-projects-joined-by-users): `get-projects-joined-by-users`
- [获取自由任务详情](https://open.dingtalk.com/document/development/queries-free-task-details): `queries-free-task-details`
- [获取项目成员](https://open.dingtalk.com/document/development/get-project-members): `get-project-members`
- [项目放入回收站](https://open.dingtalk.com/document/development/items-in-recycle-bin): `items-in-recycle-bin`

### 通讯录（49 条）

- [企业账号修改钉钉号](https://open.dingtalk.com/document/development/api-changedingtalkid): `api-changedingtalkid`
- [企业账号转交主管理员（创建者）](https://open.dingtalk.com/document/development/transfer-exclusive-account-to-main-administrator-creator): `transfer-exclusive-account-to-main-administrator-creator`
- [停用企业账号](https://open.dingtalk.com/document/development/disable-an-exclusive-account): `disable-an-exclusive-account`
- [创建上下游组织](https://open.dingtalk.com/document/development/create-a-cooperation-space): `create-a-cooperation-space`
- [删除用户属性可见性设置](https://open.dingtalk.com/document/development/delete-enterprise-employee-attribute-field-visibility-settings): `delete-enterprise-employee-attribute-field-visibility-settings`
- [删除通讯录隐藏设置](https://open.dingtalk.com/document/development/delete-hide-settings): `delete-hide-settings`
- [删除限制查看通讯录设置](https://open.dingtalk.com/document/development/delete-visible-restrictions): `delete-visible-restrictions`
- [启用企业账号](https://open.dingtalk.com/document/development/enable-a-dedicated-account): `enable-a-dedicated-account`
- [异步转译通讯录ID](https://open.dingtalk.com/document/development/asynchronous-address-book-file-content-translation): `asynchronous-address-book-file-content-translation`
- [强制登出企业账号](https://open.dingtalk.com/document/development/force-logout-from-dedicated-account): `force-logout-from-dedicated-account`
- [批量通过伙伴组织的加入申请](https://open.dingtalk.com/document/development/apply-for-batch-addition-through-upstream-and-downstream-organizations): `apply-for-batch-addition-through-upstream-and-downstream-organizations`
- [批量通过伙伴组织的加入申请](https://open.dingtalk.com/document/development/batch-through-the-application-of-partner-organizations-to-join-contact): `batch-through-the-application-of-partner-organizations-to-join-contact`
- [授权企业账号可加入多组织](https://open.dingtalk.com/document/development/authorize-a-dedicated-account-to-join-multiple-organizations): `authorize-a-dedicated-account-to-join-multiple-organizations`
- [授权其他组织查看本组织的企业账号信息](https://open.dingtalk.com/document/development/api-orgaccountmobilevisibleinotherorg): `api-orgaccountmobilevisibleinotherorg`
- [搜索用户userId](https://open.dingtalk.com/document/development/address-book-search-user-id): `address-book-search-user-id`
- [搜索部门ID](https://open.dingtalk.com/document/development/address-book-search-department-id): `address-book-search-department-id`
- [新增或修改限制查看通讯录设置](https://open.dingtalk.com/document/development/add-or-modify-visibility-settings-for-address-book-restrictions): `add-or-modify-visibility-settings-for-address-book-restrictions`
- [新增或更新通讯录隐藏设置](https://open.dingtalk.com/document/development/update-address-book-hide-settings): `update-address-book-hide-settings`
- [更新伙伴组织在上下游组织内的属性信息](https://open.dingtalk.com/document/development/update-properties-of-branches-in-alibaba-group-1): `update-properties-of-branches-in-alibaba-group-1`
- [更新分支组织在主干组织内的属性信息](https://open.dingtalk.com/document/development/updates-the-property-information-of-a-branch-organization-in-a): `updates-the-property-information-of-a-branch-organization-in-a`
- [查询企业账号拥有的组织](https://open.dingtalk.com/document/development/you-can-call-this-operation-to-query-the-organization-that): `you-can-call-this-operation-to-query-the-organization-that`
- [查询企业账号状态](https://open.dingtalk.com/document/development/query-dedicated-account-status-1): `query-dedicated-account-status-1`
- [查询离职记录列表](https://open.dingtalk.com/document/development/query-the-details-of-employees-who-have-left-office): `query-the-details-of-employees-who-have-left-office`
- [根据原dingId查询迁移后的dingId](https://open.dingtalk.com/document/development/query-the-new-dingid-based-on-the-original-dingid): `query-the-new-dingid-based-on-the-original-dingid`
- [根据原unionId查询迁移后的unionId](https://open.dingtalk.com/document/development/the-union-id-that-you-want-to-query-you-can): `the-union-id-that-you-want-to-query-you-can`
- [根据手机号查询用户](https://open.dingtalk.com/document/development/query-users-by-phone-number): `query-users-by-phone-number`
- [根据迁移后的dingId查询原dingId](https://open.dingtalk.com/document/development/query-the-original-dingid-based-on-the-dingid-after-migration): `query-the-original-dingid-based-on-the-dingid-after-migration`
- [根据迁移后的unionId查询原unionId](https://open.dingtalk.com/document/development/query-the-original-union-id-based-on-the-union-id): `query-the-original-union-id-based-on-the-union-id`
- [获取上下游组织的邀请信息](https://open.dingtalk.com/document/development/obtain-the-invitation-information-of-a-cooperation-space): `obtain-the-invitation-information-of-a-cooperation-space`
- [获取上下级组织分支授权的数据](https://open.dingtalk.com/document/development/data-authorized-by-a-branch-of-an-associated-organization): `data-authorized-by-a-branch-of-an-associated-organization`
- [获取企业最新钉钉指数信息](https://open.dingtalk.com/document/development/queries-the-latest-dingtalk-index-information): `queries-the-latest-dingtalk-index-information`
- [获取企业认证信息](https://open.dingtalk.com/document/development/obtain-enterprise-authentication-information): `obtain-enterprise-authentication-information`
- [获取企业邀请信息](https://open.dingtalk.com/document/development/obtain-invitation-information): `obtain-invitation-information`
- [获取子部门ID列表](https://open.dingtalk.com/document/development/obtain-the-list-of-sub-department-ids): `obtain-the-list-of-sub-department-ids`
- [获取异步转译任务结果](https://open.dingtalk.com/document/development/obtains-the-results-of-an-asynchronous-translation-task): `obtains-the-results-of-an-asynchronous-translation-task`
- [获取用户属性可见性设置](https://open.dingtalk.com/document/development/pull-hidden-property-field-for-enterprise-employees): `pull-hidden-property-field-for-enterprise-employees`
- [获取用户通讯录个人信息](https://open.dingtalk.com/document/development/dingtalk-retrieve-user-information): `dingtalk-retrieve-user-information`
- [获取用户高管模式设置](https://open.dingtalk.com/document/development/get-user-executive-mode-settings): `get-user-executive-mode-settings`
- [获取通讯录隐藏设置](https://open.dingtalk.com/document/development/obtains-the-hide-settings-of-the-address-book): `obtains-the-hide-settings-of-the-address-book`
- [获取部门用户基础信息](https://open.dingtalk.com/document/development/queries-the-simple-information-of-a-department-user): `queries-the-simple-information-of-a-department-user`
- [获取限制查看通讯录设置列表](https://open.dingtalk.com/document/development/gets-a-list-of-address-book-limit-visibility-settings): `gets-a-list-of-address-book-limit-visibility-settings`
- [解除关联组织](https://open.dingtalk.com/document/development/disassociate-an-organization): `disassociate-an-organization`
- [解除关联组织](https://open.dingtalk.com/document/development/disassociate-upstream-and-downstream-organizations): `disassociate-upstream-and-downstream-organizations`
- [设置伙伴组织在上下游组织内的可见范围](https://open.dingtalk.com/document/development/set-the-visible-range-of-the-branch-in-the-group-1): `set-the-visible-range-of-the-branch-in-the-group-1`
- [设置分支组织在主干组织内的可见范围](https://open.dingtalk.com/document/development/sets-the-visible-range-of-branch-organizations-within-the-group): `sets-the-visible-range-of-branch-organizations-within-the-group`
- [设置用户属性可见性](https://open.dingtalk.com/document/development/add-or-update-the-hidden-settings-of-the-employee-property): `add-or-update-the-hidden-settings-of-the-employee-property`
- [设置部门可见性优先级](https://open.dingtalk.com/document/development/set-address-book-visibility-sub-department-settings-to-take-precedence): `set-address-book-visibility-sub-department-settings-to-take-precedence`
- [设置高管模式](https://open.dingtalk.com/document/development/update-executive-settings): `update-executive-settings`
- [通讯录userId排序](https://open.dingtalk.com/document/development/address-book-userid-sorting): `address-book-userid-sorting`

### 电子签（41 条）

- [ISV服务商数据初始化](https://open.dingtalk.com/document/development/offline-isv-service-provider-data-initialization): `offline-isv-service-provider-data-initialization`
- [e签宝数据初始化](https://open.dingtalk.com/document/development/isv-service-provider-data-initialization): `isv-service-provider-data-initialization`
- [创建签署流程](https://open.dingtalk.com/document/development/use-the-api-to-initiate-a-signature-process): `use-the-api-to-initiate-a-signature-process`
- [取消企业授权](https://open.dingtalk.com/document/development/cancel-enterprise-authorization): `cancel-enterprise-authorization`
- [取消企业授权](https://open.dingtalk.com/document/development/e-sign-has-revoked-the-enterprise-authorization): `e-sign-has-revoked-the-enterprise-authorization`
- [套餐转售---分润模式](https://open.dingtalk.com/document/development/package-resale-1-distribution-mode): `package-resale-1-distribution-mode`
- [套餐转售---底价结算模式](https://open.dingtalk.com/document/development/package-resale-2-reserve-price-settlement-mode): `package-resale-2-reserve-price-settlement-mode`
- [套餐转售（分润模式）](https://open.dingtalk.com/document/development/package-resale-profit-distribution-model-1): `package-resale-profit-distribution-model-1`
- [套餐转售（底价结算模式）](https://open.dingtalk.com/document/development/package-resale-base-price-settlement-mode-1): `package-resale-base-price-settlement-mode-1`
- [查询个人是否实名认证](https://open.dingtalk.com/document/development/query-personal-information): `query-personal-information`
- [查询企业是否实名](https://open.dingtalk.com/document/development/check-whether-enterprise-registered-realname-e-sign-treasure): `check-whether-enterprise-registered-realname-e-sign-treasure`
- [查询企业是否实名认证](https://open.dingtalk.com/document/development/query-enterprise-information): `query-enterprise-information`
- [查询套餐余量](https://open.dingtalk.com/document/development/query-package-balance): `query-package-balance`
- [查询用户是否实名](https://open.dingtalk.com/document/development/query-whether-a-user-has-a-real-name): `query-whether-a-user-has-a-real-name`
- [获取个人实名的地址](https://open.dingtalk.com/document/development/obtain-the-address-that-is-redirected-to-the-user-s-real): `obtain-the-address-that-is-redirected-to-the-user-s-real`
- [获取企业e签宝微应用状态](https://open.dingtalk.com/document/development/obtain-the-status-of-enterprise-e-sign-treasure-micro-application): `obtain-the-status-of-enterprise-e-sign-treasure-micro-application`
- [获取企业实名地址](https://open.dingtalk.com/document/development/obtain-enterprise-real-name-address): `obtain-enterprise-real-name-address`
- [获取企业控制台地址](https://open.dingtalk.com/document/development/get-enterprise-console-address): `get-enterprise-console-address`
- [获取企业控制台地址](https://open.dingtalk.com/document/development/get-the-address-enterprise-console-through-e-sign): `get-the-address-enterprise-console-through-e-sign`
- [获取企业的e签宝微应用状态](https://open.dingtalk.com/document/development/obtain-the-current-status-of-the-company-s-e-sign-micro-application): `obtain-the-current-status-of-the-company-s-e-sign-micro-application`
- [获取发起签署任务地址](https://open.dingtalk.com/document/development/obtain-the-address-of-the-initiating-signing-task): `obtain-the-address-of-the-initiating-signing-task`
- [获取发起签署任务的地址](https://open.dingtalk.com/document/development/obtain-the-address-used-to-initiate-a-signed-task): `obtain-the-address-used-to-initiate-a-signed-task`
- [获取套餐余量](https://open.dingtalk.com/document/development/get-package-margin): `get-package-margin`
- [获取授权的页面地址](https://open.dingtalk.com/document/development/authurl): `authurl`
- [获取授权的页面地址](https://open.dingtalk.com/document/development/obtain-the-address-of-the-authorized-page): `obtain-the-address-of-the-authorized-page`
- [获取文件上传地址](https://open.dingtalk.com/document/development/obtain-the-file-upload-address-1): `obtain-the-file-upload-address-1`
- [获取文件上传地址](https://open.dingtalk.com/document/development/obtain-the-upload-url-of-a-file-1): `obtain-the-upload-url-of-a-file-1`
- [获取文件详情](https://open.dingtalk.com/document/development/gets-the-file-details): `gets-the-file-details`
- [获取文件详情](https://open.dingtalk.com/document/development/query-file-details): `query-file-details`
- [获取流程任务合同列表](https://open.dingtalk.com/document/development/obtain-the-process-task-contract-list): `obtain-the-process-task-contract-list`
- [获取流程任务用印审批列表](https://open.dingtalk.com/document/development/obtain-the-process-task-print-approval-list-1): `obtain-the-process-task-print-approval-list-1`
- [获取流程任务用印审批列表](https://open.dingtalk.com/document/development/obtains-the-print-approval-list-for-process-tasks): `obtains-the-print-approval-list-for-process-tasks`
- [获取流程任务的所有合同列表](https://open.dingtalk.com/document/development/get-a-list-of-all-contracts-for-the-process-task): `get-a-list-of-all-contracts-for-the-process-task`
- [获取流程任务详情](https://open.dingtalk.com/document/development/obtain-the-task-details-of-the-corresponding-process): `obtain-the-task-details-of-the-corresponding-process`
- [获取流程的签署详情](https://open.dingtalk.com/document/development/get-the-details-of-process-signing): `get-the-details-of-process-signing`
- [获取流程签署详细信息](https://open.dingtalk.com/document/development/get-process-sign-off-details): `get-process-sign-off-details`
- [获取流程详细信息及操作记录](https://open.dingtalk.com/document/development/obtains-the-task-details): `obtains-the-task-details`
- [获取用户实名地址](https://open.dingtalk.com/document/development/obtain-personal-real-name-address): `obtain-personal-real-name-address`
- [获取签署人签署地址](https://open.dingtalk.com/document/development/get-signatory-address): `get-signatory-address`
- [获取签署人签署地址](https://open.dingtalk.com/document/development/obtain-signing-address-of-signatory-1): `obtain-signing-address-of-signatory-1`
- [获取跳转到企业实名的地址](https://open.dingtalk.com/document/development/obtain-the-address-that-is-redirected-to-the-enterprise-s-real): `obtain-the-address-that-is-redirected-to-the-enterprise-s-real`

### 消息会话（37 条）

- [人与人会话中机器人发送互动卡片](https://open.dingtalk.com/document/development/send-dingtalk-interactive-cards-to-person-to-person-chat-sessions): `send-dingtalk-interactive-cards-to-person-to-person-chat-sessions`
- [修改钉钉客联互通群名称](https://open.dingtalk.com/document/development/modify-the-group-name): `modify-the-group-name`
- [修改钉钉客联互通群头像](https://open.dingtalk.com/document/development/modify-the-avatar-of-a-communication-group): `modify-the-avatar-of-a-communication-group`
- [关闭互动卡片吊顶](https://open.dingtalk.com/document/development/close-interactive-card-ceiling): `close-interactive-card-ceiling`
- [创建互通群](https://open.dingtalk.com/document/development/create-an-intercommunication-group): `create-an-intercommunication-group`
- [创建并开启互动卡片吊顶](https://open.dingtalk.com/document/development/send-group-helper-message): `send-group-helper-message`
- [创建店铺群](https://open.dingtalk.com/document/development/create-a-store-group): `create-a-store-group`
- [创建钉外两人群](https://open.dingtalk.com/document/development/create-two-people-outside-the-nail): `create-two-people-outside-the-nail`
- [创建钉钉客联两人互通群](https://open.dingtalk.com/document/development/creating-two-groups-of-people): `creating-two-groups-of-people`
- [创建钉钉客联普通互通群](https://open.dingtalk.com/document/development/create-common-group-new-version): `create-common-group-new-version`
- [创建钉钉客联钉外账号](https://open.dingtalk.com/document/development/create-bc-account-association): `create-bc-account-association`
- [发送轻量级互动卡片](https://open.dingtalk.com/document/development/send-lightweight-interactive-cards): `send-lightweight-interactive-cards`
- [发送钉钉互动卡片（高级版）](https://open.dingtalk.com/document/development/send-interactive-dynamic-cards-1): `send-interactive-dynamic-cards-1`
- [在钉钉客联互通群中使用机器人发送消息](https://open.dingtalk.com/document/development/group-robots-send-messages): `group-robots-send-messages`
- [在钉钉客联互通群中使用钉内账号发送消息](https://open.dingtalk.com/document/development/send-b2c-messages): `send-b2c-messages`
- [在钉钉客联互通群中使用钉外账号发送消息](https://open.dingtalk.com/document/development/send-c2b-messages): `send-c2b-messages`
- [批量查询跨钉两人互通群列表](https://open.dingtalk.com/document/development/queries-the-session-information-of-two-population-groups): `queries-the-session-information-of-two-population-groups`
- [批量设置企业群管理员](https://open.dingtalk.com/document/development/batch-setup-group-administrator): `batch-setup-group-administrator`
- [更换钉钉客联互通群群主](https://open.dingtalk.com/document/development/change-group-owner): `change-group-owner`
- [更新机器人发送互动卡片（普通版）](https://open.dingtalk.com/document/development/update-the-robot-to-send-interactive-cards): `update-the-robot-to-send-interactive-cards`
- [更新群成员的群昵称](https://open.dingtalk.com/document/development/update-group-nicknames): `update-group-nicknames`
- [更新群管理员](https://open.dingtalk.com/document/development/update-group-administrators): `update-group-administrators`
- [更新钉钉互动卡片](https://open.dingtalk.com/document/development/update-dingtalk-interactive-cards): `update-dingtalk-interactive-cards`
- [机器人发送互动卡片（普通版）](https://open.dingtalk.com/document/development/robots-send-interactive-cards): `robots-send-interactive-cards`
- [查询群内群模板机器人](https://open.dingtalk.com/document/development/search-group-scene-template-robot): `search-group-scene-template-robot`
- [查询群成员](https://open.dingtalk.com/document/development/query-group-members): `query-group-members`
- [查询群禁言状态](https://open.dingtalk.com/document/development/query-group-silence-status): `query-group-silence-status`
- [查询群简要信息](https://open.dingtalk.com/document/development/query-group-information): `query-group-information`
- [查询钉钉客联互通群成员列表](https://open.dingtalk.com/document/development/queries-the-group-member-list): `queries-the-group-member-list`
- [查询钉钉客联钉外账号未读消息数](https://open.dingtalk.com/document/development/querying-the-number-of-unread-messages-of-the-user): `querying-the-number-of-unread-messages-of-the-user`
- [添加钉钉客联互通群成员](https://open.dingtalk.com/document/development/add-a-group-member-1): `add-a-group-member-1`
- [移除钉钉客联互通群成员](https://open.dingtalk.com/document/development/remove-group-members): `remove-group-members`
- [获取群会话的OpenConversationId](https://open.dingtalk.com/document/development/obtain-group-openconversationid): `obtain-group-openconversationid`
- [获取钉钉客联H5页面地址](https://open.dingtalk.com/document/development/get-the-dingtalk-guest-group-session-address): `get-the-dingtalk-guest-group-session-address`
- [解散群聊](https://open.dingtalk.com/document/development/api-dsbandopenscenegroup): `api-dsbandopenscenegroup`
- [解散钉钉客联互通群](https://open.dingtalk.com/document/development/disband-bc-interconnection-group): `disband-bc-interconnection-group`
- [设置群成员禁言状态](https://open.dingtalk.com/document/development/set-group-members-access-control): `set-group-members-access-control`

### 智能 CRM（34 条）

- [创建个人或企业客户数据](https://open.dingtalk.com/document/development/add-crm-personal-customers): `add-crm-personal-customers`
- [创建客户群](https://open.dingtalk.com/document/development/create-a-customer-group): `create-a-customer-group`
- [创建客户群组](https://open.dingtalk.com/document/development/crm-create-group): `crm-create-group`
- [删除CRM自定义对象数据](https://open.dingtalk.com/document/development/delete-crm-custom-object-data): `delete-crm-custom-object-data`
- [删除个人或企业客户数据](https://open.dingtalk.com/document/development/delete-crm-personal-customer): `delete-crm-personal-customer`
- [发送服务窗单人消息](https://open.dingtalk.com/document/development/sends-a-single-message-from-the-service-window): `sends-a-single-message-from-the-service-window`
- [批量修改联系人数据](https://open.dingtalk.com/document/development/modify-contact-data-in-batches): `modify-contact-data-in-batches`
- [批量删除跟进记录数据](https://open.dingtalk.com/document/development/batch-delete-follow-up-record-data): `batch-delete-follow-up-record-data`
- [批量发送服务窗消息](https://open.dingtalk.com/document/development/batch-sending-of-service-window-messages): `batch-sending-of-service-window-messages`
- [批量新增个人或企业客户数据](https://open.dingtalk.com/document/development/add-multiple-relationship-data-in-batches): `add-multiple-relationship-data-in-batches`
- [批量新增联系人数据](https://open.dingtalk.com/document/development/add-contact-data-in-batches): `add-contact-data-in-batches`
- [批量新增跟进记录数据](https://open.dingtalk.com/document/development/batch-add-follow-up-record-data): `batch-add-follow-up-record-data`
- [批量更新个人或企业客户数据](https://open.dingtalk.com/document/development/update-multiple-relational-data-tables-at-a-time): `update-multiple-relational-data-tables-at-a-time`
- [批量更新跟进记录数据](https://open.dingtalk.com/document/development/batch-update-follow-up-record-data): `batch-update-follow-up-record-data`
- [批量查询客户群](https://open.dingtalk.com/document/development/query-customer-groups-in-batches): `query-customer-groups-in-batches`
- [批量获取个人或企业客户数据](https://open.dingtalk.com/document/development/acquire-crm-individual-customers-in-batches): `acquire-crm-individual-customers-in-batches`
- [更新个人或企业客户数据](https://open.dingtalk.com/document/development/update-crm-personal-customers): `update-crm-personal-customers`
- [更新客户群组](https://open.dingtalk.com/document/development/crm-update-group): `crm-update-group`
- [查询客户数据](https://open.dingtalk.com/document/development/querying-customer-data): `querying-customer-data`
- [查询客户群列表](https://open.dingtalk.com/document/development/query-the-list-of-customer-groups): `query-the-list-of-customer-groups`
- [查询客户群组列表](https://open.dingtalk.com/document/development/query-groups): `query-groups`
- [查询服务窗粉丝用户基础信息](https://open.dingtalk.com/document/development/queries-the-basic-information-of-fans-in-the-service-window): `queries-the-basic-information-of-fans-in-the-service-window`
- [根据指定条件查询个人或企业客户数据](https://open.dingtalk.com/document/development/obtains-crm-individual-customers-in-batches-based-on-specified-query): `obtains-crm-individual-customers-in-batches-based-on-specified-query`
- [根据指定条件查询联系人数据](https://open.dingtalk.com/document/development/api-getcontacts): `api-getcontacts`
- [根据指定条件查询自定义对象数据](https://open.dingtalk.com/document/development/api-getobjectdata): `api-getobjectdata`
- [第三方个人应用发送服务窗单人消息](https://open.dingtalk.com/document/development/a-third-party-personal-application-sends-a-message-to-a-single): `a-third-party-personal-application-sends-a-message-to-a-single`
- [获取个人或企业客户查重字段](https://open.dingtalk.com/document/development/obtain-duplicate-check-fields): `obtain-duplicate-check-fields`
- [获取个人或企业客户的元数据](https://open.dingtalk.com/document/development/obtain-the-metadata-of-individual-enterprise-customers): `obtain-the-metadata-of-individual-enterprise-customers`
- [获取全量个人或企业客户数据](https://open.dingtalk.com/document/development/crm-obtains-all-private-sea-customer-data): `crm-obtains-all-private-sea-customer-data`
- [获取单个客户群组详情](https://open.dingtalk.com/document/development/queries-the-details-of-a-single-customer-group): `queries-the-details-of-a-single-customer-group`
- [获取单个客户群详情](https://open.dingtalk.com/document/development/obtain-a-single-customer-group): `obtain-a-single-customer-group`
- [获取审批中创建与CRM客户关联的TAB表单元数据](https://open.dingtalk.com/document/development/api-getrelatedviewtabmeta): `api-getrelatedviewtabmeta`
- [获取审批里创建的与CRM客户关联的TAB表单数据实例列表](https://open.dingtalk.com/document/development/api-getrelatedviewtabdata): `api-getrelatedviewtabdata`
- [获取客户管理全局信息](https://open.dingtalk.com/document/development/get-customer-management-global-information): `get-customer-management-global-information`

### 日历（30 条）

- [修改日程](https://open.dingtalk.com/document/development/modify-event): `modify-event`
- [创建日程](https://open.dingtalk.com/document/development/create-event): `create-event`
- [创建订阅日历](https://open.dingtalk.com/document/development/create-subscription-calendar): `create-subscription-calendar`
- [创建访问控制](https://open.dingtalk.com/document/development/create-schedule-access-control): `create-schedule-access-control`
- [删除日程](https://open.dingtalk.com/document/development/delete-event): `delete-event`
- [删除日程参与者](https://open.dingtalk.com/document/development/delete-schedule-participant): `delete-schedule-participant`
- [删除订阅日历](https://open.dingtalk.com/document/development/delete-subscription-calendar): `delete-subscription-calendar`
- [删除访问控制](https://open.dingtalk.com/document/development/delete-an-access-control-list): `delete-an-access-control-list`
- [取消订阅公共日历](https://open.dingtalk.com/document/development/unsubscribe-from-a-public-calendar): `unsubscribe-from-a-public-calendar`
- [取消预定会议室](https://open.dingtalk.com/document/development/remove-a-meeting-room): `remove-a-meeting-room`
- [更新订阅日历](https://open.dingtalk.com/document/development/update-subscription-calendar): `update-subscription-calendar`
- [查看单个日程的签到详情](https://open.dingtalk.com/document/development/view-the-check-in-details-of-a-single-schedule): `view-the-check-in-details-of-a-single-schedule`
- [查看单个日程的签退详情](https://open.dingtalk.com/document/development/view-the-billing-details-of-a-single-schedule): `view-the-billing-details-of-a-single-schedule`
- [查询单个日程详情](https://open.dingtalk.com/document/development/query-details-about-an-event): `query-details-about-an-event`
- [查询单个订阅日历详情](https://open.dingtalk.com/document/development/query-a-single-subscription-calendar): `query-a-single-subscription-calendar`
- [查询日历](https://open.dingtalk.com/document/development/query-a-calendar): `query-a-calendar`
- [查询日程列表](https://open.dingtalk.com/document/development/query-an-event-list): `query-an-event-list`
- [查询日程视图](https://open.dingtalk.com/document/development/query-schedule-view): `query-schedule-view`
- [添加日程参与者](https://open.dingtalk.com/document/development/add-schedule-participant): `add-schedule-participant`
- [获取会议室忙闲信息](https://open.dingtalk.com/document/development/queries-free-and-busy-meeting-room-information): `queries-free-and-busy-meeting-room-information`
- [获取日程参与者](https://open.dingtalk.com/document/development/get-the-participants-of-a-schedule): `get-the-participants-of-a-schedule`
- [获取用户忙闲信息](https://open.dingtalk.com/document/development/free-schedule): `free-schedule`
- [获取签到链接](https://open.dingtalk.com/document/development/api-getsigninlink): `api-getsigninlink`
- [获取签退链接](https://open.dingtalk.com/document/development/api-getsignoutlink): `api-getsignoutlink`
- [获取访问控制列表](https://open.dingtalk.com/document/development/obtain-the-access-control-list-of-the-calendar): `obtain-the-access-control-list-of-the-calendar`
- [订阅公共日历](https://open.dingtalk.com/document/development/subscribe-to-a-public-calendar): `subscribe-to-a-public-calendar`
- [设置日程响应邀请状态](https://open.dingtalk.com/document/development/configure-response-status): `configure-response-status`
- [针对单个日程进行签到](https://open.dingtalk.com/document/development/sign-in-single-schedule-news): `sign-in-single-schedule-news`
- [针对单个日程进行签退](https://open.dingtalk.com/document/development/sign-off-for-a-single-schedule): `sign-off-for-a-single-schedule`
- [预定会议室](https://open.dingtalk.com/document/development/add-a-meeting-room): `add-a-meeting-room`

### 视频会议（30 条）

- [停止视频会议云录制](https://open.dingtalk.com/document/development/video-conferencing-stops-cloud-recording): `video-conferencing-stops-cloud-recording`
- [停止视频会议直播推流](https://open.dingtalk.com/document/development/videoconferencing-stops-live-stream-ingest): `videoconferencing-stops-live-stream-ingest`
- [全员静音或全员取消静音](https://open.dingtalk.com/document/development/mute-all-staff-or-unmute-all-staff): `mute-all-staff-or-unmute-all-staff`
- [关闭视频会议](https://open.dingtalk.com/document/development/close-audio-video-conference): `close-audio-video-conference`
- [创建用户专属短链](https://open.dingtalk.com/document/development/api-createcustomshortlink): `api-createcustomshortlink`
- [创建视频会议](https://open.dingtalk.com/document/development/create-a-video-conference): `create-a-video-conference`
- [创建预约会议](https://open.dingtalk.com/document/development/create-appointment-meeting): `create-appointment-meeting`
- [取消预约会议](https://open.dingtalk.com/document/development/cancel-appointment-meeting): `cancel-appointment-meeting`
- [开启视频会议云录制](https://open.dingtalk.com/document/development/video-conference-open-cloud-recording): `video-conference-open-cloud-recording`
- [开启视频会议直播推流](https://open.dingtalk.com/document/development/video-conference-enables-live-stream-ingest): `video-conference-enables-live-stream-ingest`
- [批量查询视频会议信息](https://open.dingtalk.com/document/development/batch-query-of-video-conference-information): `batch-query-of-video-conference-information`
- [指定人员静音或取消静音](https://open.dingtalk.com/document/development/specify-person-to-mute-or-unmute): `specify-person-to-mute-or-unmute`
- [更新预约会议](https://open.dingtalk.com/document/development/update-appointment-meeting): `update-appointment-meeting`
- [更新预约会议设置](https://open.dingtalk.com/document/development/api-updatescheduleconfsettings): `api-updatescheduleconfsettings`
- [查询企业进行中会议列表](https://open.dingtalk.com/document/development/api-queryorgconferencelist): `api-queryorgconferencelist`
- [查询会议录制中的文本信息](https://open.dingtalk.com/document/development/queries-the-text-information-about-cloud-recording): `queries-the-text-information-about-cloud-recording`
- [查询会议录制中的视频信息](https://open.dingtalk.com/document/development/queries-the-playback-information-about-a-recorded-cloud-video): `queries-the-playback-information-about-a-recorded-cloud-video`
- [查询会议录制的详情信息](https://open.dingtalk.com/document/development/query-recording-information): `query-recording-information`
- [查询用户进行中会议列表](https://open.dingtalk.com/document/development/api-queryuserongoingconference): `api-queryuserongoingconference`
- [查询视频会议信息](https://open.dingtalk.com/document/development/querying-video-conference-information): `querying-video-conference-information`
- [查询视频会议成员](https://open.dingtalk.com/document/development/querying-video-conference-members): `querying-video-conference-members`
- [查询预约会议](https://open.dingtalk.com/document/development/query-meeting-reservation): `query-meeting-reservation`
- [查询预约会议历史会议信息](https://open.dingtalk.com/document/development/query-appointment-meeting-history-meeting-information): `query-appointment-meeting-history-meeting-information`
- [查询预约会议设置](https://open.dingtalk.com/document/development/api-queryscheduleconfsettings): `api-queryscheduleconfsettings`
- [根据会议号查询会议信息](https://open.dingtalk.com/document/development/api-queryconferenceinfobyroomcode): `api-queryconferenceinfobyroomcode`
- [设置全员看他](https://open.dingtalk.com/document/development/set-the-whole-staff-to-see-him): `set-the-whole-staff-to-see-him`
- [设置联席主持人](https://open.dingtalk.com/document/development/set-up-co-hosts): `set-up-co-hosts`
- [踢出会议成员](https://open.dingtalk.com/document/development/kick-out-meeting-members): `kick-out-meeting-members`
- [邀请用户入会](https://open.dingtalk.com/document/development/invite-users-to-join): `invite-users-to-join`
- [锁定会议](https://open.dingtalk.com/document/development/api-lockconference): `api-lockconference`

### 阿里商旅（25 条）

- [修改成本中心](https://open.dingtalk.com/document/development/modify-basic-cost-center-information): `modify-basic-cost-center-information`
- [关联单号查询相关订单信息列表](https://open.dingtalk.com/document/development/related-order-information): `related-order-information`
- [创建日志](https://open.dingtalk.com/document/development/create-a-log): `create-a-log`
- [同步市内用车申请单](https://open.dingtalk.com/document/development/synchronize-third-party-city-vehicle-approval-form): `synchronize-third-party-city-vehicle-approval-form`
- [回传第三方超标审批结果](https://open.dingtalk.com/document/development/dingtalk-oapi-alitrip-btrip-exceedapply-sync): `dingtalk-oapi-alitrip-btrip-exceedapply-sync`
- [审批市内用车申请单](https://open.dingtalk.com/document/development/approval-of-third-party-city-car-application-form): `approval-of-third-party-city-car-application-form`
- [搜索第三方机票超标审批单](https://open.dingtalk.com/document/development/dingtalk-oapi-alitrip-btrip-exceedapply-flight): `dingtalk-oapi-alitrip-btrip-exceedapply-flight`
- [搜索第三方火车票超标审批单](https://open.dingtalk.com/document/development/dingtalk-oapi-alitrip-btrip-exceedapply-train-get): `dingtalk-oapi-alitrip-btrip-exceedapply-train-get`
- [搜索第三方酒店超标审批单](https://open.dingtalk.com/document/development/dingtalk-oapi-alitrip-btrip-exceedapply-hotel-get): `dingtalk-oapi-alitrip-btrip-exceedapply-hotel-get`
- [新建审批单](https://open.dingtalk.com/document/development/user-new-approval-form): `user-new-approval-form`
- [新建成本中心](https://open.dingtalk.com/document/development/new-cost-center): `new-cost-center`
- [查询商旅火车票结算记账数据](https://open.dingtalk.com/document/development/business-travel-train-ticket-settlement-bookkeeping-query-interface): `business-travel-train-ticket-settlement-bookkeeping-query-interface`
- [查询市内用车申请单](https://open.dingtalk.com/document/development/query-the-application-form-for-third-party-vehicles-in-the-city): `query-the-application-form-for-third-party-vehicles-in-the-city`
- [查询成员排班信息](https://open.dingtalk.com/document/development/query-scheduling-for-a-day): `query-scheduling-for-a-day`
- [查询机票结算记账数据](https://open.dingtalk.com/document/development/ticket-settlement-bookkeeping-query-interface): `ticket-settlement-bookkeeping-query-interface`
- [查询用车结算记账记录](https://open.dingtalk.com/document/development/query-interface-for-vehicle-settlement-and-bookkeeping): `query-interface-for-vehicle-settlement-and-bookkeeping`
- [查询酒店结算记账数据](https://open.dingtalk.com/document/development/hotel-settlement-bookkeeping-query-interface): `hotel-settlement-bookkeeping-query-interface`
- [获取企业机票订单数据](https://open.dingtalk.com/document/development/obtains-enterprise-ticket-order-data): `obtains-enterprise-ticket-order-data`
- [获取商旅访问地址](https://open.dingtalk.com/document/development/obtain-business-travel-access-addresses): `obtain-business-travel-access-addresses`
- [获取打卡结果](https://open.dingtalk.com/document/development/open-attendance-clock-in-data): `open-attendance-clock-in-data`
- [获取打卡详情](https://open.dingtalk.com/document/development/attendance-clock-in-record-is-open): `attendance-clock-in-record-is-open`
- [获取日志评论详情](https://open.dingtalk.com/document/development/queries-log-comment-details): `queries-log-comment-details`
- [获取月对账结算数据](https://open.dingtalk.com/document/development/obtain-monthly-reconciliation-settlement-data): `obtain-monthly-reconciliation-settlement-data`
- [获取用车订单数据](https://open.dingtalk.com/document/development/vehicle-order-query-interface): `vehicle-order-query-interface`
- [设置成本中心人员信息](https://open.dingtalk.com/document/development/set-up-cost-center-personnel-information): `set-up-cost-center-personnel-information`

### Agoal 目标（23 条）

- [Agoal业务数据查询](https://open.dingtalk.com/document/development/agoal-business-biz-data-query): `agoal-business-biz-data-query`
- [创建业务实体](https://open.dingtalk.com/document/development/api-agoalentitycreate): `api-agoalentitycreate`
- [创建目标规则下的考核任务](https://open.dingtalk.com/document/development/api-agoalperftaskcreate): `api-agoalperftaskcreate`
- [更新业务实体](https://open.dingtalk.com/document/development/api-agoalentityupdate): `api-agoalentityupdate`
- [更新目标规则下的考核任务](https://open.dingtalk.com/document/development/api-agoalperftaskupdate): `api-agoalperftaskupdate`
- [查询企业下个人目标详情](https://open.dingtalk.com/document/development/api-getobjectivedetail): `api-getobjectivedetail`
- [查询企业下单个目标规则详情](https://open.dingtalk.com/document/development/api-getobjectiveruledetail): `api-getobjectiveruledetail`
- [查询企业下指定个人目标的所有进展](https://open.dingtalk.com/document/development/api-agoalobjectiveprogresslist): `api-agoalobjectiveprogresslist`
- [查询企业下的所有考核计划](https://open.dingtalk.com/document/development/api-agoalorgperfplanquery): `api-agoalorgperfplanquery`
- [查询企业下目标规则列表](https://open.dingtalk.com/document/development/api-agoalobjectiverulelist): `api-agoalobjectiverulelist`
- [查询某个考核计划的部门得分](https://open.dingtalk.com/document/development/api-agoalorgperfdocquery): `api-agoalorgperfdocquery`
- [查询组织目标详情](https://open.dingtalk.com/document/development/api-agoalorgobjectivequery): `api-agoalorgobjectivequery`
- [获取 Agoal 组织目标列表](https://open.dingtalk.com/document/development/api-agoalorgobjectivelist): `api-agoalorgobjectivelist`
- [获取Agoal指定目标规则下的周期列表](https://open.dingtalk.com/document/development/api-agoalobjectiveruleperiodlist): `api-agoalobjectiveruleperiodlist`
- [获取Agoal指定组织下的所有目标规则列表](https://open.dingtalk.com/document/development/api-agoalorgobjectiverulelist): `api-agoalorgobjectiverulelist`
- [获取Agoal指定规则周期下负责人的目标列表](https://open.dingtalk.com/document/development/api-agoaluserobjectivelist): `api-agoaluserobjectivelist`
- [获取Agoal指定部门下的计分卡维度和指标id](https://open.dingtalk.com/document/development/api-getdeptscorecardindicator): `api-getdeptscorecardindicator`
- [获取Agoal用户管理员列表](https://open.dingtalk.com/document/development/api-agoaluseradminlist): `api-agoaluseradminlist`
- [获取Agoal目标或关键结果关联的关键行动](https://open.dingtalk.com/document/development/api-agoalobjectivekeyactionlist): `api-agoalobjectivekeyactionlist`
- [获取计分卡指标详情](https://open.dingtalk.com/document/development/api-getindicatordetail): `api-getindicatordetail`
- [通过Agoal系统账号发送消息](https://open.dingtalk.com/document/development/api-agoalsendmessage): `api-agoalsendmessage`
- [通过指标编码批量查询指标列表](https://open.dingtalk.com/document/development/api-agoalindicatorbatchquery): `api-agoalindicatorbatchquery`
- [通过指标编码推送指标时间维度数据](https://open.dingtalk.com/document/development/api-agoalindicatordatapush): `api-agoalindicatordatapush`

### 教育（23 条）

- [修改用户成员类型](https://open.dingtalk.com/document/development/api-updatecollegeuseremptype): `api-updatecollegeuseremptype`
- [创建个人账号用户](https://open.dingtalk.com/document/development/api-addcollegecontactuser): `api-addcollegecontactuser`
- [创建组织单元](https://open.dingtalk.com/document/development/api-createcollegecontactdept): `api-createcollegecontactdept`
- [创建自定义校区或部门](https://open.dingtalk.com/document/development/create-a-custom-campus-or-department): `create-a-custom-campus-or-department`
- [创建自定义部门下的班级](https://open.dingtalk.com/document/development/create-classes-in-a-custom-department): `create-classes-in-a-custom-department`
- [创建高校账号用户](https://open.dingtalk.com/document/development/api-addcollegecontactexclusive): `api-addcollegecontactexclusive`
- [删除学生](https://open.dingtalk.com/document/development/delete-student): `delete-student`
- [删除家校部门](https://open.dingtalk.com/document/development/delete-home-school-department): `delete-home-school-department`
- [删除家长关系](https://open.dingtalk.com/document/development/delete-parent-relationship): `delete-parent-relationship`
- [删除老师](https://open.dingtalk.com/document/development/delete-teacher): `delete-teacher`
- [判断用户是否是认证组织的语文老师接口](https://open.dingtalk.com/document/development/api-isyuwencertifiedteacher): `api-isyuwencertifiedteacher`
- [学生调班](https://open.dingtalk.com/document/development/shift-students): `shift-students`
- [更新个人账号用户信息](https://open.dingtalk.com/document/development/api-updatecollegecontactuser): `api-updatecollegecontactuser`
- [更新学生](https://open.dingtalk.com/document/development/api-updatestudent): `api-updatestudent`
- [更新家长](https://open.dingtalk.com/document/development/api-updateguardian): `api-updateguardian`
- [更新班级](https://open.dingtalk.com/document/development/api-updateclass): `api-updateclass`
- [更新组织单元](https://open.dingtalk.com/document/development/api-updatecollegecontactdept): `api-updatecollegecontactdept`
- [更新高校账号用户信息](https://open.dingtalk.com/document/development/api-updatecollegecontactexclusive): `api-updatecollegecontactexclusive`
- [查询用户信息详情](https://open.dingtalk.com/document/development/api-querycollegecontactuserdetail): `api-querycollegecontactuserdetail`
- [获取子组织单元列表](https://open.dingtalk.com/document/development/api-listcollegecontactsubdepts): `api-listcollegecontactsubdepts`
- [获取组织单元支持的部门类型](https://open.dingtalk.com/document/development/api-listcollegecontactdepttypeconfig): `api-listcollegecontactdepttypeconfig`
- [获取组织单元详情](https://open.dingtalk.com/document/development/api-getcollegecontactdeptdetail): `api-getcollegecontactdeptdetail`
- [获取行政组织架构部门详情](https://open.dingtalk.com/document/development/api-getcollegecontactstandardstrudeptdetail): `api-getcollegecontactstandardstrudeptdetail`

### 智能人事（17 条）

- [修改已离职员工信息](https://open.dingtalk.com/document/development/modify-resigned-employee-information): `modify-resigned-employee-information`
- [员工加入待离职](https://open.dingtalk.com/document/development/api-empstartdismission): `api-empstartdismission`
- [批量获取员工离职信息](https://open.dingtalk.com/document/development/obtain-resignation-information-of-employees-new-version): `obtain-resignation-information-of-employees-new-version`
- [撤销员工待离职](https://open.dingtalk.com/document/development/api-revoketermination): `api-revoketermination`
- [新增或删除花名册选项类型字段的选项](https://open.dingtalk.com/document/development/intelligent-personnel-roster-field-option-modification): `intelligent-personnel-roster-field-option-modification`
- [智能人事员工调岗](https://open.dingtalk.com/document/development/intelligent-personnel-staff-transfer): `intelligent-personnel-staff-transfer`
- [智能人事员工转正](https://open.dingtalk.com/document/development/intelligent-personnel-staff-to-become-regular): `intelligent-personnel-staff-to-become-regular`
- [更新待离职员工离职信息](https://open.dingtalk.com/document/development/api-updateempdismissioninfo): `api-updateempdismissioninfo`
- [查询花名册中有权限的字段列表](https://open.dingtalk.com/document/development/query-the-list-of-fields-with-permissions-in-the-roster): `query-the-list-of-fields-with-permissions-in-the-roster`
- [添加待入职员工](https://open.dingtalk.com/document/development/add-employees-to-be-hired-supports-system-and-custom-fields): `add-employees-to-be-hired-supports-system-and-custom-fields`
- [确认员工离职并删除](https://open.dingtalk.com/document/development/api-hrmprocessterminationandhandover): `api-hrmprocessterminationandhandover`
- [获取企业已有的所有离职原因](https://open.dingtalk.com/document/development/api-getalldismissionreasons): `api-getalldismissionreasons`
- [获取企业职位列表](https://open.dingtalk.com/document/development/obtain-enterprise-position-information): `obtain-enterprise-position-information`
- [获取企业职务列表](https://open.dingtalk.com/document/development/obtain-enterprise-title-information): `obtain-enterprise-title-information`
- [获取企业职级列表](https://open.dingtalk.com/document/development/obtain-enterprise-rank-information): `obtain-enterprise-rank-information`
- [获取员工花名册字段信息](https://open.dingtalk.com/document/development/api-getemployeerosterbyfield): `api-getemployeerosterbyfield`
- [获取离职员工列表](https://open.dingtalk.com/document/development/obtain-the-list-of-employees-who-have-left): `obtain-the-list-of-employees-who-have-left`

### 机器人（16 条）

- [下载机器人接收消息的文件内容](https://open.dingtalk.com/document/development/download-the-file-content-of-the-robot-receiving-message): `download-the-file-content-of-the-robot-receiving-message`
- [人与人会话中机器人发送普通消息](https://open.dingtalk.com/document/development/the-robot-sends-ordinary-messages-in-a-person-to-person-conversation): `the-robot-sends-ordinary-messages-in-a-person-to-person-conversation`
- [企业机器人撤回内部群消息](https://open.dingtalk.com/document/development/enterprise-chatbot-withdraws-internal-group-messages): `enterprise-chatbot-withdraws-internal-group-messages`
- [发送DING消息](https://open.dingtalk.com/document/development/robot-sends-nail-message): `robot-sends-nail-message`
- [批量发送人与机器人会话中机器人消息](https://open.dingtalk.com/document/development/chatbots-send-one-on-one-chat-messages-in-batches): `chatbots-send-one-on-one-chat-messages-in-batches`
- [批量撤回人与人会话中机器人消息](https://open.dingtalk.com/document/development/batch-withdrawal-of-single-chat-robot-messages-in-person-to-person-conversations): `batch-withdrawal-of-single-chat-robot-messages-in-person-to-person-conversations`
- [批量撤回人与机器人会话中机器人消息](https://open.dingtalk.com/document/development/batch-message-recall-chat): `batch-message-recall-chat`
- [批量查询人与机器人会话机器人消息是否已读](https://open.dingtalk.com/document/development/chatbot-batch-query-the-read-status-of-messages): `chatbot-batch-query-the-read-status-of-messages`
- [撤回已经发送的DING消息](https://open.dingtalk.com/document/development/robot-withdraws-pin-message): `robot-withdraws-pin-message`
- [机器人发送群聊消息](https://open.dingtalk.com/document/development/the-robot-sends-a-group-message): `the-robot-sends-a-group-message`
- [查询人与人会话中机器人消息已读列表](https://open.dingtalk.com/document/development/query-the-read-list-of-robot-messages-in-person-to-person-conversations): `query-the-read-list-of-robot-messages-in-person-to-person-conversations`
- [查询企业机器人群聊消息用户已读状态](https://open.dingtalk.com/document/development/chatbot-queries-the-read-status-of-a-message): `chatbot-queries-the-read-status-of-a-message`
- [查询单聊机器人的快捷入口](https://open.dingtalk.com/document/development/quick-entrance-of-inquiry-single-chat-robot): `quick-entrance-of-inquiry-single-chat-robot`
- [清空单聊机器人快捷入口](https://open.dingtalk.com/document/development/clear-single-chat-robot-quick-entry): `clear-single-chat-robot-quick-entry`
- [获取群内机器人列表](https://open.dingtalk.com/document/development/obtain-the-list-of-robots-in-the-group): `obtain-the-list-of-robots-in-the-group`
- [设置单聊机器人快捷入口](https://open.dingtalk.com/document/development/set-robot-quick-entrance): `set-robot-quick-entrance`

### 简知 CRM（15 条）

- [产品信息](https://open.dingtalk.com/document/development/add-or-edit-product-information): `add-or-edit-product-information`
- [入库单](https://open.dingtalk.com/document/development/add-or-edit-a-shipment-record): `add-or-edit-a-shipment-record`
- [出库单](https://open.dingtalk.com/document/development/add-or-edit-an-issue-ticket): `add-or-edit-an-issue-ticket`
- [发货单](https://open.dingtalk.com/document/development/add-or-edit-invoices): `add-or-edit-invoices`
- [合同订单](https://open.dingtalk.com/document/development/add-or-edit-contract-orders): `add-or-edit-contract-orders`
- [客户公共池](https://open.dingtalk.com/document/development/add-or-edit-customer-public-pools): `add-or-edit-customer-public-pools`
- [客户资料](https://open.dingtalk.com/document/development/add-or-edit-customer-profile): `add-or-edit-customer-profile`
- [报价记录](https://open.dingtalk.com/document/development/add-or-edit-quotation-records): `add-or-edit-quotation-records`
- [生产单](https://open.dingtalk.com/document/development/add-or-edit-a-production-order): `add-or-edit-a-production-order`
- [联系人](https://open.dingtalk.com/document/development/add-or-edit-contacts): `add-or-edit-contacts`
- [获取数据列表](https://open.dingtalk.com/document/development/obtain-the-data-list): `obtain-the-data-list`
- [获取数据详情](https://open.dingtalk.com/document/development/queries-data-details): `queries-data-details`
- [采购单](https://open.dingtalk.com/document/development/edit-purchase-order): `edit-purchase-order`
- [销售换货单](https://open.dingtalk.com/document/development/add-or-edit-a-sales-order): `add-or-edit-a-sales-order`
- [销售机会](https://open.dingtalk.com/document/development/add-or-edit-opportunities): `add-or-edit-opportunities`

### 考勤（15 条）

- [分页获取加班规则列表](https://open.dingtalk.com/document/development/retrieve-a-list-of-overtime-rules-by-page): `retrieve-a-list-of-overtime-rules-by-page`
- [分页获取补卡规则列表](https://open.dingtalk.com/document/development/retrieve-a-list-of-replenishment-rules-by-page): `retrieve-a-list-of-replenishment-rules-by-page`
- [批量查询员工假期余额变更记录](https://open.dingtalk.com/document/development/batch-query-employee-leave-balance-change-record): `batch-query-employee-leave-balance-change-record`
- [批量获取加班规则设置](https://open.dingtalk.com/document/development/batch-retrieve-overtime-rules): `batch-retrieve-overtime-rules`
- [更新假期规则](https://open.dingtalk.com/document/development/update-holiday-rules): `update-holiday-rules`
- [查询指定用户的封账规则](https://open.dingtalk.com/document/development/encapsulate-account-sealing-and-unsealing-rules): `encapsulate-account-sealing-and-unsealing-rules`
- [查询用户某段时间内是否处于封账状态](https://open.dingtalk.com/document/development/checks-whether-a-user-has-blocked-accounts-within-a-specified): `checks-whether-a-user-has-blocked-accounts-within-a-specified`
- [查询用户考勤节假日信息](https://open.dingtalk.com/document/development/obtain-user-attendance-and-holiday-information): `obtain-user-attendance-and-holiday-information`
- [查询考勤写操作权限](https://open.dingtalk.com/document/development/attendance-writing-operation-is-brand-new-query): `attendance-writing-operation-is-brand-new-query`
- [查询考勤机信息](https://open.dingtalk.com/document/development/query-attendance-machine-information): `query-attendance-machine-information`
- [根据设备ID获取员工信息](https://open.dingtalk.com/document/development/obtain-information-about-employees-based-on-device-ids): `obtain-information-about-employees-based-on-device-ids`
- [添加假期规则](https://open.dingtalk.com/document/development/add-holiday-rules): `add-holiday-rules`
- [通知审批通过](https://open.dingtalk.com/document/development/api-processapprovefinish): `api-processapprovefinish`
- [配置考勤排班附加信息](https://open.dingtalk.com/document/development/synchronization-scheduling-information): `synchronization-scheduling-information`
- [预计算时长](https://open.dingtalk.com/document/development/api-calculateduration): `api-calculateduration`

### 登录授权（11 条）

- [查询个人授权记录](https://open.dingtalk.com/document/development/query-personal-authorization-records): `query-personal-authorization-records`
- [获取jsapiTicket](https://open.dingtalk.com/document/development/create-a-jsapi-ticket): `create-a-jsapi-ticket`
- [获取企业内部应用的accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-an-internal-app): `obtain-the-access-token-of-an-internal-app`
- [获取企业开通应用后的授权信息](https://open.dingtalk.com/document/development/obtains-the-authorization-information-after-the-enterprise-activates-the-application-1): `obtains-the-authorization-information-after-the-enterprise-activates-the-application-1`
- [获取定制应用的accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-the-third-party-application-authorization-enterprise): `obtain-the-access-token-of-the-third-party-application-authorization-enterprise`
- [获取应用的 Access Token](https://open.dingtalk.com/document/development/api-gettoken): `api-gettoken`
- [获取应用管理后台免登的用户信息](https://open.dingtalk.com/document/development/obtains-the-identity-of-an-application-administrator): `obtains-the-identity-of-an-application-administrator`
- [获取微应用后台免登的accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-the-micro-application-background-without-log-on): `obtain-the-access-token-of-the-micro-application-background-without-log-on`
- [获取用户token](https://open.dingtalk.com/document/development/obtain-user-token): `obtain-user-token`
- [获取第三方企业应用的suiteAccessToken](https://open.dingtalk.com/document/development/obtains-the-suite-acess-token-of-third-party-enterprise-applications): `obtains-the-suite-acess-token-of-third-party-enterprise-applications`
- [获取第三方应用授权企业的accessToken](https://open.dingtalk.com/document/development/obtain-the-access-token-of-the-authorized-enterprise-1): `obtain-the-access-token-of-the-authorized-enterprise-1`

### 知识库（11 条）

- [匹配文本中的词条](https://open.dingtalk.com/document/development/enterprise-encyclopedia-match-entries-in-a-text): `enterprise-encyclopedia-match-entries-in-a-text`
- [批量获取知识库](https://open.dingtalk.com/document/development/batch-acquisition-of-knowledge-base): `batch-acquisition-of-knowledge-base`
- [批量获取节点](https://open.dingtalk.com/document/development/obtain-nodes-in-batch): `obtain-nodes-in-batch`
- [新建知识库](https://open.dingtalk.com/document/development/new-knowledge-base): `new-knowledge-base`
- [查询词条详情](https://open.dingtalk.com/document/development/enterprise-encyclopedia-query-entry-details-by-entry-name): `enterprise-encyclopedia-query-entry-details-by-entry-name`
- [获取我的文档知识库信息](https://open.dingtalk.com/document/development/get-my-documents): `get-my-documents`
- [获取知识库](https://open.dingtalk.com/document/development/obtain-the-knowledge-base): `obtain-the-knowledge-base`
- [获取知识库列表](https://open.dingtalk.com/document/development/get-knowledge-base-list): `get-knowledge-base-list`
- [获取节点](https://open.dingtalk.com/document/development/get-knowledge-base-acquisition-node): `get-knowledge-base-acquisition-node`
- [获取节点列表](https://open.dingtalk.com/document/development/get-node-list): `get-node-list`
- [通过链接获取节点](https://open.dingtalk.com/document/development/get-node-by-link): `get-node-by-link`

### 服务群（8 条）

- [创建场景服务群](https://open.dingtalk.com/document/development/create-a-scenario-service-group): `create-a-scenario-service-group`
- [升级云客服服务群为钉钉智能服务群](https://open.dingtalk.com/document/development/upgraded-the-cloud-customer-service-group-to-the-dingtalk-intelligent): `upgraded-the-cloud-customer-service-group-to-the-dingtalk-intelligent`
- [升级普通群为服务群](https://open.dingtalk.com/document/development/a-dingtalk-group-is-upgraded-to-one-of-the-intelligent): `a-dingtalk-group-is-upgraded-to-one-of-the-intelligent`
- [发送服务群消息](https://open.dingtalk.com/document/development/service-group-message-sending-interface): `service-group-message-sending-interface`
- [更换服务群所在的群分组](https://open.dingtalk.com/document/development/modify-a-service-group): `modify-a-service-group`
- [查询服务群活跃用户](https://open.dingtalk.com/document/development/queries-active-service-users): `queries-active-service-users`
- [添加服务群成员](https://open.dingtalk.com/document/development/add-service-group-members): `add-service-group-members`
- [群发任务](https://open.dingtalk.com/document/development/service-group-sending-task-interface): `service-group-sending-task-interface`

### Team 空间（7 条）

- [创建协作空间](https://open.dingtalk.com/document/development/api-createprojectv3): `api-createprojectv3`
- [创建协作空间任务](https://open.dingtalk.com/document/development/api-createtask): `api-createtask`
- [创建自由任务](https://open.dingtalk.com/document/development/api-createorganizationtask): `api-createorganizationtask`
- [更新协作空间](https://open.dingtalk.com/document/development/api-updateprojectv3): `api-updateprojectv3`
- [查询任务详情](https://open.dingtalk.com/document/development/api-queryalltask): `api-queryalltask`
- [获取协作空间列表](https://open.dingtalk.com/document/development/api-searchprojectsv3-1): `api-searchprojectsv3-1`
- [获取用户参与项目](https://open.dingtalk.com/document/development/api-getuserjoinedprojectsv3): `api-getuserjoinedprojectsv3`

### 智能招聘（7 条）

- [根据手机号获取候选人信息](https://open.dingtalk.com/document/development/obtain-candidate-information-based-on-mobile-phone-number): `obtain-candidate-information-based-on-mobile-phone-number`
- [添加智能招聘文件到钉盘](https://open.dingtalk.com/document/development/add-nail-disk-file): `add-nail-disk-file`
- [确认完成权益的更新](https://open.dingtalk.com/document/development/confirm-benefits): `confirm-benefits`
- [获取候选人的面试信息](https://open.dingtalk.com/document/development/query-the-interview-list): `query-the-interview-list`
- [获取招聘流程标识](https://open.dingtalk.com/document/development/get-recruitment-process-identity): `get-recruitment-process-identity`
- [获取智能招聘文件上传信息](https://open.dingtalk.com/document/development/obtain-information-about-the-dingtalk-disk-upload-file): `obtain-information-about-the-dingtalk-disk-upload-file`
- [通知完成指定的新手任务](https://open.dingtalk.com/document/development/notify-the-completion-of-the-specified-novice-task): `notify-the-completion-of-the-specified-novice-task`

### 财务（7 条）

- [创建钉工牌电子码](https://open.dingtalk.com/document/development/create-a-user-code-instance): `create-a-user-code-instance`
- [同步钉工牌码验证结果](https://open.dingtalk.com/document/development/sync-pin-badge-code-verification-result): `sync-pin-badge-code-verification-result`
- [更新钉工牌电子码](https://open.dingtalk.com/document/development/update-user-code-instance): `update-user-code-instance`
- [解码钉工牌电子码](https://open.dingtalk.com/document/development/decoding-dingtalk-payment-code): `decoding-dingtalk-payment-code`
- [通知支付结果](https://open.dingtalk.com/document/development/notify-dingtalk-payment-code-payment-result): `notify-dingtalk-payment-code-payment-result`
- [通知退款结果](https://open.dingtalk.com/document/development/dingtalk-payment-code-refund-information-synchronization-operation): `dingtalk-payment-code-refund-information-synchronization-operation`
- [配置企业钉工牌](https://open.dingtalk.com/document/development/set-up-enterprise-payment-code-configuration-interface): `set-up-enterprise-payment-code-configuration-interface`

### 待办（6 条）

- [创建钉钉个人待办任务](https://open.dingtalk.com/document/development/api-createpersonaltodotask): `api-createpersonaltodotask`
- [创建钉钉待办任务](https://open.dingtalk.com/document/development/add-dingtalk-to-do-task): `add-dingtalk-to-do-task`
- [删除钉钉待办任务](https://open.dingtalk.com/document/development/delete-dingtalk-to-do-tasks): `delete-dingtalk-to-do-tasks`
- [更新钉钉待办任务](https://open.dingtalk.com/document/development/updates-dingtalk-to-do-tasks): `updates-dingtalk-to-do-tasks`
- [更新钉钉待办执行者状态](https://open.dingtalk.com/document/development/update-dingtalk-to-do-status): `update-dingtalk-to-do-status`
- [查询企业下用户待办列表](https://open.dingtalk.com/document/development/query-the-to-do-list-of-enterprise-users): `query-the-to-do-list-of-enterprise-users`

### 百科问答（6 条）

- [分页获取企业词条信息](https://open.dingtalk.com/document/development/entry-search): `entry-search`
- [删除词条](https://open.dingtalk.com/document/development/delete-entry): `delete-entry`
- [审核词条](https://open.dingtalk.com/document/development/review-entries): `review-entries`
- [新增词条](https://open.dingtalk.com/document/development/new-entry): `new-entry`
- [更新词条](https://open.dingtalk.com/document/development/update-entry): `update-entry`
- [根据词条ID查询详情](https://open.dingtalk.com/document/development/query-entry): `query-entry`

### 直播（6 条）

- [修改培训课程](https://open.dingtalk.com/document/development/modify-the-basic-information-of-a-live-streaming-course): `modify-the-basic-information-of-a-live-streaming-course`
- [修改直播属性信息](https://open.dingtalk.com/document/development/modify-live-streaming): `modify-live-streaming`
- [创建直播](https://open.dingtalk.com/document/development/create-live-streaming): `create-live-streaming`
- [删除直播](https://open.dingtalk.com/document/development/delete-live-streaming): `delete-live-streaming`
- [查询直播信息](https://open.dingtalk.com/document/development/queries-the-live-streaming-information): `queries-the-live-streaming-information`
- [查询直播观看人员信息](https://open.dingtalk.com/document/development/queries-the-viewing-information-of-viewers): `queries-the-viewing-information-of-viewers`

### 行业（6 条）

- [保存人员扩展属性](https://open.dingtalk.com/document/development/personnel-extension-property-error): `personnel-extension-property-error`
- [创建园区项目](https://open.dingtalk.com/document/development/create-a-campus-project): `create-a-campus-project`
- [创建项目组](https://open.dingtalk.com/document/development/create-a-project-group): `create-a-project-group`
- [删除项目组](https://open.dingtalk.com/document/development/delete-the-project-group-team): `delete-the-project-group-team`
- [查询园区项目信息](https://open.dingtalk.com/document/development/query-a-project-in-a-specified-campus): `query-a-project-in-a-specified-campus`
- [查询项目组信息](https://open.dingtalk.com/document/development/query-a-project-group-in-the-specified-park): `query-a-project-group-in-the-specified-park`

### 居民服务（5 条）

- [分页查询居民积分流水](https://open.dingtalk.com/document/development/query-the-integral-flow-records-by-page): `query-the-integral-flow-records-by-page`
- [增加或减少居民积分](https://open.dingtalk.com/document/development/increase-or-decrease-resident-points): `increase-or-decrease-resident-points`
- [查询组织维度配置的所有积分规则](https://open.dingtalk.com/document/development/query-all-credit-rules): `query-all-credit-rules`
- [获取用户所在的行业角色信息](https://open.dingtalk.com/document/development/obtain-the-industry-role-information-of-the-user): `obtain-the-industry-role-information-of-the-user`
- [获取行业角色下的用户列表](https://open.dingtalk.com/document/development/obtains-a-list-of-users-under-an-industry-role): `obtains-a-list-of-users-under-an-industry-role`

### 组织文化（5 条）

- [创建荣誉勋章模板](https://open.dingtalk.com/document/development/create-medal-of-honor-template): `create-medal-of-honor-template`
- [撤销员工获得的荣誉勋章](https://open.dingtalk.com/document/development/revoke-an-employee-s-medal-of-honor): `revoke-an-employee-s-medal-of-honor`
- [查询员工已获得的组织荣誉](https://open.dingtalk.com/document/development/check-the-honors-that-an-employee-has-received): `check-the-honors-that-an-employee-has-received`
- [查询当前企业下可颁发的荣誉列表](https://open.dingtalk.com/document/development/query-the-list-of-honors-that-can-be-issued-under): `query-the-list-of-honors-that-can-be-issued-under`
- [给员工颁发荣誉](https://open.dingtalk.com/document/development/award-of-honor): `award-of-honor`

### 钉密（5 条）

- [小蜜客服机器人消息回复](https://open.dingtalk.com/document/development/xiaomi-customer-service-robot-message-reply): `xiaomi-customer-service-robot-message-reply`
- [推送小蜜机器人单聊O2O消息](https://open.dingtalk.com/document/development/push-xiaomi-customer-service-robot-single-chat-message): `push-xiaomi-customer-service-robot-single-chat-message`
- [智能问答](https://open.dingtalk.com/document/development/alimebot-intelligent-q-a-interface): `alimebot-intelligent-q-a-interface`
- [查询机器人基础指标数据](https://open.dingtalk.com/document/development/query-robot-data-indicators): `query-robot-data-indicators`
- [获取用户登录凭证](https://open.dingtalk.com/document/development/obtains-the-user-login-credential-of-the-third-party-system-of): `obtains-the-user-login-credential-of-the-third-party-system-of`

### 云盘（4 条）

- [删除空间](https://open.dingtalk.com/document/development/delete-a-space): `delete-a-space`
- [新建空间](https://open.dingtalk.com/document/development/new-space): `new-space`
- [根据spaceId获取指定空间信息](https://open.dingtalk.com/document/development/retrieves-the-space-list-on-the-management-side): `retrieves-the-space-list-on-the-management-side`
- [获取空间列表](https://open.dingtalk.com/document/development/queries-a-space-list): `queries-a-space-list`

### 会话文件（4 条）

- [以应用身份发送文件给指定用户](https://open.dingtalk.com/document/development/sends-a-storage-file-to-a-specified-user): `sends-a-storage-file-to-a-specified-user`
- [发送文件到指定会话](https://open.dingtalk.com/document/development/send-file-to-specified-session): `send-file-to-specified-session`
- [发送文件链接到指定会话](https://open.dingtalk.com/document/development/send-a-file-link-to-the-specified-session): `send-a-file-link-to-the-specified-session`
- [获取群存储空间信息](https://open.dingtalk.com/document/development/obtain-group-storage-space-information): `obtain-group-storage-space-information`

### 应用市场（4 条）

- [内购商品订单处理完成](https://open.dingtalk.com/document/development/internal-purchase-order-processing-completed): `internal-purchase-order-processing-completed`
- [查询应用市场订单详情](https://open.dingtalk.com/document/development/check-the-order-details-app-store): `check-the-order-details-app-store`
- [获取个人应用内购商品SKU页面地址](https://open.dingtalk.com/document/development/obtain-the-sku-page-address-of-goods-purchased-in-personal): `obtain-the-sku-page-address-of-goods-purchased-in-personal`
- [获取未处理的已支付订单](https://open.dingtalk.com/document/development/obtaining-isv-unfinished-processing-order): `obtaining-isv-unfinished-processing-order`

### 公告（3 条）

- [查询公告已读未读人员列表](https://open.dingtalk.com/document/development/query-bulletin-read-unread-persons-list): `query-bulletin-read-unread-persons-list`
- [获取公告详情](https://open.dingtalk.com/document/development/obtains-the-details-get-blackboard): `obtains-the-details-get-blackboard`
- [获取公告钉盘空间信息](https://open.dingtalk.com/document/development/obtain-bulletin-nail-disk-space-information): `obtain-bulletin-nail-disk-space-information`

### 制造（2 条）

- [查询计件报工数据](https://open.dingtalk.com/document/development/riqing-monthly-settlement-query-interface-for-piece-rate-reporting): `riqing-monthly-settlement-query-interface-for-piece-rate-reporting`
- [计件报工](https://open.dingtalk.com/document/development/riqing-monthly-settlement-piece-rate-reporting-interface): `riqing-monthly-settlement-piece-rate-reporting-interface`

### 碳能力（2 条）

- [写入每日用户碳数据明细信息](https://open.dingtalk.com/document/development/write-in-the-detailed-information-of-daily-user-carbon-data): `write-in-the-detailed-information-of-daily-user-carbon-data`
- [写入每日组织碳数据明细信息](https://open.dingtalk.com/document/development/third-party-applications-write-daily-organizational-carbon-data-details-1): `third-party-applications-write-daily-organizational-carbon-data-details-1`

### Agoal（1 条）

- [查询假期余额](https://open.dingtalk.com/document/development/query-holiday-balance): `query-holiday-balance`

### 会员（1 条）

- [查询用户钉钉365会员信息](https://open.dingtalk.com/document/development/api-queryvipmemberinfo): `api-queryvipmemberinfo`

### 硬件（1 条）

- [获取用户考勤数据](https://open.dingtalk.com/document/development/obtain-the-attendance-update-data): `obtain-the-attendance-update-data`

### 职业认证（1 条）

- [检查用户是否完成所有任务](https://open.dingtalk.com/document/development/docking-of-provincial-practical-exercises-for-digital-managers): `docking-of-provincial-practical-exercises-for-digital-managers`

## 连接平台（66 条）

二级 Tab：平台介绍 / 开发指南 / 连接器中心 / 连接平台自动化（公开 URL 待 `doc_url_mapping` 灌入）。

### 连接平台自动化（60 条）

- [下载审批附件](https://open.dingtalk.com/document/development/api-premiumgrantprocessinstancefordownloadfile): `api-premiumgrantprocessinstancefordownloadfile`
- [下载审批附件](https://open.dingtalk.com/document/development/download-an-approval-attachment): `download-an-approval-attachment`
- [保存流程中心外部集成审批任务](https://open.dingtalk.com/document/development/api-premiumsaveintegratedtask): `api-premiumsaveintegratedtask`
- [保存流程中心外部集成审批实例](https://open.dingtalk.com/document/development/api-premiumexternalintegrationprocessinstance): `api-premiumexternalintegrationprocessinstance`
- [保存流程中心外部集成审批模板](https://open.dingtalk.com/document/development/api-premiumsaveintegratedprocess): `api-premiumsaveintegratedprocess`
- [创建实例](https://open.dingtalk.com/document/development/create-a-ticket-approval-instance): `create-a-ticket-approval-instance`
- [创建或更新业务分组](https://open.dingtalk.com/document/development/api-premiuminsertorupdatedir): `api-premiuminsertorupdatedir`
- [创建或更新审批模板](https://open.dingtalk.com/document/development/create-orupdate-the-approval-template-new): `create-orupdate-the-approval-template-new`
- [创建或更新审批表单模板](https://open.dingtalk.com/document/development/create-an-approval-form-template): `create-an-approval-form-template`
- [创建或更新数据表单模板](https://open.dingtalk.com/document/development/api-premiumsaveform): `api-premiumsaveform`
- [创建数据表单实例](https://open.dingtalk.com/document/development/api-createdatapremiumsaveforminstance): `api-createdatapremiumsaveforminstance`
- [创建流程中心待处理任务](https://open.dingtalk.com/document/development/create-pending-tasks-in-process-center): `create-pending-tasks-in-process-center`
- [删除业务分组](https://open.dingtalk.com/document/development/api-premiumdeldir): `api-premiumdeldir`
- [删除数据表单实例](https://open.dingtalk.com/document/development/api-premiumdeleteforminstance): `api-premiumdeleteforminstance`
- [删除模板](https://open.dingtalk.com/document/development/self-owned-approval-deletion-template): `self-owned-approval-deletion-template`
- [加签审批任务](https://open.dingtalk.com/document/development/api-premiumappendtask): `api-premiumappendtask`
- [发起审批实例](https://open.dingtalk.com/document/development/create-an-approval-instance): `create-an-approval-instance`
- [同意或拒绝审批任务](https://open.dingtalk.com/document/development/approve-or-reject-the-approval-task): `approve-or-reject-the-approval-task`
- [归档审批实例](https://open.dingtalk.com/document/development/api-archiveprocessinstance): `api-archiveprocessinstance`
- [批量取消流程中心待处理任务](https://open.dingtalk.com/document/development/cancel-multiple-oa-approval-tasks): `cancel-multiple-oa-approval-tasks`
- [批量同意或拒绝审批任务](https://open.dingtalk.com/document/development/api-premiumbatchexecuteprocessinstances): `api-premiumbatchexecuteprocessinstances`
- [批量更新实例状态](https://open.dingtalk.com/document/development/self-owned-batch-update-of-instance-status): `self-owned-batch-update-of-instance-status`
- [批量获取表单模板schema（包含表单和流程配置信息）](https://open.dingtalk.com/document/development/api-premiumqueryschemaandprocessbycodelist): `api-premiumqueryschemaandprocessbycodelist`
- [授权下载审批钉盘文件](https://open.dingtalk.com/document/development/api-premiumaddapprovedentryauth): `api-premiumaddapprovedentryauth`
- [授权下载审批钉盘文件](https://open.dingtalk.com/document/development/download-the-approval-nail-file): `download-the-approval-nail-file`
- [授权预览审批附件](https://open.dingtalk.com/document/development/api-premiumgetspacewithdownloadauth): `api-premiumgetspacewithdownloadauth`
- [授权预览审批附件](https://open.dingtalk.com/document/development/official-authorized-preview-approval-attachment): `official-authorized-preview-approval-attachment`
- [撤销审批实例](https://open.dingtalk.com/document/development/revoke-an-approval-instance): `revoke-an-approval-instance`
- [更新实例状态](https://open.dingtalk.com/document/development/update-instance-status): `update-instance-status`
- [更新数据表单实例](https://open.dingtalk.com/document/development/api-premiumupdateforminstance): `api-premiumupdateforminstance`
- [更新流程中心任务状态](https://open.dingtalk.com/document/development/update-process-center-task-status): `update-process-center-task-status`
- [更新流程表单审批实例](https://open.dingtalk.com/document/development/api-premiumupdateprocessinstancevariables): `api-premiumupdateprocessinstancevariables`
- [查询审批中心用户已发起实例列表](https://open.dingtalk.com/document/development/api-premiumgetsubmittedinstances): `api-premiumgetsubmittedinstances`
- [查询审批中心用户已处理任务列表](https://open.dingtalk.com/document/development/api-premiumgetdonetasks): `api-premiumgetdonetasks`
- [查询审批中心用户已收到的实例列表](https://open.dingtalk.com/document/development/api-premiumgetnoticedinstances): `api-premiumgetnoticedinstances`
- [查询审批中心用户待处理任务列表](https://open.dingtalk.com/document/development/api-premiumgettodotasks): `api-premiumgettodotasks`
- [查询已设置为条件的表单组件](https://open.dingtalk.com/document/development/query-form-components-that-have-been-set-as-criteria-1): `query-form-components-that-have-been-set-as-criteria-1`
- [查询通过流程中心集成的OA审批任务](https://open.dingtalk.com/document/development/query-oa-approval-tasks-integrated-through-process-center): `query-oa-approval-tasks-integrated-through-process-center`
- [根据processCode分页获取审批流程数据](https://open.dingtalk.com/document/development/api-premiumgetprocessinstances): `api-premiumgetprocessinstances`
- [添加审批评论](https://open.dingtalk.com/document/development/official-approval-adds-approval-comments): `official-approval-adds-approval-comments`
- [清理OA审批数据](https://open.dingtalk.com/document/development/clear-oa-approval-data): `clear-oa-approval-data`
- [管理员批量转交指定员工的待处理任务](https://open.dingtalk.com/document/development/api-premiumredirecttasksbymanager): `api-premiumredirecttasksbymanager`
- [管理员查询指定员工的待处理任务列表](https://open.dingtalk.com/document/development/api-premiumquerytodotasksbymanager): `api-premiumquerytodotasksbymanager`
- [获取单个审批实例详情](https://open.dingtalk.com/document/development/obtains-the-details-of-a-single-approval-instance-pop): `obtains-the-details-of-a-single-approval-instance-pop`
- [获取单个数据表单实例详情](https://open.dingtalk.com/document/development/obtain-details-of-a-single-data-form-instance): `obtain-details-of-a-single-data-form-instance`
- [获取审批单流程中的节点信息](https://open.dingtalk.com/document/development/approval-process-prediction): `approval-process-prediction`
- [获取审批实例ID列表](https://open.dingtalk.com/document/development/obtain-an-approval-list-of-instance-ids): `obtain-an-approval-list-of-instance-ids`
- [获取审批表单控件字段内容修改记录](https://open.dingtalk.com/document/development/api-premiumgetfieldmodifiedhistory): `api-premiumgetfieldmodifiedhistory`
- [获取审批表单控件字段操作权限](https://open.dingtalk.com/document/development/api-premiumgetinstfieldsetting): `api-premiumgetinstfieldsetting`
- [获取审批钉盘空间信息](https://open.dingtalk.com/document/development/api-premiumgetattachmentspace): `api-premiumgetattachmentspace`
- [获取审批钉盘空间信息](https://open.dingtalk.com/document/development/obtains-the-information-about-approval-nail-disk): `obtains-the-information-about-approval-nail-disk`
- [获取当前企业所有可管理的表单](https://open.dingtalk.com/document/development/get-all-manageable-forms-for-the-current-enterprise): `get-all-manageable-forms-for-the-current-enterprise`
- [获取指定用户可见的审批表单列表](https://open.dingtalk.com/document/development/obtains-a-list-of-approval-forms-visible-to-the-specified): `obtains-a-list-of-approval-forms-visible-to-the-specified`
- [获取数据表单schema](https://open.dingtalk.com/document/development/api-premiumgetformschema): `api-premiumgetformschema`
- [获取数据表单实例列表](https://open.dingtalk.com/document/development/api-premiumgetforminstances): `api-premiumgetforminstances`
- [获取模板code](https://open.dingtalk.com/document/development/obtain-the-template-code): `obtain-the-template-code`
- [获取用户待审批数量](https://open.dingtalk.com/document/development/queries-the-number-of-requests-to-be-approved-by-users): `queries-the-number-of-requests-to-be-approved-by-users`
- [获取表单 schema](https://open.dingtalk.com/document/development/obtain-the-form-schema): `obtain-the-form-schema`
- [转交OA审批任务](https://open.dingtalk.com/document/development/transfer-the-oa-approval-task): `transfer-the-oa-approval-task`
- [退回审批任务](https://open.dingtalk.com/document/development/api-premiumreverttask): `api-premiumreverttask`

### 企业互通（6 条）

- [批量获取关注服务窗用户信息](https://open.dingtalk.com/document/development/obtains-the-follower-information-from-the-service-window): `obtains-the-follower-information-from-the-service-window`
- [获取企业下服务窗列表](https://open.dingtalk.com/document/development/queries-the-list-of-services-under-an-enterprise): `queries-the-list-of-services-under-an-enterprise`
- [获取关注服务窗用户信息](https://open.dingtalk.com/document/development/queries-the-follower-information-of-the-service-window): `queries-the-follower-information-of-the-service-window`
- [获取用户关注服务窗状态](https://open.dingtalk.com/document/development/third-party-enterprise-application-obtains-user-attention-service-window-status): `third-party-enterprise-application-obtains-user-attention-service-window-status`
- [获取用户服务窗关注状态](https://open.dingtalk.com/document/development/obtain-the-attention-status-of-the-user-service-window): `obtain-the-attention-status-of-the-user-service-window`
- [获取组织服务窗账号列表](https://open.dingtalk.com/document/development/the-third-party-enterprise-application-obtains-the-account-list-of-the): `the-third-party-enterprise-application-obtains-the-account-list-of-the`

## AI PaaS（48 条）

二级 Tab：平台介绍 / 炼丹炉大模型平台 / AI 助理创建平台 / AI 客服助理（公开 URL 待 `doc_url_mapping` 灌入）。

### 组织大脑（41 条）

- [人员标签数据查询](https://open.dingtalk.com/document/development/api-stafflabelrecordsquery): `api-stafflabelrecordsquery`
- [人才档案基础数据查询](https://open.dingtalk.com/document/development/api-hrbraintalentprofilebasicquery): `api-hrbraintalentprofilebasicquery`
- [人才档案照片查询](https://open.dingtalk.com/document/development/api-hrbraintalentprofileattachmentquery): `api-hrbraintalentprofileattachmentquery`
- [人才池信息查询](https://open.dingtalk.com/document/development/api-hrbrainemppoolquery): `api-hrbrainemppoolquery`
- [人才池在池人员列表](https://open.dingtalk.com/document/development/api-hrbrainemppooluser): `api-hrbrainemppooluser`
- [数据集成专业技能删除](https://open.dingtalk.com/document/development/api-hrbraindeletelabelprofskill): `api-hrbraindeletelabelprofskill`
- [数据集成专业技能同步](https://open.dingtalk.com/document/development/api-hrbrainimportlabelprofskill): `api-hrbrainimportlabelprofskill`
- [数据集成人员信息删除](https://open.dingtalk.com/document/development/api-hrbraindeleteempinfo): `api-hrbraindeleteempinfo`
- [数据集成人员信息同步](https://open.dingtalk.com/document/development/api-hrbrainimportempinfo): `api-hrbrainimportempinfo`
- [数据集成人员标签删除](https://open.dingtalk.com/document/development/api-hrbraindeletetlabelbase): `api-hrbraindeletetlabelbase`
- [数据集成入职信息同步](https://open.dingtalk.com/document/development/api-hrbrainimportregist): `api-hrbrainimportregist`
- [数据集成入职记录删除](https://open.dingtalk.com/document/development/api-hrbraindeleteregist): `api-hrbraindeleteregist`
- [数据集成删除自定义模型数据](https://open.dingtalk.com/document/development/api-hrbraindeletecustom): `api-hrbraindeletecustom`
- [数据集成培训学习数据删除](https://open.dingtalk.com/document/development/api-hrbraindeletetraining): `api-hrbraindeletetraining`
- [数据集成培训学习记录同步](https://open.dingtalk.com/document/development/api-hrbrainimporttraining): `api-hrbrainimporttraining`
- [数据集成基础标签同步](https://open.dingtalk.com/document/development/api-hrbrainimportlabelbase): `api-hrbrainimportlabelbase`
- [数据集成处分记录删除](https://open.dingtalk.com/document/development/api-hrbraindeletepundetail): `api-hrbraindeletepundetail`
- [数据集成处分记录同步](https://open.dingtalk.com/document/development/api-hrbrainimportpundetail): `api-hrbrainimportpundetail`
- [数据集成奖励信息删除](https://open.dingtalk.com/document/development/api-hrbraindeleteawardrecords): `api-hrbraindeleteawardrecords`
- [数据集成奖励记录同步](https://open.dingtalk.com/document/development/api-hrbrainimportawarddetail): `api-hrbrainimportawarddetail`
- [数据集成工作经历删除](https://open.dingtalk.com/document/development/api-hrbraindeleteworkexp): `api-hrbraindeleteworkexp`
- [数据集成工作经历同步](https://open.dingtalk.com/document/development/api-hrbrainimportworkexp): `api-hrbrainimportworkexp`
- [数据集成异动记录同步](https://open.dingtalk.com/document/development/api-hrbrainimporttransfereval): `api-hrbrainimporttransfereval`
- [数据集成教育经历删除](https://open.dingtalk.com/document/development/api-hrbraindeleteeduexp): `api-hrbraindeleteeduexp`
- [数据集成教育经历同步](https://open.dingtalk.com/document/development/api-hrbrainimporteduexp): `api-hrbrainimporteduexp`
- [数据集成晋升记录同步](https://open.dingtalk.com/document/development/api-hrbrainimportpromeval): `api-hrbrainimportpromeval`
- [数据集成盘点数据删除](https://open.dingtalk.com/document/development/api-hrbraindeletelabelinventory): `api-hrbraindeletelabelinventory`
- [数据集成盘点数据同步](https://open.dingtalk.com/document/development/api-hrbrainimportlabelinventory): `api-hrbrainimportlabelinventory`
- [数据集成离职信息同步](https://open.dingtalk.com/document/development/api-hrbrainimportdimission): `api-hrbrainimportdimission`
- [数据集成离职记录删除](https://open.dingtalk.com/document/development/api-hrbraindeletedimission): `api-hrbraindeletedimission`
- [数据集成组织架构同步](https://open.dingtalk.com/document/development/api-hrbrainimportdeptinfo): `api-hrbrainimportdeptinfo`
- [数据集成组织架构数据删除](https://open.dingtalk.com/document/development/api-hrbraindeletedeptinfo): `api-hrbraindeletedeptinfo`
- [数据集成绩效记录删除](https://open.dingtalk.com/document/development/api-hrbraindeleteperfeval): `api-hrbraindeleteperfeval`
- [数据集成绩效记录同步](https://open.dingtalk.com/document/development/api-hrbrainimportperfeval): `api-hrbrainimportperfeval`
- [数据集成自定义标签同步](https://open.dingtalk.com/document/development/api-hrbrainimportlabelcustom): `api-hrbrainimportlabelcustom`
- [数据集成调岗记录删除](https://open.dingtalk.com/document/development/api-hrbraindeletetransfereval): `api-hrbraindeletetransfereval`
- [数据集成转正数据删除](https://open.dingtalk.com/document/development/api-hrbraindeleteregular): `api-hrbraindeleteregular`
- [数据集成转正记录同步](https://open.dingtalk.com/document/development/api-hrbrainimportregular): `api-hrbrainimportregular`
- [数据集成领域经验删除](https://open.dingtalk.com/document/development/api-hrbraindeletelabelindustry): `api-hrbraindeletelabelindustry`
- [数据集成领域经验同步](https://open.dingtalk.com/document/development/api-hrbrainimportlabelindustry): `api-hrbrainimportlabelindustry`
- [自定义模型数据同步](https://open.dingtalk.com/document/development/api-hrbrainimportcustom): `api-hrbrainimportcustom`

### 智能客服（4 条）

- [分页查询工单](https://open.dingtalk.com/document/development/intelligent-customer-service-paging-query-work-order): `intelligent-customer-service-paging-query-work-order`
- [创建自助单](https://open.dingtalk.com/document/development/create-a-self-service-ticket): `create-a-self-service-ticket`
- [执行工单活动](https://open.dingtalk.com/document/development/intelligent-customer-service-execute-work-order-activities): `intelligent-customer-service-execute-work-order-activities`
- [查询动作记录](https://open.dingtalk.com/document/development/intelligent-customer-service-query-action-records): `intelligent-customer-service-query-action-records`

### AI PaaS（3 条）

- [大模型推理服务（多模态模型）](https://open.dingtalk.com/document/development/large-model-reasoning-service-interface-multimodal-model-1): `large-model-reasoning-service-interface-multimodal-model-1`
- [大模型推理服务（文生文模型）](https://open.dingtalk.com/document/development/api-exclusivemodelcompleteservice): `api-exclusivemodelcompleteservice`
- [炼丹炉专属模型服务](https://open.dingtalk.com/document/development/api-liandanluexclusivemodel): `api-liandanluexclusivemodel`

## 硬件开发（28 条）

二级 Tab：智能硬件（公开 URL 待 `doc_url_mapping` 灌入）。

### 智能会议室（22 条）

- [创建会议室](https://open.dingtalk.com/document/development/create-meeting-room): `create-meeting-room`
- [创建会议室分组](https://open.dingtalk.com/document/development/create-meeting-room-groups): `create-meeting-room-groups`
- [创建会议室预定黑名单](https://open.dingtalk.com/document/development/api-createbookingblacklist): `api-createbookingblacklist`
- [创建自定义屏幕模板](https://open.dingtalk.com/document/development/api-createdevicecustomtemplate): `api-createdevicecustomtemplate`
- [删除会议室](https://open.dingtalk.com/document/development/delete-a-meeting-room): `delete-a-meeting-room`
- [删除会议室分组](https://open.dingtalk.com/document/development/delete-a-conference-room-group): `delete-a-conference-room-group`
- [删除会议室预定黑名单](https://open.dingtalk.com/document/development/api-deletebookingblacklist): `api-deletebookingblacklist`
- [删除自定义屏幕模板](https://open.dingtalk.com/document/development/api-deletedevicecustomtemplate): `api-deletedevicecustomtemplate`
- [发送Rooms中控API信令](https://open.dingtalk.com/document/development/api-sendcentralcontrol): `api-sendcentralcontrol`
- [更新会议室信息](https://open.dingtalk.com/document/development/update-meeting-room-information): `update-meeting-room-information`
- [更新会议室分组信息](https://open.dingtalk.com/document/development/update-meeting-room-groups): `update-meeting-room-groups`
- [更新自定义屏幕模板](https://open.dingtalk.com/document/development/api-updatedevicecustomtemplate): `api-updatedevicecustomtemplate`
- [查询会议室分组信息](https://open.dingtalk.com/document/development/query-meeting-room-groups): `query-meeting-room-groups`
- [查询会议室分组列表](https://open.dingtalk.com/document/development/query-meeting-rooms-groups): `query-meeting-rooms-groups`
- [查询会议室列表](https://open.dingtalk.com/document/development/check-the-meeting-room-list): `check-the-meeting-room-list`
- [查询会议室详情](https://open.dingtalk.com/document/development/check-meeting-room-details): `check-meeting-room-details`
- [查询自定义屏幕信息](https://open.dingtalk.com/document/development/api-querydevicecustomtemplate): `api-querydevicecustomtemplate`
- [查询自定义屏幕模板列表](https://open.dingtalk.com/document/development/api-querylistdevicecustomscreentemplate): `api-querylistdevicecustomscreentemplate`
- [查询视频会议设备信息](https://open.dingtalk.com/document/development/querying-video-conference-device-information): `querying-video-conference-device-information`
- [查询视频会议设备属性信息](https://open.dingtalk.com/document/development/querying-video-conference-device-attribute-information): `querying-video-conference-device-attribute-information`
- [获取用户发出的日志列表](https://open.dingtalk.com/document/development/query-logs-sent-by-an-employee): `query-logs-sent-by-an-employee`
- [获取用户可见的日志模板](https://open.dingtalk.com/document/development/obtains-the-list-of-visible-log-templates-based-on-the): `obtains-the-list-of-visible-log-templates-based-on-the`

### 设备管理（5 条）

- [批量注册与激活设备](https://open.dingtalk.com/document/development/register-and-activate-devices-in-batches): `register-and-activate-devices-in-batches`
- [查询已经注册的设备信息](https://open.dingtalk.com/document/development/query-information-about-a-registered-device): `query-information-about-a-registered-device`
- [注册设备到钉钉](https://open.dingtalk.com/document/development/pin-registration-interface): `pin-registration-interface`
- [获取巡检或保养记录](https://open.dingtalk.com/document/development/obtain-inspection-and-maintenance-records): `obtain-inspection-and-maintenance-records`
- [获取报修记录](https://open.dingtalk.com/document/development/obtain-the-repair-report-record): `obtain-the-repair-report-record`

### 智能设备（1 条）

- [变更智能考勤机员工](https://open.dingtalk.com/document/development/change-intelligent-attendance-machine-staff): `change-intelligent-attendance-machine-staff`

## 互动卡片（9 条）

二级 Tab：开发指南 / 使用教程 / 卡片模板搭建器 / 互动卡片搭建平台 / 卡片规范设计（公开 URL 待 `doc_url_mapping` 灌入）。

### 互动卡片（9 条）

- [AI卡片流式更新](https://open.dingtalk.com/document/development/api-streamingupdate): `api-streamingupdate`
- [关闭吊顶卡片](https://open.dingtalk.com/document/development/api-closetopcard): `api-closetopcard`
- [创建卡片](https://open.dingtalk.com/document/development/interface-for-creating-a-card-instance): `interface-for-creating-a-card-instance`
- [创建并投放卡片](https://open.dingtalk.com/document/development/create-and-deliver-cards): `create-and-deliver-cards`
- [卡片平台模板复制](https://open.dingtalk.com/document/development/api-copytemplate): `api-copytemplate`
- [投放卡片](https://open.dingtalk.com/document/development/delivery-card-interface): `delivery-card-interface`
- [新增或者更新卡片的场域信息](https://open.dingtalk.com/document/development/add-field-interface): `add-field-interface`
- [更新卡片](https://open.dingtalk.com/document/development/interactive-card-update-interface): `interactive-card-update-interface`
- [注册卡片回调地址](https://open.dingtalk.com/document/development/register-card-callback-address): `register-card-callback-address`

## 专属版客户端插件（37 条）

二级 Tab：功能介绍 / 插件开发（公开 URL 待 `doc_url_mapping` 灌入）。

### 专属版（37 条）

- [DING服务](https://open.dingtalk.com/document/development/send-in-application-ding): `send-in-application-ding`
- [专属小红点推送](https://open.dingtalk.com/document/development/push-a-red-dot-to-the-micro-application): `push-a-red-dot-to-the-micro-application`
- [企业内部群禁言或解除禁言](https://open.dingtalk.com/document/development/exclusive-dingtalk-group-ban): `exclusive-dingtalk-group-ban`
- [企业员工专属安全管控功能命中查询](https://open.dingtalk.com/document/development/api-checkcontrolhitstatus): `api-checkcontrolhitstatus`
- [修改伙伴类型可见性](https://open.dingtalk.com/document/development/modify-partner-type-visibility): `modify-partner-type-visibility`
- [修改角色可见性](https://open.dingtalk.com/document/development/modify-role-visibility): `modify-role-visibility`
- [删除可信设备](https://open.dingtalk.com/document/development/delete-trusted-devices): `delete-trusted-devices`
- [发送文件更改的评论](https://open.dingtalk.com/document/development/send-comments-on-file-changes): `send-comments-on-file-changes`
- [发送电话DING](https://open.dingtalk.com/document/development/outgoing-phone-ding): `outgoing-phone-ding`
- [发送邀请函](https://open.dingtalk.com/document/development/send-invitations): `send-invitations`
- [同步存储数据](https://open.dingtalk.com/document/development/api-datasync): `api-datasync`
- [批量新增可信设备](https://open.dingtalk.com/document/development/create-multiple-trusted-devices): `create-multiple-trusted-devices`
- [新增可信设备信息](https://open.dingtalk.com/document/development/add-information-about-a-trusted-device): `add-information-about-a-trusted-device`
- [更新发送文件的检测状态](https://open.dingtalk.com/document/development/update-the-detection-status-of-a-sent-file): `update-the-detection-status-of-a-sent-file`
- [查询人脸录入状态](https://open.dingtalk.com/document/development/query-face-entry-status): `query-face-entry-status`
- [查询企业内部群信息](https://open.dingtalk.com/document/development/obtain-group-info): `obtain-group-info`
- [查询伙伴角色列表](https://open.dingtalk.com/document/development/query-the-list-of-partners): `query-the-list-of-partners`
- [查询公共设备](https://open.dingtalk.com/document/development/query-public-equipment): `query-public-equipment`
- [查询可信设备详细信息](https://open.dingtalk.com/document/development/query-trusted-device-details): `query-trusted-device-details`
- [查询实人认证状态](https://open.dingtalk.com/document/development/queries-the-id-verification-status): `queries-the-id-verification-status`
- [查询群发消息列表](https://open.dingtalk.com/document/development/service-account-query-msgsend-records): `service-account-query-msgsend-records`
- [查询群发消息详情](https://open.dingtalk.com/document/development/service-account-msg-record-detail): `service-account-msg-record-detail`
- [根据userId查询人员的标签信息](https://open.dingtalk.com/document/development/you-can-call-this-operation-to-retrieve-the-user-tag): `you-can-call-this-operation-to-retrieve-the-user-tag`
- [根据会议逻辑ID查询会议基本信息](https://open.dingtalk.com/document/development/query-basic-meeting-information-using-a-logical-id): `query-basic-meeting-information-using-a-logical-id`
- [消息群发](https://open.dingtalk.com/document/development/api-sendmessage): `api-sendmessage`
- [获取专属存储文件路径](https://open.dingtalk.com/document/development/api-getprivatestorefilepath): `api-getprivatestorefilepath`
- [获取人脸对比接口调用记录](https://open.dingtalk.com/document/development/you-can-call-this-operation-to-query-the-call-records): `you-can-call-this-operation-to-query-the-call-records`
- [获取企业专属钉钉权益列表](https://open.dingtalk.com/document/development/api-queryexclusivebenefits): `api-queryexclusivebenefits`
- [获取可打标部门列表](https://open.dingtalk.com/document/development/obtains-a-list-of-departments-that-can-be-marked): `obtains-a-list-of-departments-that-can-be-marked`
- [获取子标签列表](https://open.dingtalk.com/document/development/obtain-child-tags-from-a-parent-tag): `obtain-child-tags-from-a-parent-tag`
- [获取实人认证接口调用记录](https://open.dingtalk.com/document/development/obtains-the-call-record-of-the-id-authentication-api): `obtains-the-call-record-of-the-id-authentication-api`
- [获取审计协议签署人员信息](https://open.dingtalk.com/document/development/obtains-the-information-about-the-persons-who-sign-the-audit-1): `obtains-the-information-about-the-persons-who-sign-the-audit-1`
- [获取文件操作记录](https://open.dingtalk.com/document/development/obtain-file-operation-records): `obtain-file-operation-records`
- [获取群活跃明细列表](https://open.dingtalk.com/document/development/obtains-the-group-activity-details-list): `obtains-the-group-activity-details-list`
- [获取视频会议详情](https://open.dingtalk.com/document/development/get-video-meeting-details): `get-video-meeting-details`
- [获取防截屏操作记录](https://open.dingtalk.com/document/development/obtain-anti-screen-capture-operation-records): `obtain-anti-screen-capture-operation-records`
- [设置部门伙伴类型和伙伴编码](https://open.dingtalk.com/document/development/set-department-partner-type-and-partner-code): `set-department-partner-type-and-partner-code`

## 数据资产（102 条）

二级 Tab：平台介绍 / 宜数（智能问数）（公开 URL 待 `doc_url_mapping` 灌入）。

### 宜搭（61 条）

- [保存表单数据](https://open.dingtalk.com/document/development/api-saveformdata-v2): `api-saveformdata-v2`
- [分页获取集成自动化日志列表](https://open.dingtalk.com/document/development/api-pageautoflowlog): `api-pageautoflowlog`
- [删除流程实例](https://open.dingtalk.com/document/development/delete-the-process-instance): `delete-the-process-instance`
- [删除表单数据](https://open.dingtalk.com/document/development/delete-form-data): `delete-form-data`
- [发起宜搭审批流程](https://open.dingtalk.com/document/development/api-startinstance-v2): `api-startinstance-v2`
- [同意或拒绝宜搭审批任务](https://open.dingtalk.com/document/development/execute-approval-tasks): `execute-approval-tasks`
- [应用授权校验](https://open.dingtalk.com/document/development/application-authorization-verification): `application-authorization-verification`
- [批量创建表单实例](https://open.dingtalk.com/document/development/create-multiple-form-instances): `create-multiple-form-instances`
- [批量删除宜搭角色成员](https://open.dingtalk.com/document/development/batch-deleterolemembers): `batch-deleterolemembers`
- [批量删除指定矩阵的明细数据](https://open.dingtalk.com/document/development/api-deletematrixdatabyrowids): `api-deletematrixdatabyrowids`
- [批量删除表单实例](https://open.dingtalk.com/document/development/delete-multiple-form-instances): `delete-multiple-form-instances`
- [批量执行宜搭审批任务](https://open.dingtalk.com/document/development/batch-execution-should-take-the-lead-of-approval-tasks): `batch-execution-should-take-the-lead-of-approval-tasks`
- [批量更新宜搭角色成员](https://open.dingtalk.com/document/development/batch-rolemembers): `batch-rolemembers`
- [批量更新表单实例内的组件值](https://open.dingtalk.com/document/development/batch-update-of-component-values-in-form-instances): `batch-update-of-component-values-in-form-instances`
- [批量查询宜搭表单实例的评论](https://open.dingtalk.com/document/development/batch-query-of-comments-appropriate-for-form-instances): `batch-query-of-comments-appropriate-for-form-instances`
- [批量获取流程实例列表](https://open.dingtalk.com/document/development/queries-multiple-process-instances): `queries-multiple-process-instances`
- [批量获取表单实例数据](https://open.dingtalk.com/document/development/obtain-multiple-form-instance-data): `obtain-multiple-form-instance-data`
- [提交评论](https://open.dingtalk.com/document/development/submit-comment): `submit-comment`
- [新增或更新表单实例](https://open.dingtalk.com/document/development/api-createorupdateformdata-v2): `api-createorupdateformdata-v2`
- [新购授权订单](https://open.dingtalk.com/document/development/new-order-authorization): `new-order-authorization`
- [更新宜搭子表单数据](https://open.dingtalk.com/document/development/api-updatesubtable): `api-updatesubtable`
- [更新指定矩阵的明细数据](https://open.dingtalk.com/document/development/api-saveandupdatematrixdata): `api-saveandupdatematrixdata`
- [更新表单数据](https://open.dingtalk.com/document/development/api-updateformdata-v2): `api-updateformdata-v2`
- [查询宜搭应用列表](https://open.dingtalk.com/document/development/query-the-application-list): `query-the-application-list`
- [查询宜搭表单服务调用执行记录](https://open.dingtalk.com/document/development/the-query-should-be-based-on-the-execution-records-of): `the-query-should-be-based-on-the-execution-records-of`
- [查询抄送我的任务列表（应用维度）](https://open.dingtalk.com/document/development/query-copied-my-task-list-application-dimension): `query-copied-my-task-list-application-dimension`
- [查询流程运行任务（VPC）](https://open.dingtalk.com/document/development/query-process-running-tasks-vpc): `query-process-running-tasks-vpc`
- [查询用户详情](https://open.dingtalk.com/document/development/query-user-details): `query-user-details`
- [查询表单实例数据](https://open.dingtalk.com/document/development/api-searchformdatas-v2): `api-searchformdatas-v2`
- [查询表单数据](https://open.dingtalk.com/document/development/api-getformdatabyid-v2): `api-getformdatabyid-v2`
- [查询表单的变更记录](https://open.dingtalk.com/document/development/change-records-of-query-forms): `change-records-of-query-forms`
- [校验订单](https://open.dingtalk.com/document/development/verification-order): `verification-order`
- [根据流程实例ID获取流程实例](https://open.dingtalk.com/document/development/api-getinstancebyid-v2): `api-getinstancebyid-v2`
- [终止流程实例](https://open.dingtalk.com/document/development/terminate-a-process-instance): `terminate-a-process-instance`
- [获取任务列表（组织维度）](https://open.dingtalk.com/document/development/query-tasks-from-the-organization-dimension): `query-tasks-from-the-organization-dimension`
- [获取发送给用户的通知](https://open.dingtalk.com/document/development/get-notifications-sent-to-users): `get-notifications-sent-to-users`
- [获取员工组件的值](https://open.dingtalk.com/document/development/gets-the-value-of-the-employee-component): `gets-the-value-of-the-employee-component`
- [获取多个表单实例ID](https://open.dingtalk.com/document/development/api-searchformdataidlist-v2): `api-searchformdataidlist-v2`
- [获取子表组件数据](https://open.dingtalk.com/document/development/obtain-child-table-component-data): `obtain-child-table-component-data`
- [获取宜搭附件临时免登地址](https://open.dingtalk.com/document/development/obtain-the-temporary-free-access-address-of-yixian-accessories): `obtain-the-temporary-free-access-address-of-yixian-accessories`
- [获取实例ID列表](https://open.dingtalk.com/document/development/api-getinstanceidlist-v2): `api-getinstanceidlist-v2`
- [获取审批记录](https://open.dingtalk.com/document/development/queries-an-approval-record): `queries-an-approval-record`
- [获取指定宜搭角色的角色详情](https://open.dingtalk.com/document/development/get-roledetailbyid): `get-roledetailbyid`
- [获取指定应用下的表单列表](https://open.dingtalk.com/document/development/depending-on-the-application-id-to-get-the-form-list): `depending-on-the-application-id-to-get-the-form-list`
- [获取指定权限矩阵的明细数据](https://open.dingtalk.com/document/development/api-getmatrixdetailbyid): `api-getmatrixdetailbyid`
- [获取流程实例](https://open.dingtalk.com/document/development/api-getinstances-v2): `api-getinstances-v2`
- [获取流程设计结构](https://open.dingtalk.com/document/development/api-getprocessdesign): `api-getprocessdesign`
- [获取组件别名列表](https://open.dingtalk.com/document/development/api-getformcomponentaliaslist): `api-getformcomponentaliaslist`
- [获取组织内已完成的审批任务](https://open.dingtalk.com/document/development/obtains-the-completed-approval-tasks-in-an-organization): `obtains-the-completed-approval-tasks-in-an-organization`
- [获取组织内某人提交的任务](https://open.dingtalk.com/document/development/obtains-the-tasks-submitted-by-someone-in-an-organization): `obtains-the-tasks-submitted-by-someone-in-an-organization`
- [获取表单内的组件信息](https://open.dingtalk.com/document/development/get-form-field-information-based-on-form-uuid): `get-form-field-information-based-on-form-uuid`
- [获取表单组件定义列表](https://open.dingtalk.com/document/development/get-a-list-of-form-component-definitions): `get-a-list-of-form-component-definitions`
- [获取集成自动化日志详情](https://open.dingtalk.com/document/development/api-getautoflowlogdetail): `api-getautoflowlogdetail`
- [转交任务](https://open.dingtalk.com/document/development/transfer-tasks): `transfer-tasks`
- [退还商品](https://open.dingtalk.com/document/development/refund-of-goods): `refund-of-goods`
- [通过RowId更新子表单数据](https://open.dingtalk.com/document/development/update-the-subform-data-via-rowid): `update-the-subform-data-via-rowid`
- [通过流程code获取流程定义](https://open.dingtalk.com/document/development/obtain-definition-through-process-code): `obtain-definition-through-process-code`
- [通过表单实例数据批量更新表单实例](https://open.dingtalk.com/document/development/update-multiple-form-instances-with-the-form-instance-data): `update-multiple-form-instances-with-the-form-instance-data`
- [通过高级查询条件获取表单实例数据（不包括子表单组件数据）](https://open.dingtalk.com/document/development/obtain-form-instance-data-using-advanced-query-conditions-excluding-subform): `obtain-form-instance-data-using-advanced-query-conditions-excluding-subform`
- [通过高级查询条件获取表单实例数据（包括子表单组件数据）](https://open.dingtalk.com/document/development/api-searchformdatasecondgeneration-v2): `api-searchformdatasecondgeneration-v2`
- [预览审批流程](https://open.dingtalk.com/document/development/api-previewpublishedprocess): `api-previewpublishedprocess`

### 氚云（20 条）

- [修改表单业务对象数据](https://open.dingtalk.com/document/development/modify-form-business-object-data): `modify-form-business-object-data`
- [创建流程实例](https://open.dingtalk.com/document/development/create-a-process-instance): `create-a-process-instance`
- [创建表单业务数据](https://open.dingtalk.com/document/development/create-form-business-data): `create-form-business-data`
- [删除业务对象](https://open.dingtalk.com/document/development/delete-a-business-object): `delete-a-business-object`
- [删除流程实例数据](https://open.dingtalk.com/document/development/delete-process-instance-data): `delete-process-instance-data`
- [取消流程实例](https://open.dingtalk.com/document/development/cancel-a-process-instance): `cancel-a-process-instance`
- [批量新增表单业务数据](https://open.dingtalk.com/document/development/batch-add-form-business-data): `batch-add-form-business-data`
- [查询流程实例](https://open.dingtalk.com/document/development/query-flow-instances): `query-flow-instances`
- [查询流程实例节点工作项](https://open.dingtalk.com/document/development/query-flow-instance-node-work-items): `query-flow-instance-node-work-items`
- [查询表单业务数据列表](https://open.dingtalk.com/document/development/querying-form-business-data): `querying-form-business-data`
- [获取业务实例信息](https://open.dingtalk.com/document/development/queries-business-instance-information): `queries-business-instance-information`
- [获取应用列表](https://open.dingtalk.com/document/development/queries-applications): `queries-applications`
- [获取应用功能节点](https://open.dingtalk.com/document/development/queries-the-application-feature-nodes): `queries-the-application-feature-nodes`
- [获取文件上传地址](https://open.dingtalk.com/document/development/obtain-the-upload-url-of-a-file-2): `obtain-the-upload-url-of-a-file-2`
- [获取用户数据](https://open.dingtalk.com/document/development/obtain-user-data): `obtain-user-data`
- [获取组织数据](https://open.dingtalk.com/document/development/queries-organization-data): `queries-organization-data`
- [获取表单对象结构](https://open.dingtalk.com/document/development/gets-the-form-object-structure): `gets-the-form-object-structure`
- [获取角色数据](https://open.dingtalk.com/document/development/obtain-role-data): `obtain-role-data`
- [获取角色用户数据](https://open.dingtalk.com/document/development/historical-acquisition-of-role-user-data): `historical-acquisition-of-role-user-data`
- [获取附件临时免登地址](https://open.dingtalk.com/document/development/obtain-the-temporary-attachment-free-address): `obtain-the-temporary-attachment-free-address`

### 多维表（14 条）

- [列出多行记录](https://open.dingtalk.com/document/development/api-notable-listrecords): `api-notable-listrecords`
- [创建字段](https://open.dingtalk.com/document/development/api-noatable-createfield): `api-noatable-createfield`
- [创建数据表](https://open.dingtalk.com/document/development/api-createsheet): `api-createsheet`
- [删除多行记录](https://open.dingtalk.com/document/development/api-noatable-deleterecords): `api-noatable-deleterecords`
- [删除字段](https://open.dingtalk.com/document/development/api-noatable-deletefield): `api-noatable-deletefield`
- [删除数据表](https://open.dingtalk.com/document/development/api-noatable-deletesheet): `api-noatable-deletesheet`
- [新增记录](https://open.dingtalk.com/document/development/api-notable-insertrecords): `api-notable-insertrecords`
- [更新多行记录](https://open.dingtalk.com/document/development/api-noatable-updaterecords): `api-noatable-updaterecords`
- [更新字段](https://open.dingtalk.com/document/development/api-noatable-updatefield): `api-noatable-updatefield`
- [更新数据表](https://open.dingtalk.com/document/development/api-noatable-updatesheet): `api-noatable-updatesheet`
- [获取所有字段](https://open.dingtalk.com/document/development/api-noatable-getallfields): `api-noatable-getallfields`
- [获取所有数据表](https://open.dingtalk.com/document/development/api-notable-getallsheets): `api-notable-getallsheets`
- [获取数据表](https://open.dingtalk.com/document/development/api-notable-getsheet): `api-notable-getsheet`
- [获取记录](https://open.dingtalk.com/document/development/api-getrecord): `api-getrecord`

### 数据中台 AMDP（4 条）

- [组织变革主数据人员任职数据推送](https://open.dingtalk.com/document/development/api-amdpjobpositiondatapush): `api-amdpjobpositiondatapush`
- [组织变革主数据人员数据推送](https://open.dingtalk.com/document/development/api-amdpemployeedatapush): `api-amdpemployeedatapush`
- [组织变革主数据人员角色数据推送](https://open.dingtalk.com/document/development/api-amdpemproledatapush): `api-amdpemproledatapush`
- [组织变革主数据部门数据推送](https://open.dingtalk.com/document/development/api-amdporganizationdatapush): `api-amdporganizationdatapush`

### 表单（3 条）

- [获取单条填表实例详情](https://open.dingtalk.com/document/development/obtains-the-instance-details-of-a-single-fill-table): `obtains-the-instance-details-of-a-single-fill-table`
- [获取填表实例列表](https://open.dingtalk.com/document/development/obtain-the-table-filling-instance-list-data): `obtain-the-table-filling-instance-list-data`
- [获取用户创建的填表模板列表](https://open.dingtalk.com/document/development/new-obtains-the-template-that-a-user-creates): `new-obtains-the-template-that-a-user-creates`

## 工作台（24 条）

二级 Tab：平台介绍 / 使用教程（公开 URL 待 `doc_url_mapping` 灌入）。

### 微应用（13 条）

- [创建企业内部应用](https://open.dingtalk.com/document/development/create-an-h5-application-for-your-enterprise): `create-an-h5-application-for-your-enterprise`
- [删除企业内部应用](https://open.dingtalk.com/document/development/delete-an-internal-h5-application): `delete-an-internal-h5-application`
- [发布企业内部小程序版本](https://open.dingtalk.com/document/development/release-internal-applet-version): `release-internal-applet-version`
- [回滚企业内部小程序版本](https://open.dingtalk.com/document/development/rollback-of-enterprise-internal-applet-version): `rollback-of-enterprise-internal-applet-version`
- [更新企业内部应用](https://open.dingtalk.com/document/development/update-internal-h5-applications): `update-internal-h5-applications`
- [更新企业内部应用的可使用范围](https://open.dingtalk.com/document/development/update-the-visible-range-of-micro-applications): `update-the-visible-range-of-micro-applications`
- [查询管理员是否有应用管理权限](https://open.dingtalk.com/document/development/check-whether-the-administrator-has-application-management-permissions): `check-whether-the-administrator-has-application-management-permissions`
- [获取企业内部小程序历史版本列表](https://open.dingtalk.com/document/development/obtain-the-list-of-historical-versions-of-enterprise-internal-applets): `obtain-the-list-of-historical-versions-of-enterprise-internal-applets`
- [获取企业内部小程序的版本列表](https://open.dingtalk.com/document/development/get-the-version-list-of-the-enterprise-internal-applet): `get-the-version-list-of-the-enterprise-internal-applet`
- [获取企业内部应用的可使用范围](https://open.dingtalk.com/document/development/obtains-the-application-visible-range): `obtains-the-application-visible-range`
- [获取企业内部所有应用列表](https://open.dingtalk.com/document/development/get-a-list-of-all-applications-inside-the-enterprise): `get-a-list-of-all-applications-inside-the-enterprise`
- [获取企业所有应用列表](https://open.dingtalk.com/document/development/obtains-a-list-of-all-enterprise-applications): `obtains-a-list-of-all-enterprise-applications`
- [获取用户可见的企业应用列表](https://open.dingtalk.com/document/development/obtains-the-list-of-enterprise-applications-visible-to-a-user): `obtains-the-list-of-enterprise-applications-visible-to-a-user`

### 应用角标（8 条）

- [创建钉工牌电子码](https://open.dingtalk.com/document/development/create-a-badge-user-instance): `create-a-badge-user-instance`
- [同步钉工牌码验证结果](https://open.dingtalk.com/document/development/notification-dingtalk-badge-verification-result): `notification-dingtalk-badge-verification-result`
- [更新钉工牌电子码](https://open.dingtalk.com/document/development/update-dingtalk-user-instance): `update-dingtalk-user-instance`
- [解码钉工牌电子码](https://open.dingtalk.com/document/development/stack-dingtalk-badge): `stack-dingtalk-badge`
- [通知支付结果](https://open.dingtalk.com/document/development/sync-dingtalk-badge-code-payment-result): `sync-dingtalk-badge-code-payment-result`
- [通知退款结果](https://open.dingtalk.com/document/development/notification-dingtalk-badge-code-refund-result): `notification-dingtalk-badge-code-refund-result`
- [配置企业钉工牌](https://open.dingtalk.com/document/development/save-dingtalk-enterprise-instance): `save-dingtalk-enterprise-instance`
- [钉工牌通知消息](https://open.dingtalk.com/document/development/dingtalk-badge-notification-message): `dingtalk-badge-notification-message`

### 工作台（3 条）

- [批量添加最近使用应用](https://open.dingtalk.com/document/development/add-recently-used-apps-in-bulk): `add-recently-used-apps-in-bulk`
- [获取工作台插件权限点](https://open.dingtalk.com/document/development/obtain-the-permissions-of-the-workbench-plug-in): `obtain-the-permissions-of-the-workbench-plug-in`
- [获取工作台插件检验的规则信息](https://open.dingtalk.com/document/development/you-can-call-this-operation-to-obtain-the-information-about): `you-can-call-this-operation-to-obtain-the-information-about`

## 其它 / 未归类（557 条）

### 未归类（557 条）

- [**获取指定用户的所有父部门列表**](https://open.dingtalk.com/document/development/queries-all-parent-departments-of-a-specified-user): `queries-all-parent-departments-of-a-specified-user`
- [**获取第三方企业应用的suite_access_token**](https://open.dingtalk.com/document/development/obtain-application-suite-ticket): `obtain-application-suite-ticket`
- [AI 助理发消息（回复消息模式）](https://open.dingtalk.com/document/development/ai-assistant-messages-reply-mode): `ai-assistant-messages-reply-mode`
- [AI 助理更新消息（主动发送模式）](https://open.dingtalk.com/document/development/the-ai-assistant-updates-active-message-sending-mode): `the-ai-assistant-updates-active-message-sending-mode`
- [AI 助理结束发消息（主动发送模式）](https://open.dingtalk.com/document/development/api-finish): `api-finish`
- [AI 助理预备发消息（主动发送模式）](https://open.dingtalk.com/document/development/api-prepare): `api-prepare`
- [AI助理发消息（主动发送模式）](https://open.dingtalk.com/document/development/ai-assistant-active-sends-messages-mode): `ai-assistant-active-sends-messages-mode`
- [AI助理响应接口-个人权限](https://open.dingtalk.com/document/development/api-assistantmeresponse): `api-assistantmeresponse`
- [AI助理响应接口-应用权限](https://open.dingtalk.com/document/development/api-assistantresponse): `api-assistantresponse`
- [ASR 一句话语音识别](https://open.dingtalk.com/document/development/asr-short-sentence-recognition): `asr-short-sentence-recognition`
- [OCR文字识别](https://open.dingtalk.com/document/development/structured-image-recognition-api): `structured-image-recognition-api`
- [groupId转换为groupKey](https://open.dingtalk.com/document/development/groupid-to-groupkey): `groupid-to-groupkey`
- [groupKey转换为groupId](https://open.dingtalk.com/document/development/convert-groupkey-to-groupid): `convert-groupkey-to-groupid`
- [上传媒体文件](https://open.dingtalk.com/document/development/upload-media-files): `upload-media-files`
- [上传打卡记录](https://open.dingtalk.com/document/development/upload-punch-records): `upload-punch-records`
- [上传文件块](https://open.dingtalk.com/document/development/upload-file-blocks): `upload-file-blocks`
- [下载审批附件](https://open.dingtalk.com/document/development/grants-the-permission-to-download-the-approval-file): `grants-the-permission-to-download-the-approval-file`
- [事件推送](https://open.dingtalk.com/document/development/dingtalk-iot-push-events): `dingtalk-iot-push-events`
- [企业活跃用户统计列表（部门维度）](https://open.dingtalk.com/document/development/query-the-statistics-of-active-users-in-a-department-of): `query-the-statistics-of-active-users-in-a-department-of`
- [使商品过期](https://open.dingtalk.com/document/development/make-goods-expire): `make-goods-expire`
- [使用服务助手推送消息](https://open.dingtalk.com/document/development/the-message-pushing-interface-of-the-assistant): `the-message-pushing-interface-of-the-assistant`
- [使用模板发送工作通知消息](https://open.dingtalk.com/document/development/work-notification-templating-send-notification-interface): `work-notification-templating-send-notification-interface`
- [保存文件到自定义或审批钉盘空间](https://open.dingtalk.com/document/development/add-file-to-user-s-dingtalk-disk): `add-file-to-user-s-dingtalk-disk`
- [保存日志内容](https://open.dingtalk.com/document/development/save-custom-log-content): `save-custom-log-content`
- [保存表单数据](https://open.dingtalk.com/document/development/save-form-data): `save-form-data`
- [修改发票配置](https://open.dingtalk.com/document/development/modify-invoice-configuration): `modify-invoice-configuration`
- [修改打卡时段设置](https://open.dingtalk.com/document/development/modify-card-settings): `modify-card-settings`
- [修改文件（夹）名](https://open.dingtalk.com/document/development/modify-the-file-and-folder-name): `modify-the-file-and-folder-name`
- [修改日程](https://open.dingtalk.com/document/development/schedule-2-0-modify-schedule): `schedule-2-0-modify-schedule`
- [修改日程参与者](https://open.dingtalk.com/document/development/schedule-2-0-participant-modification): `schedule-2-0-participant-modification`
- [修改权限](https://open.dingtalk.com/document/development/modify-pin-disk-permission-click): `modify-pin-disk-permission-click`
- [修改申请单](https://open.dingtalk.com/document/development/user-modify-approval-form): `user-modify-approval-form`
- [修改直播课程的可观看白名单](https://open.dingtalk.com/document/development/modify-the-whitelist-for-live-streaming-courses): `modify-the-whitelist-for-live-streaming-courses`
- [修改设备昵称](https://open.dingtalk.com/document/development/intelligent-hardware-device-nickname-modification): `intelligent-hardware-device-nickname-modification`
- [修改课程](https://open.dingtalk.com/document/development/modify-course): `modify-course`
- [修改项目](https://open.dingtalk.com/document/development/project-change): `project-change`
- [停用群模板](https://open.dingtalk.com/document/development/disable-a-group-template): `disable-a-group-template`
- [关闭互动卡片实例置顶](https://open.dingtalk.com/document/development/disable-the-sticky-card-setting): `disable-the-sticky-card-setting`
- [列出插件信息](https://open.dingtalk.com/document/development/query-plug-in-information): `query-plug-in-information`
- [创建AI助理](https://open.dingtalk.com/document/development/assistant-management-create-an-ai-assistant): `assistant-management-create-an-ai-assistant`
- [创建AI助理的运行任务](https://open.dingtalk.com/document/development/api-createassistantrun): `api-createassistantrun`
- [创建CRM自定义对象数据](https://open.dingtalk.com/document/development/dingtalk-paas-master-create-custom-crm-object-data): `dingtalk-paas-master-create-custom-crm-object-data`
- [创建SSO企业账号](https://open.dingtalk.com/document/development/create-an-sso-account): `create-an-sso-account`
- [创建专属帐号用户](https://open.dingtalk.com/document/development/create-dedicated-accounts): `create-dedicated-accounts`
- [创建互动卡片实例](https://open.dingtalk.com/document/development/create-an-interactive-card-instance-2): `create-an-interactive-card-instance-2`
- [创建企业内部应用H5微应用](https://open.dingtalk.com/document/development/create-an-h5-microapplication): `create-an-h5-microapplication`
- [创建企业客户数据](https://open.dingtalk.com/document/development/dingtalk-paas-master-data-create-crm-customer-data): `dingtalk-paas-master-data-create-crm-customer-data`
- [创建公告](https://open.dingtalk.com/document/development/create-an-enterprise-announcement): `create-an-enterprise-announcement`
- [创建培训课程](https://open.dingtalk.com/document/development/create-live-courses): `create-live-courses`
- [创建学段](https://open.dingtalk.com/document/development/create-a-learning-segment): `create-a-learning-segment`
- [创建学科实例](https://open.dingtalk.com/document/development/create-dingtalk-education-subject-instance): `create-dingtalk-education-subject-instance`
- [创建实例](https://open.dingtalk.com/document/development/initiate-an-approval-process-without-a-process): `initiate-an-approval-process-without-a-process`
- [创建年级](https://open.dingtalk.com/document/development/create-grade): `create-grade`
- [创建待办事项](https://open.dingtalk.com/document/development/create-a-to-do-task): `create-a-to-do-task`
- [创建或更新审批模板](https://open.dingtalk.com/document/development/create-or-update-approval-templates): `create-or-update-approval-templates`
- [创建或更新审批模板](https://open.dingtalk.com/document/development/save-approval-template): `save-approval-template`
- [创建日程](https://open.dingtalk.com/document/development/schedule-2-0-creation-interface): `schedule-2-0-creation-interface`
- [创建消息](https://open.dingtalk.com/document/development/api-createassistantmessage): `api-createassistantmessage`
- [创建班次](https://open.dingtalk.com/document/development/create-modify-shifts): `create-modify-shifts`
- [创建班级](https://open.dingtalk.com/document/development/create-a-class): `create-a-class`
- [创建用户](https://open.dingtalk.com/document/development/create-user): `create-user`
- [创建用户](https://open.dingtalk.com/document/development/user-information-creation): `user-information-creation`
- [创建线程](https://open.dingtalk.com/document/development/api-createassistantthread): `api-createassistantthread`
- [创建群](https://open.dingtalk.com/document/development/create-a-scene-group-v2): `create-a-scene-group-v2`
- [创建群](https://open.dingtalk.com/document/development/session-management-creates-groups): `session-management-creates-groups`
- [创建考勤组](https://open.dingtalk.com/document/development/attendance-group-write): `attendance-group-write`
- [创建联系人数据](https://open.dingtalk.com/document/development/dingtalk-paas-master-data-create-crm-contact-data): `dingtalk-paas-master-data-create-crm-contact-data`
- [创建角色](https://open.dingtalk.com/document/development/address-book-add-role): `address-book-add-role`
- [创建角色组](https://open.dingtalk.com/document/development/add-a-role-group): `add-a-role-group`
- [创建课程](https://open.dingtalk.com/document/development/create-course): `create-course`
- [创建部门](https://open.dingtalk.com/document/development/address-book-creation-department-established-department): `address-book-creation-department-established-department`
- [创建部门](https://open.dingtalk.com/document/development/create-a-department): `create-a-department`
- [创建部门](https://open.dingtalk.com/document/development/industry-connection-department-creation): `industry-connection-department-creation`
- [创建钉钉自建企业账号](https://open.dingtalk.com/document/development/create-dingtalk-user-created-dedicated-account): `create-dingtalk-user-created-dedicated-account`
- [初始化假期余额](https://open.dingtalk.com/document/development/initialize-holiday-balance): `initialize-holiday-balance`
- [初始化家校架构](https://open.dingtalk.com/document/development/initialize-the-home-school-architecture): `initialize-the-home-school-architecture`
- [删除AI助理](https://open.dingtalk.com/document/development/assistant-management-deletes-ai-assistants): `assistant-management-deletes-ai-assistants`
- [删除AI助理的消息体](https://open.dingtalk.com/document/development/api-deleteassistantmessage): `api-deleteassistantmessage`
- [删除AI助理线程](https://open.dingtalk.com/document/development/api-deleteassistantthread): `api-deleteassistantthread`
- [删除H5微应用](https://open.dingtalk.com/document/development/delete-an-h5-microapplication): `delete-an-h5-microapplication`
- [删除企业客户数据](https://open.dingtalk.com/document/development/delete-crm-customer): `delete-crm-customer`
- [删除假期规则](https://open.dingtalk.com/document/development/api-for-deleting-holiday-types): `api-for-deleting-holiday-types`
- [删除公告](https://open.dingtalk.com/document/development/delete-announcements-based-on-the-announcement-id): `delete-announcements-based-on-the-announcement-id`
- [删除发票信息](https://open.dingtalk.com/document/development/delete-invoice-information): `delete-invoice-information`
- [删除回收站文件（夹）](https://open.dingtalk.com/document/development/delete-recycle-bin-files-folders): `delete-recycle-bin-files-folders`
- [删除图文卡片](https://open.dingtalk.com/document/development/delete-message-card): `delete-message-card`
- [删除培训课程](https://open.dingtalk.com/document/development/delete-live-training-courses): `delete-live-training-courses`
- [删除外部联系人](https://open.dingtalk.com/document/development/delete-external-contact): `delete-external-contact`
- [删除学科实例](https://open.dingtalk.com/document/development/delete-dingtalk-education-disciplines): `delete-dingtalk-education-disciplines`
- [删除序列](https://open.dingtalk.com/document/development/delete-sequence): `delete-sequence`
- [删除成本中心](https://open.dingtalk.com/document/development/delete-cost-center): `delete-cost-center`
- [删除成本中心人员信息](https://open.dingtalk.com/document/development/delete-the-personnel-information-of-the-cost-center): `delete-the-personnel-information-of-the-cost-center`
- [删除文件（夹）](https://open.dingtalk.com/document/development/delete-objects): `delete-objects`
- [删除文章](https://open.dingtalk.com/document/development/delete-article-1): `delete-article-1`
- [删除权限](https://open.dingtalk.com/document/development/delete-the-pin-disk-permission): `delete-the-pin-disk-permission`
- [删除模板](https://open.dingtalk.com/document/development/delete-a-template): `delete-a-template`
- [删除班次](https://open.dingtalk.com/document/development/delete-shift): `delete-shift`
- [删除用户](https://open.dingtalk.com/document/development/delete-a-member): `delete-a-member`
- [删除用户](https://open.dingtalk.com/document/development/delete-a-user): `delete-a-user`
- [删除群成员](https://open.dingtalk.com/document/development/scene-group-delete): `scene-group-delete`
- [删除考勤组](https://open.dingtalk.com/document/development/delete-attendance-group): `delete-attendance-group`
- [删除联系人数据](https://open.dingtalk.com/document/development/delete-crm-contact): `delete-crm-contact`
- [删除角色](https://open.dingtalk.com/document/development/delete-role-information): `delete-role-information`
- [删除设备](https://open.dingtalk.com/document/development/delete-a-device): `delete-a-device`
- [删除课程](https://open.dingtalk.com/document/development/delete-course): `delete-course`
- [删除部门](https://open.dingtalk.com/document/development/address-book-deletion-department): `address-book-deletion-department`
- [删除部门](https://open.dingtalk.com/document/development/delete-a-department): `delete-a-department`
- [删除项目](https://open.dingtalk.com/document/development/delete-a-project): `delete-a-project`
- [剪辑直播课程回放](https://open.dingtalk.com/document/development/clip-live-course-playback): `clip-live-course-playback`
- [加入课程](https://open.dingtalk.com/document/development/join-course): `join-course`
- [助理删除知识](https://open.dingtalk.com/document/development/api-deleteknowledge): `api-deleteknowledge`
- [助理学习知识](https://open.dingtalk.com/document/development/api-learnknowledge): `api-learnknowledge`
- [助理重新学习](https://open.dingtalk.com/document/development/api-relearnknowledge): `api-relearnknowledge`
- [单步文件上传](https://open.dingtalk.com/document/development/single-step-file-upload): `single-step-file-upload`
- [发布商品](https://open.dingtalk.com/document/development/release-products): `release-products`
- [发布文章](https://open.dingtalk.com/document/development/article-publishing-interface-1): `article-publishing-interface-1`
- [发起宜搭审批流程](https://open.dingtalk.com/document/development/initiate-the-approval-process): `initiate-the-approval-process`
- [发起审批实例](https://open.dingtalk.com/document/development/oa-approval-initiates-approval-instances): `oa-approval-initiates-approval-instances`
- [发送工作通知](https://open.dingtalk.com/document/development/asynchronous-sending-of-enterprise-session-messages): `asynchronous-sending-of-enterprise-session-messages`
- [发送普通消息](https://open.dingtalk.com/document/development/send-normal-messages-1): `send-normal-messages-1`
- [发送消息到企业群](https://open.dingtalk.com/document/development/send-group-messages): `send-group-messages`
- [发送群助手消息](https://open.dingtalk.com/document/development/group-template-robot-message): `group-template-robot-message`
- [发送钉盘文件给指定用户](https://open.dingtalk.com/document/development/sends-a-file-to-a-specified-user): `sends-a-file-to-a-specified-user`
- [取消日程](https://open.dingtalk.com/document/development/schedule-2-0-cancel-schedule): `schedule-2-0-cancel-schedule`
- [变配租户信息](https://open.dingtalk.com/document/development/change-tenant-information-1): `change-tenant-information-1`
- [同意或拒绝审批任务](https://open.dingtalk.com/document/development/execute-approval-operation-with-attachment): `execute-approval-operation-with-attachment`
- [启用群模板](https://open.dingtalk.com/document/development/enable-a-group-template): `enable-a-group-template`
- [商旅成本中心转换为外部成本中心](https://open.dingtalk.com/document/development/business-travel-cost-center-converted-to-external-cost-center): `business-travel-cost-center-converted-to-external-cost-center`
- [回放课程](https://open.dingtalk.com/document/development/replay-course): `replay-course`
- [多渠道新购校验](https://open.dingtalk.com/document/development/multi-channel-new-purchase-verification): `multi-channel-new-purchase-verification`
- [学习推荐数据回流](https://open.dingtalk.com/document/development/learn-to-recommend-data-backflow): `learn-to-recommend-data-backflow`
- [家庭Feed同步](https://open.dingtalk.com/document/development/dingtalk-education-family-feed-synchronization): `dingtalk-education-family-feed-synchronization`
- [将助理技能发布到组织技能库](https://open.dingtalk.com/document/development/api-addtoorgskillrepository): `api-addtoorgskillrepository`
- [应用内购商品核销](https://open.dingtalk.com/document/development/application-of-in-house-purchase-verification): `application-of-in-house-purchase-verification`
- [开启互动卡片实例置顶](https://open.dingtalk.com/document/development/enable-the-interactive-card-setting): `enable-the-interactive-card-setting`
- [开启分块上传事务](https://open.dingtalk.com/document/development/enable-upload-transaction): `enable-upload-transaction`
- [开始课程](https://open.dingtalk.com/document/development/start-course): `start-course`
- [执行宜搭的审批任务](https://open.dingtalk.com/document/development/execute-appropriate-approval-tasks): `execute-appropriate-approval-tasks`
- [执行自定义API](https://open.dingtalk.com/document/development/run-custom-api): `run-custom-api`
- [批量修改设备](https://open.dingtalk.com/document/development/batch-modify-devices): `batch-modify-devices`
- [批量删除参与考勤人员](https://open.dingtalk.com/document/development/batch-delete-employees-under-the-attendance-group): `batch-delete-employees-under-the-attendance-group`
- [批量删除员工角色](https://open.dingtalk.com/document/development/delete-the-color-information-of-employee-corners-in-batches): `delete-the-color-information-of-employee-corners-in-batches`
- [批量删除地点](https://open.dingtalk.com/document/development/delete-position-in-batches-under-the-attendance-group): `delete-position-in-batches-under-the-attendance-group`
- [批量发起回调](https://open.dingtalk.com/document/development/initiate-multiple-callbacks): `initiate-multiple-callbacks`
- [批量取消待办](https://open.dingtalk.com/document/development/cancel-multiple-tasks): `cancel-multiple-tasks`
- [批量增加员工角色](https://open.dingtalk.com/document/development/add-role-information-to-employees-in-batches): `add-role-information-to-employees-in-batches`
- [批量新增Wi-Fi信息](https://open.dingtalk.com/document/development/batch-add-wifi-under-attendance-group): `batch-add-wifi-under-attendance-group`
- [批量新增参与考勤人员](https://open.dingtalk.com/document/development/batch-add-employees-under-the-attendance-group): `batch-add-employees-under-the-attendance-group`
- [批量新增地点](https://open.dingtalk.com/document/development/atch-add-position-under-attendance-group): `atch-add-position-under-attendance-group`
- [批量更新假期余额](https://open.dingtalk.com/document/development/bulk-update-holiday-balance): `bulk-update-holiday-balance`
- [批量更新实例状态](https://open.dingtalk.com/document/development/update-the-status-of-multiple-instances-at-a-time): `update-the-status-of-multiple-instances-at-a-time`
- [批量查询Wi-Fi信息](https://open.dingtalk.com/document/development/batch-query-wifi-under-attendance-group): `batch-query-wifi-under-attendance-group`
- [批量查询人员排班信息](https://open.dingtalk.com/document/development/query-batch-scheduling-information): `query-batch-scheduling-information`
- [批量查询员工假期余额变更记录](https://open.dingtalk.com/document/development/query-holiday-consumption-records): `query-holiday-consumption-records`
- [批量查询地点](https://open.dingtalk.com/document/development/batch-query-position-under-attendance-group): `batch-query-position-under-attendance-group`
- [批量查询成员排班概要信息](https://open.dingtalk.com/document/development/query-scheduling-summary-information): `query-scheduling-summary-information`
- [批量注册事件类型](https://open.dingtalk.com/document/development/registration-event-type): `registration-event-type`
- [批量注册设备](https://open.dingtalk.com/document/development/batchregister-devices): `batchregister-devices`
- [批量注册设备](https://open.dingtalk.com/document/development/industry-connection-device-batch-registration): `industry-connection-device-batch-registration`
- [批量移除Wi-Fi信息](https://open.dingtalk.com/document/development/batch-remove-wifi-under-attendance-group): `batch-remove-wifi-under-attendance-group`
- [批量获取企业客户数据](https://open.dingtalk.com/document/development/obtains-customer-data-in-batches-based-on-the-id-list): `obtains-customer-data-in-batches-based-on-the-id-list`
- [批量获取客户数据](https://open.dingtalk.com/document/development/dingtalk-paas-master-data-customer-data-search-and-query-interface): `dingtalk-paas-master-data-customer-data-search-and-query-interface`
- [批量获取应用信息](https://open.dingtalk.com/document/development/queries-the-information-about-multiple-applications): `queries-the-information-about-multiple-applications`
- [批量获取服务窗联系人数据](https://open.dingtalk.com/document/development/obtain-contact-data-from-the-service-window): `obtain-contact-data-from-the-service-window`
- [批量获取考勤组摘要](https://open.dingtalk.com/document/development/batch-query-of-simple-information-of-the-attendance-group): `batch-query-of-simple-information-of-the-attendance-group`
- [批量获取考勤组详情](https://open.dingtalk.com/document/development/batch-obtain-attendance-group-details): `batch-obtain-attendance-group-details`
- [批量获取钉钉运动数据](https://open.dingtalk.com/document/development/queries-the-number-of-dingtalk-movement-steps-of-multiple-users): `queries-the-number-of-dingtalk-movement-steps-of-multiple-users`
- [拷贝文件（夹）](https://open.dingtalk.com/document/development/copy-files-folders): `copy-files-folders`
- [按名称搜索班次](https://open.dingtalk.com/document/development/search-shifts-by-rank): `search-shifts-by-rank`
- [按照ID列表批量获取CRM自定义表单数据](https://open.dingtalk.com/document/development/retrieves-custom-crm-forms-from-the-id-list): `retrieves-custom-crm-forms-from-the-id-list`
- [按照ID列表批量获取联系人数据](https://open.dingtalk.com/document/development/retrieves-contact-data-in-batches-based-on-the-id-list): `retrieves-contact-data-in-batches-based-on-the-id-list`
- [授权下载审批钉盘文件](https://open.dingtalk.com/document/development/approve-nail-disk-file-authorization): `approve-nail-disk-file-authorization`
- [授权用户访问企业的自定义空间](https://open.dingtalk.com/document/development/authorize-a-user-to-access-a-custom-workspace-of-an): `authorize-a-user-to-access-a-custom-workspace-of-an`
- [授权预览审批附件](https://open.dingtalk.com/document/development/preview-authorization-attachment): `preview-authorization-attachment`
- [排班制考勤组排班](https://open.dingtalk.com/document/development/scheduling-system-attendance-group-scheduling): `scheduling-system-attendance-group-scheduling`
- [提交文件上传事务](https://open.dingtalk.com/document/development/submit-a-file-upload-transaction): `submit-a-file-upload-transaction`
- [搜索考勤组摘要](https://open.dingtalk.com/document/development/attendance-group-search): `attendance-group-search`
- [撤回工作通知消息](https://open.dingtalk.com/document/development/notification-of-work-withdrawal): `notification-of-work-withdrawal`
- [撤销审批实例](https://open.dingtalk.com/document/development/terminate-a-workflow-by-using-an-instance-id): `terminate-a-workflow-by-using-an-instance-id`
- [新增发票配置](https://open.dingtalk.com/document/development/new-invoice-configuration): `new-invoice-configuration`
- [新增图文卡片](https://open.dingtalk.com/document/development/new-message-card-1): `new-message-card-1`
- [新增或更新表单实例](https://open.dingtalk.com/document/development/add-or-update-form-instances): `add-or-update-form-instances`
- [新增文章](https://open.dingtalk.com/document/development/new-article-1): `new-article-1`
- [新增服务号](https://open.dingtalk.com/document/development/added-service-number): `added-service-number`
- [新增群成员](https://open.dingtalk.com/document/development/add-people-to-scene-group): `add-people-to-scene-group`
- [新增钉钉待办任务](https://open.dingtalk.com/document/development/new-to-do-items): `new-to-do-items`
- [新购宜搭产品](https://open.dingtalk.com/document/development/suitable-for-new-purchase): `suitable-for-new-purchase`
- [更新 AI 助理的使用范围](https://open.dingtalk.com/document/development/api-updateassistantscope): `api-updateassistantscope`
- [更新AI助理基础信息](https://open.dingtalk.com/document/development/api-updateassistantbasicinfo): `api-updateassistantbasicinfo`
- [更新企业内部应用微应用的可使用范围](https://open.dingtalk.com/document/development/set-the-visible-range-of-the-microapplication): `set-the-visible-range-of-the-microapplication`
- [更新企业客户数据](https://open.dingtalk.com/document/development/dingtalk-paas-master-data-update-crm-customer-data): `dingtalk-paas-master-data-update-crm-customer-data`
- [更新企业账号用户信息](https://open.dingtalk.com/document/development/update-dedicated-accounts-information): `update-dedicated-accounts-information`
- [更新假期规则](https://open.dingtalk.com/document/development/holiday-type-update): `holiday-type-update`
- [更新公告](https://open.dingtalk.com/document/development/modify-the-announcement-according-to-the-announcement-id): `modify-the-announcement-according-to-the-announcement-id`
- [更新参与考勤人员](https://open.dingtalk.com/document/development/attendance-group-member-update): `attendance-group-member-update`
- [更新员工花名册](https://open.dingtalk.com/document/development/update-employee-roster): `update-employee-roster`
- [更新员工花名册信息](https://open.dingtalk.com/document/development/intelligent-personnel-update-employee-file-information): `intelligent-personnel-update-employee-file-information`
- [更新图文卡片](https://open.dingtalk.com/document/development/message-card-material-update-interface): `message-card-material-update-interface`
- [更新外部联系人](https://open.dingtalk.com/document/development/update-enterprise-external-contacts): `update-enterprise-external-contacts`
- [更新学科实例](https://open.dingtalk.com/document/development/update-dingtalk-education-instance): `update-dingtalk-education-instance`
- [更新实例状态](https://open.dingtalk.com/document/development/to-do-instance-status): `to-do-instance-status`
- [更新工作通知状态栏](https://open.dingtalk.com/document/development/update-work-notification-status-bar): `update-work-notification-status-bar`
- [更新待办状态](https://open.dingtalk.com/document/development/update-to-do-task-status): `update-to-do-task-status`
- [更新文章](https://open.dingtalk.com/document/development/save-article-details-1): `save-article-details-1`
- [更新服务号](https://open.dingtalk.com/document/development/service-number-update-1): `service-number-update-1`
- [更新流程实例](https://open.dingtalk.com/document/development/update-process-instance-yida): `update-process-instance-yida`
- [更新状态](https://open.dingtalk.com/document/development/update-status): `update-status`
- [更新用户信息](https://open.dingtalk.com/document/development/update-user-details): `update-user-details`
- [更新用户信息](https://open.dingtalk.com/document/development/user-information-update): `user-information-update`
- [更新申请单状态](https://open.dingtalk.com/document/development/update-approval-form): `update-approval-form`
- [更新群](https://open.dingtalk.com/document/development/modify-a-group-session): `modify-a-group-session`
- [更新群](https://open.dingtalk.com/document/development/scene-group-update): `scene-group-update`
- [更新群成员的群昵称](https://open.dingtalk.com/document/development/set-a-group-nickname): `set-a-group-nickname`
- [更新群管理员](https://open.dingtalk.com/document/development/set-chat-admin): `set-chat-admin`
- [更新考勤组](https://open.dingtalk.com/document/development/attendance-group-update-interface): `attendance-group-update-interface`
- [更新联系人数据](https://open.dingtalk.com/document/development/dingtalk-paas-master-data-update-crm-contact-data): `dingtalk-paas-master-data-update-crm-contact-data`
- [更新自定义对象数据](https://open.dingtalk.com/document/development/crm-master-data-opens-interface-for-updating-custom-object-data): `crm-master-data-opens-interface-for-updating-custom-object-data`
- [更新表单数据](https://open.dingtalk.com/document/development/update-form-data): `update-form-data`
- [更新角色名称](https://open.dingtalk.com/document/development/update-the-character-name): `update-the-character-name`
- [更新部门](https://open.dingtalk.com/document/development/address-book-update-department): `address-book-update-department`
- [更新部门](https://open.dingtalk.com/document/development/update-a-department): `update-a-department`
- [更新部门扩展信息](https://open.dingtalk.com/document/development/department-update-extension-information): `department-update-extension-information`
- [更新钉钉待办任务](https://open.dingtalk.com/document/development/update-to-do-status): `update-to-do-status`
- [服务号菜单更新](https://open.dingtalk.com/document/development/service-number-menu-update): `service-number-menu-update`
- [服务商获取第三方应用授权企业的access_token](https://open.dingtalk.com/document/development/obtain-isvapp-token): `obtain-isvapp-token`
- [机票城市搜索](https://open.dingtalk.com/document/development/air-ticket-city-search): `air-ticket-city-search`
- [查询 AI 助理基本信息](https://open.dingtalk.com/document/development/api-retrieveassistantbasicinfo): `api-retrieveassistantbasicinfo`
- [查询企业下用户待办列表](https://open.dingtalk.com/document/development/get-the-user-s-to-do-items): `get-the-user-s-to-do-items`
- [查询企业个人待办数量](https://open.dingtalk.com/document/development/query-the-number-of-to-do-tasks-of-the-enterprise): `query-the-number-of-to-do-tasks-of-the-enterprise`
- [查询企业级别](https://open.dingtalk.com/document/development/query-enterprise-level): `query-enterprise-level`
- [查询企业考勤排班详情](https://open.dingtalk.com/document/development/interface-for-daily-full-query-of-attendance-scheduling-information): `interface-for-daily-full-query-of-attendance-scheduling-information`
- [查询企业账号用户详情](https://open.dingtalk.com/document/development/queries-the-details-of-a-dedicated-account): `queries-the-details-of-a-dedicated-account`
- [查询企业通讯录未激活用户列表](https://open.dingtalk.com/document/development/queries-the-list-of-inactive-accounts-in-the-key-account): `queries-the-list-of-inactive-accounts-in-the-key-account`
- [查询假期规则列表](https://open.dingtalk.com/document/development/holiday-type-query): `holiday-type-query`
- [查询历史班次](https://open.dingtalk.com/document/development/query-history-shifts): `query-history-shifts`
- [查询参与考勤人员列表](https://open.dingtalk.com/document/development/batch-query-of-employees-in-the-attendance-group): `batch-query-of-employees-in-the-attendance-group`
- [查询可用发票列表](https://open.dingtalk.com/document/development/query-available-invoices): `query-available-invoices`
- [查询员工智能考勤机列表](https://open.dingtalk.com/document/development/query-the-list-of-employee-intelligent-attendance-machines): `query-the-list-of-employee-intelligent-attendance-machines`
- [查询商品列表](https://open.dingtalk.com/document/development/query-product-lists): `query-product-lists`
- [查询回收站文件（夹）列表](https://open.dingtalk.com/document/development/obtain-the-recycle-bin-folder-list): `obtain-the-recycle-bin-folder-list`
- [查询图文卡片列表](https://open.dingtalk.com/document/development/query-message-card-list): `query-message-card-list`
- [查询家庭孩子信息](https://open.dingtalk.com/document/development/query-family-child-information): `query-family-child-information`
- [查询应用信息列表](https://open.dingtalk.com/document/development/queries-application-information): `queries-application-information`
- [查询待办列表](https://open.dingtalk.com/document/development/query-a-user-s-to-do-items): `query-a-user-s-to-do-items`
- [查询成本中心](https://open.dingtalk.com/document/development/query-cost-center): `query-cost-center`
- [查询排班打卡结果](https://open.dingtalk.com/document/development/query-the-results-of-a-batch-of-tasks): `query-the-results-of-a-batch-of-tasks`
- [查询插件信息列表](https://open.dingtalk.com/document/development/query-plug-ins): `query-plug-ins`
- [查询文件（夹）信息](https://open.dingtalk.com/document/development/obtain-file-information): `obtain-file-information`
- [查询文件（夹）列表](https://open.dingtalk.com/document/development/obtain-the-file-list): `obtain-the-file-list`
- [查询文档模板](https://open.dingtalk.com/document/development/query-a-document-template): `query-a-document-template`
- [查询文章列表](https://open.dingtalk.com/document/development/query-the-article-list): `query-the-article-list`
- [查询是否启用智能统计报表](https://open.dingtalk.com/document/development/determine-whether-to-enable-attendance-intelligent-report): `determine-whether-to-enable-attendance-intelligent-report`
- [查询服务号列表](https://open.dingtalk.com/document/development/query-service-number-list): `query-service-number-list`
- [查询服务号菜单](https://open.dingtalk.com/document/development/query-service-number-menu-1): `query-service-number-menu-1`
- [查询服务号详情](https://open.dingtalk.com/document/development/inquire-about-service-number-details): `inquire-about-service-number-details`
- [查询激活码](https://open.dingtalk.com/document/development/query-activation-code): `query-activation-code`
- [查询用户是否参与企业步数排行榜](https://open.dingtalk.com/document/development/check-whether-dingtalk-is-enabled): `check-whether-dingtalk-is-enabled`
- [查询用户详情](https://open.dingtalk.com/document/development/queries-user-details): `queries-user-details`
- [查询直播的观看数据](https://open.dingtalk.com/document/development/queries-the-playback-data-of-a-live-stream): `queries-the-playback-data-of-a-live-stream`
- [查询直播课程的可观看白名单](https://open.dingtalk.com/document/development/query-the-whitelist-of-live-courses): `query-the-whitelist-of-live-courses`
- [查询群信息](https://open.dingtalk.com/document/development/obtain-a-group-session): `obtain-a-group-session`
- [查询群信息](https://open.dingtalk.com/document/development/queries-the-basic-information-of-a-scenario-group): `queries-the-basic-information-of-a-scenario-group`
- [查询群消息已读人员列表](https://open.dingtalk.com/document/development/queries-the-list-of-people-who-have-read-a-group): `queries-the-list-of-people-who-have-read-a-group`
- [查询表单实例数据](https://open.dingtalk.com/document/development/querying-form-instance-data): `querying-form-instance-data`
- [查询表单数据](https://open.dingtalk.com/document/development/query-form-data): `query-form-data`
- [查询设备列表](https://open.dingtalk.com/document/development/intelligent-hardware-list-query): `intelligent-hardware-list-query`
- [查询设备详情](https://open.dingtalk.com/document/development/intelligent-hardware-device-query): `intelligent-hardware-device-query`
- [查询请假状态](https://open.dingtalk.com/document/development/query-status): `query-status`
- [查询销售用户信息](https://open.dingtalk.com/document/development/query-sales-user-information): `query-sales-user-information`
- [查询项目中文件操作日志](https://open.dingtalk.com/document/development/query-file-operation-logs-of-a-project): `query-file-operation-logs-of-a-project`
- [查询预估价](https://open.dingtalk.com/document/development/query-estimated-price): `query-estimated-price`
- [校验变配](https://open.dingtalk.com/document/development/verify-configuration): `verify-configuration`
- [校验用户是否在当前考勤组](https://open.dingtalk.com/document/development/query-members-by-id): `query-members-by-id`
- [校验订单的升级](https://open.dingtalk.com/document/development/upgrade-the-verification-order): `upgrade-the-verification-order`
- [校验订单的升级状态](https://open.dingtalk.com/document/development/upgrade-status-of-the-verification-order): `upgrade-status-of-the-verification-order`
- [根据ID列表批量获取跟进记录数据](https://open.dingtalk.com/document/development/dingtalk-the-primary-data-of-apsara-stack-agility-paas-allows-you): `dingtalk-the-primary-data-of-apsara-stack-agility-paas-allows-you`
- [根据groupKey查询考勤组信息](https://open.dingtalk.com/document/development/queries-attendance-group-information-by-id): `queries-attendance-group-information-by-id`
- [根据sns临时授权码获取用户信息](https://open.dingtalk.com/document/development/obtain-the-user-information-based-on-the-sns-temporary-authorization): `obtain-the-user-information-based-on-the-sns-temporary-authorization`
- [根据unionid获取用户userid](https://open.dingtalk.com/document/development/query-a-user-by-the-union-id): `query-a-user-by-the-union-id`
- [根据unionid获取用户userid](https://open.dingtalk.com/document/development/you-can-call-this-operation-to-retrieve-the-userids-of): `you-can-call-this-operation-to-retrieve-the-userids-of`
- [根据手机号查询企业账号用户](https://open.dingtalk.com/document/development/obtain-the-userid-of-your-mobile-phone-number): `obtain-the-userid-of-your-mobile-phone-number`
- [根据手机号查询用户](https://open.dingtalk.com/document/development/retrieve-userid-from-mobile-phone-number): `retrieve-userid-from-mobile-phone-number`
- [根据指定条件查询联系人数据](https://open.dingtalk.com/document/development/dingtalk-the-contact-data-query-api): `dingtalk-the-contact-data-query-api`
- [根据指定条件查询自定义对象数据](https://open.dingtalk.com/document/development/retrieve-custom-crm-object-data): `retrieve-custom-crm-object-data`
- [根据指定条件查询跟进记录数据](https://open.dingtalk.com/document/development/query-and-dingtalk-data-of-track-records-in-apsara-stack): `query-and-dingtalk-data-of-track-records-in-apsara-stack`
- [根据流程实例ID获取流程实例](https://open.dingtalk.com/document/development/queries-a-process-instance-based-on-its-id): `queries-a-process-instance-based-on-its-id`
- [根据设备ID查询设备](https://open.dingtalk.com/document/development/the-smart-hardware-can-query-details-based-on-the-device): `the-smart-hardware-can-query-details-based-on-the-device`
- [检索AI助理线程](https://open.dingtalk.com/document/development/api-retrieveassistantthread): `api-retrieveassistantthread`
- [注册互动卡片回调地址](https://open.dingtalk.com/document/development/registration-card-interaction-callback-address-1): `registration-card-interaction-callback-address-1`
- [注册单个设备](https://open.dingtalk.com/document/development/industry-connection-single-device-registration): `industry-connection-single-device-registration`
- [注册设备](https://open.dingtalk.com/document/development/register-devices): `register-devices`
- [注册账号](https://open.dingtalk.com/document/development/register-an-account): `register-an-account`
- [消息撤回](https://open.dingtalk.com/document/development/service-number-message-withdrawal): `service-number-message-withdrawal`
- [消息群发](https://open.dingtalk.com/document/development/interactive-service-window-group-message-sending-interface): `interactive-service-window-group-message-sending-interface`
- [添加企业待入职员工](https://open.dingtalk.com/document/development/add-employees-to-be-hired-through-intelligent-personnel): `add-employees-to-be-hired-through-intelligent-personnel`
- [添加假期规则](https://open.dingtalk.com/document/development/holiday-type-added): `holiday-type-added`
- [添加外部联系人](https://open.dingtalk.com/document/development/add-enterprise-external-contacts): `add-enterprise-external-contacts`
- [添加学生](https://open.dingtalk.com/document/development/add-student): `add-student`
- [添加审批评论](https://open.dingtalk.com/document/development/add-an-approval-comment): `add-an-approval-comment`
- [添加家长](https://open.dingtalk.com/document/development/add-parent): `add-parent`
- [添加文件（夹）](https://open.dingtalk.com/document/development/add-file-and-folder): `add-file-and-folder`
- [添加权限](https://open.dingtalk.com/document/development/add-pin-disk-permission): `add-pin-disk-permission`
- [添加老师](https://open.dingtalk.com/document/development/add-teacher): `add-teacher`
- [添加自定义空间权限](https://open.dingtalk.com/document/development/add-custom-workspace-permissions): `add-custom-workspace-permissions`
- [添加课程参与方](https://open.dingtalk.com/document/development/add-course-participants): `add-course-participants`
- [添加项目](https://open.dingtalk.com/document/development/add-a-project): `add-a-project`
- [清理审批数据](https://open.dingtalk.com/document/development/clean-up-workflow-data): `clean-up-workflow-data`
- [清空回收站](https://open.dingtalk.com/document/development/empty-recycle-bin-files-folders): `empty-recycle-bin-files-folders`
- [火车票城市搜索](https://open.dingtalk.com/document/development/train-ticket-city-search): `train-ticket-city-search`
- [移动文件（夹）](https://open.dingtalk.com/document/development/move-file-and-folder): `move-file-and-folder`
- [移除助理组织技能库技能](https://open.dingtalk.com/document/development/api-removefromorgskillrepository): `api-removefromorgskillrepository`
- [移除租户资源](https://open.dingtalk.com/document/development/remove-tenant-resources): `remove-tenant-resources`
- [移除课程参与方](https://open.dingtalk.com/document/development/remove-course-participants): `remove-course-participants`
- [终止阿里云授权](https://open.dingtalk.com/document/development/terminate-authorization-for-alibaba-cloud-services): `terminate-authorization-for-alibaba-cloud-services`
- [绑定设备](https://open.dingtalk.com/document/development/establishing-a-binding-relationship-between-intelligent-hardware-and-cloud): `establishing-a-binding-relationship-between-intelligent-hardware-and-cloud`
- [结束课程](https://open.dingtalk.com/document/development/end-course): `end-course`
- [统计企业活跃用户](https://open.dingtalk.com/document/development/query-for-dau-statistics): `query-for-dau-statistics`
- [续费服务订单](https://open.dingtalk.com/document/development/renewal-service-order): `renewal-service-order`
- [续费租户](https://open.dingtalk.com/document/development/renewal-tenant): `renewal-tenant`
- [自定义机器人发送群消息](https://open.dingtalk.com/document/development/custom-robots-send-group-messages): `custom-robots-send-group-messages`
- [获取 AI 助理的使用范围](https://open.dingtalk.com/document/development/api-retrieveassistantscope): `api-retrieveassistantscope`
- [获取AI助理对话明细列表](https://open.dingtalk.com/document/development/api-loglist): `api-loglist`
- [获取AI助理的消息体](https://open.dingtalk.com/document/development/api-retrieveassistantmessage): `api-retrieveassistantmessage`
- [获取AI助理的消息列表](https://open.dingtalk.com/document/development/api-listassistantmessage): `api-listassistantmessage`
- [获取AI助理的运行任务](https://open.dingtalk.com/document/development/api-retrieveassistantrun): `api-retrieveassistantrun`
- [获取AI助理的运行任务的列表](https://open.dingtalk.com/document/development/api-listassistantrun): `api-listassistantrun`
- [获取jsapi_ticket](https://open.dingtalk.com/document/development/obtain-jsapi-ticket): `obtain-jsapi-ticket`
- [获取个人或部门钉钉运动数据](https://open.dingtalk.com/document/development/queries-individual-or-department-dingtalk-exercise-steps): `queries-individual-or-department-dingtalk-exercise-steps`
- [获取主干组织列表](https://open.dingtalk.com/document/development/obtain-backbone-organization-list): `obtain-backbone-organization-list`
- [获取互动服务窗相关数据](https://open.dingtalk.com/document/development/queries-the-data-about-the-interactive-service-window): `queries-the-data-about-the-interactive-service-window`
- [获取人员列表](https://open.dingtalk.com/document/development/obtains-a-list-of-home-school-user-identities): `obtains-a-list-of-home-school-user-identities`
- [获取人员详情](https://open.dingtalk.com/document/development/obtain-the-identity-details-of-home-school-personnel): `obtain-the-identity-details-of-home-school-personnel`
- [获取企业DING使用数据](https://open.dingtalk.com/document/development/enterprise-ding-quantity-statistics): `enterprise-ding-quantity-statistics`
- [获取企业DING使用数据（部门维度）](https://open.dingtalk.com/document/development/query-the-departmental-transmission-status-of-key-clients): `query-the-departmental-transmission-status-of-key-clients`
- [获取企业DING发送统计数据](https://open.dingtalk.com/document/development/obtain-sending-statistics-of-an-enterprise-ding): `obtain-sending-statistics-of-an-enterprise-ding`
- [获取企业DING接收及评论统计数据](https://open.dingtalk.com/document/development/obtain-statistics-on-receiving-and-comments-of-enterprise-ding): `obtain-statistics-on-receiving-and-comments-of-enterprise-ding`
- [获取企业下的自定义空间](https://open.dingtalk.com/document/development/obtain-user-space-under-the-enterprise): `obtain-user-space-under-the-enterprise`
- [获取企业信息](https://open.dingtalk.com/document/development/obtain-enterprise-information): `obtain-enterprise-information`
- [获取企业全员圈统计数据](https://open.dingtalk.com/document/development/obtains-the-statistical-data-of-all-employees-of-an-enterprise): `obtains-the-statistical-data-of-all-employees-of-an-enterprise`
- [获取企业公告统计数据](https://open.dingtalk.com/document/development/queries-corporate-announcement-statistics): `queries-corporate-announcement-statistics`
- [获取企业内部应用微应用的可使用范围](https://open.dingtalk.com/document/development/gets-the-microapplication-visible-range-set-by-the-enterprise): `gets-the-microapplication-visible-range-set-by-the-enterprise`
- [获取企业内部应用的access_token](https://open.dingtalk.com/document/development/obtain-orgapp-token): `obtain-orgapp-token`
- [获取企业单聊统计数据](https://open.dingtalk.com/document/development/queries-the-statistics-on-one-time-enterprise-chats): `queries-the-statistics-on-one-time-enterprise-chats`
- [获取企业发布智能填表数（组织维度）](https://open.dingtalk.com/document/development/queries-the-number-of-tables-published-in-an-organization): `queries-the-number-of-tables-published-in-an-organization`
- [获取企业发红包统计数据](https://open.dingtalk.com/document/development/obtains-the-statistics-on-red-packets-issued-by-enterprises): `obtains-the-statistics-on-red-packets-issued-by-enterprises`
- [获取企业各类群组创建情况](https://open.dingtalk.com/document/development/api-for-obtaining-the-creation-status-of-various-groups): `api-for-obtaining-the-creation-status-of-various-groups`
- [获取企业员工类型统计数据](https://open.dingtalk.com/document/development/obtains-statistics-on-employee-types): `obtains-statistics-on-employee-types`
- [获取企业商旅酒店订单数据](https://open.dingtalk.com/document/development/enterprises-obtain-order-data-for-business-hotels): `enterprises-obtain-order-data-for-business-hotels`
- [获取企业审批统计数据](https://open.dingtalk.com/document/development/obtains-enterprise-approval-statistics): `obtains-enterprise-approval-statistics`
- [获取企业客户的元数据](https://open.dingtalk.com/document/development/get-metadata-description-of-crm-customer-object): `get-metadata-description-of-crm-customer-object`
- [获取企业已经加入的或申请加入中的上下游组织的信息](https://open.dingtalk.com/document/development/obtains-information-about-the-workspaces-that-the-enterprise-has-joined): `obtains-information-about-the-workspaces-that-the-enterprise-has-joined`
- [获取企业应用访问情况](https://open.dingtalk.com/document/development/queries-the-daily-usage-summary-of-microapplications-in-an-enterprise): `queries-the-daily-usage-summary-of-microapplications-in-an-enterprise`
- [获取企业待办统计数据](https://open.dingtalk.com/document/development/obtains-the-to-do-statistics-of-an-enterprise): `obtains-the-to-do-statistics-of-an-enterprise`
- [获取企业所有应用列表](https://open.dingtalk.com/document/development/manager-microapplications-api-permission): `manager-microapplications-api-permission`
- [获取企业授权信息](https://open.dingtalk.com/document/development/obtains-the-basic-information-of-an-enterprise): `obtains-the-basic-information-of-an-enterprise`
- [获取企业接收红包统计数据](https://open.dingtalk.com/document/development/queries-the-red-envelope-receiving-statistics-of-an-enterprise): `queries-the-red-envelope-receiving-statistics-of-an-enterprise`
- [获取企业文档统计数据](https://open.dingtalk.com/document/development/get-enterprise-document-statistics): `get-enterprise-document-statistics`
- [获取企业日志统计数据](https://open.dingtalk.com/document/development/obtain-enterprise-log-statistics): `obtain-enterprise-log-statistics`
- [获取企业日程统计数据](https://open.dingtalk.com/document/development/queries-enterprise-schedule-statistics): `queries-enterprise-schedule-statistics`
- [获取企业某天的所有部门电话会议统计列表](https://open.dingtalk.com/document/development/major-customer-department-dimension-teleconference-statistics): `major-customer-department-dimension-teleconference-statistics`
- [获取企业某天的所有部门视频会议统计数据](https://open.dingtalk.com/document/development/video-conferencing-statistics-list-for-key-accounts-and-departments): `video-conferencing-statistics-list-for-key-accounts-and-departments`
- [获取企业某天的电话会议数据](https://open.dingtalk.com/document/development/major-customer-teleconference-statistics-interface): `major-customer-teleconference-statistics-interface`
- [获取企业某天的视频会议统计数据](https://open.dingtalk.com/document/development/video-conferencing-statistics-query-v2-for-key-accounts): `video-conferencing-statistics-query-v2-for-key-accounts`
- [获取企业火车票订单数据](https://open.dingtalk.com/document/development/obtains-the-enterprise-train-ticket-order-data): `obtains-the-enterprise-train-ticket-order-data`
- [获取企业用户在线统计数据](https://open.dingtalk.com/document/development/retrieve-online-statistics-of-enterprise-users): `retrieve-online-statistics-of-enterprise-users`
- [获取企业用户激活状态统计数据](https://open.dingtalk.com/document/development/obtains-statistics-on-user-activation-status): `obtains-statistics-on-user-activation-status`
- [获取企业电话会议明细列表](https://open.dingtalk.com/document/development/major-account-conference-call-details-list): `major-account-conference-call-details-list`
- [获取企业电话会议统计数据](https://open.dingtalk.com/document/development/get-enterprise-teleconference-statistics): `get-enterprise-teleconference-statistics`
- [获取企业签到统计数据](https://open.dingtalk.com/document/development/queries-enterprise-check-in-statistics): `queries-enterprise-check-in-statistics`
- [获取企业群直播统计数据](https://open.dingtalk.com/document/development/obtains-the-live-stream-statistics-for-an-enterprise-group): `obtains-the-live-stream-statistics-for-an-enterprise-group`
- [获取企业群聊统计数据](https://open.dingtalk.com/document/development/obtain-enterprise-group-chat-statistics): `obtain-enterprise-group-chat-statistics`
- [获取企业考勤统计数据](https://open.dingtalk.com/document/development/queries-enterprise-attendance-statistics): `queries-enterprise-attendance-statistics`
- [获取企业聊天数据](https://open.dingtalk.com/document/development/chat-data-statistics-query-for-key-accounts): `chat-data-statistics-query-for-key-accounts`
- [获取企业视频会议明细列表](https://open.dingtalk.com/document/development/video-conference-details-for-key-accounts): `video-conference-details-for-key-accounts`
- [获取企业视频会议统计数据](https://open.dingtalk.com/document/development/get-enterprise-video-conference-statistics): `get-enterprise-video-conference-statistics`
- [获取企业视频直播统计列表（部门维度）](https://open.dingtalk.com/document/development/live-broadcast-summary-statistics-of-key-account-departments): `live-broadcast-summary-statistics-of-key-account-departments`
- [获取企业视频直播统计数据](https://open.dingtalk.com/document/development/query-live-streaming-statistics): `query-live-streaming-statistics`
- [获取企业邮箱统计数据](https://open.dingtalk.com/document/development/queries-enterprise-email-statistics): `queries-enterprise-email-statistics`
- [获取企业部门聊天数据（部门维度）](https://open.dingtalk.com/document/development/dingtalk-chat-information-in-key-accounts): `dingtalk-chat-information-in-key-accounts`
- [获取企业钉盘统计数据](https://open.dingtalk.com/document/development/obtains-the-statistics-on-enterprise-dingtalk-trays): `obtains-the-statistics-on-enterprise-dingtalk-trays`
- [获取企业钉钉运动统计数据](https://open.dingtalk.com/document/development/queries-dingtalk-movement-statistics): `queries-dingtalk-movement-statistics`
- [获取入群二维码链接](https://open.dingtalk.com/document/development/obtain-a-qr-code-link): `obtain-a-qr-code-link`
- [获取公告ID列表](https://open.dingtalk.com/document/development/obtains-the-id-list-of-announcements-that-are-not-deleted): `obtains-the-id-list-of-announcements-that-are-not-deleted`
- [获取公告分类列表](https://open.dingtalk.com/document/development/obtains-the-list-of-categories-not-deleted-for-enterprise-announcements): `obtains-the-list-of-categories-not-deleted-for-enterprise-announcements`
- [获取公告详情](https://open.dingtalk.com/document/development/obtains-the-details-of-a-bulletin-that-is-not-deleted): `obtains-the-details-of-a-bulletin-that-is-not-deleted`
- [获取内购商品SKU页面地址](https://open.dingtalk.com/document/development/obtain-the-address-of-the-product-sku-details-page): `obtain-the-address-of-the-product-sku-details-page`
- [获取内购订单信息](https://open.dingtalk.com/document/development/obtain-information-about-internal-purchase-orders): `obtain-information-about-internal-purchase-orders`
- [获取分支组织列表](https://open.dingtalk.com/document/development/obtains-the-branch-organization-list): `obtains-the-branch-organization-list`
- [获取助理技能信息](https://open.dingtalk.com/document/development/api-getassistantactioninfo): `api-getassistantactioninfo`
- [获取单个审批实例详情](https://open.dingtalk.com/document/development/get-details-single-approval-instance): `get-details-single-approval-instance`
- [获取参与考勤人员](https://open.dingtalk.com/document/development/batch-query-of-attendance-group-members): `batch-query-of-attendance-group-members`
- [获取参与考勤人员的userid](https://open.dingtalk.com/document/development/query-attendance-group-personnel-information-in-batches): `query-attendance-group-personnel-information-in-batches`
- [获取发布智能填表数量和使用智能填表人数（部门维度）](https://open.dingtalk.com/document/development/obtains-the-number-of-tables-published-by-the-enterprise-from): `obtains-the-number-of-tables-published-by-the-enterprise-from`
- [获取员工人数](https://open.dingtalk.com/document/development/obtain-the-number-of-employees): `obtain-the-number-of-employees`
- [获取员工人数](https://open.dingtalk.com/document/development/user-management-acquires-number-employees): `user-management-acquires-number-employees`
- [获取员工离职信息](https://open.dingtalk.com/document/development/obtain-multiple-employee-demission-information): `obtain-multiple-employee-demission-information`
- [获取员工花名册字段信息](https://open.dingtalk.com/document/development/intelligent-personnel-obtain-employee-roster-information): `intelligent-personnel-obtain-employee-roster-information`
- [获取员工花名册字段信息](https://open.dingtalk.com/document/development/obtain-employee-roster-field-information-in-batches): `obtain-employee-roster-field-information-in-batches`
- [获取图文卡片详情](https://open.dingtalk.com/document/development/get-message-card-details): `get-message-card-details`
- [获取在职员工列表](https://open.dingtalk.com/document/development/intelligent-personnel-query-the-list-of-on-the-job-employees-of-the): `intelligent-personnel-query-the-list-of-on-the-job-employees-of-the`
- [获取培训观看数据](https://open.dingtalk.com/document/development/obtains-the-playback-data-of-a-live-stream): `obtains-the-playback-data-of-a-live-stream`
- [获取培训课程的基本信息](https://open.dingtalk.com/document/development/get-basic-information-about-the-course): `get-basic-information-about-the-course`
- [获取填表实例数据](https://open.dingtalk.com/document/development/obtains-multiple-form-filling-records): `obtains-multiple-form-filling-records`
- [获取外部联系人列表](https://open.dingtalk.com/document/development/obtain-the-external-contact-list): `obtain-the-external-contact-list`
- [获取外部联系人标签列表](https://open.dingtalk.com/document/development/obtains-a-list-of-external-contact-tags): `obtains-a-list-of-external-contact-tags`
- [获取外部联系人详情](https://open.dingtalk.com/document/development/obtains-the-external-contact-details-of-an-enterprise): `obtains-the-external-contact-details-of-an-enterprise`
- [获取多个表单实例ID](https://open.dingtalk.com/document/development/obtain-the-ids-of-multiple-form-instances): `obtain-the-ids-of-multiple-form-instances`
- [获取子部门ID列表](https://open.dingtalk.com/document/development/obtain-a-sub-department-id-list): `obtain-a-sub-department-id-list`
- [获取学习的知识列表](https://open.dingtalk.com/document/development/api-getknowledgelist): `api-getknowledgelist`
- [获取学段元数据列表](https://open.dingtalk.com/document/development/dingtalk-the-main-data-of-the-education-ecosystem-to-query): `dingtalk-the-main-data-of-the-education-ecosystem-to-query`
- [获取学生ID列表](https://open.dingtalk.com/document/development/retrieve-student-based-class): `retrieve-student-based-class`
- [获取学生信息](https://open.dingtalk.com/document/development/obtain-student-information): `obtain-student-information`
- [获取学生监护人详情](https://open.dingtalk.com/document/development/obtain-the-relationship-between-home-and-school-personnel): `obtain-the-relationship-between-home-and-school-personnel`
- [获取学科元数据列表](https://open.dingtalk.com/document/development/dingtalk-the-main-data-of-the-education-ecosystem-query-the-subject): `dingtalk-the-main-data-of-the-education-ecosystem-query-the-subject`
- [获取学科实例列表](https://open.dingtalk.com/document/development/get-the-list-of-subject-examples): `get-the-list-of-subject-examples`
- [获取学科实例详情](https://open.dingtalk.com/document/development/query-dingtalk-education-subject-instances): `query-dingtalk-education-subject-instances`
- [获取定制应用的access_token](https://open.dingtalk.com/document/development/obtains-the-enterprise-authorized-credential): `obtains-the-enterprise-authorized-credential`
- [获取实例ID列表](https://open.dingtalk.com/document/development/obtains-a-list-of-instance-ids): `obtains-a-list-of-instance-ids`
- [获取实例详情](https://open.dingtalk.com/document/development/query-collection-form-instance-details): `query-collection-form-instance-details`
- [获取审批实例ID列表](https://open.dingtalk.com/document/development/operation-to-retrieve-a-list-of): `operation-to-retrieve-a-list-of`
- [获取审批钉盘空间信息](https://open.dingtalk.com/document/development/query-the-space-of-an-approval-nail): `query-the-space-of-an-approval-nail`
- [获取工作通知消息的发送结果](https://open.dingtalk.com/document/development/gets-the-result-of-sending-messages-asynchronously-to-the-enterprise): `gets-the-result-of-sending-messages-asynchronously-to-the-enterprise`
- [获取工作通知消息的发送进度](https://open.dingtalk.com/document/development/obtain-the-sending-progress-of-asynchronous-sending-of-enterprise-session): `obtain-the-sending-progress-of-asynchronous-sending-of-enterprise-session`
- [获取已加入或正在申请加入上下游组织的组织和个人信息](https://open.dingtalk.com/document/development/obtains-the-information-about-how-to-join-or-apply-to): `obtains-the-information-about-how-to-join-or-apply-to`
- [获取平台服务资源](https://open.dingtalk.com/document/development/obtain-platform-service-resources): `obtain-platform-service-resources`
- [获取平台资源](https://open.dingtalk.com/document/development/obtain-platform-resources): `obtain-platform-resources`
- [获取年报数据](https://open.dingtalk.com/document/development/obtain-annual-report-data): `obtain-annual-report-data`
- [获取应用下的页面列表](https://open.dingtalk.com/document/development/obtains-the-page-list-under-an-application): `obtains-the-page-list-under-an-application`
- [获取应用未激活的企业列表](https://open.dingtalk.com/document/development/obtains-a-list-of-enterprises-whose-applications-are-not-activated): `obtains-a-list-of-enterprises-whose-applications-are-not-activated`
- [获取应用管理后台免登的用户信息](https://open.dingtalk.com/document/development/exchange-code-for-the-identity-information-of-a-microapplication-administrator): `exchange-code-for-the-identity-information-of-a-microapplication-administrator`
- [获取应用自定义空间使用详情](https://open.dingtalk.com/document/development/queries-the-usage-details-of-a-custom-application-space): `queries-the-usage-details-of-a-custom-application-space`
- [获取当前企业所有可管理的表单](https://open.dingtalk.com/document/development/obtains-the-information-about-all-manageable-templates-of-the-current): `obtains-the-information-about-all-manageable-templates-of-the-current`
- [获取待入职员工列表](https://open.dingtalk.com/document/development/intelligent-personnel-query-the-list-of-employees-to-be-hired): `intelligent-personnel-query-the-list-of-employees-to-be-hired`
- [获取微应用后台免登的access_token](https://open.dingtalk.com/document/development/obtain-the-ssotoken-for-micro-application-background-logon-free): `obtain-the-ssotoken-for-micro-application-background-logon-free`
- [获取报表假期数据](https://open.dingtalk.com/document/development/obtains-the-holiday-data-from-the-smart-attendance-report): `obtains-the-holiday-data-from-the-smart-attendance-report`
- [获取指定用户可见的审批表单列表](https://open.dingtalk.com/document/development/you-can-call-this-operation-to-retrieve-a-list-of): `you-can-call-this-operation-to-retrieve-a-list-of`
- [获取指定用户的所有父部门列表](https://open.dingtalk.com/document/development/queries-the-list-of-all-parent-departments-of-a-user): `queries-the-list-of-all-parent-departments-of-a-user`
- [获取指定角色的员工列表](https://open.dingtalk.com/document/development/obtain-the-list-of-employees-of-a-role): `obtain-the-list-of-employees-of-a-role`
- [获取指定部门的所有父部门列表](https://open.dingtalk.com/document/development/queries-all-parent-departments-of-a-department): `queries-all-parent-departments-of-a-department`
- [获取指定部门的所有父部门列表](https://open.dingtalk.com/document/development/query-the-list-of-all-parent-departments-of-a-department): `query-the-list-of-all-parent-departments-of-a-department`
- [获取授权应用的基本信息](https://open.dingtalk.com/document/development/obtains-application-information-of-an-enterprise): `obtains-application-information-of-an-enterprise`
- [获取数字化证书](https://open.dingtalk.com/document/development/obtain-digital-certificate): `obtain-digital-certificate`
- [获取数字区县组织信息](https://open.dingtalk.com/document/development/querydigitaldistrictorginfo-api-reference): `querydigitaldistrictorginfo-api-reference`
- [获取文件上传信息](https://open.dingtalk.com/document/development/obtain-upload-information): `obtain-upload-information`
- [获取文件下载信息](https://open.dingtalk.com/document/development/obtain-download-file-info): `obtain-download-file-info`
- [获取文章详情](https://open.dingtalk.com/document/development/get-article): `get-article`
- [获取日志接收人员列表](https://open.dingtalk.com/document/development/queries-log-sharing-personnel): `queries-log-sharing-personnel`
- [获取日志相关人员列表](https://open.dingtalk.com/document/development/obtains-a-list-of-log-related-personnel-by-type): `obtains-a-list-of-log-related-personnel-by-type`
- [获取日志统计数据](https://open.dingtalk.com/document/development/query-log-statistics): `query-log-statistics`
- [获取服务窗联系人信息](https://open.dingtalk.com/document/development/obtains-the-contact-information-of-the-service-window): `obtains-the-contact-information-of-the-service-window`
- [获取未活跃用户登录明细](https://open.dingtalk.com/document/development/obtains-the-logon-details-of-inactive-users): `obtains-the-logon-details-of-inactive-users`
- [获取未登录钉钉的员工列表](https://open.dingtalk.com/document/development/queries-the-inactive-users-or-active-users-under-an-enterprise): `queries-the-inactive-users-or-active-users-under-an-enterprise`
- [获取未登录钉钉的员工列表](https://open.dingtalk.com/document/development/query-data-of-inactive-users): `query-data-of-inactive-users`
- [获取权限列表](https://open.dingtalk.com/document/development/obtain-a-permission-list): `obtain-a-permission-list`
- [获取模板code](https://open.dingtalk.com/document/development/obtains-the-template-code-based-on-the-template-name): `obtains-the-template-code-based-on-the-template-name`
- [获取模板详情](https://open.dingtalk.com/document/development/query-template-details): `query-template-details`
- [获取流程定义](https://open.dingtalk.com/document/development/obtain-process-definition): `obtain-process-definition`
- [获取流程实例](https://open.dingtalk.com/document/development/obtain-process-instance): `obtain-process-instance`
- [获取流程节点按钮列表](https://open.dingtalk.com/document/development/obtain-a-list-of-process-node-buttons): `obtain-a-list-of-process-node-buttons`
- [获取流程设计的节点信息](https://open.dingtalk.com/document/development/obtain-the-information-about-the-nodes-in-process-design): `obtain-the-information-about-the-nodes-in-process-design`
- [获取班次摘要信息](https://open.dingtalk.com/document/development/enterprise-shift-query-in-batches): `enterprise-shift-query-in-batches`
- [获取班次详情](https://open.dingtalk.com/document/development/shift-query): `shift-query`
- [获取班级内学生的关系列表](https://open.dingtalk.com/document/development/queries-the-list-of-relationships): `queries-the-list-of-relationships`
- [获取班级圈动态列表](https://open.dingtalk.com/document/development/dynamic-list-opening-of-class-circle): `dynamic-list-opening-of-class-circle`
- [获取班级圈话题列表](https://open.dingtalk.com/document/development/obtain-a-topic-list-of-class-circles): `obtain-a-topic-list-of-class-circles`
- [获取用户创建文档数和创建文档人数（组织维度）](https://open.dingtalk.com/document/development/queries-the-number-dingtalk-documents-created-per-day-in-an): `queries-the-number-dingtalk-documents-created-per-day-in-an`
- [获取用户创建的AI助理列表](https://open.dingtalk.com/document/development/assistant-acquires-list-assistants-created-by-users): `assistant-acquires-list-assistants-created-by-users`
- [获取用户创建的填表模板](https://open.dingtalk.com/document/development/obtains-the-template-that-a-user-creates): `obtains-the-template-that-a-user-creates`
- [获取用户发送日志的概要信息](https://open.dingtalk.com/document/development/view-log-summary-data): `view-log-summary-data`
- [获取用户可查看的公告](https://open.dingtalk.com/document/development/list-the-user-s-announcement-list): `list-the-user-s-announcement-list`
- [获取用户可见的企业应用列表](https://open.dingtalk.com/document/development/list-the-microapplications-visible-to-employees): `list-the-microapplications-visible-to-employees`
- [获取用户可见范围的AI助理列表](https://open.dingtalk.com/document/development/api-listvisibleassistant): `api-listvisibleassistant`
- [获取用户基本信息](https://open.dingtalk.com/document/development/queries-basic-user-information): `queries-basic-user-information`
- [获取用户待审批数量](https://open.dingtalk.com/document/development/obtain-the-number-of-tasks-to-be-approved-by-me): `obtain-the-number-of-tasks-to-be-approved-by-me`
- [获取用户授权的持久授权码](https://open.dingtalk.com/document/development/persistent-authorization-code): `persistent-authorization-code`
- [获取用户日志未读数](https://open.dingtalk.com/document/development/querying-the-employee-s-log-is-not-reading): `querying-the-employee-s-log-is-not-reading`
- [获取用户月活跃数据](https://open.dingtalk.com/document/development/retrieves-the-user-s-monthly-active-data): `retrieves-the-user-s-monthly-active-data`
- [获取用户版本分布情况](https://open.dingtalk.com/document/development/queries-the-distribution-of-user-versions): `queries-the-distribution-of-user-versions`
- [获取用户签到记录](https://open.dingtalk.com/document/development/obtain-the-check-in-records-of-multiple-users): `obtain-the-check-in-records-of-multiple-users`
- [获取用户考勤组](https://open.dingtalk.com/document/development/queries-a-user-attendance-group): `queries-a-user-attendance-group`
- [获取申请单列表](https://open.dingtalk.com/document/development/search-enterprise-approval-form-data): `search-enterprise-approval-form-data`
- [获取申请单详情](https://open.dingtalk.com/document/development/obtains-the-detailed-data-of-a-single-request): `obtains-the-detailed-data-of-a-single-request`
- [获取离职员工列表](https://open.dingtalk.com/document/development/intelligent-personnel-query-company-turnover-list): `intelligent-personnel-query-company-turnover-list`
- [获取空间信息](https://open.dingtalk.com/document/development/obtain-space-information): `obtain-space-information`
- [获取第三方个人应用的access_token](https://open.dingtalk.com/document/development/obtain-personal-application): `obtain-personal-application`
- [获取管理员列表](https://open.dingtalk.com/document/development/obtains-a-list-of-administrators): `obtains-a-list-of-administrators`
- [获取管理员列表](https://open.dingtalk.com/document/development/query-the-administrator-list): `query-the-administrator-list`
- [获取管理员的应用管理权限](https://open.dingtalk.com/document/development/obtains-the-administrator-s-microapplication-management-permission): `obtains-the-administrator-s-microapplication-management-permission`
- [获取管理员通讯录权限范围](https://open.dingtalk.com/document/development/query-permissions-of-the-administrator-address-book): `query-permissions-of-the-administrator-address-book`
- [获取考勤报表列值](https://open.dingtalk.com/document/development/queries-the-column-value-of-the-attendance-report): `queries-the-column-value-of-the-attendance-report`
- [获取考勤报表列定义](https://open.dingtalk.com/document/development/queries-the-enterprise-attendance-report-column): `queries-the-enterprise-attendance-report-column`
- [获取考勤组详情](https://open.dingtalk.com/document/development/query-a-single-attendance-group): `query-a-single-attendance-group`
- [获取联系人的元数据](https://open.dingtalk.com/document/development/gets-the-metadata-description-of-a-crm-contact-object): `gets-the-metadata-description-of-a-crm-contact-object`
- [获取自定义对象的元数据](https://open.dingtalk.com/document/development/get-metadata-description-of-crm-custom-object): `get-metadata-description-of-crm-custom-object`
- [获取花名册元数据](https://open.dingtalk.com/document/development/intelligent-personnel-roster-metadata-query): `intelligent-personnel-roster-metadata-query`
- [获取花名册字段组详情](https://open.dingtalk.com/document/development/get-roster-field-group-details): `get-roster-field-group-details`
- [获取视频会议详情](https://open.dingtalk.com/document/development/get-details-of-the-video-conference): `get-details-of-the-video-conference`
- [获取视频直播明细列表](https://open.dingtalk.com/document/development/queries-the-details-list-of-apsaravideo-live): `queries-the-details-list-of-apsaravideo-live`
- [获取视频直播观看人员列表](https://open.dingtalk.com/document/development/query-users-of-apsaravideo-live): `query-users-of-apsaravideo-live`
- [获取角色列表](https://open.dingtalk.com/document/development/obtains-a-list-of-enterprise-roles): `obtains-a-list-of-enterprise-roles`
- [获取角色组列表](https://open.dingtalk.com/document/development/obtains-the-role-group-information): `obtains-the-role-group-information`
- [获取角色详情](https://open.dingtalk.com/document/development/queries-role-details): `queries-role-details`
- [获取课堂明细数据](https://open.dingtalk.com/document/development/obtain-course-detail-data): `obtain-course-detail-data`
- [获取课堂概要数据](https://open.dingtalk.com/document/development/get-course-summary-data): `get-course-summary-data`
- [获取课程列表](https://open.dingtalk.com/document/development/get-course-list): `get-course-list`
- [获取课程参与方列表](https://open.dingtalk.com/document/development/get-a-list-of-course-participants): `get-a-list-of-course-participants`
- [获取课程详情](https://open.dingtalk.com/document/development/get-course-details): `get-course-details`
- [获取跟进记录对象的元数据](https://open.dingtalk.com/document/development/obtains-the-metadata-description-of-the-crm-follow-up-record-object): `obtains-the-metadata-description-of-the-crm-follow-up-record-object`
- [获取通讯录权限范围](https://open.dingtalk.com/document/development/obtain-corpsecret-authorization-scope): `obtain-corpsecret-authorization-scope`
- [获取部门下人员列表](https://open.dingtalk.com/document/development/obtains-the-list-of-people-under-a-department): `obtains-the-list-of-people-under-a-department`
- [获取部门企业账号用户详情](https://open.dingtalk.com/document/development/queries-account-details): `queries-account-details`
- [获取部门列表](https://open.dingtalk.com/document/development/obtain-the-department-list): `obtain-the-department-list`
- [获取部门列表](https://open.dingtalk.com/document/development/obtains-a-list-of-industry-departments): `obtains-a-list-of-industry-departments`
- [获取部门列表](https://open.dingtalk.com/document/development/obtains-the-department-node-list): `obtains-the-department-node-list`
- [获取部门列表](https://open.dingtalk.com/document/development/user-management-acquires-the-list-departments): `user-management-acquires-the-list-departments`
- [获取部门扩展信息](https://open.dingtalk.com/document/development/obtain-department-extension-information): `obtain-department-extension-information`
- [获取部门用户userid列表](https://open.dingtalk.com/document/development/obtain-the-list-of-employee-ids-by-department-id): `obtain-the-list-of-employee-ids-by-department-id`
- [获取部门用户userid列表](https://open.dingtalk.com/document/development/query-the-list-of-department-userids): `query-the-list-of-department-userids`
- [获取部门用户基础信息](https://open.dingtalk.com/document/development/obtain-the-basic-information-of-department-users): `obtain-the-basic-information-of-department-users`
- [获取部门用户签到记录](https://open.dingtalk.com/document/development/get-check-in-data): `get-check-in-data`
- [获取部门用户详情](https://open.dingtalk.com/document/development/obtain-department-members-details): `obtain-department-members-details`
- [获取部门用户详情](https://open.dingtalk.com/document/development/queries-department-user-details): `queries-department-user-details`
- [获取部门用户详情](https://open.dingtalk.com/document/development/queries-the-complete-information-of-a-department-user): `queries-the-complete-information-of-a-department-user`
- [获取部门的扩展字段定义](https://open.dingtalk.com/document/development/gets-the-extended-field-definition-of-a-department): `gets-the-extended-field-definition-of-a-department`
- [获取部门详情](https://open.dingtalk.com/document/development/industry-address-book-api-for-obtaining-department-information): `industry-address-book-api-for-obtaining-department-information`
- [获取部门详情](https://open.dingtalk.com/document/development/obtains-queries-department-details): `obtains-queries-department-details`
- [获取部门详情](https://open.dingtalk.com/document/development/queries-department-details): `queries-department-details`
- [获取部门详情](https://open.dingtalk.com/document/development/query-department-details0-v2): `query-department-details0-v2`
- [获取问答明细](https://open.dingtalk.com/document/development/api-getaskdetail): `api-getaskdetail`
- [获得企业创建日志相关信息（组织维度）](https://open.dingtalk.com/document/development/obtains-information-about-a-created-enterprise-log-from-the-organization): `obtains-information-about-a-created-enterprise-log-from-the-organization`
- [获得企业创建日志相关信息（部门维度）](https://open.dingtalk.com/document/development/obtains-information-about-a-created-enterprise-log-from-the-department): `obtains-information-about-a-created-enterprise-log-from-the-department`
- [获得用户创建文档数和创建文档人数（部门维度）](https://open.dingtalk.com/document/development/obtain-departmental-dimensions-documents-created-people-creating): `obtain-departmental-dimensions-documents-created-people-creating`
- [获得组织维度日程相关信息](https://open.dingtalk.com/document/development/queries-the-number-of-people-who-have-created-an-event): `queries-the-number-of-people-who-have-created-an-event`
- [解绑设备](https://open.dingtalk.com/document/development/unbind-a-smart-hardware-device): `unbind-a-smart-hardware-device`
- [计算请假时长](https://open.dingtalk.com/document/development/calculate-leave-duration): `calculate-leave-duration`
- [设备入会](https://open.dingtalk.com/document/development/equipment-membership): `equipment-membership`
- [设备账号向目标用户发送DING消息](https://open.dingtalk.com/document/development/device-publishing): `device-publishing`
- [设定角色成员管理范围](https://open.dingtalk.com/document/development/update-role-member-management-department-scope): `update-role-member-management-department-scope`
- [设置禁止群成员私聊](https://open.dingtalk.com/document/development/set-private-chat): `set-private-chat`
- [还原回收站文件（夹）](https://open.dingtalk.com/document/development/restore-recycle-bin-files-folder): `restore-recycle-bin-files-folder`
- [通知审批撤销](https://open.dingtalk.com/document/development/notify-the-attendance-to-modify-the-punch-result-when-the): `notify-the-attendance-to-modify-the-punch-result-when-the`
- [通知审批通过](https://open.dingtalk.com/document/development/notice-of-approval): `notice-of-approval`
- [通知换班通过](https://open.dingtalk.com/document/development/shift-change-operation-after-approval): `shift-change-operation-after-approval`
- [通知授权结果](https://open.dingtalk.com/document/development/notify-the-authorization-result): `notify-the-authorization-result`
- [通知补卡通过](https://open.dingtalk.com/document/development/make-up-the-card-after-approval): `make-up-the-card-after-approval`
- [通过免登码获取用户信息](https://open.dingtalk.com/document/development/obtain-the-userid-of-a-user-by-using-the-log-free): `obtain-the-userid-of-a-user-by-using-the-log-free`
- [通过免登码获取用户信息（不推荐）](https://open.dingtalk.com/document/development/get-user-userid-through-login-free-code): `get-user-userid-through-login-free-code`
- [通过调用者unionId获取激活码](https://open.dingtalk.com/document/development/obtain-the-activation-code-by-calling-the-union-id-of): `obtain-the-activation-code-by-calling-the-union-id-of`
- [通过高级查询条件获取表单实例数据（包括子表单组件数据）](https://open.dingtalk.com/document/development/query-form-instances-using-advanced-search-conditions): `query-form-instances-using-advanced-search-conditions`
- [邀请其他组织企业账号加入](https://open.dingtalk.com/document/development/invite-other-organization-specific-accounts-to-join): `invite-other-organization-specific-accounts-to-join`
- [配置发票适用人群](https://open.dingtalk.com/document/development/configure-invoice-users): `configure-invoice-users`
- [重新授权未激活应用的企业](https://open.dingtalk.com/document/development/re-authorize-enterprises-whose-applications-are-not-activated): `re-authorize-enterprises-whose-applications-are-not-activated`
- [针对单个日程进行签到](https://open.dingtalk.com/document/development/sign-in-for-a-single-schedule): `sign-in-for-a-single-schedule`
- [钉钉文本翻译](https://open.dingtalk.com/document/development/dingtalk-translation): `dingtalk-translation`
- [静态推荐数据同步](https://open.dingtalk.com/document/development/statically-recommended-data-synchronization): `statically-recommended-data-synchronization`
- [预计算时长](https://open.dingtalk.com/document/development/calculate-duration-based-on-attendance-scheduling): `calculate-duration-based-on-attendance-scheduling`
- [验证激活结果](https://open.dingtalk.com/document/development/verify-activation-results): `verify-activation-results`

> 已知 Gap：7 条 slug 在 development 段为空壳（<50000B），可能属别的 section（如 solution），待补正确 section，本目录暂不收录。

<!-- END GENERATED: capability_domains -->

## OpenAPI Specs

钉钉对外暂无机器可读 openapi.json：`TODO(gap #7)`，待平台侧提供，不伪造。
