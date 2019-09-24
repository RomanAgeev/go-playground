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
		root.Left = InsertBST(root.Left, val)
	} else if val > data {
		root.Right = InsertBST(root.Right, val)
	}

	return root
}

func RemoveBST(root *structs.BTNode, val int) *structs.BTNode {
	if root == nil {
		return nil
	}

	data := root.Data.(int)

	if val < data {
		root.Left = RemoveBST(root.Left, val)
		return root
	}
	if val > data {
		root.Right = RemoveBST(root.Right, val)
		return root
	}

	if root.Left != nil && root.Right != nil {
		root.Left = ReplaceMinMaxBST(root, root.Left)
		return root
	}

	if root.Left != nil {
		return root.Left
	}
	if root.Right != nil {
		return root.Right
	}

	return nil
}

func ReplaceMinMaxBST(root *structs.BTNode, node *structs.BTNode) *structs.BTNode {
	if node.Right != nil {
		node.Right = ReplaceMinMaxBST(root, node.Right)
		return node
	}
	if node.Left != nil {
		node.Left = ReplaceMinMaxBST(root, node.Left)
		return node
	}

	root.Data = node.Data

	return nil
}
