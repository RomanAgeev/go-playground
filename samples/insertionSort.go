package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}

func InsertionSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	res := insertionSort(arr)

	return utils.ToStringArray(res...), nil
}
