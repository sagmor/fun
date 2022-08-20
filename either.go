package fun

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

// IsRight tells you if the object has a Right value.
func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

// Right gets the right value of the either.
func (e Either[L, R]) Right() R {
	return e.right
}
