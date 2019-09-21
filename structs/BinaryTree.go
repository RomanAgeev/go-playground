package structs

type BTNode struct {
	Data  interface{}
	Left  *BTNode
	Right *BTNode
}

func NewBTNode(data interface{}) *BTNode {
	return &BTNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func NewBTFromArray(data []interface{}) *BTNode {
	return newBTFromArrayInternal(data, 0)
}

func newBTFromArrayInternal(data []interface{}, i int) *BTNode {
	if i >= len(data) {
		return nil
	}

	node := NewBTNode(data[i])

	node.Left = newBTFromArrayInternal(data, i*2+1)
	node.Right = newBTFromArrayInternal(data, i*2+2)

	return node
}

func NewBTFromList(llNode *LLNode) *BTNode {
	return newBTFromListInternal(llNode, 0)
}

func newBTFromListInternal(llNode *LLNode, i int) *BTNode {
	if llNode == nil {
		return nil
	}

	node := NewBTNode(llNode.Data)

	llLeft := llNode
	j := i
	for ; j < i*2+1; j++ {
		if llLeft == nil {
			break
		}
		llLeft = llLeft.Next
	}

	var llRight *LLNode
	if llLeft != nil {
		llRight = llLeft.Next
	}

	node.Left = newBTFromListInternal(llLeft, j)
	node.Right = newBTFromListInternal(llRight, j+1)

	return node
}

func TraversePreorder(root *BTNode) []interface{} {
	var res []interface{}
	preorderInternal(root, &res)
	return res
}

func preorderInternal(node *BTNode, res *[]interface{}) {
	if node == nil {
		return
	}

	*res = append(*res, node.Data)
	preorderInternal(node.Left, res)
	preorderInternal(node.Right, res)
}

func TraverseInorder(root *BTNode) []interface{} {
	var res []interface{}
	inorderInternal(root, &res)
	return res
}

func inorderInternal(node *BTNode, res *[]interface{}) {
	if node == nil {
		return
	}
	inorderInternal(node.Left, res)
	*res = append(*res, node.Data)
	inorderInternal(node.Right, res)
}

func TraversePostorder(node *BTNode) []interface{} {
	var res []interface{}
	postorderInternal(node, &res)
	return res
}

func postorderInternal(node *BTNode, res *[]interface{}) {
	if node == nil {
		return
	}

	postorderInternal(node.Left, res)
	postorderInternal(node.Right, res)
	*res = append(*res, node.Data)
}
