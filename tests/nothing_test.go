package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
)

func TestNil(t *testing.T) {
	assert.Equal(t, 0, fun.Nil[int]())
	assert.Equal(t, "", fun.Nil[string]())
	assert.Equal(t, nil, fun.Nil[interface{}]())
	assert.Equal(t, struct{}{}, fun.Nil[struct{}]())
}
