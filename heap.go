package genericheap

import "container/heap"

// Heap is a generic heap implementation that works with any comparable type
type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// New creates a new heap with the given comparison function
// For max heap: func(a, b T) bool { return a > b }
// For min heap: func(a, b T) bool { return a < b }
func New[T any](less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{less: less}
	heap.Init(h)
	return h
}

// Len returns the number of elements in the heap
func (h *Heap[T]) Len() int {
	return len(h.data)
}

// Less compares two elements using the provided comparison function
func (h *Heap[T]) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

// Swap swaps two elements in the heap
func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push adds an element to the heap
func (h *Heap[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

// Pop removes and returns the top element from the heap
func (h *Heap[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

// PushItem adds an item to the heap (type-safe wrapper)
func (h *Heap[T]) PushItem(item T) {
	heap.Push(h, item)
}

// PopItem removes and returns the top item from the heap (type-safe wrapper)
func (h *Heap[T]) PopItem() T {
	return heap.Pop(h).(T)
}

// Peek returns the top item without removing it
func (h *Heap[T]) Peek() T {
	if len(h.data) == 0 {
		var zero T
		return zero
	}
	return h.data[0]
}

// IsEmpty returns true if the heap is empty
func (h *Heap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

// Size returns the number of elements in the heap
func (h *Heap[T]) Size() int {
	return len(h.data)
}

// Clear removes all elements from the heap
func (h *Heap[T]) Clear() {
	h.data = h.data[:0]
}

// ToSlice returns a copy of all elements in the heap
func (h *Heap[T]) ToSlice() []T {
	result := make([]T, len(h.data))
	copy(result, h.data)
	return result
}
