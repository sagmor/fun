package iterator

import (
	"github.com/sagmor/fun"
	"github.com/sagmor/fun/maybe"
)

// Map applies a mapper function to every element of an iterator an returns a slice.
func Map[T, R any](iter fun.Iterator[T], mapper func(T) R) []R {
	return ToSlice(WithTransform(iter, mapper))
}

// Reduce collects over all the elements of an iterator.
func Reduce[T, R any](iter fun.Iterator[T], start R, collector func(R, T) R) R {
	result := start

	for iter.Next() {
		result = collector(result, iter.Value())
	}

	return result
}

// Any return any value provided by the iterator if there is any.
func Any[T any](iter fun.Iterator[T]) fun.Maybe[T] {
	if iter.Next() {
		return maybe.Just(iter.Value())
	}
	return maybe.Empty[T]()
}
