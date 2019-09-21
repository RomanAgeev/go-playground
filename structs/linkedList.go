package structs

type LLNode struct {
	Data interface{}
	Next *LLNode
}

func NewLLNode(next *LLNode, data interface{}) *LLNode {
	return &LLNode{
		Data: data,
		Next: next,
	}
}

func NewLLFromArray(data []interface{}) *LLNode {
	if len(data) == 0 {
		return nil
	}

	head := NewLLNode(nil, data[0])

	node := head
	for i := 1; i < len(data); i++ {
		node.Next = NewLLNode(nil, data[i])
		node = node.Next
	}

	return head
}
