package fun_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/builder"
)

type testBuildableObject struct {
	Name string
}

func TestBuildReturnsAnObject(t *testing.T) {
	t.Parallel()

	o := builder.Build[testBuildableObject]()
	assert.NotNil(t, o)
	assert.True(t, o.IsSuccess())
}

func TestBuildExecutesOptions(t *testing.T) {
	t.Parallel()

	o := builder.Build(func(o *testBuildableObject) error {
		o.Name = "other_name"

		return nil
	})
	assert.Equal(t, "other_name", o.RequireValue().Name)
}

func TestBuildPassThroughOptionErrors(t *testing.T) {
	t.Parallel()

	o := builder.Build(func(o *testBuildableObject) error {
		return assert.AnError
	})
	assert.NotNil(t, o)
	assert.True(t, o.IsFailure())
	assert.Equal(t, assert.AnError, o.Error())
}

func TestFinishCanRunOnlyOnce(t *testing.T) {
	t.Parallel()

	b := builder.Start[testBuildableObject]()

	_ = b.Finish()
	o := b.Finish()

	assert.NotNil(t, o)
	assert.True(t, o.IsFailure())
}

func TestApplyShortCircuitsOnError(t *testing.T) {
	t.Parallel()

	counter := 0
	option := func(*testBuildableObject) error {
		counter++

		return assert.AnError
	}
	o := builder.Build(option, option, option)

	assert.True(t, o.IsFailure())
	assert.Equal(t, 1, counter)
}

func TestWithOptions(t *testing.T) {
	t.Parallel()

	counter := 0
	option := func(*testBuildableObject) error {
		counter++

		return nil
	}

	o := builder.Build(
		option,
		builder.WithOptions([]builder.Option[testBuildableObject]{
			option,
			option,
		}),
		option,
	)

	assert.True(t, o.IsSuccess())
	assert.Equal(t, 4, counter)
}

func TestBuilderWithFuncion(t *testing.T) {
	t.Parallel()

	counter := 0

	o := builder.Build(
		builder.WithFunction(func() builder.Option[testBuildableObject] {
			counter++

			return func(*testBuildableObject) error {
				counter++

				return nil
			}
		}),
	)

	assert.True(t, o.IsSuccess())
	assert.Equal(t, 2, counter)
}
