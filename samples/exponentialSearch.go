package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func exponentialSearch(arr []int, val int) int {
	lo, hi := 0, 0
	for val > arr[hi] && hi < len(arr)-1 {
		lo = hi
		hi = (hi+1)*2 - 1
		if hi >= len(arr) {
			hi = len(arr) - 1
		}
	}

	// Binary Search
	for lo <= hi {
		m := lo + (hi-lo)/2

		switch {
		case val > arr[m]:
			lo = m + 1
		case val < arr[m]:
			hi = m - 1
		default:
			return m
		}
	}
	return -1
}

func ExponentialSearch(params []string) ([]string, error) {
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

	res := exponentialSearch(arr, val)

	return utils.ToStringArray(res), nil
}
