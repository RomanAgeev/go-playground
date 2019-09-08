package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func interpolationSearch(arr []int, val int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi && val >= arr[lo] && val <= arr[hi] {
		m := lo + int((float64(hi-lo) * float64(val-arr[lo]) / float64(arr[hi]-arr[lo])))

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

func InterpolationSearch(params []string) ([]string, error) {
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

	res := interpolationSearch(arr, val)

	return utils.ToStringArray(res), nil
}
