package shortcut

import (
	"go/ast"
)

func NewFieldList(values ...string) *ast.FieldList {
	fields := []*ast.Field{}

	for _, value := range values {
		fields = append(fields, &ast.Field{
			Type: NewIdent(value),
		})
	}

	return &ast.FieldList{
		List: fields,
	}
}
