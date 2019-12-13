package container

import (
	"strings"
	"unicode"
)

func isPublicProperty(name string) bool {
	return unicode.IsUpper(firstRune(name))
}

func firstRune(str string) (first rune) {
	for _, c := range str {
		first = c
		break
	}

	return
}

func firstUpper(name string) string {
	runes := []rune(name)
	return string(unicode.ToUpper(runes[0])) + string(runes[1:])
}

func methodName(path []string) string {
	var upper []string
	for _, name := range path {
		upper = append(upper, firstUpper(name))
	}
	return strings.Join(upper, "")
}
