# Heap Package - Quick Usage Guide

## Installation
```bash
go get github.com/SachinHg/go-generic-heap
```

## Basic Usage

### Integer Heaps
```go
import "github.com/SachinHg/go-generic-heap"

// Max heap (highest values first)
maxHeap := genericheap.IntHeap(true)
maxHeap.PushItem(100)
maxHeap.PushItem(70)
maxHeap.PushItem(90)
fmt.Println(maxHeap.Peek()) // 100

// Min heap (lowest values first)
minHeap := genericheap.IntHeap(false)
minHeap.PushItem(100)
minHeap.PushItem(70)
minHeap.PushItem(90)
fmt.Println(minHeap.Peek()) // 70
```

### Custom Struct Heaps
```go
type PriceEntry struct {
    Commodity string
    Price     int
}

// Max heap by price
priceHeap := genericheap.New[PriceEntry](func(a, b PriceEntry) bool {
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
taskHeap := genericheap.New[Task](func(a, b Task) bool {
    if a.Priority != b.Priority {
        return a.Priority > b.Priority
    }
    return a.Deadline < b.Deadline
})
```

## API Reference
- `IntHeap(max bool)` - Integer heap (true=max, false=min)
- `FloatHeap(max bool)` - Float heap
- `StringHeap(max bool)` - String heap
- `New[T](less func(a, b T) bool)` - Custom heap with comparison function
- `PushItem(item T)` - Add item
- `PopItem() T` - Remove and return top item
- `Peek() T` - Get top item without removing
- `IsEmpty() bool` - Check if empty
- `Size() int` - Get size