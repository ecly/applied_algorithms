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

func createRandomInput(size int) []int {
	// Create array A[a] and fill with random numbers
	var A []int
	A = make([]int, size)
	for i := 0; i < size; i++ {
		A[i] = rand.Intn(size)
	}
	return A
}

func sortWithBinaryHeap(input []int) {
	// Initialize the heap
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < len(input); i++ {
		heap.Push(h, input[i])
	}
	for i := 0; i < len(input); i++ {
		input[i] = heap.Pop(h).(int)
	}
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// Seed our RNG
	rand.Seed(time.Now().UTC().UnixNano())
	const a = 10000000
	const repetitions = 100

	/*
		fmt.Println("Sorting with PQ, format: <#iteration,time/ms>")
		for i := 1; i <= repetitions; i++ {
			// Ensure no garbage collection during next experiment
			// by doing it beforehand
			runtime.GC()
			A := createRandomInput(a)

			before := makeTimestamp()
			sortWithBinaryHeap(A)
			after := makeTimestamp()
			fmt.Println("%d,%d", i, after-before)
		}
	*/
	fmt.Println("Sorting with sort.Ints, format: <#iteration,time/ms>")
	for i := 1; i <= repetitions; i++ {
		// Ensure no garbage collection during next experiment
		// by doing it beforehand
		runtime.GC()
		A := createRandomInput(a)

		before := makeTimestamp()
		sort.Ints(A)
		after := makeTimestamp()
		fmt.Println("%d,%d", i, after-before)
	}
}
