package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func shellSort(arr []int) {
	h := 1
	for h < len(arr)/3 {
		h = h*3 + 1
	}

	for h > 0 {
		for i := h; i < len(arr); i++ {
			for j := i; j >= h; j -= h {
				if arr[j] >= arr[j-h] {
					break
				}
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
}

func ShellSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	shellSort(arr)

	return utils.ToStringArray(arr...), nil
}
