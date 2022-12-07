package fun_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun/task"
)

func TestTaskSuccess(t *testing.T) {
	t.Parallel()

	r := task.Success()

	assert.True(t, r.IsSuccess())
	assert.False(t, r.IsFailure())
	assert.Nil(t, r.Error())
}

func TestTaskFailure(t *testing.T) {
	t.Parallel()

	r := task.Failure(assert.AnError)

	assert.False(t, r.IsSuccess())
	assert.True(t, r.IsFailure())
	assert.Equal(t, assert.AnError, r.Error())

	assert.Panics(t, func() {
		task.Failure(nil)
	})
}

func TestTaskFromReturn(t *testing.T) {
	t.Parallel()

	t1 := task.FromReturn(assert.AnError)
	assert.True(t, t1.IsFailure())

	t2 := task.FromReturn(nil)
	assert.True(t, t2.IsSuccess())
}
