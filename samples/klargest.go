package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

type kLargestImpl = func(arr []int, k int) []int

func kLargestImpl_1(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	res := make([]int, k)

	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		res[i] = arr[len(arr)-1-i]
	}

	return res
}

func kLargestImpl_2(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	res := make([]int, k)

	copy(res, arr[:k])

	for i := 0; i < k; i++ {
		iMin := i
		for j := i + 1; j < k; j++ {
			if res[j] > res[iMin] {
				iMin = j
			}
		}
		res[i], res[iMin] = res[iMin], res[i]

		for j := k; j < len(arr); j++ {
			if arr[j] > res[i] {
				res[i], arr[j] = arr[j], res[i]
			}
		}
	}

	return res
}

func kLargest(impl kLargestImpl, params []string) ([]string, error) {
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

	res := impl(arr, k)

	return utils.ToStringArray(res...), nil
}

func KLargest_1(params []string) ([]string, error) {
	return kLargest(kLargestImpl_1, params)
}

func KLargest_2(params []string) ([]string, error) {
	return kLargest(kLargestImpl_2, params)
}
