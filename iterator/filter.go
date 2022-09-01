package iterator

import (
	"github.com/sagmor/fun"
)

type filteredIterator[T any] struct {
	iterator fun.Iterator[T]
	filter   func(T) bool
}

// Next implements fun.Iterator.
func (iter *filteredIterator[T]) Next() bool {
	for {
		if !iter.iterator.Next() {
			return false
		}

		if iter.filter(iter.iterator.Value()) {
			return true
		}
	}
}

// Value implements fun.Iterator.
func (iter *filteredIterator[T]) Value() T {
	return iter.iterator.Value()
}

// WithFilter creates an iterator that filter values as it's called.
func WithFilter[T any](iter fun.Iterator[T], filter func(T) bool) fun.Iterator[T] {
	return &filteredIterator[T]{
		iterator: iter,
		filter:   filter,
	}
}
