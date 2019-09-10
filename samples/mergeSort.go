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

	leftCopy := make([]int, len(left))
	copy(leftCopy, left)

	rightCopy := make([]int, len(right))
	copy(rightCopy, right)

	l, r, s := 0, 0, 0
	for l < len(leftCopy) && r < len(rightCopy) {
		if leftCopy[l] <= rightCopy[r] {
			arr[s] = leftCopy[l]
			l++
		} else {
			arr[s] = rightCopy[r]
			r++
		}
		s++
	}

	for l < len(leftCopy) {
		arr[s] = leftCopy[l]
		l++
		s++
	}

	for r < len(rightCopy) {
		arr[s] = rightCopy[r]
		r++
		s++
	}

	return arr
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
