package fun

// Nil returns a null value for any type.
func Nil[T any]() T {
	var e T
	return e
}
