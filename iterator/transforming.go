package iterator

import (
	"github.com/sagmor/fun"
	"github.com/sagmor/fun/result"
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

// Clone implements fun.Iterator.
func (iter *transformingIterator[From, To]) Clone() fun.Result[fun.Iterator[To]] {
	return result.Step(
		iter.iterator.Clone(),
		func(clone fun.Iterator[From]) (fun.Iterator[To], error) {
			return WithTransform(clone, iter.transform), nil
		},
	)
}

// WithTransform creates an iterator that transforms values as it's called.
func WithTransform[From, To any](iter fun.Iterator[From], transform func(From) To) fun.Iterator[To] {
	return &transformingIterator[From, To]{
		iterator:  iter,
		transform: transform,
	}
}
