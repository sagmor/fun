// Package iterator implements functions to work with fun.Iterator[T] types.
package iterator

import "github.com/sagmor/fun"

type sliceIterator[T any] struct {
	values  []T
	current int
}

// Next implements fun.Iterator.
func (i *sliceIterator[T]) Next() bool {
	i.current++

	return i.current < len(i.values)
}

// Value implements fun.Iterator.
func (i *sliceIterator[T]) Value() T {
	return i.values[i.current]
}

// FromSlice builds an iterator for a slice.
func FromSlice[T any](slice []T) fun.Iterator[T] {
	return &sliceIterator[T]{
		values:  slice,
		current: -1,
	}
}

type revertSliceIterator[T any] struct {
	values  []T
	current int
}

// Next implements fun.Iterator.
func (i *revertSliceIterator[T]) Next() bool {
	i.current--

	return i.current >= 0
}

// Value implements fun.Iterator.
func (i *revertSliceIterator[T]) Value() T {
	return i.values[i.current]
}

// FromRevertSlice builds an iterator for a slice that iterates backwards.
func FromRevertSlice[T any](slice []T) fun.Iterator[T] {
	return &revertSliceIterator[T]{
		values:  slice,
		current: len(slice),
	}
}

// ToSlice collects all values and returns a slice.
func ToSlice[T any](iter fun.Iterator[T]) []T {
	results := []T{}

	for iter.Next() {
		results = append(results, iter.Value())
	}

	return results
}
