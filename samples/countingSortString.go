package samples

import (
	"strings"

	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/utils"
)

func countingSortString(str string, l rune, h rune) string {
	cnt := make([]int, h-l+1)

	count := 0
	for _, rn := range str {
		cnt[rn-l]++
		count++
	}

	var tmp int
	for i, acc := 0, 0; i < len(cnt); i++ {
		tmp = cnt[i]
		cnt[i] = acc
		acc += tmp
	}

	res := make([]rune, count)

	for _, val := range str {
		res[cnt[val-l]] = val
		cnt[val-l]++
	}

	return string(res)
}

func CountingSortString(params []string) ([]string, error) {
	if len(params) != 3 {
		return nil, sampleError.InvalidParamCount(3, len(params))
	}

	str := strings.Join(strings.Fields(params[0]), ``)
	lo := utils.FirstRune(params[1])
	hi := utils.FirstRune(params[2])

	res := countingSortString(str, lo, hi)

	return []string{res}, nil
}
