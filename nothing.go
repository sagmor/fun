// Package fun exposes the core interfaces for doing functional programming in go.
package fun

// Nil returns a null value for any type.
func Nil[T any]() T {
	var e T

	return e
}

// Default returns a static value.
func Default[T any](v T) func() T {
	return func() T { return v }
}

// Identity returns the same object.
func Identity[T any](v T) T {
	return v
}

// Nothing represents nothing.
type Nothing struct{}
