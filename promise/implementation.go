package promise

import (
	"context"
	"sync"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/maybe"
)

type promise[T any] struct {
	mu         sync.Mutex
	context    context.Context
	cancelFunc context.CancelFunc
	resolution fun.Maybe[fun.Result[T]]
	signal     chan fun.Nothing
}

// Compile time check to ensure promise is a fun.Promise.
var _ fun.Promise[int] = &promise[int]{}

// Resolver is a functiuon to resolve a promise.
type Resolver[T any] func(T, error)

// Handler is a function that handles a promise.
type Handler[T any] func(context.Context, Resolver[T])

func newPromise[T any](ctx context.Context, cancelFunc context.CancelFunc, handler Handler[T]) *promise[T] {
	p := &promise[T]{
		context:    ctx,
		cancelFunc: cancelFunc,
		signal:     make(chan fun.Nothing),
	}

	go handler(p.context, p.resolver)

	return p
}

// Cancel implements Promise.
func (p *promise[T]) Cancel() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.resolution.IsEmpty() {
		p.cancelFunc()
	}
}

// Collect implements Promise.
func (p *promise[T]) Collect() (T, error) {
	panic("unimplemented")
}

// IsResolved implements Promise.
func (p *promise[T]) IsResolved() bool {
	return p.resolution.HasValue()
}

// IsFailure implements Promise.
func (p *promise[T]) IsFailure() bool {
	return p.Result().IsFailure()
}

// IsSuccess implements Promise.
func (p *promise[T]) IsSuccess() bool {
	return p.Result().IsSuccess()
}

func (p *promise[T]) Wait() {
	if p.IsResolved() {
		return
	}

	select {
	case <-p.context.Done():
		p.resolver(fun.Nil[T](), p.context.Err())
	case <-p.signal:
		return
	}
}

// Result implements Promise.
func (p *promise[T]) Result() fun.Result[T] {
	p.Wait()
	return p.resolution.RequireValue()
}

func (p *promise[T]) resolver(val T, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.resolution = maybe.Just(fun.NewResult(val, err))
	close(p.signal)
}
