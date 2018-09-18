package util

import "strings"

// TrimDoubleQuotes は、ダブルクォーテーションを削除する。
func TrimDoubleQuotes(target string) string {
	return strings.Trim(target, "\"")
}
