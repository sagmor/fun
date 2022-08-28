package promise

import (
	"github.com/sagmor/fun"
)

type promiseIterator[T any] struct {
	promises  []fun.Promise[T]
	current   int
	collected int
	collector chan int
}

// Next implements fun.Iterator.
func (iter *promiseIterator[T]) Next() bool {
	if iter.collected >= len(iter.promises) {
		return false
	}

	iter.current = <-iter.collector
	iter.collected++

	return true
}

// Value implements fun.Iterator.
func (iter *promiseIterator[T]) Value() fun.Result[T] {
	return iter.promises[iter.current].Result()
}

// All iterates over all promises as they are resolved.
func All[T any](promises ...fun.Promise[T]) fun.Iterator[fun.Result[T]] {
	iter := &promiseIterator[T]{
		promises:  promises,
		current:   -1,
		collected: 0,
		collector: make(chan int, len(promises)),
	}

	for index, promise := range promises {
		go func(index int, promise fun.Promise[T]) {
			promise.Wait()
			iter.collector <- index
		}(index, promise)
	}

	return iter
}
