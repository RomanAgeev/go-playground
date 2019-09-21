package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func TraversePostorder(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	res := traversePostorder(root)

	return utils.ToStringArray(res...), nil
}

func traversePostorder(root *structs.BTNode) []int {
	var res []int
	postorderInternal(root, &res)
	return res
}

func postorderInternal(node *structs.BTNode, res *[]int) {
	if node == nil {
		return
	}

	postorderInternal(node.Left, res)
	postorderInternal(node.Right, res)
	*res = append(*res, node.Data.(int))
}
