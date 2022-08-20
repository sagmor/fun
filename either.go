package fun

type Either[L, R any] struct {
	isLeft bool
	left   L
	right  R
}

func NewEither[L, R any](isLeft bool, left L, right R) Either[L, R] {
	return Either[L, R]{
		isLeft: isLeft,
		left:   left,
		right:  right,
	}
}

func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

func (e Either[L, R]) Left() L {
	return e.left
}

func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

func (e Either[L, R]) Right() R {
	return e.right
}
