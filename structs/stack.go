package structs

type Stack struct {
	Head *LLNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(data interface{}) {
	stack.Head = NewLLNode(stack.Head, data)
}

func (stack *Stack) Pop() interface{} {
	head := stack.Head

	if head == nil {
		return nil
	}

	stack.Head = head.Next

	return head.Data
}
