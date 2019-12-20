package container

import (
	"go/ast"
	"go/token"

	"github.com/go-tee/di/utils/shortcut"
)

func (b *Builder) astNewContainerFunc(paths map[string][]string) *ast.FuncDecl {
	containerTempVariable := "container"
	statements := []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: []ast.Expr{shortcut.NewIdent(containerTempVariable)},
			Rhs: []ast.Expr{
				&ast.CompositeLit{
					Type: shortcut.NewIdent("&Container"),
				},
			},
		},
	}

	for name, def := range b.definitions {
		if def.defScope.isPrototype() {
			statements = append(statements,
				&ast.AssignStmt{
					Tok: token.ASSIGN,
					Lhs: []ast.Expr{shortcut.NewIdent(containerTempVariable + ".services." + name)},
					Rhs: []ast.Expr{
						&ast.FuncLit{
							Type: def.astFunctionPrototype(b),
							Body: def.astFunctionBody(b.fset, b.file, b, "", name, paths),
						},
					},
				},
			)
		}
	}

	statements = append(statements, shortcut.NewReturn(shortcut.NewIdent(containerTempVariable)))

	return shortcut.NewFunc("NewContainer", nil, []string{"*Container"}, shortcut.NewBlock(statements...))
}
