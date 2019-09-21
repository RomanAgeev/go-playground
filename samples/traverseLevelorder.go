package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func TraverseLevelorder(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	res := traverseLevelorder(root)

	return utils.ToStringArray(res...), nil
}

func traverseLevelorder(root *structs.BTNode) []int {
	var res []int

	if root == nil {
		return res
	}

	level := structs.NewQueue()
	level.Enqueue(root)

	for level.Length() > 0 {
		nextLevel := structs.NewQueue()

		for item := level.Dequeue(); item != nil; item = level.Dequeue() {
			node := item.(*structs.BTNode)

			res = append(res, node.Data.(int))

			if node.Left != nil {
				nextLevel.Enqueue(node.Left)
			}
			if node.Right != nil {
				nextLevel.Enqueue(node.Right)
			}

			level = nextLevel
		}
	}

	return res
}
