package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func partition(arr []int, lo int, hi int) int {
	i := lo

	for j := lo; j < hi; j++ {
		if arr[j] < arr[hi] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[hi] = arr[hi], arr[i]

	return i
}

func quickSort(arr []int, lo int, hi int) {
	if lo < hi {
		m := partition(arr, lo, hi)

		quickSort(arr, lo, m-1)
		quickSort(arr, m+1, hi)
	}
}

func QuickSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	quickSort(arr, 0, len(arr)-1)

	return utils.ToStringArray(arr...), nil
}
