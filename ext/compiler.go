package ext

import (
	"log"

	"github.com/gooff/di/container"
	"github.com/gooff/di/ds"
)

func NewCompiler(config ds.Config) *Compiler {
	return &Compiler{
		config:     &config,
		builder:    container.NewBuilder(),
		extensions: map[string]CompilerExtension{},
	}
}

type Compiler struct {
	config     *ds.Config
	builder    *container.Builder
	sections   []string
	extensions map[string]CompilerExtension
}

func (c *Compiler) AddExtension(name string, extension CompilerExtension) {
	c.sections = append(c.sections, name)
	c.extensions[name] = extension
	extension.Init(name)
}

func (c *Compiler) Compile(packageName string, outputFile string) error {
	for _, section := range c.sections {
		sectionConfig := (*c.config)[section]
		c.extensions[section].SetConfig(sectionConfig.(ds.Config))
	}

	for _, section := range c.sections {
		log.Println("Preparing section", section)
		c.extensions[section].Prepare(c.builder)
	}

	return c.builder.Build(packageName, outputFile)
}
