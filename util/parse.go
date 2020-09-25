package util

import (
	"bytes"
	"strconv"
	"strings"
)

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

func RetrieveIntFromDisplayText(s string) (int, error) {
	var buf bytes.Buffer

	for _, r := range s {
		if r >= '0' && r <= '9' {
			buf.WriteRune(r)
		}
	}

	return strconv.Atoi(buf.String())
}
