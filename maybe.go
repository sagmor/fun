package fun

import "errors"

// Maybe can eithe contin a value or not.
type Maybe[T any] Either[T, Nothing]

// NewMaybe builds a new Maybe object.
//
// Note: Considered an internal implementation.
// Prefer using maybe.Just() or maybe.Nothing().
func NewMaybe[T any](hasValue bool, value T) Maybe[T] {
	return Maybe[T]{
		isLeft: hasValue,
		left:   value,
		right:  Nothing{},
	}
}

// IsEmpty checks if this is a Nothing/Empty Maybe.
func (m Maybe[T]) IsEmpty() bool {
	return m.Either().IsRight()
}

// HasValue checks if the maybe has something in it.
func (m Maybe[T]) HasValue() bool {
	return m.Either().IsLeft()
}

// Either converts into a Maybe.
func (m Maybe[T]) Either() Either[T, Nothing] {
	return Either[T, Nothing](m)
}

// ErrMaybeMissingValue represents a failure to extract a value from a Maybe.
var ErrMaybeMissingValue = errors.New("maybe had no value")

// RequireValue gets the value or panics.
func (m Maybe[T]) RequireValue() T {
	if m.IsEmpty() {
		panic(ErrMaybeMissingValue)
	}

	return m.left
}
