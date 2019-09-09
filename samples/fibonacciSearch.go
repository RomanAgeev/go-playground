package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func fibonacciSearch(arr []int, val int) int {
	fibm1 := 1
	fibm2 := 0
	fib := fibm1 + fibm2

	for fib < len(arr) {
		fibm1 = fibm2
		fibm2 = fib
		fib = fibm1 + fibm2
	}

	off := 0

	for fib > 1 {
		i := off + fibm1
		if i > len(arr) {
			i = len(arr)
		}
		switch {
		case val > arr[i-1]:
			off = i
			fib = fibm2
			fibm2 = fibm1
			fibm1 = fib - fibm2
		case val < arr[i-1]:
			fib = fibm1
			fibm2 = fibm2 - fibm1
			fibm1 = fib - fibm2
		default:
			return i - 1
		}
	}
	if fibm2 == 1 && arr[off-1] == val {
		return off - 1
	}
	return -1
}

func FibonacciSearch(params []string) ([]string, error) {
	if len(params) != 2 {
		return nil, sampleError.InvalidParamCount(2, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	val, err := utils.ToInteger(params[1])
	if err != nil {
		return nil, err
	}

	res := fibonacciSearch(arr, val)

	return utils.ToStringArray(res), nil
}
