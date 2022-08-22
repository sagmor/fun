package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sagmor/fun"
	"github.com/sagmor/fun/promise"
	"github.com/sagmor/fun/result"
)

func TestPromiseFromValue(t *testing.T) {
	p := promise.FromValue(3)

	assert.Equal(t, result.Success(3), p.Result())
}

func TestPromiseFromError(t *testing.T) {
	p := promise.FromError[string](assert.AnError)

	assert.Equal(t, result.Failure[string](assert.AnError), p.Result())
}

func TestPromiseFromResult(t *testing.T) {
	r := result.Success("hello")
	p := promise.FromResult(r)

	assert.Equal(t, r, p.Result())
}

func TestPromiseWithTimeout(t *testing.T) {
	p := promise.WithTimeout(time.Millisecond*10, func(ctx context.Context, resolver promise.Resolver[bool]) {
		time.Sleep(time.Second)
		resolver(true, nil)
	})
	assert.False(t, p.IsResolved())
	assert.Error(t, p.Result().Error())
}

func TestPromiseCancel(t *testing.T) {
	p := promise.New(func(ctx context.Context, resolver promise.Resolver[bool]) {
		time.Sleep(time.Second)
		resolver(true, nil)
	})
	p.Cancel()
	assert.Error(t, p.Result().Error())
}

func TestPromiseAll(t *testing.T) {
	build := func(i int) fun.Promise[int] {
		return promise.New(func(ctx context.Context, r promise.Resolver[int]) {
			time.Sleep(time.Millisecond * time.Duration(10*i))
			r(i, nil)
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
