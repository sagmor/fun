package result

import "github.com/sagmor/fun"

// Success returns a successful result.
func Success[T any](val T) fun.Result[T] {
	return fun.NewResult(val, nil)
}

// Failure returns a failed result.
func Failure[T any](err error) fun.Result[T] {
	return fun.NewResult(fun.Nil[T](), err)
}
