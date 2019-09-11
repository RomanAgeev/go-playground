package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

type tripleSortImpl = func(arr []int)

func normLeft(arr []int, l int) int {
	for l < len(arr) && arr[l] == 0 {
		l++
	}
	return l
}

func normRight(arr []int, r int) int {
	for r >= 0 && arr[r] == 2 {
		r--
	}
	return r
}

func tripleSort_1(arr []int) {
	l := normLeft(arr, 0)
	r := normRight(arr, len(arr)-1)

	i := l
	for l < r && i <= r {
		if arr[i] == 1 {
			i++
			continue
		}

		if arr[i] == 0 {
			arr[l], arr[i] = arr[i], arr[l]
		} else {
			arr[r], arr[i] = arr[i], arr[r]
		}

		l = normLeft(arr, l)
		r = normRight(arr, r)
		if i < l {
			i = l
		}
	}
}

func tripleSort_2(arr []int) {
	l, m, r, i := -1, -1, len(arr), 0

	for i < r {
		switch arr[i] {
		case 0:
			l++
			m++
			arr[l], arr[i] = arr[i], arr[l]
			i++
		case 2:
			r--
			arr[r], arr[i] = arr[i], arr[r]
		default:
			m++
			i++
		}
	}
}

func tripleSort(impl tripleSortImpl, params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	impl(arr)

	return utils.ToStringArray(arr...), nil
}

func TripleSort1(params []string) ([]string, error) {
	return tripleSort(tripleSort_1, params)
}

func TripleSort2(params []string) ([]string, error) {
	return tripleSort(tripleSort_2, params)
}
