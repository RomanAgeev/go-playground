package structs

// TODO: remove test comment

type Stack struct {
	Head   *LLNode
	length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack Stack) Length() int {
	return stack.length
}

func (stack *Stack) Push(data interface{}) {
	stack.Head = NewLLNode(stack.Head, data)
	stack.length++
}

func (stack *Stack) Pop() interface{} {
	head := stack.Head

	if head == nil {
		return nil
	}

	stack.Head = head.Next
	stack.length--

	return head.Data
}

func (stack Stack) Peak() interface{} {
	head := stack.Head

	if head == nil {
		return nil
	}

	return head.Data
}
