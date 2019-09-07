package samples

import (
	"math"

	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func jumpSearch(arr []int, val int) int {
	jump := int(math.Sqrt(float64(len(arr))))

	lo := 0

	for lo < len(arr) {
		hi := lo + jump
		if hi >= len(arr) {
			hi = len(arr) - 1
		}

		switch {
		case val > arr[hi]:
			lo = hi + 1
		case val < arr[hi]:
			for i := lo; i < hi; i++ {
				if val == arr[i] {
					return i
				}
				if val < arr[i] {
					return -1
				}
			}
			return -1
		default:
			return hi
		}
	}
	return -1
}

func JumpSearch(params []string) ([]string, error) {
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

	res := jumpSearch(arr, val)

	return utils.ToStringArray(res), nil
}
