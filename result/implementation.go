package result

import "github.com/sagmor/fun"

// Result represents a value of something that can fail.
type result[T any] fun.Either[T, error]

// Compile time check to ensure result implements fun.Result.
var _ fun.Result[int] = result[int]{}

// Error gets the Result error.
func (r result[T]) Error() error {
	return r.Either().Right()
}

// IsFailure checks if the result is a failure.
func (r result[T]) IsFailure() bool {
	return r.Either().IsRight()
}

// IsSuccess check if the Result is successful.
func (r result[T]) IsSuccess() bool {
	return r.Either().IsLeft()
}

// RequireValue gets the result value or panics.
func (r result[T]) RequireValue() T {
	if r.IsSuccess() {
		return r.Either().Left()
	}

	panic(r.Error())
}

// Tuple extract both value and error.
func (r result[T]) Tuple() (T, error) {
	return r.Either().Tuple()
}

// Either converts a result to an Either.
func (r result[T]) Either() fun.Either[T, error] {
	return fun.Either[T, error](r)
}

// Maybe converts a result to a Maybe.
func (r result[T]) Maybe() fun.Maybe[T] {
	return fun.NewMaybe(r.IsSuccess(), r.Either().Left())
}
