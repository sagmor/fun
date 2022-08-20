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

func (r Result[T]) ToEither() Either[T, error] {
	return r.either
}

func (r Result[T]) IsSuccess() bool {
	return r.either.isLeft
}

func (r Result[T]) RequireValue() T {
	if r.either.isLeft {
		return r.either.left
	}

	panic(r.either.right)
}

func (r Result[T]) Get() (T, error) {
	return r.either.Get()
}

func (r Result[T]) IsFailure() bool {
	return !r.either.isLeft
}

func (r Result[T]) Error() error {
	return r.either.right
}
