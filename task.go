package fun

// A Task represents an action that can fail but has no return value.
type Task interface {
	Error() error
	IsFailure() bool
	IsSuccess() bool
}
