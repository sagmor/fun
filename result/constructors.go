package result

import (
	"github.com/sagmor/fun"
)

// Success returns a successful result.
func Success[T any](val T) fun.Result[T] {
	return result[T](fun.NewEither[T, error](true, val, nil))
}

// Failure returns a failed result.
func Failure[T any](err error) fun.Result[T] {
	return result[T](fun.NewEither(false, fun.Nil[T](), err))
}

// FromTuple returns a Result from a value error pair.
func FromTuple[T any](val T, err error) fun.Result[T] {
	return result[T](fun.NewEither(err == nil, val, err))
}
