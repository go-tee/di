package ext

import (
	"github.com/gooff/di/container"
)

func NewNavigationExtension() CompilerExtension {
	return &NavigationExtension{}
}

type NavigationExtension struct {
	BaseExtension
}

func (e *NavigationExtension) Prepare(builder *container.Builder) {
	factory := builder.AddDefinition(e.Prefix("factory.test.deep")).
		SetType("github.com/gooff/di/ext/navigation.NavigationControl")
	for name, _ := range e.config {
		factory.AddSetup("item0 = NewCategoryItem(" + name + ")")
	}
}

func (e *NavigationExtension) Compile(builder *container.Builder) {
	panic("implement me")
}
