package ext

import (
	"github.com/gooff/di/container"
	"github.com/gooff/di/ds"
)

func NewParametersExtension() CompilerExtension {
	return &ParametersExtension{}
}

type ParametersExtension struct {

}

func (p *ParametersExtension) Init(name string) {
	panic("implement me")
}

func (p *ParametersExtension) SetConfig(config ds.Config) {
	panic("implement me")
}

func (p *ParametersExtension) Prepare(builder *container.Builder) {
	panic("implement me")
}

func (p *ParametersExtension) Compile(builder *container.Builder) {
	panic("implement me")
}

