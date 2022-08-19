package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	assert.Equal(t, 0, Nil[int]())
	assert.Equal(t, "", Nil[string]())
	assert.Equal(t, nil, Nil[interface{}]())
	assert.Equal(t, struct{}{}, Nil[struct{}]())
}
