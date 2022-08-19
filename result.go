package fun

type Result[T any] struct {
	Either[T, error]
}