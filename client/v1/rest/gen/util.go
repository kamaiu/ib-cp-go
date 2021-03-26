package gen

import "strings"

func toGoFieldName(s string) string {
	s = capitalize(s)
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, "+", "_plus_")
	s = strings.ReplaceAll(s, "-", "_minus_")
	return s
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[0] == '_' {
		s = s[1:]
	}
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
