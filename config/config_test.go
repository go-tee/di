package config

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

var testData = "../test/data"

func TestIncludes(t *testing.T) {
	conf, err := ParseConfig(path.Join(testData, "config.includes.yml"))
	assert.NoError(t, err)

	merged, err := yaml.Marshal(conf)
	assert.NoError(t, err)

	expected := normalizeYaml(t, path.Join(testData, "config.includes.expected.yml"))
	assert.Equal(t, expected, string(merged))
}

func TestIncludesRecursive(t *testing.T) {
	filePath := path.Join(testData, "config.includes.recursive.yml")
	_, err := ParseConfig(filePath)
	assert.Errorf(t, err, "file '%s' already loaded", filePath)
}

func normalizeYaml(t *testing.T, filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		assert.NoError(t, err)
	}
	var m Config
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		assert.NoError(t, err)
	}
	expected, err := yaml.Marshal(m)
	if err != nil {
		assert.NoError(t, err)
	}
	return string(expected)
}
