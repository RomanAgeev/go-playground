package samples

import "github.com/RomanAgeev/playground/structs"

func TreeFromArray(data []int) *structs.BTNode {
	return treeFromArrayInternal(data, 0)
}

func treeFromArrayInternal(data []int, i int) *structs.BTNode {
	if i >= len(data) {
		return nil
	}

	node := structs.NewBTNode(data[i])

	node.Left = treeFromArrayInternal(data, i*2+1)
	node.Right = treeFromArrayInternal(data, i*2+2)

	return node
}
