// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Utility function
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// Seed our RNG
	rand.Seed(time.Now().UTC().UnixNano())

	// Create array A[a] and fill with random numbers
	const a = 10000000
	var A []int
	A = make([]int, a)
	for i := 0; i < a; i++ {
		A[i] = rand.Intn(a)
	}

	// Initialize the heap
	h := &IntHeap{}
	heap.Init(h)

	before := makeTimestamp()
	for i := 0; i < a; i++ {
		heap.Push(h, A[i])
	}
	fmt.Printf("Minimum item in PQ: %d\n", heap.Pop(h))

	after := makeTimestamp()
	fmt.Printf("Elapsed time for priority queue: %dms \n", after-before)

	// Ensure no garbage collection during next experiment
	// by doing it beforehand
	runtime.GC()

	before = makeTimestamp()
	sort.Ints(A)
	after = makeTimestamp()
	fmt.Printf("Minimum item in sorted array: %d\n", A[0])
	fmt.Printf("Elapsed time for library sort: %dms \n", after-before)
}
