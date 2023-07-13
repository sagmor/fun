package iterator

import (
	"github.com/sagmor/fun"
	"github.com/sagmor/fun/result"
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

// Clone implements fun.Iterator.
func (iter *filteredIterator[T]) Clone() fun.Result[fun.Iterator[T]] {
	return result.Step(
		iter.iterator.Clone(),
		func(cloned fun.Iterator[T]) (fun.Iterator[T], error) {
			return WithFilter(cloned, iter.filter), nil
		},
	)
}

// WithFilter creates an iterator that filter values as it's called.
func WithFilter[T any](iter fun.Iterator[T], filter func(T) bool) fun.Iterator[T] {
	return &filteredIterator[T]{
		iterator: iter,
		filter:   filter,
	}
}
