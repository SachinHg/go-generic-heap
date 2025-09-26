package main

import (
	"fmt"
	"github.com/SachinHg/go-generic-heap"
)

func main() {
	// Test basic functionality
	maxHeap := genericheap.IntHeap(true)
	maxHeap.PushItem(100)
	maxHeap.PushItem(70)
	maxHeap.PushItem(90)

	fmt.Printf("Max value: %d\n", maxHeap.Peek())
	fmt.Printf("Size: %d\n", maxHeap.Size())

	// Test custom struct
	type PriceEntry struct {
		Commodity string
		Price     int
	}

	priceHeap := genericheap.New[PriceEntry](func(a, b PriceEntry) bool {
		return a.Price > b.Price
	})

	priceHeap.PushItem(PriceEntry{"Gold", 100})
	priceHeap.PushItem(PriceEntry{"Silver", 70})

	fmt.Printf("Highest price: %+v\n", priceHeap.Peek())
}
