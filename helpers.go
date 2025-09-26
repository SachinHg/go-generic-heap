package genericheap

import (
	"cmp"
)

// IntHeap creates a heap for integers
// For max heap: IntHeap(true) - higher values first
// For min heap: IntHeap(false) - lower values first
func IntHeap(max bool) *Heap[int] {
	if max {
		return New[int](func(a, b int) bool { return a > b })
	}
	return New[int](func(a, b int) bool { return a < b })
}

// FloatHeap creates a heap for floats
// For max heap: FloatHeap(true) - higher values first
// For min heap: FloatHeap(false) - lower values first
func FloatHeap(max bool) *Heap[float64] {
	if max {
		return New[float64](func(a, b float64) bool { return a > b })
	}
	return New[float64](func(a, b float64) bool { return a < b })
}

// StringHeap creates a heap for strings
// For max heap: StringHeap(true) - lexicographically larger first
// For min heap: StringHeap(false) - lexicographically smaller first
func StringHeap(max bool) *Heap[string] {
	if max {
		return New[string](func(a, b string) bool { return a > b })
	}
	return New[string](func(a, b string) bool { return a < b })
}

// ComparableHeap creates a heap for any comparable type
// Uses Go's built-in cmp.Compare for comparison
// For max heap: ComparableHeap[T](true) - larger values first
// For min heap: ComparableHeap[T](false) - smaller values first
func ComparableHeap[T cmp.Ordered](max bool) *Heap[T] {
	if max {
		return New[T](func(a, b T) bool { return cmp.Compare(a, b) > 0 })
	}
	return New[T](func(a, b T) bool { return cmp.Compare(a, b) < 0 })
}

// PriorityHeap creates a heap for items with priority
// The comparison function should return true if a has higher priority than b
func PriorityHeap[T any](priority func(a, b T) bool) *Heap[T] {
	return New[T](priority)
}
