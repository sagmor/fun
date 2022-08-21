package fun

import "fmt"

// Either represents one of two possible results.
type Either[L, R any] struct {
	isLeft bool
	left   L
	right  R
}

// NewEither builds a new Either object.
//
// Note: Considered an internal implementation.
// Prefer using either.Left() or either.Right().
func NewEither[L, R any](isLeft bool, left L, right R) Either[L, R] {
	return Either[L, R]{
		isLeft: isLeft,
		left:   left,
		right:  right,
	}
}

// IsLeft tells you if the object has a Left value.
func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

// Left gets the left value of the either.
func (e Either[L, R]) Left() L {
	return e.left
}

var errMissingLeftValue = fmt.Errorf("no Left value available")

// RequireLeft gets the left value of the either or panics.
func (e Either[L, R]) RequireLeft() L {
	if !e.IsLeft() {
		panic(errMissingLeftValue)
	}

	return e.left
}

// IsRight tells you if the object has a Right value.
func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

// Right gets the right value of the either.
func (e Either[L, R]) Right() R {
	return e.right
}

var errMissingRightValue = fmt.Errorf("no Right value available")

// RequireRight gets the right value of the either or panics.
func (e Either[L, R]) RequireRight() R {
	if !e.IsRight() {
		panic(errMissingRightValue)
	}
	return e.right
}

// ToTuple both values regardless of the type.
func (e Either[L, R]) ToTuple() (L, R) {
	return e.left, e.right
}
