package iterator

import (
	"github.com/sagmor/fun"
)

type transformingIterator[From, To any] struct {
	iterator  fun.Iterator[From]
	transform func(From) To
}

// Next implements fun.Iterator.
func (iter *transformingIterator[From, To]) Next() bool {
	return iter.iterator.Next()
}

// Value implements fun.Iterator.
func (iter *transformingIterator[From, To]) Value() To {
	return iter.transform(iter.iterator.Value())
}

// WithTransform creates an iterator that transforms values as it's called.
func WithTransform[From, To any](iter fun.Iterator[From], transform func(From) To) fun.Iterator[To] {
	return &transformingIterator[From, To]{
		iterator:  iter,
		transform: transform,
	}
}
