package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		sw := false

		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sw = true
			}
		}

		if !sw {
			break
		}
	}
	return arr
}

func BubbleSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	res := bubbleSort(arr)

	return utils.ToStringArray(res...), nil
}
