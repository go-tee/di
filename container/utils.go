package container

import (
	"strings"
	"unicode"

	"github.com/go-tee/di/utils"
)

func isPublicProperty(name string) bool {
	return unicode.IsUpper(utils.FirstRune(name))
}

func methodName(path []string) string {
	var upper []string
	for _, name := range path {
		upper = append(upper, utils.FirstUpper(name))
	}
	return strings.Join(upper, "")
}
