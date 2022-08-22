package promise

import (
	"context"
	"time"

	"github.com/sagmor/fun"
)

// New builds a new promise.
func New[T any](handler Handler[T]) fun.Promise[T] {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return newPromise(ctx, cancelFunc, handler)
}

// WithTimeout creates a promise with a timeout.
func WithTimeout[T any](timeout time.Duration, handler Handler[T]) fun.Promise[T] {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)

	return newPromise(ctx, cancelFunc, handler)
}

// WithDeadline creates a promise with a deadline.
func WithDeadline[T any](d time.Time, handler Handler[T]) fun.Promise[T] {
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)

	return newPromise(ctx, cancelFunc, handler)
}

// FromValue returns a promise that fulfills with a static value.
func FromValue[T any](value T) fun.Promise[T] {
	return New(func(ctx context.Context, resolve Resolver[T]) {
		resolve(value, nil)
	})
}

// FromResult returns a promise that fulfills with a static result.
func FromResult[T any](value fun.Result[T]) fun.Promise[T] {
	return New(func(ctx context.Context, resolve Resolver[T]) {
		resolve(value.ToTuple())
	})
}

// FromError returns a promise that fulfills with a static error.
func FromError[T any](err error) fun.Promise[T] {
	return New(func(ctx context.Context, resolve Resolver[T]) {
		resolve(fun.Nil[T](), err)
	})
}
