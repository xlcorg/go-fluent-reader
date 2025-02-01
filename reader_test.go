package fluent_reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFluentReader_Next(t *testing.T) {
	t.Run("next", func(t *testing.T) {
		reader := NewFluentReader("привет")

		assert.Equal(t, 'п', reader.Next())

		assert.Equal(t, 'р', reader.SkipOne().Next())

		assert.Equal(t, 'т', reader.SkipOne().Last())
	})
}

func TestFluentReader_Read(t *testing.T) {
	t.Run("read until", func(t *testing.T) {
		input := "абв:вба"
		reader := NewFluentReader(input)

		got := reader.ReadUntil('а')
		assert.Equal(t, "", got.String())

		got = reader.ReadUntil(':')
		assert.Equal(t, "абв", got.String())

		got = reader.ReadUntil('б')
		assert.Equal(t, ":в", got.String())

		got = reader.ReadUntil('!')
		assert.Equal(t, "ба", got.String())
	})
}

func TestFluentParser_Skip(t *testing.T) {
	t.Run("skip one", func(t *testing.T) {
		input := "абв"
		reader := NewFluentReader(input)

		reader.SkipOne()
		assert.Equal(t, "бв", reader.String())

		reader.SkipOne()
		assert.Equal(t, "в", reader.String())

		reader.SkipOne()
		assert.Equal(t, "", reader.String())

		reader.SkipOne()
		assert.Equal(t, "", reader.String())
	})

	t.Run("skip", func(t *testing.T) {
		input := "абв"
		reader := NewFluentReader(input)

		reader.Skip(0)
		assert.Equal(t, "абв", reader.String())

		reader.Skip(2)
		assert.Equal(t, "в", reader.String())

		reader.Skip(10)
		assert.Equal(t, "", reader.String())

		reader.Skip(-5)
		assert.Equal(t, "", reader.String())
	})

	t.Run("skip until", func(t *testing.T) {
		input := "абв:вба"
		reader := NewFluentReader(input)

		reader.SkipUntil('а')
		assert.Equal(t, "абв:вба", reader.String())

		reader.SkipUntil(':')
		assert.Equal(t, ":вба", reader.String())

		r := reader.Clone()

		reader.SkipUntil('!')
		assert.Equal(t, "", reader.String())

		got := r.SkipUntil('а').String()
		want := "а"
		assert.Equal(t, want, got)
	})

	t.Run("skip after", func(t *testing.T) {
		input := "абв:вба"
		reader := NewFluentReader(input)

		reader.SkipAfter('а')
		assert.Equal(t, "бв:вба", reader.String())

		reader.SkipAfter(':')
		assert.Equal(t, "вба", reader.String())

		reader.SkipAfter('!')
		assert.Equal(t, "", reader.String())
	})
}

func TestFluentParser_Parse(t *testing.T) {
	input := "gnsbgyг:200260797,jsiczw:434278153,qkuyencyih:88313277"
	reader := NewFluentReader(input)

	val := reader.ReadUntil(':')
	assert.Equal(t, "gnsbgyг", val.String())
	assert.True(t, val.HasLettersOnly())

	val = reader.SkipOne().ReadUntil(',')
	assert.True(t, val.HasDigitsOnly())
	assert.Equal(t, "200260797", val.String())

	val = reader.SkipOne().ReadUntil(':')
	assert.True(t, val.HasLettersOnly())
	assert.Equal(t, "jsiczw", val.String())

}
