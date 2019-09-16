package samples

import (
	"math"

	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func digit(val int, exp int) (dig int) {
	for i := 1; i <= exp; i++ {
		dig = val % 10
		val = val / 10
	}
	return
}

func lsdSort(arr []int) {
	aux := make([]int, len(arr))
	dig := make([]int, len(arr))

	for exp := 1; ; exp++ {
		stop := true

		threshhold := int(math.Pow10(exp))

		count := make([]int, 11)

		for i := 0; i < len(arr); i++ {
			if arr[i] >= threshhold {
				stop = false
			}
			dig[i] = digit(arr[i], exp)
			count[dig[i]+1]++
		}

		for i := 0; i < len(count)-1; i++ {
			count[i+1] += count[i]
		}

		for i := 0; i < len(arr); i++ {
			aux[count[dig[i]]] = arr[i]
			count[dig[i]]++
		}

		for i := 0; i < len(arr); i++ {
			arr[i] = aux[i]
		}

		if stop {
			break
		}
	}
}

func LSDSort(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	lsdSort(arr)

	return utils.ToStringArray(arr...), nil
}
