package ext

import (
	"github.com/go-tee/di/container"
)

func NewNavigationExtension() CompilerExtension {
	return &NavigationExtension{}
}

type NavigationExtension struct {
	BaseExtension
}

func (e *NavigationExtension) Prepare(builder *container.Builder) error {
	factory := builder.MustAddDefinition(e.Prefix("factory.test.deep")).
		SetType("github.com/go-tee/di/ext/navigation.NavigationControl")
	for name, _ := range e.config {
		factory.AddSetup("item0 = NewCategoryItem(" + name + ")")
	}
	return nil
}

func (e *NavigationExtension) Compile(builder *container.Builder) {
	panic("implement me")
}
