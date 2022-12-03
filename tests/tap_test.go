package fun_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
)

func TestTapInt(t *testing.T) {
	t.Parallel()

	s := 0
	fun.Tap(&s, func(object *int) {
		*object = 5
	})

	assert.Equal(t, 5, s)
}

func TestTapStruct(t *testing.T) {
	t.Parallel()

	type test struct {
		value string
	}

	o := &test{
		value: "foo",
	}

	fun.Tap(o, func(object *test) {
		object.value = "bar"
	})

	assert.Equal(t, "bar", o.value)
}
