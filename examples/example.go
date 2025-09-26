package main

import (
	"fmt"
	"github.com/SachinHg/go-generic-heap"
)

// Example struct for custom heap usage
type PriceEntry struct {
	Commodity string
	Price     int
}

type Task struct {
	Name     string
	Priority int
	Deadline int
}

func main() {
	fmt.Println("=== Heap Package Examples ===\n")

	// 1. Simple integer heap (max heap)
	fmt.Println("1. Integer Max Heap:")
	intHeap := genericheap.IntHeap(true) // true = max heap
	intHeap.PushItem(100)
	intHeap.PushItem(70)
	intHeap.PushItem(90)
	
	fmt.Printf("Max value: %d\n", intHeap.Peek())
	fmt.Printf("Size: %d\n", intHeap.Size())
	
	// 2. Simple integer heap (min heap)
	fmt.Println("\n2. Integer Min Heap:")
	minHeap := genericheap.IntHeap(false) // false = min heap
	minHeap.PushItem(100)
	minHeap.PushItem(70)
	minHeap.PushItem(90)
	
	fmt.Printf("Min value: %d\n", minHeap.Peek())
	
	// 3. Custom struct heap
	fmt.Println("\n3. Custom Struct Heap (by price):")
	priceHeap := genericheap.New[PriceEntry](func(a, b PriceEntry) bool {
		return a.Price > b.Price // Max heap by price
	})
	
	priceHeap.PushItem(PriceEntry{"Gold", 100})
	priceHeap.PushItem(PriceEntry{"Silver", 70})
	priceHeap.PushItem(PriceEntry{"Oil", 90})
	
	fmt.Printf("Highest price commodity: %+v\n", priceHeap.Peek())
	
	// 4. Complex struct with multiple fields
	fmt.Println("\n4. Task Priority Heap:")
	taskHeap := genericheap.New[Task](func(a, b Task) bool {
		// Higher priority first, then earlier deadline
		if a.Priority != b.Priority {
			return a.Priority > b.Priority
		}
		return a.Deadline < b.Deadline
	})
	
	taskHeap.PushItem(Task{"Fix bug", 5, 10})
	taskHeap.PushItem(Task{"Code review", 3, 5})
	taskHeap.PushItem(Task{"Write docs", 1, 15})
	
	fmt.Printf("Next task: %+v\n", taskHeap.Peek())
	
	// 5. Using helper functions
	fmt.Println("\n5. Using Helper Functions:")
	stringHeap := genericheap.StringHeap(true) // Max heap for strings
	stringHeap.PushItem("apple")
	stringHeap.PushItem("banana")
	stringHeap.PushItem("cherry")
	
	fmt.Printf("Lexicographically largest: %s\n", stringHeap.Peek())
	
	// 6. Comparable types (Go 1.21+)
	fmt.Println("\n6. Comparable Types:")
	floatHeap := genericheap.FloatHeap(false) // Min heap for floats
	floatHeap.PushItem(3.14)
	floatHeap.PushItem(2.71)
	floatHeap.PushItem(1.41)
	
	fmt.Printf("Smallest float: %.2f\n", floatHeap.Peek())
	
	// 7. Heap operations
	fmt.Println("\n7. Heap Operations:")
	fmt.Printf("Is empty: %t\n", intHeap.IsEmpty())
	fmt.Printf("Size: %d\n", intHeap.Size())
	
	// Pop all items
	fmt.Println("Popping all items:")
	for !intHeap.IsEmpty() {
		item := intHeap.PopItem()
		fmt.Printf("Popped: %d\n", item)
	}
	
	fmt.Printf("Is empty after popping: %t\n", intHeap.IsEmpty())
}