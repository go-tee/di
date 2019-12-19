package ext

import (
	"github.com/go-tee/di/container"
)

func NewParametersExtension() CompilerExtension {
	return &ParametersExtension{}
}

type ParametersExtension struct {
	BaseExtension
}

func (p *ParametersExtension) Prepare(builder *container.Builder) error {
	panic("implement me")
}
