package container

import (
	"fmt"
	"sort"
)

type defArguments map[string]defType

func (args defArguments) names() (names []string) {
	for arg := range args {
		names = append(names, arg)
	}

	sort.Strings(names)

	return
}

func (args defArguments) arguments() (ss []string) {
	for _, argName := range args.names() {
		ss = append(ss, fmt.Sprintf("%s %s", argName, args[argName].localEntityType()))
	}

	return
}
