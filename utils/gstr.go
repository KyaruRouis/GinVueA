package utils

import "strings"

func Split(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

func Equal(a, b string) bool {
	return strings.EqualFold(a, b)
}
