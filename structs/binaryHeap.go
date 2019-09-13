package structs

type BinaryHeap []int

func NewBinaryHeap(capacity int) BinaryHeap {
	if capacity <= 0 {
		panic("Binary Heap capacity should be a positive number")
	}

	return make([]int, 0, capacity)
}

func (heap BinaryHeap) Size() int {
	return len(heap)
}

func (heap BinaryHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (heap BinaryHeap) Left(i int) int {
	return i*2 + 1
}

func (heap BinaryHeap) Right(i int) int {
	return i*2 + 2
}

func (heap BinaryHeap) Min() int {
	return heap[0]
}

func (pHeap *BinaryHeap) Insert(val int) {
	heap := *pHeap

	if len(heap) == cap(heap) {
		tmp := BinaryHeap(make([]int, len(heap), (cap(heap)+1)*2))
		copy(tmp, heap)
		heap = tmp
	}

	heap = append(heap, val)

	for i := len(heap) - 1; i > 0; {
		iParent := heap.Parent(i)
		if heap[iParent] <= heap[i] {
			break
		}

		heap[iParent], heap[i] = heap[i], heap[iParent]
		i = iParent
	}

	*pHeap = heap
}
