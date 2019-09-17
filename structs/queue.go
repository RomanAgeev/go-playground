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
	if queue.Head == nil {
		return nil
	}

	data := queue.Head.Data

	queue.Head = queue.Head.Next
	if queue.Head == nil {
		queue.Tail = nil
	}

	queue.length--

	return data
}

func (queue Queue) Length() int {
	return queue.length
}
