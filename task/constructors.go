// Package task implements functions to work with fun.Task types.
package task

import (
	"github.com/sagmor/fun"
	"github.com/sagmor/fun/result"
)

// Success returns a successful task.
func Success() fun.Task {
	return result.Success(fun.Nothing{})
}

// Failure returns a failed task.
// Remark: Panics if the error is nil.
func Failure(err error) fun.Task {
	if err == nil {
		panic("no error was provided as a task failure")
	}

	return result.Failure[fun.Nothing](err)
}

// FromReturn returns a task from an error return.
func FromReturn(err error) fun.Task {
	return result.FromTuple(fun.Nothing{}, err)
}
