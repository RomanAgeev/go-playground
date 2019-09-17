package structs

type Queue struct {
	Head   *LLNode
	Tail   *LLNode
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Enqueue(data interface{}) {
	node := NewLLNode(nil, data)
	if queue.Tail != nil {
		queue.Tail.Next = node
	}
	queue.Tail = node
	if queue.Head == nil {
		queue.Head = node
	}

	queue.length++
}

func (queue *Queue) Dequeue() interface{} {
	head := queue.Head
	if head == nil {
		return nil
	}

	queue.Head = queue.Head.Next
	if queue.Head == nil {
		queue.Tail = nil
	}

	queue.length--

	return head.Data
}

func (queue Queue) Length() int {
	return queue.length
}
