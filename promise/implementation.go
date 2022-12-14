package promise

import (
	"context"
	"sync"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/maybe"
	"github.com/sagmor/fun/result"
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

// Handler is a function that handles a promise.
type Handler[T any] func(context.Context) (T, error)

func newPromise[T any](ctx context.Context, cancelFunc context.CancelFunc, handler Handler[T]) *promise[T] {
	p := &promise[T]{
		context:    ctx,
		cancelFunc: cancelFunc,
		signal:     make(chan fun.Nothing),
	}

	go func() {
		p.resolver(handler(p.context))
	}()

	return p
}

func (p *promise[T]) resolver(val T, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.IsResolved() {
		return
	}

	p.resolution = maybe.Just(result.FromTuple(val, err))
	close(p.signal)
}

// Cancel implements Promise.
func (p *promise[T]) Cancel() {
	p.cancelFunc()
}

// IsResolved implements Promise.
func (p *promise[T]) IsResolved() bool {
	return p.resolution.HasValue()
}

// Result implements Promise.
func (p *promise[T]) Result() fun.Result[T] {
	p.Wait()

	return p.resolution.RequireValue()
}

// Wait implements Promise.
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

// IsFailure implements Result.
func (p *promise[T]) Error() error {
	return p.Result().Error()
}

// IsFailure implements Result.
func (p *promise[T]) IsFailure() bool {
	return p.Result().IsFailure()
}

// IsSuccess implements Result.
func (p *promise[T]) IsSuccess() bool {
	return p.Result().IsSuccess()
}

// RequireValue implements Result.
func (p *promise[T]) RequireValue() T {
	return p.Result().RequireValue()
}

// Either implements Result.
func (p *promise[T]) Either() fun.Either[T, error] {
	return p.Result().Either()
}

// Maybe implements Result.
func (p *promise[T]) Maybe() fun.Maybe[T] {
	return p.Result().Maybe()
}

// Tuple implements Result.
func (p *promise[T]) Tuple() (T, error) {
	return p.Result().Tuple()
}
