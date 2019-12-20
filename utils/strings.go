package utils

import (
	"unicode"
)

func FirstRune(str string) (first rune) {
	for _, c := range str {
		first = c
		break
	}

	return
}

func FirstUpper(name string) string {
	runes := []rune(name)
	return string(unicode.ToUpper(runes[0])) + string(runes[1:])
}
