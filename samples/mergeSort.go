package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	m := len(arr) / 2

	left := mergeSort(arr[:m])
	right := mergeSort(arr[m:])

	sum := make([]int, len(left)+len(right))

	l, r, s := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			sum[s] = left[l]
			l++
		} else {
			sum[s] = right[r]
			r++
		}
		s++
	}

	for l < len(left) {
		sum[s] = left[l]
		l++
		s++
	}

	for r < len(right) {
		sum[s] = right[r]
		r++
		s++
	}

	return sum
}

func MergeSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	res := mergeSort(arr)

	return utils.ToStringArray(res...), nil
}
