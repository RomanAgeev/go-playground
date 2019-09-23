package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func TreeSize(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)
	size := treeSize(root)

	return utils.ToStringArray(size), nil
}

func treeSize(node *structs.BTNode) int {
	if node == nil {
		return 0
	}

	lSize := treeSize(node.Left)
	rSize := treeSize(node.Right)

	var maxSize int
	if lSize > rSize {
		maxSize = lSize
	} else {
		maxSize = rSize
	}

	return maxSize + 1
}
