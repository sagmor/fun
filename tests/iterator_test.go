package fun_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/iterator"
)

func TestIteratorFromSlice(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	it := iterator.FromSlice([]int{1, 2, 3})

	assert.Equal(t, []string{"1", "2", "3"}, iterator.Map(it, strconv.Itoa))
}

func TestIteratorReduce(t *testing.T) {
	t.Parallel()

	it := iterator.FromSlice([]int{1, 2, 3})

	assert.Equal(t, 6, iterator.Reduce(it, 0, func(r, i int) int { return r + i }))
}

func TestIteratorAny(t *testing.T) {
	t.Parallel()

	assert.True(t, iterator.Any(iterator.FromSlice([]int{1, 2, 3})).HasValue())
	assert.False(t, iterator.Any(iterator.FromSlice([]int{})).HasValue())
}

func TestIteratorToSlice(t *testing.T) {
	t.Parallel()

	it := iterator.FromRevertSlice([]int{1, 2, 3})
	assert.Equal(t, []int{3, 2, 1}, iterator.ToSlice(it))
}

func TestIteratorWithTransform(t *testing.T) {
	t.Parallel()

	it := iterator.WithTransform(iterator.FromSlice([]int{1, 2, 3}), strconv.Itoa)

	assert.True(t, it.Next())
	assert.Equal(t, "1", it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, "2", it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, "3", it.Value())

	assert.False(t, it.Next())
}

func TestIteratorWithFilter(t *testing.T) {
	t.Parallel()

	it := iterator.WithFilter(iterator.FromSlice([]int{1, 2, 3, 4, 5}), func(i int) bool { return i%2 == 0 })
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())

	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())

	assert.False(t, it.Next())
}
