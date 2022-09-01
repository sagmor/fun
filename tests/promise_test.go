package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/either"
	"github.com/sagmor/fun/maybe"
	"github.com/sagmor/fun/promise"
	"github.com/sagmor/fun/result"
)

func TestPromiseFromValue(t *testing.T) {
	p := promise.FromValue(3)

	assert.Equal(t, result.Success(3), p.Result())
	assert.Equal(t, maybe.Just(3), p.Maybe())
	assert.Equal(t, either.Left[error](3), p.Either())
	assert.Equal(t, 3, p.RequireValue())
	assert.True(t, p.IsSuccess())
	assert.False(t, p.IsFailure())
}

func TestPromiseFromError(t *testing.T) {
	p := promise.FromError[string](assert.AnError)

	assert.Equal(t, result.Failure[string](assert.AnError), p.Result())
	assert.Equal(t, either.Right[string](assert.AnError), p.Either())
	assert.False(t, p.IsSuccess())
	assert.True(t, p.IsFailure())
}

func TestPromiseFromResult(t *testing.T) {
	r := result.Success("hello")
	p := promise.FromResult(r)

	assert.Equal(t, r, p.Result())
}

func TestPromiseWithTimeout(t *testing.T) {
	p := promise.WithTimeout(time.Millisecond*10, func(ctx context.Context) (bool, error) {
		time.Sleep(time.Second)
		return true, nil
	})
	assert.False(t, p.IsResolved())
	assert.Error(t, p.Error())
}

func TestPromiseWithDeadline(t *testing.T) {
	p := promise.WithDeadline(time.Now().Add(time.Millisecond), func(ctx context.Context) (bool, error) {
		time.Sleep(time.Second)
		return true, nil
	})
	assert.False(t, p.IsResolved())
	assert.Error(t, p.Error())
}

func TestPromiseCancel(t *testing.T) {
	p := promise.New(func(ctx context.Context) (bool, error) {
		time.Sleep(time.Second)
		return true, nil
	})
	p.Cancel()
	assert.Error(t, p.Result().Error())
}

func TestPromiseAll(t *testing.T) {
	build := func(i int) fun.Promise[int] {
		return promise.New(func(ctx context.Context) (int, error) {
			time.Sleep(time.Millisecond * time.Duration(10*i))
			return i, nil
		})
	}

	// Promises should be returned in the order they are fulfilled.
	all := promise.All(build(3), build(1), build(2))
	assert.True(t, all.Next())
	assert.Equal(t, 1, all.Value().RequireValue())

	assert.True(t, all.Next())
	assert.Equal(t, 2, all.Value().RequireValue())

	assert.True(t, all.Next())
	assert.Equal(t, 3, all.Value().RequireValue())

	assert.False(t, all.Next())
}

func TestPromiseCancelRace(t *testing.T) {
	i := 0
	for i < 100 {
		p := promise.FromValue(i)
		go p.Cancel()
		val, err := p.Tuple()
		if err == nil {
			assert.Equal(t, i, val)
		}
		i++
	}
}
