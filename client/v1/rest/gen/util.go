package gen

import "strings"

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
