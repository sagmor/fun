package fun

// Promise represents the expectation that something will happen in the future.
type Promise[T any] interface {
	Cancel()
	Collect() (T, error)
	IsFailure() bool
	IsResolved() bool
	IsSuccess() bool
	Result() Result[T]
	Wait()
}
