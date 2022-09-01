package tests

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/either"
	"github.com/sagmor/fun/maybe"
	"github.com/sagmor/fun/result"
)

func TestSuccess(t *testing.T) {
	r := result.Success(5)

	assert.True(t, r.IsSuccess())
	assert.False(t, r.IsFailure())
	assert.Equal(t, 5, r.RequireValue())
	assert.Equal(t, either.Left[error](5), r.Either())
	assert.Equal(t, maybe.Just(5), r.Maybe())

	v, err := r.Tuple()
	assert.Equal(t, 5, v)
	assert.Nil(t, err)
}

func TestFailure(t *testing.T) {
	r := result.Failure[int](assert.AnError)

	assert.False(t, r.IsSuccess())
	assert.True(t, r.IsFailure())
	assert.Equal(t, assert.AnError, r.Error())
	assert.Panics(t, func() {
		r.RequireValue()
	})
	assert.Equal(t, either.Right[int](assert.AnError), r.Either())
	assert.Equal(t, maybe.Empty[int](), r.Maybe())

	_, err := r.Tuple()
	assert.Equal(t, assert.AnError, err)
}

func TestStep(t *testing.T) {
	assert.Equal(t, 5, result.Step(
		result.Success("5"),
		strconv.Atoi,
	).RequireValue())

	assert.Equal(t, "3", result.Step(
		result.Success(3),
		result.Stepper(strconv.Itoa),
	).RequireValue())

	assert.Error(t, result.Step(
		result.Failure[string](assert.AnError),
		strconv.Atoi,
	).Error())

	assert.Error(t, result.Step(
		result.Success(3),
		func(int) (string, error) { return "", assert.AnError },
	).Error())
}

func TestSteps(t *testing.T) {
	assert.Equal(t, "5", result.Steps2(
		result.Success("5"),
		strconv.Atoi,
		result.Stepper(strconv.Itoa),
	).RequireValue())

	assert.Equal(t, 5, result.Steps3(
		result.Success("5"),
		strconv.Atoi,
		result.Stepper(strconv.Itoa),
		strconv.Atoi,
	).RequireValue())

	assert.Equal(t, "5", result.Steps4(
		result.Success("5"),
		strconv.Atoi,
		result.Stepper(strconv.Itoa),
		strconv.Atoi,
		result.Stepper(strconv.Itoa),
	).RequireValue())

	assert.Error(t, result.Steps2(
		result.Success("something"),
		strconv.Atoi,
		result.Stepper(strconv.Itoa),
	).Error())
}
