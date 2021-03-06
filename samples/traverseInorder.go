package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func TraverseInorder(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	res := traverseInorder(root)

	return utils.ToStringArray(res...), nil
}

func traverseInorder(root *structs.BTNode) []int {
	var res []int
	inorderInternal(root, &res)
	return res
}

func inorderInternal(node *structs.BTNode, res *[]int) {
	if node == nil {
		return
	}
	inorderInternal(node.Left, res)
	*res = append(*res, node.Data.(int))
	inorderInternal(node.Right, res)
}
