package heaps

import "fmt"

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) heapifyUp(index int) {
	// if index == 0 {
	// 	return
	// }
	// parent := parent(index)
	// if h.data[index] > h.data[parent] {
	// 	h.swap(index, parent)
	// }
	// h.heapifyUp(parent)
	for h.data[parent(index)] < h.data[index] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) swap(first, last int) {
	h.data[first], h.data[last] = h.data[last], h.data[first]
}

func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) heapifyDown(index int) {
	// for h.data[index] < h.data[leftChild(index)] || h.data[index] < h.data[rightChild(index)] {
	// 	if h.data[index] < h.data[leftChild(index)] {
	// 		h.swap(index, leftChild(index))
	// 		index = leftChild(index)
	// 	}
	// 	h.swap(index, rightChild(index))
	// 	index = rightChild(index)
	// }
}

// Extract returns the largest value and removes it from the heap
func (h *MaxHeap) Extract(val int) int {
	if len(h.data) == 0 {
		return -1
	}
	value := h.data[0]
	// promote the last value to be the first
	h.data[0] = h.data[len(h.data)-1]
	// remove the last value from the array
	h.data = h.data[:len(h.data)-2]
	// heapify downwards
	h.heapifyDown(0)
	return value
}

func main() {
	h := []int{10, 20, 30, 40, 50, 60, 70, 80}
	m := &MaxHeap{}
	for _, v := range h {
		m.Insert(v)
	}
	fmt.Println(m)
}
