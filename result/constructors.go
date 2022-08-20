// Package result implements functions for working with the Result type
package result

import "github.com/sagmor/fun"

// Success returns a successful result.
func Success[T any](val T) fun.Result[T] {
	return fun.NewResult(val, nil)
}

// Failure returns a failed result.
func Failure[T any](err error) fun.Result[T] {
	return fun.NewResult(fun.Nil[T](), err)
}

// StepFun is a function to transform from one type to the other.
type StepFun[From, To any] func(From) (To, error)

// Step transitions from one result to another.
func Step[From, To any](
	from fun.Result[From],
	step StepFun[From, To]) fun.Result[To] {

	val, err := from.Get()
	if err != nil {
		return Failure[To](err)
	}

	return fun.NewResult(step(val))
}

// Steps2 runs Step twice.
func Steps2[T1, T2, T3 any](
	from fun.Result[T1],
	step1 StepFun[T1, T2],
	step2 StepFun[T2, T3]) fun.Result[T3] {

	return Step(
		Step(
			from,
			step1,
		),
		step2,
	)
}

// Steps3 runs Step 3 times.
func Steps3[T1, T2, T3, T4 any](
	from fun.Result[T1],
	step1 StepFun[T1, T2],
	step2 StepFun[T2, T3],
	step3 StepFun[T3, T4]) fun.Result[T4] {

	return Step(
		Steps2(
			from,
			step1,
			step2,
		),
		step3,
	)
}

// Steps4 runs Step 3 times.
func Steps4[T1, T2, T3, T4, T5 any](
	from fun.Result[T1],
	step1 StepFun[T1, T2],
	step2 StepFun[T2, T3],
	step3 StepFun[T3, T4],
	step4 StepFun[T4, T5]) fun.Result[T5] {

	return Step(
		Steps3(
			from,
			step1,
			step2,
			step3,
		),
		step4,
	)
}
