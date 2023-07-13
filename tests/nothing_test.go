package fun_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
)

func TestNil(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, fun.Nil[int]())
	assert.Equal(t, "", fun.Nil[string]())
	assert.Equal(t, nil, fun.Nil[any]())
	assert.Equal(t, struct{}{}, fun.Nil[struct{}]())
}
