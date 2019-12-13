package utilsast

import (
	"go/ast"
)

func NewFunc(name string, params []string, returns []string, body *ast.BlockStmt) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: NewIdent(name),
		Type: &ast.FuncType{
			Params:  NewFieldList(params...),
			Results: NewFieldList(returns...),
		},
		Body: body,
	}
}
