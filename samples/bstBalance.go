package samples

import (
	"github.com/RomanAgeev/playground/structs"
)

func BSTBalance(root *structs.BTNode) (res *structs.BTNode) {
	n := bstCountNodes(root)
	res = bstComplete(n)
	if res == nil {
		return
	}

	inOrd1 := bstInorder(root)
	inOrd2 := bstInorder(res)

	node1, ok1 := inOrd1()
	node2, ok2 := inOrd2()
	for ok1 && ok2 {
		node2.Data = node1.Data

		node1, ok1 = inOrd1()
		node2, ok2 = inOrd2()
	}

	return
}

func bstInorder(root *structs.BTNode) func() (*structs.BTNode, bool) {
	stack := structs.NewStack()

	for node := root; node != nil; node = node.Left {
		stack.Push(node)
	}

	return func() (res *structs.BTNode, ok bool) {
		res, ok = stack.Pop().(*structs.BTNode)
		if !ok {
			return
		}

		for node := res.Right; node != nil; node = node.Left {
			stack.Push(node)
		}

		return
	}
}

func bstCountNodes(root *structs.BTNode) int {
	if root == nil {
		return 0
	}

	return 1 + bstCountNodes(root.Left) + bstCountNodes(root.Right)
}

func bstComplete(n int) *structs.BTNode {
	return bstComplete_rec(n, 0)
}

func bstComplete_rec(n int, i int) *structs.BTNode {
	if i >= n {
		return nil
	}

	node := structs.NewBTNode(nil)

	node.Left = bstComplete_rec(n, 2*i+1)
	node.Right = bstComplete_rec(n, 2*i+2)

	return node
}
