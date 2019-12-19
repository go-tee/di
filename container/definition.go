package container

import (
	"unicode"

	"github.com/go-tee/di/utils"
)

func NewServiceDefinition() *Definition {
	return &Definition{}
}

type Definition struct {
	defName       string
	defType       defType
	defInterface  defType
	defProperties map[string]defExpression
	defReturns    defExpression
	defArguments  defArguments
	defError      string
	defImports    []string
	defScope      defScope
	defSetup      []defStatement
}

func (d *Definition) GetName() string {
	return d.defName
}

func (d *Definition) SetName(n string) *Definition {
	d.defName = n
	return d
}

func (d *Definition) GetType() string {
	return string(d.defType)
}

func (d *Definition) SetType(t string) *Definition {
	d.defType = defType(t)
	return d
}

func (d *Definition) GetInterface() string {
	return string(d.defInterface)
}

func (d *Definition) SetInterface(interf string) *Definition {
	d.defInterface = defType(interf)
	return d
}

func (d *Definition) GetProperties() map[string]string {
	props := map[string]string{}
	for k, v := range d.defProperties {
		props[k] = string(v)
	}
	return props
}

func (d *Definition) SetProperties(properties map[string]string) *Definition {
	d.defProperties = map[string]defExpression{}
	for k, v := range properties {
		d.defProperties[k] = defExpression(v)
	}
	return d
}

func (d *Definition) AddProperty(name string, value string) *Definition {
	d.defProperties[name] = defExpression(value)
	return d
}

func (d *Definition) GetReturns() string {
	return string(d.defReturns)
}

func (d *Definition) SetReturns(returns string) *Definition {
	d.defReturns = defExpression(returns)
	return d
}

func (d *Definition) GetArguments() map[string]string {
	args := map[string]string{}
	for k, v := range d.defArguments {
		args[k] = string(v)
	}
	return args
}

func (d *Definition) SetArguments(args map[string]string) *Definition {
	d.defArguments = defArguments{}
	for k, v := range args {
		d.defArguments[k] = defType(v)
	}
	return d
}

func (d *Definition) GetError() string {
	return d.defError
}

func (d *Definition) SetError(err string) *Definition {
	d.defError = err
	return d
}

func (d *Definition) GetImports() []string {
	return d.defImports
}

func (d *Definition) SetImports(imports []string) *Definition {
	d.defImports = imports
	return d
}

func (d *Definition) AddImport(imp string) *Definition {
	d.defImports = append(d.defImports, imp)
	return d
}

func (d *Definition) GetScope() string {
	return string(d.defScope)
}

func (d *Definition) SetScope(scope string) *Definition {
	d.defScope = defScope(scope)
	return d
}

func (d *Definition) AddSetup(entity string, args ...string) {
	d.defSetup = append(d.defSetup, defStatement{
		entity: entity,
		args:   args,
	})
}

func (d *Definition) isPublic() bool {
	return unicode.IsUpper(utils.FirstRune(d.defName))
}
