// Package either implements functions for working with the Either type
package either

import "github.com/sagmor/fun"

// Left builds an Either with it's Left value set.
func Left[R, L any](val L) fun.Either[L, R] {
	return fun.NewEither(true, val, fun.Nil[R]())
}

// Right builds an Either with it's Right value set.
func Right[L, R any](val R) fun.Either[L, R] {
	return fun.NewEither(false, fun.Nil[L](), val)
}
