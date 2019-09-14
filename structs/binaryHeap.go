package structs

type BinaryHeap []int

func NewBinaryHeap() BinaryHeap {
	return make([]int, 1)
}

func (heap BinaryHeap) Size() int {
	return len(heap) - 1
}

func (heap BinaryHeap) swim(i int) {
	for i > 1 && heap[i/2] > heap[i] {
		j := i / 2

		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}

func (heap BinaryHeap) sink(i int) {
	for i*2 <= heap.Size() {
		j := i * 2

		if j < heap.Size() && heap[j+1] < heap[j] {
			j++
		}

		if heap[i] <= heap[j] {
			break
		}

		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}

func (heap BinaryHeap) Min() (int, bool) {
	var min int
	ok := heap.Size() > 0
	if ok {
		min = heap[1]
	}
	return min, ok
}

func (pHeap *BinaryHeap) Insert(val int) {
	heap := *pHeap

	if len(heap) == cap(heap) {
		tmp := make([]int, len(heap), cap(heap)*2)
		copy(tmp, heap)
		heap = tmp
	}

	heap = append(heap, val)
	heap.swim(heap.Size())

	*pHeap = heap
}

func (pHeap *BinaryHeap) RemoveMin() (int, bool) {
	heap := *pHeap

	min, ok := heap.Min()
	if ok {
		heap[1] = heap[heap.Size()]
		heap = heap[:heap.Size()]
		heap.sink(1)

		*pHeap = heap
	}

	return min, ok
}
