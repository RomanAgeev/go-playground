package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

type mergeItem struct {
	Slice []int
	Level int
}

func mergeSortIterative(arr []int) []int {
	stack := structs.NewStack()
	mergeStack := structs.NewStack()

	stack.Push(arr)

	for {
		slice, ok := stack.Peak().([]int)
		if !ok {
			return mergeStack.Pop().(mergeItem).Slice
		}

		m := len(slice) / 2

		left := slice[:m]
		right := slice[m:]

		if len(left) > 1 || len(right) > 1 {
			stack.Push(left)
			stack.Push(right)
			continue
		}

		mergeStack.Push(mergeItem{
			Slice: left,
			Level: 0,
		})
		mergeStack.Push(mergeItem{
			Slice: right,
			Level: 0,
		})

		for {
			if mergeStack.Length <= 1 {
				break
			}

			fst := mergeStack.Pop().(mergeItem)
			snd := mergeStack.Pop().(mergeItem)

			if fst.Level != snd.Level {
				mergeStack.Push(snd)
				mergeStack.Push(fst)
				break
			}

			res := stack.Pop().([]int)

			merge(res, fst.Slice, snd.Slice)

			mergeStack.Push(mergeItem{
				Slice: res,
				Level: fst.Level + 1,
			})
		}
	}
}

func MergeSortIterative(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	res := mergeSortIterative(arr)

	return utils.ToStringArray(res...), nil
}
