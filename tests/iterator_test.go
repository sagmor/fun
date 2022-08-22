package tests

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/iterator"
)

func TestIteratorFromSlice(t *testing.T) {
	it := iterator.FromSlice([]int{1, 2, 3})

	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())

	assert.False(t, it.Next())
}

func TestIteratorFromRevertSlice(t *testing.T) {
	it := iterator.FromRevertSlice([]int{1, 2, 3})

	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())

	assert.False(t, it.Next())
}

func TestIteratorMap(t *testing.T) {
	it := iterator.FromSlice([]int{1, 2, 3})

	assert.Equal(t, []string{"1", "2", "3"}, iterator.Map(it, strconv.Itoa))
}

func TestIteratorReduce(t *testing.T) {
	it := iterator.FromSlice([]int{1, 2, 3})

	assert.Equal(t, 6, iterator.Reduce(it, 0, func(r, i int) int { return r + i }))
}

func TestIteratorAny(t *testing.T) {
	assert.True(t, iterator.Any(iterator.FromSlice([]int{1, 2, 3})).HasValue())
	assert.False(t, iterator.Any(iterator.FromSlice([]int{})).HasValue())
}

func TestIteratorToSlice(t *testing.T) {
	it := iterator.FromRevertSlice([]int{1, 2, 3})
	assert.Equal(t, []int{3, 2, 1}, iterator.ToSlice(it))
}
