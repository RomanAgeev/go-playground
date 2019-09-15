package samples

import (
	"strings"

	"github.com/RomanAgeev/playground/sampleError"
)

// NOTE: Requires all the strings to be the same length

func lsdSortAZ(arr []string) {
	if len(arr) == 0 {
		return
	}

	length := len(arr[0])

	const l = byte('A')
	const r = byte('Z')

	aux := make([]string, len(arr))

	for i := length - 1; i >= 0; i-- {
		count := make([]int, r-l+2)

		for j := 0; j < len(arr); j++ {
			ch := []byte(arr[j])[i]
			count[ch-l+1]++
		}

		for j := 0; j < len(count)-1; j++ {
			count[j+1] += count[j]
		}

		for j := 0; j < len(arr); j++ {
			ch := []byte(arr[j])[i]
			aux[count[ch-l]] = arr[j]
			count[ch-l]++
		}

		for j := 0; j < len(arr); j++ {
			arr[j] = aux[j]
		}
	}
}

func LSDSortAZ(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr := strings.Fields(params[0])

	lsdSortAZ(arr)

	return arr, nil
}
