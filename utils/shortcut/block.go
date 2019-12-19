package shortcut

import (
	"go/ast"
)

func NewBlock(stmts ...ast.Stmt) *ast.BlockStmt {
	return &ast.BlockStmt{
		List: stmts,
	}
}
