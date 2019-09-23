package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func CheckBST(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	var res string
	if isBST := checkBST(root); isBST {
		res = "BST"
	} else {
		res = "Not BST"
	}

	return []string{res}, nil
}

var CheckBSTProxy = checkBST

func checkBST(root *structs.BTNode) bool {
	return checkBST_rec(root, 0, 0, false, false)
}

func checkBST_rec(node *structs.BTNode, min int, max int, minDef bool, maxDef bool) bool {
	if node == nil {
		return true
	}

	val := node.Data.(int)

	if minDef && val <= min {
		return false
	}

	if maxDef && val >= max {
		return false
	}

	checkLeft := checkBST_rec(node.Left, min, val, minDef, true)
	checkRight := checkBST_rec(node.Right, val, max, true, maxDef)

	return checkLeft && checkRight
}
