package fun

// Result represents a value of something that can fail.
type Result[T any] Either[T, error]

// NewResult builds a new Either object.
func NewResult[T any](val T, err error) Result[T] {
	return Result[T]{
		isLeft: err == nil,
		left:   val,
		right:  err,
	}
}

// IsSuccess check if the Result is successfull.
func (r Result[T]) IsSuccess() bool {
	return r.isLeft
}

// RequireValue get's the result value or panics.
func (r Result[T]) RequireValue() T {
	if r.isLeft {
		return r.left
	}

	panic(r.right)
}

// Get extract both value and error.
func (r Result[T]) Get() (T, error) {
	return r.left, r.right
}

// IsFailure checks if the result is a failure.
func (r Result[T]) IsFailure() bool {
	return !r.isLeft
}

// Error gets the Result error.
func (r Result[T]) Error() error {
	return r.right
}

// ToEither converts a result to an Either.
func (r Result[T]) ToEither() Either[T, error] {
	return Either[T, error](r)
}

// ToMaybe converts a result to a Maybe.
func (r Result[T]) ToMaybe() Maybe[T] {
	return NewMaybe(r.isLeft, r.left)
}
