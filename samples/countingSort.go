package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func countingSort(arr []int, l int, h int) []int {
	cnt := make([]int, h-l+1)

	for _, val := range arr {
		cnt[val-l]++
	}

	var tmp int
	for i, acc := 0, 0; i < len(cnt); i++ {
		tmp = cnt[i]
		cnt[i] = acc
		acc += tmp
	}

	res := make([]int, len(arr))

	for _, val := range arr {
		res[cnt[val-l]] = val
		cnt[val-l]++
	}

	return res
}

func CountingSort(params []string) ([]string, error) {
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

	res := countingSort(arr, lo, hi)

	return utils.ToStringArray(res...), nil
}
