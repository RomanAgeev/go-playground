package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

type direction bool

const (
	left  direction = false
	right direction = true
)

type lowestCommonAncestorImpl func(root *structs.BTNode, val1 int, val2 int) *structs.BTNode

func LowestCommonAncestor1(params []string) ([]string, error) {
	return lowestCommonAncestor(lowestCommonAncestor_impl1, params)
}

func LowestCommonAncestor2(params []string) ([]string, error) {
	return lowestCommonAncestor(lowestCommonAncestor_impl2, params)
}

func lowestCommonAncestor(impl lowestCommonAncestorImpl, params []string) ([]string, error) {
	if len(params) != 3 {
		return nil, sampleError.InvalidParamCount(3, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	val1, err := utils.ToInteger(params[1])
	if err != nil {
		return nil, err
	}

	val2, err := utils.ToInteger(params[2])
	if err != nil {
		return nil, err
	}

	root := TreeFromArray(arr)

	node := impl(root, val1, val2)

	var res []string
	if node == nil {
		res = []string{"<NO>"}
	} else {
		res = utils.ToStringArray(node.Data.(int))
	}

	return res, nil
}

func lowestCommonAncestor_impl1(root *structs.BTNode, val1 int, val2 int) *structs.BTNode {
	path1 := structs.NewStack()
	if !buildPath(root, val1, path1) {
		return nil
	}

	path2 := structs.NewStack()
	if !buildPath(root, val2, path2) {
		return nil
	}

	stack1 := structs.NewStack()
	for path1.Length() > 0 {
		stack1.Push(path1.Pop())
	}

	stack2 := structs.NewStack()
	for path2.Length() > 0 {
		stack2.Push(path2.Pop())
	}

	node := root
	for {
		dir1 := stack1.Pop()
		dir2 := stack2.Pop()

		if dir1 == nil || dir2 == nil || dir1 != dir2 {
			break
		}

		switch dir1 {
		case left:
			node = node.Left

		case right:
			node = node.Right
		}
	}
	return node
}

func buildPath(node *structs.BTNode, val int, path *structs.Stack) bool {
	if node == nil {
		return false
	}

	if node.Data == val {
		return true
	}

	path.Push(left)
	if buildPath(node.Left, val, path) {
		return true
	}

	path.Pop()

	path.Push(right)
	if buildPath(node.Right, val, path) {
		return true
	}

	path.Pop()

	return false
}

func lowestCommonAncestor_impl2(root *structs.BTNode, val1 int, val2 int) *structs.BTNode {
	found1, found2 := false, false
	var ant *structs.BTNode
	lowestCommonAncestor_impl2_rec(root, val1, val2, &found1, &found2, &ant)
	return ant
}

func lowestCommonAncestor_impl2_rec(node *structs.BTNode, val1 int, val2 int, found1 *bool, found2 *bool, ant **structs.BTNode) {
	if node == nil {
		return
	}

	if node.Data == val1 {
		*found1 = true
		if !*found2 {
			*ant = node
		}
	}

	if node.Data == val2 {
		*found2 = true
		if !*found1 {
			*ant = node
		}
	}

	if *found1 && *found2 {
		return
	}

	foundUpper := *found1 != *found2

	lowestCommonAncestor_impl2_rec(node.Left, val1, val2, found1, found2, ant)

	if *found1 && *found2 {
		return
	}

	if *found1 != *found2 && !foundUpper {
		*ant = node
	}

	lowestCommonAncestor_impl2_rec(node.Right, val1, val2, found1, found2, ant)

	if *found1 && *found2 {
		return
	}

	if *found1 != *found2 && !foundUpper {
		*ant = node
	}
}
