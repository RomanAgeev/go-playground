package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func TreeFromList(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	arr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	list := ListFromArray(arr)
	root := treeFromList(list)
	res := traversePreorder(root)

	return utils.ToStringArray(res...), nil
}

func treeFromList(llNode *structs.LLNode) *structs.BTNode {
	return treeFromListInternal(llNode, 0)
}

func treeFromListInternal(llNode *structs.LLNode, i int) *structs.BTNode {
	if llNode == nil {
		return nil
	}

	node := structs.NewBTNode(llNode.Data)

	llLeft := llNode
	j := i
	for ; j < i*2+1; j++ {
		if llLeft == nil {
			break
		}
		llLeft = llLeft.Next
	}

	var llRight *structs.LLNode
	if llLeft != nil {
		llRight = llLeft.Next
	}

	node.Left = treeFromListInternal(llLeft, j)
	node.Right = treeFromListInternal(llRight, j+1)

	return node
}
