package ds

import (
	"io/ioutil"
	"log"
	"path"
	"reflect"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
)

type Config map[string]interface{}

func ParseConfig(files ...string) Config {
	config := Config{}
	for _, file := range files {
		p := parseFile(path.Dir(file), path.Base(file))
		if err := mergo.Merge(&config, p, mergo.WithOverride); err != nil {
			log.Fatalf("Error in merging: %s", err)
		}
	}

	return config
}

func parseFile(dir string, fileName string) Config {
	file, err := ioutil.ReadFile(path.Join(dir, fileName))
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
	}

	var m Config

	err = yaml.Unmarshal(file, &m)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s", err)
	}

	if includes, ok := m["includes"]; ok {
		t := reflect.TypeOf(includes)
		if t.Kind() != reflect.Slice {
			log.Fatalf("Unexpected type in includes section in file %s", path.Join(dir, fileName))
		}
		if v, ok := includes.([]interface{}); ok {
			for _, include := range v {
				if v, ok := include.(string); ok {
					sub := parseFile(dir, v)
					if err := mergo.Merge(&m, sub, mergo.WithOverride); err != nil {
						log.Fatalf("Error in merging: %s", err)
					}
				}
			}
		} else {
			log.Fatalf("Unexpected type %s in includes section in file %s", t.String(), path.Join(dir, fileName))
		}

		delete(m, "includes")
	}

	return m
}
