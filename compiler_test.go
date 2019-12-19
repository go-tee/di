package di

import (
	"testing"

	"github.com/elliotchance/testify-stats/assert"

	"github.com/go-tee/di/config"
	"github.com/go-tee/di/container"
	"github.com/go-tee/di/ext"
)

type FooExtension struct {
	ext.BaseExtension
}

func (p *FooExtension) Prepare(builder *container.Builder) error {
	panic("prepare not implemented")
}

func TestAddExtension(t *testing.T) {
	t.Run("error-already-used", func(t *testing.T) {
		compiler := NewCompiler(config.Config{})
		err1 := compiler.AddExtension("foo", &FooExtension{})
		assert.NoError(t, err1)
		err2 := compiler.AddExtension("foo", &FooExtension{})
		assert.EqualError(t, err2, "name 'foo' is already used")
	})
	t.Run("error-already-used-insensitive", func(t *testing.T) {
		compiler := NewCompiler(config.Config{})
		err1 := compiler.AddExtension("foo", &FooExtension{})
		assert.NoError(t, err1)
		err2 := compiler.AddExtension("Foo", &FooExtension{})
		assert.EqualError(t, err2, "name 'Foo' is already used in a case-insensitive manner")
	})
	t.Run("success-prefix", func(t *testing.T) {
		compiler := NewCompiler(config.Config{})
		extension := &FooExtension{}
		assert.NoError(t, compiler.AddExtension("foo", extension))
		assert.Equal(t, "foo.test", extension.Prefix("test"))
		assert.Equal(t, "@foo.test", extension.Prefix("@test"))
	})
}
