package container

import (
	"go/ast"
	"log"

	"golang.org/x/tools/go/ast/astutil"

	. "github.com/gooff/di/container/utilsast"
)

func (b *Builder) astGetMethods(tree Tree, paths map[string][]string) []ast.Decl {
	var methods []ast.Decl
	for name, def := range b.definitions {
		log.Println("Getter method for", name)

		for packageName, shortName := range def.imports() {
			astutil.AddNamedImport(b.fset, b.file, shortName, packageName)
		}

		methodPrefix := "get"
		if def.isPublic() {
			methodPrefix = "Get"
		}

		methods = append(methods, &ast.FuncDecl{
			Name: NewIdent(methodPrefix + methodName(paths[name])),
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							NewIdent("container"),
						},
						Type: NewIdent("*Container"),
					},
				},
			},
			Type: &ast.FuncType{
				Params:  def.astArguments(),
				Results: NewFieldList(def.interfaceOrLocalEntityType(b, false)),
			},
			Body: def.astFunctionBody(b.fset, b.file, b, name, name),
		})
	}
	return methods
}
