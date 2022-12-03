package fun_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/either"
	"github.com/sagmor/fun/maybe"
	"github.com/sagmor/fun/result"
)

func TestSuccess(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
		func(int) (string, error) {
			return "", assert.AnError
		},
	).Error())
}

func TestSteps(t *testing.T) {
	t.Parallel()

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

type testResultAsHelper struct{}

func (testResultAsHelper) GoString() string {
	return "hello"
}

func TestResultAs(t *testing.T) {
	t.Parallel()

	// Struct/Interface transition
	t1 := result.Success(testResultAsHelper{})
	t2 := result.Success[fmt.GoStringer](testResultAsHelper{})
	assert.NotEqual(t, t1, t2)
	assert.Equal(t, t1, result.As[testResultAsHelper](t1))
	assert.Equal(t, t1, result.As[testResultAsHelper](t2))
	assert.Equal(t, t2, result.As[fmt.GoStringer](t1))
	assert.Equal(t, t2, result.As[fmt.GoStringer](t2))
	assert.True(t, result.As[string](t2).IsFailure())

	// string/any transition
	t3 := result.Success("a string")
	t4 := result.Success[any]("a string")
	assert.Equal(t, t3, result.As[string](t4))
	assert.Equal(t, t4, result.As[any](t3))
	assert.True(t, result.As[int](t3).IsFailure())

	// number/any transition
	t5 := result.Success(5)
	t6 := result.Success[any](5)
	assert.Equal(t, t5, result.As[int](t6))
	assert.Equal(t, t6, result.As[any](t5))

	// As doesn't cast
	assert.True(t, result.As[uint](t5).IsFailure())

	// Original error pass through
	assert.Equal(t, result.Failure[any](assert.AnError), result.As[any](result.Failure[string](assert.AnError)))
}

func TestResultCollect(t *testing.T) {
	t.Parallel()

	a1 := []fun.Result[int]{result.Success(1), result.Success(2), result.Success(3)}
	assert.Equal(t, []int{1, 2, 3}, result.Collect(a1).RequireValue())

	a2 := []fun.Result[int]{result.Success(1), result.Failure[int](assert.AnError), result.Success(3)}
	assert.Equal(t, assert.AnError, result.Collect(a2).Error())
}
