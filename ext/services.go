package ext

import (
	"log"
	"reflect"

	"github.com/gooff/di/container"
	"github.com/gooff/di/ds"
)

func NewServicesExtension() CompilerExtension {
	return &ServicesExtension{}
}

type ServicesExtension struct {
	name   string
	config ds.Config
}

func (e *ServicesExtension) Init(name string) {
	e.name = name
}

func (e *ServicesExtension) SetConfig(config ds.Config) {
	e.config = config
}

func (e *ServicesExtension) Prepare(builder *container.Builder) {
	for name, config := range e.config {
		loadDefinition(builder, name, config)
	}
}

func loadDefinition(builder *container.Builder, name string, config interface{}) {
	t := reflect.TypeOf(config)
	if t.Kind() == reflect.String {
		builder.AddDefinition(name).
			SetType(config.(string))
		return
	}
	if t.Kind() == reflect.Map {
		m := config.(ds.Config)
		def := builder.AddDefinition(name)
		if v, ok := m["type"]; ok {
			def.SetType(v.(string))
		}
		if v, ok := m["interface"]; ok {
			def.SetInterface(v.(string))
		}
		if v, ok := m["properties"]; ok {
			def.SetProperties(toMap(v))
		}
		if v, ok := m["returns"]; ok {
			def.SetReturns(v.(string))
		}
		if v, ok := m["arguments"]; ok {
			def.SetArguments(toMap(v))
		}
		if v, ok := m["error"]; ok {
			def.SetError(v.(string))
		}
		if v, ok := m["imports"]; ok {
			def.SetImports(toSlice(v))
		}
		if v, ok := m["scope"]; ok {
			def.SetScope(v.(string))
		}
		return
	}

	log.Fatalln("Undetected service definition type from config type =", t.Kind())
}

func toMap(value interface{}) map[string]string {
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Map {
		out := map[string]string{}
		m := value.(ds.Config)
		for k, v := range m {
			out[k] = v.(string)
		}
		return out
	}

	log.Fatalln("Undetected value toMap =", t.Kind())
	return nil
}

func toSlice(value interface{}) []string {
	t := reflect.ValueOf(value)
	if t.Kind() == reflect.Slice {
		ret := make([]string, t.Len())
		for i := 0; i < t.Len(); i++ {
			ret[i] = t.Index(i).Interface().(string)
		}

		return ret
	}

	log.Fatalln("Undetected value toSlice =", t.Kind())
	return nil
}
