package container

import (
	"go/ast"
	"go/token"

	. "github.com/gooff/di/container/utilsast"
)

func (b *Builder) astNewContainerFunc() *ast.FuncDecl {
	containerTempVariable := "container"
	statements := []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: []ast.Expr{NewIdent(containerTempVariable)},
			Rhs: []ast.Expr{
				&ast.CompositeLit{
					Type: NewIdent("&Container"),
				},
			},
		},
	}

	for name, def := range b.definitions {
		if def.defScope.isPrototype() {
			statements = append(statements,
				&ast.AssignStmt{
					Tok: token.ASSIGN,
					Lhs: []ast.Expr{NewIdent(containerTempVariable + ".services." + name)},
					Rhs: []ast.Expr{
						&ast.FuncLit{
							Type: def.astFunctionPrototype(b),
							Body: def.astFunctionBody(b.fset, b.file, b, "", name),
						},
					},
				},
			)
		}
	}

	statements = append(statements, NewReturn(NewIdent(containerTempVariable)))

	return NewFunc("NewContainer", nil, []string{"*Container"}, NewBlock(statements...))
}
