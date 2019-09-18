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
