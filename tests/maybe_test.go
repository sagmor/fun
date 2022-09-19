package fun_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/either"
	"github.com/sagmor/fun/maybe"
)

func TestJust(t *testing.T) {
	t.Parallel()

	m := maybe.Just(5)

	assert.True(t, m.HasValue())
	assert.False(t, m.IsEmpty())
	assert.Equal(t, 5, m.RequireValue())
	assert.Equal(t, either.Left[fun.Nothing](5), m.Either())
}

func TestMaybeEmpty(t *testing.T) {
	t.Parallel()

	m := maybe.Empty[int]()

	assert.True(t, m.IsEmpty())
	assert.False(t, m.HasValue())
	assert.Panics(t, func() {
		m.RequireValue()
	})

	assert.Equal(t, either.Right[int](fun.Nothing{}), m.Either())
}
