package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func TraversePreorder(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	res := traversePreorder(root)

	return utils.ToStringArray(res...), nil
}

func traversePreorder(root *structs.BTNode) []int {
	var res []int
	preorderInternal(root, &res)
	return res
}

func preorderInternal(node *structs.BTNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Data.(int))
	preorderInternal(node.Left, res)
	preorderInternal(node.Right, res)
}
