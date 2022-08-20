package fun

import "fmt"

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
	return m.ToEither().IsRight()
}

// HasValue checks if the maybe has something in it.
func (m Maybe[T]) HasValue() bool {
	return m.ToEither().IsLeft()
}

// ToEither converts into a Maybe.
func (m Maybe[T]) ToEither() Either[T, Nothing] {
	return Either[T, Nothing](m)
}

var errMaybeMissingValue = fmt.Errorf("maybe has no value")

// RequireValue gets the value or panics.
func (m Maybe[T]) RequireValue() T {
	if m.IsEmpty() {
		panic(errMaybeMissingValue)
	}

	return m.left
}

// ToResult convert into a Result.
func (m Maybe[T]) ToResult() Result[T] {
	var err error
	if m.IsEmpty() {
		err = errMaybeMissingValue
	}

	return NewResult(m.left, err)
}
