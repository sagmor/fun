package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/either"
)

func TestLeft(t *testing.T) {
	e := either.Left[string](3)

	assert.True(t, e.IsLeft())
	assert.False(t, e.IsRight())
	assert.Equal(t, 3, e.Left())
	assert.Equal(t, 3, e.RequireLeft())

	v, _ := e.Tuple()
	assert.Equal(t, 3, v)

	assert.Panics(t, func() {
		e.RequireRight()
	})
}

func TestRight(t *testing.T) {
	e := either.Right[int]("hello")

	assert.True(t, e.IsRight())
	assert.False(t, e.IsLeft())
	assert.Equal(t, "hello", e.Right())
	assert.Equal(t, "hello", e.RequireRight())

	_, v := e.Tuple()
	assert.Equal(t, "hello", v)

	assert.Panics(t, func() {
		e.RequireLeft()
	})
}
