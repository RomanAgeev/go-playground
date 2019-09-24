package samples

import "github.com/RomanAgeev/playground/structs"

func SearchBST(root *structs.BTNode, val int) *structs.BTNode {
	if root == nil {
		return nil
	}

	data := root.Data.(int)

	if val < data {
		return SearchBST(root.Left, val)
	}

	if val > data {
		return SearchBST(root.Right, val)
	}

	return root
}

func InsertBST(root *structs.BTNode, val int) *structs.BTNode {
	if root == nil {
		return structs.NewBTNode(val)
	}

	if data := root.Data.(int); val < data {
		root.Right = InsertBST(root.Left, val)
	} else if val > data {
		root.Left = InsertBST(root.Right, val)
	}

	return root
}
