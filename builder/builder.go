// Package generic provides tools to build different XDS object types
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
		b.object = result.Step(b.object, func(o *T) (*T, error) { return o, opt(o) })
	}
	return b
}

// Finish closes the builder and returns the generated object.
func (b *Builder[T]) Finish() fun.Result[*T] {
	object := b.object
	b.object = result.Failure[*T](errors.New("a builder can only be finished only once"))
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
