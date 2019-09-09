package structs

type LLNode struct {
	Data interface{}
	Next *LLNode
}

func NewLLNode(head *LLNode, data interface{}) *LLNode {
	return &LLNode{
		Data: data,
		Next: head,
	}
}
