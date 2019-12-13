package container

import (
	"fmt"
	"regexp"
	"strings"
)

type defType string

func (t defType) string() string {
	if t.isFunction() {
		args, returns := t.parseFunctionType()
		switch len(returns) {
		case 0:
			return fmt.Sprintf("func (%s)", args)
		case 1:
			return fmt.Sprintf("func (%s) %s", args, strings.Join(returns, ","))
		default:
			return fmt.Sprintf("func (%s) (%s)", args, strings.Join(returns, ","))
		}
	}

	return string(t)
}

func (t defType) isFunction() bool {
	return strings.HasPrefix(string(t), "func")
}

func (t defType) isPointer() bool {
	return strings.HasPrefix(string(t), "*") || t.isFunction()
}

func (t defType) packageName() string {
	if t.isFunction() || !strings.Contains(string(t), ".") {
		return ""
	}

	parts := strings.Split(strings.TrimLeft(string(t), "*"), ".")
	return strings.Join(parts[:len(parts)-1], ".")
}

func (t defType) unversionedPackageName() string {
	if t.isFunction() {
		return ""
	}

	packageName := strings.Split(t.packageName(), "/")
	if regexp.MustCompile(`^v\d+$`).MatchString(packageName[len(packageName)-1]) {
		packageName = packageName[:len(packageName)-1]
	}

	return strings.Join(packageName, "/")
}

func (t defType) localPackageName() string {
	if t.isFunction() {
		return ""
	}

	pkgNameParts := strings.Split(t.unversionedPackageName(), "/")
	lastPart := pkgNameParts[len(pkgNameParts)-1]

	return strings.Replace(lastPart, "-", "_", -1)
}

func (t defType) entityName() string {
	if t.isFunction() {
		return t.string()
	}

	parts := strings.Split(string(t), ".")

	return strings.TrimLeft(parts[len(parts)-1], "*")
}

func (t defType) localEntityName() string {
	if t.isFunction() {
		return t.string()
	}

	name := t.localPackageName() + "." + t.entityName()

	return strings.TrimLeft(name, ".")
}

func (t defType) localEntityType() string {
	if t.isFunction() {
		return t.string()
	}

	name := t.localEntityName()
	if t.isPointer() {
		name = "*" + name
	}

	return name
}

func (t defType) createLocalEntityType() string {
	if t.isFunction() {
		return t.string()
	}

	name := t.localEntityName()
	if t.isPointer() {
		name = "&" + name
	}

	return name
}

func (t defType) localEntityPointerType() string {
	if t.isFunction() {
		return t.string()
	}

	name := t.localEntityName()
	if !strings.HasPrefix(name, "*") {
		name = "*" + name
	}

	return name
}

var (
	functionRegexp1 = regexp.MustCompile(`func\s*\((.*?)\)\s*\((.*)\)`)
	functionRegexp2 = regexp.MustCompile(`func\s*\((.*?)\)\s*(.*)`)
)

func (t defType) splitArgs(s string) []string {
	if s == "" {
		return nil
	}

	return strings.Split(s, ",")
}

func (t defType) parseFunctionType() (string, []string) {
	matches := functionRegexp1.FindStringSubmatch(string(t))
	if len(matches) > 0 {
		return matches[1], t.splitArgs(matches[2])
	}

	matches = functionRegexp2.FindStringSubmatch(string(t))

	return matches[1], t.splitArgs(matches[2])
}
