package fun

import "errors"

// Iterator allows you to iterate over collections.
type Iterator[T any] interface {
	Next() bool
	Value() T
	Clone() Result[Iterator[T]]
}

// ErrIteratorNotCloneable is returned by iterators that can't be cloned.
var ErrIteratorNotCloneable = errors.New("iterator is not cloneable")

// Map applies a mapper function to every element of an iterator an returns a slice.
func Map[T, R any](iter Iterator[T], mapper func(T) R) []R {
	results := []R{}

	for iter.Next() {
		results = append(results, mapper(iter.Value()))
	}

	return results
}

// Reduce collects over all the elements of an iterator.
func Reduce[T, R any](iter Iterator[T], start R, collector func(R, T) R) R {
	result := start

	for iter.Next() {
		result = collector(result, iter.Value())
	}

	return result
}

// Any return any value provided by the iterator if there is any.
func Any[T any](iter Iterator[T]) Maybe[T] {
	var anyValue T

	hasAny := iter.Next()
	if hasAny {
		anyValue = iter.Value()
	}

	return NewMaybe(hasAny, anyValue)
}
