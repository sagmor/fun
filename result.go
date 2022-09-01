package fun

// Result represents a value of something that can fail.
type Result[T any] interface {
	Error() error
	IsFailure() bool
	IsSuccess() bool
	RequireValue() T
	Either() Either[T, error]
	Maybe() Maybe[T]
	Tuple() (T, error)
}
