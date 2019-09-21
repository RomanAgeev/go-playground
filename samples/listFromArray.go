package samples

import "github.com/RomanAgeev/playground/structs"

func ListFromArray(data []int) *structs.LLNode {
	if len(data) == 0 {
		return nil
	}

	head := structs.NewLLNode(nil, data[0])

	node := head
	for i := 1; i < len(data); i++ {
		node.Next = structs.NewLLNode(nil, data[i])
		node = node.Next
	}

	return head
}
