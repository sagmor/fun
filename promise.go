package fun

// Promise represents the expectation that something will happen in the future.
type Promise[T any] interface {
	Result[T]

	Cancel()
	IsResolved() bool
	Result() Result[T]
	Wait()
}
