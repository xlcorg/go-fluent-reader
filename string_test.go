package fluent_reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_HasDigitsOnly(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert.True(t, String("1").HasDigitsOnly())

		assert.True(t, String("1234567890").HasDigitsOnly())

		assert.True(t, String("0").HasDigitsOnly())
	})

	t.Run("empty", func(t *testing.T) {
		assert.False(t, String("").HasDigitsOnly())
		assert.True(t, String("").Empty())
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, String("1").Int(), 1)

		assert.Equal(t, String("42").Int(), 42)

		assert.Equal(t, String("-7").Int(), -7)
	})
}
