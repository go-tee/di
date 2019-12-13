package container

import (
	"go/ast"

	. "github.com/gooff/di/container/utilsast"
)

func (b *Builder) astBlocks(tree Tree) []ast.Decl {
	return append([]ast.Decl{
		b.astContainerStruct(),
		b.astParameterBlockStruct(),
	}, b.astTreeBlockStructs(tree, "")...)
}

func (b *Builder) astContainerStruct() ast.Decl {
	return NewStruct("Container", []*ast.Field{
		{
			Names: []*ast.Ident{{Name: "parameters"}},
			Type:  NewIdent("ParametersBlock"),
		},
		{
			Names: []*ast.Ident{{Name: "services"}},
			Type:  NewIdent("ServicesBlock"),
		},
	})
}

func (b *Builder) astParameterBlockStruct() ast.Decl {
	return NewStruct("ParametersBlock", []*ast.Field{})
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
				Type: NewIdent(prefix + blockName + "ServicesBlock"),
			})
			structs = append(structs, b.astTreeBlockStructs(subtree, prefix + blockName)...)
		}
	}

	return append([]ast.Decl{NewStruct(prefix + "ServicesBlock", fields)}, structs...)
}






// func (b *Builder) astContainerStruct_2() ast.Decl {
// 	var containerFields []*ast.Field
// 	tree, _ := prepareTree(b.definitions)
//
// 	for name, dt := range tree {
// 		containerFields = append(containerFields, &ast.Field{
// 			Names: []*ast.Ident{{Name: name}},
// 			Type:  b.prepareContainerType(dt),
// 		})
// 	}
//
// 	// for name, def := range b.definitions {
// 	// 	log.Println("Create container field", name)
// 	// 	containerFields = append(containerFields, &ast.Field{
// 	// 		Names: []*ast.Ident{
// 	// 			NewIdent(strings.Replace(name, ".", "_", -1)),
// 	// 		},
// 	// 		Type: def.astContainerFieldType(b),
// 	// 	})
// 	// }
//
// 	return &ast.GenDecl{
// 		Tok: token.TYPE,
// 		Specs: []ast.Spec{
// 			&ast.TypeSpec{
// 				Name: NewIdent("Container"),
// 				Type: &ast.StructType{
// 					Fields: &ast.FieldList{
// 						List: containerFields,
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func (b *Builder) astContainerStruct() ast.Decl {
// 	tree, _ := prepareTree(b.definitions)
//
// 	var containerFields []*ast.Field

//
// 	// for _, serviceName := range services.ServiceNames() {
// 	// 	service := services[serviceName]
// 	//
// 	// 	containerFields = append(containerFields, &ast.Field{
// 	// 		Names: []*ast.Ident{
// 	// 			{Name: serviceName},
// 	// 		},
// 	// 		Type: service.ContainerFieldType(services),
// 	// 	})
// 	// }
//
// 	return &ast.GenDecl{
// 		Tok: token.TYPE,
// 		Specs: []ast.Spec{
// 			&ast.TypeSpec{
// 				Name: NewIdent("Container"),
// 				Type: &ast.StructType{
// 					Fields: &ast.FieldList{
// 						List: containerFields,
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func (b *Builder) prepareContainerType(dt interface{}) ast.Expr {
// 	if tree, ok := dt.(Tree); ok {
// 		var fields []*ast.Field
// 		for name, dt := range tree {
// 			fields = append(fields, &ast.Field{
// 				Names: []*ast.Ident{
// 					NewIdent(name),
// 				},
// 				Type: b.prepareContainerType(dt),
// 			})
// 		}
//
// 		return &ast.StructType{
// 			Fields: &ast.FieldList{
// 				List: fields,
// 			},
// 		}
// 	}
//
// 	return dt.(*Definition).astContainerFieldType(b)
// }
