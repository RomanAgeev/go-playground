package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func kLargest_1(arr []int, k int) (res []int) {
	if k >= len(arr) {
		return arr
	}

	res = make([]int, k)

	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		res[i] = arr[len(arr)-1-i]
	}

	return
}

func KLargest(params []string) ([]string, error) {
	if len(params) != 2 {
		return nil, sampleError.InvalidParamCount(2, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	k, err := utils.ToInteger(params[1])
	if err != nil {
		return nil, err
	}

	res := kLargest_1(arr, k)

	return utils.ToStringArray(res...), nil
}
