package container

import (
	"go/ast"

	"github.com/go-tee/di/utils/shortcut"
)

func (b *Builder) astBlocks(tree Tree) []ast.Decl {
	return append([]ast.Decl{
		b.astContainerStruct(),
		b.astParameterBlockStruct(),
	}, b.astTreeBlockStructs(tree, "")...)
}

func (b *Builder) astContainerStruct() ast.Decl {
	return shortcut.NewStruct("Container", []*ast.Field{
		{
			Names: []*ast.Ident{{Name: "parameters"}},
			Type:  shortcut.NewIdent("ParametersBlock"),
		},
		{
			Names: []*ast.Ident{{Name: "services"}},
			Type:  shortcut.NewIdent("ServicesBlock"),
		},
	})
}

func (b *Builder) astParameterBlockStruct() ast.Decl {
	return shortcut.NewStruct("ParametersBlock", []*ast.Field{})
}

func (b *Builder) astTreeBlockStructs(tree Tree, prefix string) []ast.Decl {
	var structs []ast.Decl
	var fields []*ast.Field
	for name, dt := range tree {
		if def, ok := dt.(*Definition); ok {
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{{Name: name}},
				Type: def.astContainerFieldType(b),
			})
		}
		if subtree, ok := dt.(Tree); ok {
			blockName := firstUpper(name)
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{{Name: name}},
				Type:  shortcut.NewIdent(prefix + blockName + "ServicesBlock"),
			})
			structs = append(structs, b.astTreeBlockStructs(subtree, prefix + blockName)...)
		}
	}

	return append([]ast.Decl{shortcut.NewStruct(prefix + "ServicesBlock", fields)}, structs...)
}
