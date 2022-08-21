package fun

// Result represents a value of something that can fail.
type Result[T any] interface {
	Error() error
	IsFailure() bool
	IsSuccess() bool
	RequireValue() T
	ToEither() Either[T, error]
	ToMaybe() Maybe[T]
	ToTuple() (T, error)
}
