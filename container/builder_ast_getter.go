package container

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/go-tee/di/utils/shortcut"
)

func (b *Builder) astGetMethods(tree Tree, paths map[string][]string) []ast.Decl {
	var methods []ast.Decl
	for name, def := range b.definitions {

		for packageName, shortName := range def.imports() {
			astutil.AddNamedImport(b.fset, b.file, shortName, packageName)
		}

		methodPrefix := "get"
		if def.isPublic() {
			methodPrefix = "Get"
		}

		methods = append(methods, &ast.FuncDecl{
			Name: shortcut.NewIdent(methodPrefix + methodName(paths[name])),
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							shortcut.NewIdent("container"),
						},
						Type: shortcut.NewIdent("*Container"),
					},
				},
			},
			Type: &ast.FuncType{
				Params:  def.astArguments(),
				Results: shortcut.NewFieldList(def.interfaceOrLocalEntityType(b, false)),
			},
			Body: def.astFunctionBody(b.fset, b.file, b, name, name, paths),
		})
	}
	return methods
}
