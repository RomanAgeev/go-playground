package samples

import "github.com/RomanAgeev/playground/structs"

func SearchBST(root *structs.BTNode, data int) (res *structs.BTNode) {
	_, res = searchBST_rec(nil, root, data)
	return
}

func searchBST_rec(parent *structs.BTNode, node *structs.BTNode, data int) (resP *structs.BTNode, resN *structs.BTNode) {
	if node == nil {
		resP, resN = parent, nil
		return
	}

	switch nodeData := node.Data.(int); {

	case data < nodeData:
		resP, resN = searchBST_rec(node, node.Left, data)

	case data > nodeData:
		resP, resN = searchBST_rec(node, node.Right, data)

	default:
		resP, resN = parent, node
	}
	return
}

func InsertBST(root *structs.BTNode, data int) *structs.BTNode {
	parent, node := searchBST_rec(nil, root, data)

	if node != nil {
		return root
	}

	newNode := structs.NewBTNode(data)

	if parent == nil {
		return newNode
	}

	if parentData := parent.Data.(int); parentData < data {
		parent.Right = newNode
	} else {
		parent.Left = newNode
	}

	return root
}
