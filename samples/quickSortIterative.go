package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func quickSortIterative(arr []int, lo int, hi int) {
	stack := structs.NewStack()

	stack.Push(structs.NewPair(lo, hi))

	for {
		pair, ok := stack.Pop().(*structs.Pair)
		if !ok {
			return
		}

		lo := pair.Fst.(int)
		hi := pair.Snd.(int)

		if lo >= hi {
			continue
		}

		m := partition(arr, lo, hi)

		stack.Push(structs.NewPair(lo, m-1))
		stack.Push(structs.NewPair(m+1, hi))
	}
}

func QuickSortIterative(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	quickSortIterative(arr, 0, len(arr)-1)

	return utils.ToStringArray(arr...), nil
}
