package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

type traverseLevelorderImpl func(root *structs.BTNode) []int

func TraverseLevelorder1(params []string) ([]string, error) {
	return traverseLevelorder(traverseLevelorder_impl1, params)
}

func TraverseLevelorder2(params []string) ([]string, error) {
	return traverseLevelorder(traverseLevelorder_impl2, params)
}

func traverseLevelorder(impl traverseLevelorderImpl, params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	res := impl(root)

	return utils.ToStringArray(res...), nil
}

func traverseLevelorder_impl1(root *structs.BTNode) []int {
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

func traverseLevelorder_impl2(root *structs.BTNode) (res []int) {
	targetLevel := 0

	for {
		var level []int
		traverseLevelorder_impl2_rec(root, 0, targetLevel, &level)

		if len(level) <= 0 {
			break
		}

		res = append(res, level...)
		targetLevel++
	}

	return
}

func traverseLevelorder_impl2_rec(node *structs.BTNode, level int, targetLevel int, res *[]int) {
	if node == nil {
		return
	}

	if level == targetLevel {
		*res = append(*res, node.Data.(int))
		return
	}

	traverseLevelorder_impl2_rec(node.Left, level+1, targetLevel, res)
	traverseLevelorder_impl2_rec(node.Right, level+1, targetLevel, res)
}
