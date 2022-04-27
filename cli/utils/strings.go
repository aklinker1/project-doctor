package utils

import "strings"

func RemoveFinalNewline(str string) string {
	if strings.HasSuffix(str, "\n") {
		return str[0:(len(str) - 1)]
	}
	return str
}
