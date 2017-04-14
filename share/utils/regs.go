package utils

import "strings"

//FormatRegexStr replace regexp escape string
func FormatRegexpStr(str string) string {
	oldnews := []string{
		"\\", "\\\\",
		"*", "\\*",
		".", "\\.",
		"?", "\\?",
		"+", "\\+",
		"$", "\\$",
		"^", "\\^",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"{", "\\{",
		"}", "\\}",
		"|", "\\|",
		"/", "\\/",
	}
	return strings.NewReplacer(oldnews...).Replace(str)
}
