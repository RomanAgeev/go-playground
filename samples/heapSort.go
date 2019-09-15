package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func swim(arr []int, i int) {
	for i > 0 && arr[(i-1)/2] < arr[i] {
		j := (i - 1) / 2

		arr[i], arr[j] = arr[j], arr[i]
		i = j
	}
}

func sink(arr []int, r int) {
	for i := 0; i*2+1 <= r; {
		j := i*2 + 1

		if j < r && arr[j+1] > arr[j] {
			j++
		}

		if arr[i] >= arr[j] {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i = j
	}
}

func heapSort(arr []int) {
	for i := len(arr) - 1; i > 0; {
		val := arr[i]
		swim(arr, i)
		if arr[i] == val {
			i--
		}
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		sink(arr, i-1)
	}
}

func HeapSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	heapSort(arr)

	return utils.ToStringArray(arr...), nil
}
