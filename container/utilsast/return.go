package utilsast

import (
	"go/ast"
)

func NewReturn(expressions ...ast.Expr) *ast.ReturnStmt {
	var results []ast.Expr

	for _, expr := range expressions {
		results = append(results, expr)
	}

	return &ast.ReturnStmt{
		Results: results,
	}
}
