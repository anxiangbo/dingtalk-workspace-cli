// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"
	"strings"
	"unicode"
)

func mcpJSONType(prop map[string]any) string {
	t, _ := prop["type"].(string)
	switch t {
	case "string", "integer", "number", "boolean", "array", "object":
		return t
	default:
		return "string"
	}
}

func mcpDefault(prop map[string]any) (string, bool) {
	if prop == nil {
		return "", false
	}
	v, ok := prop["default"]
	if !ok || v == nil {
		return "", false
	}
	switch value := v.(type) {
	case string:
		return value, true
	case float64:
		if value == float64(int64(value)) {
			return fmt.Sprintf("%d", int64(value)), true
		}
		return fmt.Sprintf("%v", value), true
	default:
		return fmt.Sprintf("%v", value), true
	}
}

func kebabCase(name string) string {
	runes := []rune(name)
	var b strings.Builder
	for i, r := range runes {
		if unicode.IsUpper(r) {
			prevLowerOrDigit := i > 0 && (unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1]))
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if i > 0 && (prevLowerOrDigit || nextLower) {
				b.WriteByte('-')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	out := strings.ReplaceAll(b.String(), "_", "-")
	for strings.Contains(out, "--") {
		out = strings.ReplaceAll(out, "--", "-")
	}
	return strings.Trim(out, "-")
}
