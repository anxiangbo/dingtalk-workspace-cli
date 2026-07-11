// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cli

func init() {
	registerRequireOneOf("aitable.table_update", "name", "description", "record-name-key")
	registerRequireTogether("calendar.create_calendar_event", "recurrence-type", "recurrence-interval", "recurrence-range-type")
	registerRequireOneOf("calendar.query_busy_status", "users", "rooms")
	registerRequireTogether("calendar.update_calendar_event", "recurrence-type", "recurrence-interval", "recurrence-range-type")
	registerExclusiveOneOf("chat.search_messages_by_sender", "sender-user-id", "sender-open-dingtalk-id")
	registerExclusiveOneOf("chat.create_and_send_card", "group", "receiver")
	registerExclusiveOneOf("chat.get_conversation_info", "group", "user", "open-dingtalk-id")
	registerExclusiveOneOf("chat.list_conversation_message_v2", "group", "user", "open-dingtalk-id")
	registerExclusiveOneOf("chat.list_individual_chat_message", "user", "open-dingtalk-id")
	registerRequireOneOf("chat.send_personal_message", "text", "content", "msg-type")
	registerExclusiveOneOf("chat.send_robot_message", "group", "users")
	registerRequireOneOf("chat.set_group_member_mute_list", "users", "user")
	registerExclusiveOneOf("chat.transfer_group_owner", "new-owner", "user")
	registerRequireOneOf("chat.update_conv_member_roles", "users", "user")
	registerRequireTogether("contact.query_dismission_employee_list", "start", "end")
	registerRequireOneOf("dev.connect_status", "robot-client-id", "unified-app-id")
	registerRequireOneOf("dev.connect_stop", "robot-client-id", "unified-app-id")
	registerRequireOneOf("devdoc.search_open_error_code_rag", "query", "request-id", "error-code", "error-message", "context")
	registerRequireOneOf("doc.insert_document_block", "text", "heading", "element")
	registerExclusiveOneOf("doc.update_document", "content", "content-file")
	registerRequireOneOf("doc.update_document_block", "text", "heading", "element")
	registerRequireOneOf("pat.batch_grant", "scope", "product", "products", "domain", "domains", "recommend")
	registerRequireOneOf("report.create_report", "contents", "contents-file")
	RegisterRuntimeSchemaConstraints("sheet.range_set_style", RuntimeSchemaConstraints{
		MutuallyExclusive: [][]string{
			{"bg-color", "bg-colors-json"},
			{"font-size", "font-sizes-json"},
			{"h-align", "h-aligns-json"},
			{"v-align", "v-aligns-json"},
			{"font-color", "font-colors-json"},
			{"font-weight", "font-weights-json"},
		},
		RequireOneOf: [][]string{{
			"bg-color", "bg-colors-json", "font-size", "font-sizes-json",
			"h-align", "h-aligns-json", "v-align", "v-aligns-json",
			"font-color", "font-colors-json", "font-weight", "font-weights-json",
			"word-wrap", "number-format",
		}},
	})
	registerRequireOneOf("sheet.update_cond_format", "ranges", "condition", "cell-style", "data-bar-style")
	registerRequireOneOf("sheet.update_dimension", "hidden", "pixel-size")
	registerRequireOneOf("sheet.update_filter_view", "name", "range", "criteria")
	registerRequireOneOf("sheet.update_float_image", "src", "range", "width", "height", "offset-x", "offset-y")
	registerRequireOneOf("sheet.update_sheet", "name", "index", "hidden", "frozen-row-count", "frozen-column-count", "tab-color")
	registerRequireOneOf("wiki.search_wikiSpaces", "query", "type")
}

func registerRequireOneOf(canonicalPath string, names ...string) {
	RegisterRuntimeSchemaConstraints(canonicalPath, RuntimeSchemaConstraints{RequireOneOf: [][]string{names}})
}

func registerExclusiveOneOf(canonicalPath string, names ...string) {
	RegisterRuntimeSchemaConstraints(canonicalPath, RuntimeSchemaConstraints{
		MutuallyExclusive: [][]string{names},
		RequireOneOf:      [][]string{names},
	})
}

func registerRequireTogether(canonicalPath string, names ...string) {
	RegisterRuntimeSchemaConstraints(canonicalPath, RuntimeSchemaConstraints{RequireTogether: [][]string{names}})
}
