package util

import "strings"

func GetStringInBetween(str, start, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}

	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return ""
	}

	return str[s : s+e]
}
