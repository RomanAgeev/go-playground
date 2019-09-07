package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func binarySearch(arr []int, val int) int {
	lo := 0
	hi := len(arr) - 1

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

func BinarySearch(params []string) ([]string, error) {
	if len(params) != 2 {
		return nil, sampleError.New("invalid params count")
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	val, err := utils.ToInteger(params[1])
	if err != nil {
		return nil, err
	}

	res := binarySearch(arr, val)

	return utils.ToStringArray(res), nil
}
