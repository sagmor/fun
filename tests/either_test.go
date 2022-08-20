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
}

func TestRight(t *testing.T) {
	e := either.Right[int]("hello")

	assert.True(t, e.IsRight())
	assert.False(t, e.IsLeft())
	assert.Equal(t, "hello", e.Right())
}
