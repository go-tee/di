package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDefinition(t *testing.T) {
	t.Run("error-already-used", func(t *testing.T) {
		builder := NewBuilder()
		_, err1 := builder.AddDefinition("foo")
		assert.NoError(t, err1)
		_, err2 := builder.AddDefinition("foo")
		assert.EqualError(t, err2, "service 'foo' already registered")
	})
	t.Run("error-already-used-insensitive", func(t *testing.T) {
		builder := NewBuilder()
		_, err1 := builder.AddDefinition("foo")
		assert.NoError(t, err1)
		_, err2 := builder.AddDefinition("Foo")
		assert.EqualError(t, err2, "service 'Foo' already registered in a case-insensitive manner")
	})
}
