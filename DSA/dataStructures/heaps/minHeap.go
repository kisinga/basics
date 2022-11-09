package heaps

import "errors"

type MinHeap struct {
	values []int
}

func isEven(number int) bool {
	return number%2 == 0
}

func (heap *MinHeap) Peek() int {
	return heap.values[0]
}

func (heap *MinHeap) getParent(pos int) int {
	return (pos - 1) / 2
}
func (heap *MinHeap) Pop() (int, error) {
	heap.values[0], heap.values[len(heap.values)-1] = heap.values[len(heap.values)-1], heap.values[0]
	heap.values = heap.values[:len(heap.values)-1]
	if len(heap.values) < 1 {
		return 0, nil
	}
	heap.heapify(len(heap.values) - 1)
	return 0, nil
}

func (heap *MinHeap) GetValue(int) (int, error) {
	return 0, nil
}
func (heap *MinHeap) GetValues() []int {
	return heap.values
}

// returns not implemented error
func (heap *MinHeap) RemoveValue(int) (int, error) {
	return 0, errors.New("not implemented")
}
func (heap *MinHeap) swap(from, to int) {
	heap.values[to], heap.values[from] = heap.values[from], heap.values[to]
}

func (heap *MinHeap) heapify(position int) {
	if position == 0 {
		return
	}
	parent := heap.getParent(position)
	if heap.values[parent] > heap.values[position] {
		heap.swap(position, parent)
	}
	heap.heapify(parent)
}

func (heap *MinHeap) AddValue(val int) (int, error) {
	heap.values = append(heap.values, val)
	heap.heapify(len(heap.values) - 1)
	return 0, nil
}

func (heap *MinHeap) SearchValue(int) (int, error) {
	return 0, nil
}

func (heap *MinHeap) AddValues(values ...int) (int, error) {
	for _, val := range values {
		heap.AddValue(val)
	}
	return 0, nil
}
