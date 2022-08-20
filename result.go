package fun

// Result represents a value of something that can fail.
type Result[T any] struct {
	either Either[T, error]
}

// NewResult builds a new Either object.
//
// Note: Considered an internal implementation.
// Prefer using result.Success() or result.Failure().
func NewResult[T any](val T, err error) Result[T] {
	return Result[T]{
		either: Either[T, error]{
			isLeft: err == nil,
			left:   val,
			right:  err,
		},
	}
}

// ToEither converts a result to an Either.
func (r Result[T]) ToEither() Either[T, error] {
	return r.either
}

// IsSuccess check if the Result is successfull.
func (r Result[T]) IsSuccess() bool {
	return r.either.isLeft
}

// RequireValue get's the result value or panics.
func (r Result[T]) RequireValue() T {
	if r.either.isLeft {
		return r.either.left
	}

	panic(r.either.right)
}

// Get extract both value and error.
func (r Result[T]) Get() (T, error) {
	return r.either.Get()
}

// IsFailure checks if the result is a failure.
func (r Result[T]) IsFailure() bool {
	return !r.either.isLeft
}

// Error gets the Result error.
func (r Result[T]) Error() error {
	return r.either.right
}
