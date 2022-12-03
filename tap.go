package fun

// Tap let's you modify objects in place.
func Tap[T any](object *T, action func(object *T)) *T {
	action(object)

	return object
}
