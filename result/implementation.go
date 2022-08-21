package result

import "github.com/sagmor/fun"

// Result represents a value of something that can fail.
type result[T any] fun.Either[T, error]

// Compile time check to ensure result implements fun.Result.
var _ fun.Result[int] = result[int]{}

// Error gets the Result error.
func (r result[T]) Error() error {
	return r.ToEither().Right()
}

// IsFailure checks if the result is a failure.
func (r result[T]) IsFailure() bool {
	return r.ToEither().IsRight()
}

// IsSuccess check if the Result is successfull.
func (r result[T]) IsSuccess() bool {
	return r.ToEither().IsLeft()
}

// RequireValue get's the result value or panics.
func (r result[T]) RequireValue() T {
	if r.IsSuccess() {
		return r.ToEither().Left()
	}

	panic(r.Error())
}

// Get extract both value and error.
func (r result[T]) ToTuple() (T, error) {
	return r.ToEither().ToTuple()
}

// ToEither converts a result to an Either.
func (r result[T]) ToEither() fun.Either[T, error] {
	return fun.Either[T, error](r)
}

// ToMaybe converts a result to a Maybe.
func (r result[T]) ToMaybe() fun.Maybe[T] {
	return fun.NewMaybe(r.IsSuccess(), r.ToEither().Left())
}
