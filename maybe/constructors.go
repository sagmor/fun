package maybe

import "github.com/sagmor/fun"

// Just returns a Maybe with a value.
func Just[T any](val T) fun.Maybe[T] {
	return fun.NewMaybe(true, val)
}

// Empty returns an empty Maybe.
func Empty[T any]() fun.Maybe[T] {
	return fun.NewMaybe(false, fun.Nil[T]())
}
