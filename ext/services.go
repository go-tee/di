package ext

import (
	"fmt"
	"reflect"

	"github.com/gooff/di/config"
	"github.com/gooff/di/container"
)

func NewServicesExtension() CompilerExtension {
	return &ServicesExtension{}
}

type ServicesExtension struct {
	name   string
	config config.Config
}

func (e *ServicesExtension) Init(name string) {
	e.name = name
}

func (e *ServicesExtension) SetConfig(config config.Config) {
	e.config = config
}

func (e *ServicesExtension) Prepare(builder *container.Builder) error {
	for name, conf := range e.config {
		err := loadDefinition(builder, name, conf)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadDefinition(builder *container.Builder, name string, conf interface{}) error {
	t := reflect.TypeOf(conf)
	if t.Kind() == reflect.String {
		def, err := builder.AddDefinition(name)
		if err != nil {
			return err
		}
		def.SetType(conf.(string))
		return nil
	}
	if t.Kind() == reflect.Map {
		m := conf.(config.Config)
		def, err := builder.AddDefinition(name)
		if err != nil {
			return err
		}
		if v, ok := m["type"]; ok {
			def.SetType(v.(string))
		}
		if v, ok := m["interface"]; ok {
			def.SetInterface(v.(string))
		}
		if v, ok := m["properties"]; ok {
			converted, err := toMap(v)
			if err != nil {
				return fmt.Errorf("cannot set properties for service '%s', error: %s", name, err)
			}
			def.SetProperties(converted)
		}
		if v, ok := m["returns"]; ok {
			def.SetReturns(v.(string))
		}
		if v, ok := m["arguments"]; ok {
			converted, err := toMap(v)
			if err != nil {
				return fmt.Errorf("cannot set arguments for service '%s', error: %s", name, err)
			}
			def.SetArguments(converted)
		}
		if v, ok := m["error"]; ok {
			def.SetError(v.(string))
		}
		if v, ok := m["imports"]; ok {
			converted, err := toSlice(v)
			if err != nil {
				return fmt.Errorf("cannot set imports for service '%s', error: %s", name, err)
			}
			def.SetImports(converted)
		}
		if v, ok := m["scope"]; ok {
			def.SetScope(v.(string))
		}
		return nil
	}

	return fmt.Errorf("service named '%s' have unsupported config type '%s'", name, t.Kind())
}

func toMap(value interface{}) (map[string]string, error) {
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Map {
		out := map[string]string{}
		m := value.(config.Config)
		for k, v := range m {
			out[k] = v.(string)
		}
		return out, nil
	}

	return nil, fmt.Errorf("undetected value for map conversion: %s", t.Kind())
}

func toSlice(value interface{}) ([]string, error) {
	t := reflect.ValueOf(value)
	if t.Kind() == reflect.Slice {
		ret := make([]string, t.Len())
		for i := 0; i < t.Len(); i++ {
			ret[i] = t.Index(i).Interface().(string)
		}

		return ret, nil
	}

	return nil, fmt.Errorf("undetected value for slice conversion: %s", t.Kind())
}
