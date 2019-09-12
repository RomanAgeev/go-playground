package structs

type Stack struct {
	Head   *LLNode
	Length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(data interface{}) {
	stack.Head = NewLLNode(stack.Head, data)
	stack.Length++
}

func (stack *Stack) Pop() interface{} {
	head := stack.Head

	if head == nil {
		return nil
	}

	stack.Head = head.Next
	stack.Length--

	return head.Data
}

func (stack Stack) Peak() interface{} {
	head := stack.Head

	if head == nil {
		return nil
	}

	return head.Data
}
