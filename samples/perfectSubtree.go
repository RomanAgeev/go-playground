package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

type PerfectParams struct {
	Size    int
	SizeMax int
}

func PerfectSubtree(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	res := perfectSubtree(root)

	return utils.ToStringArray(res.SizeMax), nil
}

func perfectSubtree(node *structs.BTNode) *PerfectParams {
	res := &PerfectParams{}

	if node == nil {
		return res
	}

	left := perfectSubtree(node.Left)
	right := perfectSubtree(node.Right)

	min := utils.MinInt(left.Size, right.Size)
	max := utils.MaxInt(left.SizeMax, right.SizeMax)

	res.Size = min + min + 1
	res.SizeMax = utils.MaxInt(max, res.Size)

	return res
}
