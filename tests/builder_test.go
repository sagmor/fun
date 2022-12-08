package fun_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/builder"
	"github.com/sagmor/fun/result"
)

type testBuildableObject struct {
	Name string
}

func withName(name string) builder.Option[testBuildableObject] {
	return func(t *testBuildableObject) error {
		t.Name = name

		return nil
	}
}

func TestBuildReturnsAnObject(t *testing.T) {
	t.Parallel()

	o := builder.Build[testBuildableObject]()
	assert.NotNil(t, o)
	assert.True(t, o.IsSuccess())
}

func TestBuildExecutesOptions(t *testing.T) {
	t.Parallel()

	o := builder.Build(withName("other_name"))
	assert.Equal(t, "other_name", o.RequireValue().Name)
}

func TestBuildPassThroughOptionErrors(t *testing.T) {
	t.Parallel()

	o := builder.Build(builder.WithError[testBuildableObject](assert.AnError))
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

	o1 := builder.Build(
		option,
		builder.WithOptions([]builder.Option[testBuildableObject]{
			option,
			option,
		}),
		option,
	)

	assert.True(t, o1.IsSuccess())
	assert.Equal(t, 4, counter)

	o2 := builder.Build(
		option,
		builder.WithOptions([]builder.Option[testBuildableObject]{
			option,
			builder.WithError[testBuildableObject](assert.AnError),
		}),
		option,
	)

	assert.True(t, o2.IsFailure())
	assert.Equal(t, 6, counter)
	assert.Equal(t, assert.AnError, o2.Error())
}

func TestBuilderFromFunction(t *testing.T) {
	t.Parallel()

	counter := 0

	o := builder.Build(
		builder.FromFunction(func() builder.Option[testBuildableObject] {
			counter++

			return withName(fmt.Sprintf("%d", counter))
		}),
	)

	assert.True(t, o.IsSuccess())
	assert.Equal(t, 1, counter)
	assert.Equal(t, "1", o.RequireValue().Name)
}

func TestBuilderFromResult(t *testing.T) {
	t.Parallel()

	o1 := builder.Build(
		builder.FromResult(result.Success("some_name"), withName),
	)

	assert.Equal(t, "some_name", o1.RequireValue().Name)

	o2 := builder.Build(
		builder.FromResult(result.Failure[string](assert.AnError), withName),
	)

	assert.True(t, o2.IsFailure())
	assert.Equal(t, assert.AnError, o2.Error())
}
