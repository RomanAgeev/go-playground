package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func countingSortInplace(arr []int, l int, h int) {
	w := h - l + 1
	cnt := make([]int, w)
	acc := make([]int, w)

	for _, val := range arr {
		cnt[val-l]++
	}

	for i, sum := 0, 0; i < len(cnt); i++ {
		acc[i] = sum
		sum += cnt[i]
	}

	for i := 0; i < len(arr); {
		index := arr[i] - l
		if cnt[index] == 0 {
			i++
			continue
		}
		j := acc[index]
		if j > i {
			arr[i], arr[acc[index]] = arr[acc[index]], arr[i]
		} else {
			i++
		}
		acc[index]++
		cnt[index]--
	}
}

func CountingSortInplace(params []string) ([]string, error) {
	if len(params) != 3 {
		return nil, sampleError.InvalidParamCount(3, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	lo, err := utils.ToInteger(params[1])
	if err != nil {
		return nil, err
	}

	hi, err := utils.ToInteger(params[2])
	if err != nil {
		return nil, err
	}

	countingSortInplace(arr, lo, hi)

	return utils.ToStringArray(arr...), nil
}
