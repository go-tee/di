package container

import (
	"fmt"
	"go/ast"
	"go/token"
	"sort"
	"strings"

	. "github.com/gooff/di/container/utilsast"
)

func (d *Definition) astContainerFieldType(builder *Builder) ast.Expr {
	if d.defScope.isContainer() && len(d.defArguments) == 0 {
		return NewIdent(d.interfaceOrLocalEntityPointerType())
	}

	return d.astFunctionPrototype(builder)
}

func (d *Definition) astFunctionPrototype(builder *Builder) *ast.FuncType {
	t := defType(d.interfaceOrLocalEntityType(builder, true))
	if t.isFunction() {
		args, returns := t.parseFunctionType()

		return &ast.FuncType{
			Params:  NewFieldList(args),
			Results: NewFieldList(returns...),
		}
	}

	return &ast.FuncType{
		Params:  d.astAllArguments(builder),
		Results: NewFieldList(string(t)),
	}
}

func (d *Definition) astAllArguments(builder *Builder) *ast.FieldList {
	deps := d.astDependencyArguments(builder)
	args := d.astArguments()

	return &ast.FieldList{
		List: append(deps.List, args.List...),
	}
}

func (d *Definition) astArguments() *ast.FieldList {
	funcParams := &ast.FieldList{
		List: []*ast.Field{},
	}

	for arg, t := range d.defArguments {
		funcParams.List = append(funcParams.List, &ast.Field{
			Type: &ast.Ident{
				Name: arg + " " + t.string(),
			},
		})
	}

	return funcParams
}

func (d *Definition) astDependencyArguments(builder *Builder) *ast.FieldList {
	funcParams := &ast.FieldList{
		List: []*ast.Field{},
	}

	for _, dep := range d.defReturns.dependencyNames() {
		funcParams.List = append(funcParams.List, &ast.Field{
			Type: NewIdent(dep + " " + builder.GetDefinition(dep).interfaceOrLocalEntityType(builder, false)),
		})
	}

	return funcParams
}

func (d *Definition) astFunctionBody(fset *token.FileSet, file *ast.File, builder *Builder, name, serviceName string) *ast.BlockStmt {
	if name != "" && !d.defScope.isContainer() {
		var arguments []string
		for _, dep := range d.defReturns.dependencies() {
			arguments = append(arguments, fmt.Sprintf("container.Get%s", dep))
		}
		arguments = append(arguments, d.defArguments.names()...)

		return NewBlock(
			NewReturn(NewIdent("container.services." + serviceName + "(" + strings.Join(arguments, ", ") + ")")),
		)
	}

	var stmts, instantiation []ast.Stmt
	serviceVariable := "container.services." + name
	serviceTempVariable := "service"

	// Instantiation
	if d.defReturns == "" {
		instantiation = []ast.Stmt{
			&ast.AssignStmt{
				Tok: token.DEFINE,
				Lhs: []ast.Expr{NewIdent(serviceTempVariable)},
				Rhs: []ast.Expr{
					&ast.CompositeLit{
						Type: NewIdent(d.defType.createLocalEntityType()),
					},
				},
			},
		}
	} else {
		lhs := []ast.Expr{NewIdent(serviceTempVariable)}

		if d.defError != "" {
			lhs = append(lhs, NewIdent("err"))
		}

		instantiation = []ast.Stmt{
			&ast.AssignStmt{
				Tok: token.DEFINE,
				Lhs: lhs,
				Rhs: []ast.Expr{
					NewIdent(d.defReturns.performSubstitutions(fset, file, builder, name == "")),
				},
			},
		}

		if d.defError != "" {
			instantiation = append(instantiation, &ast.IfStmt{
				Cond: NewIdent("err != nil"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: NewIdent(d.defError),
						},
					},
				},
			})
		}
	}

	// Properties
	for _, property := range d.sortedProperties() {
		instantiation = append(instantiation, &ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{&ast.Ident{Name: serviceTempVariable + "." + property.Name}},
			Rhs: []ast.Expr{&ast.Ident{Name: property.Value.performSubstitutions(fset, file, builder, name == "")}},
		})
	}

	// Scope
	if d.defScope.isContainer() {
		if d.defType.isPointer() || d.defInterface != "" {
			instantiation = append(instantiation, &ast.AssignStmt{
				Tok: token.ASSIGN,
				Lhs: []ast.Expr{&ast.Ident{Name: serviceVariable}},
				Rhs: []ast.Expr{&ast.Ident{Name: serviceTempVariable}},
			})
		} else {
			instantiation = append(instantiation, &ast.AssignStmt{
				Tok: token.ASSIGN,
				Lhs: []ast.Expr{&ast.Ident{Name: serviceVariable}},
				Rhs: []ast.Expr{&ast.Ident{Name: "&" + serviceTempVariable}},
			})
		}

		stmts = append(stmts, &ast.IfStmt{
			Cond: &ast.Ident{Name: serviceVariable + " == nil"},
			Body: &ast.BlockStmt{
				List: instantiation,
			},
		})

		// Returns
		if d.defType.isPointer() || d.defInterface != "" {
			stmts = append(stmts, NewReturn(NewIdent(serviceVariable)))
		} else {
			stmts = append(stmts, NewReturn(NewIdent("*"+serviceVariable)))
		}
	} else {
		stmts = append(stmts, instantiation...)
		stmts = append(stmts, NewReturn(NewIdent("service")))
	}

	return NewBlock(stmts...)
}

func (d *Definition) imports() map[string]string {
	imports := map[string]string{}

	for _, packageName := range d.defImports {
		imports[packageName] = ""
	}

	if d.defType.packageName() != "" {
		imports[d.defType.packageName()] = d.defType.localPackageName()
	}

	if d.defInterface.packageName() != "" {
		imports[d.defInterface.packageName()] = d.defInterface.localPackageName()
	}

	return imports
}

func (d *Definition) interfaceOrLocalEntityPointerType() string {
	if d.defInterface != "" {
		return d.defInterface.localEntityType()
	}

	return d.defType.localEntityPointerType()
}

func (d *Definition) interfaceOrLocalEntityType(builder *Builder, recurse bool) string {
	localEntityType := d.defType.localEntityType()
	if d.defInterface != "" {
		localEntityType = d.defInterface.localEntityType()
	}

	if len(d.defArguments) > 0 && recurse {
		var args []string

		for _, dep := range d.defReturns.dependencies() {
			ty := builder.GetDefinition(dep).interfaceOrLocalEntityType(builder, false)
			args = append(args, fmt.Sprintf("%s %s", dep, ty))
		}

		args = append(args, d.defArguments.arguments()...)

		return fmt.Sprintf("func(%v) %s", strings.Join(args, ", "), localEntityType)
	}

	return localEntityType
}

func (d *Definition) sortedProperties() (sortedProperties []*defProperty) {
	var propertyNames []string
	for propertyName := range d.defProperties {
		propertyNames = append(propertyNames, propertyName)
	}

	sort.Strings(propertyNames)

	for _, propertyName := range propertyNames {
		sortedProperties = append(sortedProperties, &defProperty{
			Name:  propertyName,
			Value: d.defProperties[propertyName],
		})
	}

	return
}
