package heaps

//test min heap

import (
	"testing"
)

// TestMinHeap tests the minHeap package
func TestMinHeap(t *testing.T) {
	// create a new minHeap
	minHeap := MinHeap{
		values: []int{},
	}

	// insert some elements
	minHeap.AddValue(10)
	minHeap.AddValue(20)
	minHeap.AddValue(30)
	minHeap.AddValue(40)
	minHeap.AddValue(50)
	minHeap.AddValue(60)

	// check the size of the heap
	if len(minHeap.values) != 6 {
		t.Errorf("Expected size of 6, got %d", len(minHeap.values))
	}

	// check the top element
	if minHeap.Peek() != 10 {
		t.Errorf("Expected 10, got %d", minHeap.Peek())
	}

	// remove the top element
	minHeap.Pop()

	// check the top element
	if minHeap.Peek() != 20 {
		t.Errorf("Expected 20, got %d", minHeap.Peek())
	}

	// check the size of the heap
	if len(minHeap.values) != 5 {
		t.Errorf("Expected size of 5, got %d", len(minHeap.values))
	}

	// remove all the elements
	minHeap.Pop()
	minHeap.Pop()
	minHeap.Pop()
	minHeap.Pop()
	minHeap.Pop()

	// check the size of the heap
	if len(minHeap.values) != 0 {
		t.Errorf("Expected size of 0, got %d", len(minHeap.values))
	}
}
