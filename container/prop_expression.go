package container

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"strings"

	"github.com/elliotchance/pie/pie"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/go-tee/di/utils/shortcut"
)

type defExpression string

func (e defExpression) dependencyNames() (deps []string) {
	for _, v := range regexp.MustCompile(`@{(.*?)}`).FindAllStringSubmatch(string(e), -1) {
		parts := strings.Split(v[1], "(")
		deps = append(deps, parts[0])
	}

	return pie.Strings(deps).Unique()
}

func (e defExpression) dependencies() (deps []string) {
	for _, v := range regexp.MustCompile(`@{(.*?)}`).FindAllStringSubmatch(string(e), -1) {
		deps = append(deps, v[1])
	}

	return pie.Strings(deps).Unique()
}

func (e defExpression) performSubstitutions(fset *token.FileSet, file *ast.File, builder *Builder, fromArgs bool) string {
	stmt := string(e)

	// Replace environment variables.
	stmt = shortcut.ReplaceAllStringSubmatchFunc(
		regexp.MustCompile(`\${(.*?)}`), stmt, func(i []string) string {
			astutil.AddImport(fset, file, "os")

			return fmt.Sprintf("os.Getenv(\"%s\")", i[1])
		})

	// Replace service names.
	stmt = shortcut.ReplaceAllStringSubmatchFunc(
		regexp.MustCompile(`@{(.*?)}`), stmt, func(match []string) string {
			if fromArgs {
				return strings.Split(match[1], "(")[0]
			}

			methodPrefix := "get"
			if isPublicProperty(match[1]) {
				methodPrefix = "Get"
			}

			if strings.Contains(match[1], "(") {
				return fmt.Sprintf("container.%s%s", methodPrefix, firstUpper(match[1]))
			}

			if !builder.HasDefinition(match[1]) {
				panic(fmt.Sprintf("service does not exist: %s", match[1]))
			}

			if _, ok := builder.GetDefinition(match[1]).astContainerFieldType(builder).(*ast.FuncType); ok {
				return fmt.Sprintf("container.%s", match[1])
			}

			return fmt.Sprintf("container.%s%s()", methodPrefix, firstUpper(match[1]))
		})

	return stmt
}
