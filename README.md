# Go Generic Heap

A clean, generic heap implementation for Go that makes heap operations simple and type-safe. Perfect for coding interviews and production use.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/SachinHg/go-generic-heap)](https://goreportcard.com/report/github.com/SachinHg/go-generic-heap)

## Features

- **🚀 Generic**: Works with any type using Go 1.21+ generics
- **🔒 Type-safe**: No `interface{}` casting needed
- **⚡ Simple API**: Just define your comparison function
- **🎯 Interview-friendly**: Perfect for coding interviews
- **📦 Zero dependencies**: Uses only Go standard library
- **🧪 Well-tested**: Comprehensive test coverage

## Installation

```bash
go get github.com/SachinHg/go-generic-heap
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/SachinHg/go-generic-heap"
)

func main() {
    // Max heap (highest values first)
    maxHeap := heap.IntHeap(true)
    maxHeap.PushItem(100)
    maxHeap.PushItem(70)
    maxHeap.PushItem(90)
    
    fmt.Println(maxHeap.Peek()) // 100
    
    // Min heap (lowest values first)
    minHeap := heap.IntHeap(false)
    minHeap.PushItem(100)
    minHeap.PushItem(70)
    minHeap.PushItem(90)
    
    fmt.Println(minHeap.Peek()) // 70
}
```

## Usage Examples

### Custom Struct Heaps

```go
type PriceEntry struct {
    Commodity string
    Price     int
}

// Max heap by price
priceHeap := heap.New[PriceEntry](func(a, b PriceEntry) bool {
    return a.Price > b.Price
})

priceHeap.PushItem(PriceEntry{"Gold", 100})
priceHeap.PushItem(PriceEntry{"Silver", 70})

fmt.Println(priceHeap.Peek()) // {Commodity: "Gold", Price: 100}
```

### Complex Priority Heaps

```go
type Task struct {
    Name     string
    Priority int
    Deadline int
}

// Higher priority first, then earlier deadline
taskHeap := heap.New[Task](func(a, b Task) bool {
    if a.Priority != b.Priority {
        return a.Priority > b.Priority
    }
    return a.Deadline < b.Deadline
})
```

### Built-in Type Heaps

```go
// Integer heaps
intMaxHeap := heap.IntHeap(true)   // Max heap
intMinHeap := heap.IntHeap(false)  // Min heap

// Float heaps
floatMaxHeap := heap.FloatHeap(true)   // Max heap
floatMinHeap := heap.FloatHeap(false) // Min heap

// String heaps
stringMaxHeap := heap.StringHeap(true)   // Max heap
stringMinHeap := heap.StringHeap(false) // Min heap
```

## API Reference

### Core Functions

| Function | Description |
|----------|-------------|
| `New[T](less func(a, b T) bool) *Heap[T]` | Create new heap with comparison function |
| `IntHeap(max bool) *Heap[int]` | Create integer heap (true=max, false=min) |
| `FloatHeap(max bool) *Heap[float64]` | Create float heap |
| `StringHeap(max bool) *Heap[string]` | Create string heap |
| `ComparableHeap[T cmp.Ordered](max bool) *Heap[T]` | Create heap for comparable types |

### Heap Operations

| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `PushItem(item T)` | Add item to heap | O(log n) |
| `PopItem() T` | Remove and return top item | O(log n) |
| `Peek() T` | Get top item without removing | O(1) |
| `IsEmpty() bool` | Check if heap is empty | O(1) |
| `Size() int` | Get number of elements | O(1) |
| `Clear()` | Remove all elements | O(1) |
| `ToSlice() []T` | Get copy of all elements | O(n) |

## Interview Examples

### Commodity Prices (Max Heap)
```go
priceHeap := heap.New[PriceEntry](func(a, b PriceEntry) bool {
    return a.Price > b.Price
})
```

### Task Scheduling (Priority + Deadline)
```go
taskHeap := heap.New[Task](func(a, b Task) bool {
    if a.Priority != b.Priority {
        return a.Priority > b.Priority
    }
    return a.Deadline < b.Deadline
})
```

### Court Booking (Min Heap by Next Free Time)
```go
courtHeap := heap.New[Court](func(a, b Court) bool {
    return a.NextFree < b.NextFree
})
```

## Why This Package?

### The Problem
Go's standard `container/heap` package requires manual implementation of the heap interface:

```go
// Traditional approach - verbose and error-prone
type MaxHeap []int
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{}   { /* complex implementation */ }
```

### The Solution
With this package, you get clean, type-safe heap operations:

```go
// Clean approach - simple and type-safe
maxHeap := heap.IntHeap(true)
maxHeap.PushItem(100)
maxHeap.PushItem(70)
fmt.Println(maxHeap.Peek()) // 100
```

## Performance

- **Time Complexity**: O(log n) for insert/delete, O(1) for peek
- **Space Complexity**: O(n) for storage
- **Zero Runtime Overhead**: Generics are compile-time only
- **Memory Efficient**: No unnecessary allocations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go 1.21+ generics
- Uses Go's standard `container/heap` package internally
- Inspired by the need for cleaner heap implementations in coding interviews

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=SachinHg/go-generic-heap&type=Date)](https://star-history.com/#SachinHg/go-generic-heap&Date)