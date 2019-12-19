package ext

import (
	"fmt"

	"github.com/gooff/di/config"
	"github.com/gooff/di/container"
	"github.com/gooff/di/utils"
)

type CompilerExtension interface {
	Init(name string)
	SetConfig(config config.Config)
	Prepare(builder *container.Builder) error
}

type BaseExtension struct {
	name   string
	config config.Config
}

func (e *BaseExtension) Init(name string) {
	e.name = name
}

func (e *BaseExtension) SetConfig(config config.Config) {
	e.config = config
}

func (e *BaseExtension) Prefix(name string) string {
	if utils.FirstRune(name) == '@' {
		return fmt.Sprintf("@%s.%s", e.name, name[1:])
	}
	return fmt.Sprintf("%s.%s", e.name, name)
}
