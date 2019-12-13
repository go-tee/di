package container

import (
	"strings"
)

type Tree map[string]interface{}

func prepareTree(definitions map[string]*Definition) (Tree, map[string][]string) {
	tree := Tree{}

	paths := map[string][]string{}
	for name, def := range definitions {
		path := strings.Split(name, ".")
		prepareSubTree(tree, path[1:], path[0], def)
		paths[name] = path
	}

	return tree, paths
}

func prepareSubTree(tree Tree, path []string, name string, def *Definition) {
	if len(path) == 0 {
		tree[name] = def
	} else {
		if _, ok := tree[name]; !ok {
			tree[name] = Tree{}
		}
		if subtree, ok := tree[name].(Tree); ok {
			prepareSubTree(subtree, path[1:], path[0], def)
		}
	}
}
