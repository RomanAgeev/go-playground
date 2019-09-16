package samples

import (
	"strings"

	"github.com/RomanAgeev/playground/sampleError"
)

const L = int('A')
const H = int('Z')
const W = H - L + 3

func msdInsertionSort(arr []string, index int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			sub1 := arr[j][index:len(arr[j])]
			sub2 := arr[j-1][index:len(arr[j-1])]
			if sub1 >= sub2 {
				break
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func charAt(str string, index int) int {
	if index >= len(str) {
		return -1
	}
	return int([]byte(str)[index]) - L

}

func msdSortAZ(arr []string, index int) {
	if len(arr) <= 2 {
		msdInsertionSort(arr, index)
		return
	}

	count := make([]int, W)

	for i := 0; i < len(arr); i++ {
		j := charAt(arr[i], index) + 2
		count[j]++
	}

	for i := 0; i < len(count)-1; i++ {
		count[i+1] += count[i]
	}

	aux := make([]string, len(arr))

	for i := 0; i < len(arr); i++ {
		j := charAt(arr[i], index) + 1
		aux[count[j]] = arr[i]
		count[j]++
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = aux[i]
	}

	for i := 0; i < len(count)-2; i++ {
		newLo := count[i]
		newHi := count[i+1]
		if newHi > newLo {
			msdSortAZ(arr[newLo:newHi], index+1)
		}
	}
}

func MSDSortAZ(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr := strings.Fields(params[0])

	msdSortAZ(arr, 0)

	return arr, nil
}
