package utilsast

import (
	"go/ast"
	"go/token"
)

func NewVar(name string, value string) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{
					{Name: name},
				},
				Values: []ast.Expr{
					NewIdent(value),
				},
			},
		},
	}
}
