package ext

import (
	"fmt"

	"github.com/gooff/di/container"
	"github.com/gooff/di/ds"
)

type CompilerExtension interface {
	Init(name string)
	SetConfig(config ds.Config)
	Prepare(builder *container.Builder)
}

type BaseExtension struct {
	name   string
	config ds.Config
}

func (e *BaseExtension) Init(name string) {
	e.name = name
}

func (e *BaseExtension) SetConfig(config ds.Config) {
	e.config = config
}

func (e *BaseExtension) Prefix(name string) string {
	return fmt.Sprintf("%s.%s", e.name, name)
}
