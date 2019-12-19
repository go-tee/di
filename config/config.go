package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"reflect"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
)

type Config map[string]interface{}

func ParseConfig(files ...string) (Config, error) {
	config := Config{}

	closed := map[string]bool{}
	for _, file := range files {
		p, err := parseFile(path.Dir(file), path.Base(file), closed)
		if err != nil {
			return nil, err
		}
		if err := mergo.Merge(&config, p, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	return config, nil
}

func parseFile(dir string, fileName string, closed map[string]bool) (Config, error) {
	filePath := path.Join(dir, fileName)
	if _, ok := closed[filePath]; ok {
		return nil, fmt.Errorf("file '%s' already loaded", filePath)
	}
	closed[filePath] = true

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %s", err)
	}

	var m Config

	err = yaml.Unmarshal(file, &m)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %s", err)
	}

	if includes, ok := m["includes"]; ok {
		t := reflect.TypeOf(includes)
		if t.Kind() != reflect.Slice {
			return nil, fmt.Errorf("unexpected type in includes section in file %s", filePath)
		}
		if v, ok := includes.([]interface{}); ok {
			for _, include := range v {
				if v, ok := include.(string); ok {
					sub, err := parseFile(dir, v, closed)
					if err != nil {
						return nil, err
					}
					if err := mergo.Merge(&m, sub, mergo.WithAppendSlice, mergo.WithOverride); err != nil {
						return nil, fmt.Errorf("error in merging: %s", err)
					}
				}
			}
		} else {
			return nil, fmt.Errorf("unexpected type %s in includes section in file %s", t.String(), filePath)
		}

		delete(m, "includes")
	}

	return m, nil
}
