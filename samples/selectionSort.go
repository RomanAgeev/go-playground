package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		index := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				index = j
			}
		}

		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}

func SelectionSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	res := selectionSort(arr)

	return utils.ToStringArray(res...), nil
}
