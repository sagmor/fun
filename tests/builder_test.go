package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/builder"
	"github.com/sagmor/fun/ensure"
)

type testBuildableObject struct {
	Name string
}

func TestBuildReturnsAnObject(t *testing.T) {
	o := builder.Build[testBuildableObject]()
	assert.NotNil(t, o)
	assert.True(t, o.IsSuccess())
}

func TestBuildExecutesOptions(t *testing.T) {
	o := ensure.Value(t,
		builder.Build(func(o *testBuildableObject) error {
			o.Name = "other_name"
			return nil
		}),
	)
	
	assert.Equal(t, "other_name", o.Name)
}

func TestBuildPassThroughOptionErrors(t *testing.T) {
	o := builder.Build(func(o *testBuildableObject) error {
		return assert.AnError
	})
	assert.NotNil(t, o)
	assert.True(t, o.IsFailure())
	assert.Equal(t, assert.AnError, o.Error())
}

func TestFinishCanRunOnlyOnce(t *testing.T) {
	b := builder.Start[testBuildableObject]()

	_ = b.Finish()
	o := b.Finish()

	assert.NotNil(t, o)
	assert.True(t, o.IsFailure())
}

func TestApplyShortCircuitsOnError(t *testing.T) {
	counter := 0
	option := func(*testBuildableObject) error {
		counter++
		return assert.AnError
	}
	o := builder.Build(option, option, option)

	assert.True(t, o.IsFailure())
	assert.Equal(t, 1, counter)
}
