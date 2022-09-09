package result

import (
	"errors"
	"fmt"

	"github.com/sagmor/fun"
)

// StepFun is a function to transform from one type to the other.
type StepFun[From, To any] func(From) (To, error)

// Stepper converts a function into a Result Step functions.
func Stepper[From, To any](f func(From) To) StepFun[From, To] {
	return func(val From) (To, error) {
		return f(val), nil
	}
}

// Step transitions from one result to another.
func Step[From, To any](
	from fun.Result[From],
	step StepFun[From, To],
) fun.Result[To] {
	val, err := from.Tuple()
	if err != nil {
		return Failure[To](err)
	}

	return FromTuple(step(val))
}

// Steps2 runs Step twice.
func Steps2[T1, T2, T3 any](
	from fun.Result[T1],
	step1 StepFun[T1, T2],
	step2 StepFun[T2, T3],
) fun.Result[T3] {
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
	step3 StepFun[T3, T4],
) fun.Result[T4] {
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
	step4 StepFun[T4, T5],
) fun.Result[T5] {
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

// ErrTypeConversionFailed represent errors when transitioning between types.
var ErrTypeConversionFailed = errors.New("failed to convert between types")

// AsFun converts from one type to another.
func AsFun[T, S any](source S) (T, error) {
	var result T

	result, ok := interface{}(source).(T)
	if !ok {
		return result, fmt.Errorf("[%w] from %v to %T", ErrTypeConversionFailed, source, result)
	}

	return result, nil
}

// As converts a result from one type to another or fails.
// Useful for transitioning between an interface and it's implementation.
func As[T, S any](source fun.Result[S]) fun.Result[T] {
	return Step(source, AsFun[T, S])
}
