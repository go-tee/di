package utilsast

import (
	"go/ast"
	"sort"
)

func NewCompositeLit(ty string, m map[string]ast.Expr) *ast.CompositeLit {
	var exprs []ast.Expr
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		exprs = append(exprs, &ast.KeyValueExpr{
			Key:   NewIdent(k),
			Value: m[k],
		})
	}

	return &ast.CompositeLit{
		Type: NewIdent(ty),
		Elts: exprs,
	}
}
