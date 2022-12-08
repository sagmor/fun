// Package builder provides tools to build different XDS object types
package builder

import (
	"errors"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/result"
)

// Option represents configuration functions to customize a built object.
type Option[T any] func(*T) error

// Builder allows you to build new objects.
type Builder[T any] struct {
	object fun.Result[*T]
}

// Apply invokes the provided builder options against the object being built.
func (b *Builder[T]) Apply(options ...Option[T]) *Builder[T] {
	for _, opt := range options {
		b.object = result.Step(
			b.object,
			func(o *T) (*T, error) {
				return o, opt(o)
			},
		)
	}

	return b
}

// ErrDoubleFinishAttempted indicates a builder was finished more than once.
var ErrDoubleFinishAttempted = errors.New("a builder can be finished only once")

// Finish closes the builder and returns the generated object.
func (b *Builder[T]) Finish() fun.Result[*T] {
	object := b.object
	b.object = result.Failure[*T](ErrDoubleFinishAttempted)

	return object
}

// Start the build of a new T object.
func Start[T any]() *Builder[T] {
	builder := new(Builder[T])
	builder.object = result.Success(new(T))

	return builder
}

// Build an object of type T.
func Build[T any](options ...Option[T]) fun.Result[*T] {
	return Start[T]().Apply(options...).Finish()
}

// WithOptions allows the mixing of Option arrays with regular Options when building objects.
func WithOptions[T any](options []Option[T]) Option[T] {
	return func(t *T) error {
		for _, option := range options {
			err := option(t)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// WithFunction allows to insert an arbitrary function within a build chain.
func WithFunction[T any](function func() Option[T]) Option[T] {
	return func(t *T) error {
		return function()(t)
	}
}
