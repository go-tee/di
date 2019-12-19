package di

import (
	"fmt"
	"strings"

	"github.com/gooff/di/config"
	"github.com/gooff/di/container"
	"github.com/gooff/di/ext"
)

func NewCompiler(config config.Config) *Compiler {
	return &Compiler{
		config:     &config,
		builder:    container.NewBuilder(),
		extensions: map[string]ext.CompilerExtension{},
	}
}

type Compiler struct {
	config     *config.Config
	builder    *container.Builder
	sections   []string
	extensions map[string]ext.CompilerExtension
}

func (c *Compiler) AddExtension(name string, extension ext.CompilerExtension) error {
	if _, ok := c.extensions[name]; ok {
		return fmt.Errorf("name '%s' is already used", name)
	}
	for n, _ := range c.extensions {
		if strings.ToLower(n) == strings.ToLower(name) {
			return fmt.Errorf("name '%s' is already used in a case-insensitive manner", name)
		}
	}

	c.sections = append(c.sections, name)
	c.extensions[name] = extension
	extension.Init(name)

	return nil
}

func (c *Compiler) MustAddExtension(name string, extension ext.CompilerExtension) {
	err := c.AddExtension(name, extension)
	if err != nil {
		panic(err)
	}
}

func (c *Compiler) Compile(packageName string, outputFile string) error {
	for _, section := range c.sections {
		sectionConfig := (*c.config)[section]
		c.extensions[section].SetConfig(sectionConfig.(config.Config))
	}

	for _, section := range c.sections {
		err := c.extensions[section].Prepare(c.builder)
		if err != nil {
			return err
		}
	}

	return c.builder.Build(packageName, outputFile)
}
