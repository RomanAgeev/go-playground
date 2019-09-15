package samples

import (
	"strings"

	"github.com/RomanAgeev/playground/sampleError"
)

func countingSortAZ(str string) string {
	const l = byte('A')
	const h = byte('z')

	arr := []byte(str)

	cnt := make([]int, h-l+1)

	count := 0
	for _, val := range arr {
		cnt[val-l]++
		count++
	}

	var tmp int
	for i, acc := 0, 0; i < len(cnt); i++ {
		tmp = cnt[i]
		cnt[i] = acc
		acc += tmp
	}

	res := make([]byte, count)

	for _, val := range arr {
		res[cnt[val-l]] = val
		cnt[val-l]++
	}

	return string(res)
}

func CountingSortAZ(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	str := strings.Join(strings.Fields(params[0]), ``)

	res := countingSortAZ(str)

	return []string{res}, nil
}
