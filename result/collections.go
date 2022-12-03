package result

import (
	"github.com/sagmor/fun"
)

// Collect converts an array of results into a single result.
func Collect[T any](results []fun.Result[T]) fun.Result[[]T] {
	array := make([]T, len(results))

	for i, r := range results {
		if r.IsFailure() {
			return Failure[[]T](r.Error())
		}

		array[i] = r.RequireValue()
	}

	return Success(array)
}
