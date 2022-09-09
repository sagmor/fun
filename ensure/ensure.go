package ensure

import (
	"testing"

	"github.com/sagmor/fun"
)

// Value extracts a result Value or halts the test.
func Value[T any](t *testing.T, r fun.Result[T]) T {
	if r.IsFailure() {
		t.Fatal(r.Error())
	}

	return r.RequireValue()
}
