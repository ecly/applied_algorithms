package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// Globals
var Seed int32
var VertexAmount int
var Vertices []Vertex
var Unvisited map[int]bool

// Seed: the seed of the vertex.
// Edges: pointers to edges connected to the vertex.
// Unvisited: whether is has already been visited by MST().
// FringeBy: -1 if no index in fringe pointing to, otherwise the index
type Vertex struct {
	Seed     int32
	Edges    []*Edge
	FringeBy int
}

// Adds the pointer to Edge 'edge' to the Slice of Edges for Vertex 'v'.
func (v *Vertex) AddEdge(edge *Edge) {
	v.Edges = append(v.Edges, edge)
}

// An Edge has a pointer to the Vertex it points at
// and a int32 Weight.
type Edge struct {
	To     int
	Weight int32
}

// An item holding its index in the heap
// and a pointer to its designated edge.
type Item struct {
	Index int
	Edge  *Edge
}

// https://golang.org/pkg/container/heap/#example__intHeap
type Fringe []*Item

// Return current length of fringe.
func (f Fringe) Len() int { return len(f) }

// We want items pointing to edges with lowest possible weight.
func (f Fringe) Less(i, j int) bool { return f[i].Edge.Weight < f[j].Edge.Weight }

// Swap f[i] and f[j] and update the Item's Index.
func (f Fringe) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
	f[i].Index = i
	Vertices[f[i].Edge.To].FringeBy = i
	f[j].Index = j
	Vertices[f[j].Edge.To].FringeBy = j
}

// Add an item to the fringe.
func (f *Fringe) Push(x interface{}) {
	n := len(*f)
	item := x.(*Item)
	item.Index = n
	*f = append(*f, item)
}

// Removes lowest weight item from fringe and returns it.
func (f *Fringe) Pop() interface{} {
	old := *f
	n := len(old)
	item := old[n-1]
	item.Index = -1 //for safety
	*f = old[0 : n-1]
	return item
}

// Calculate a minimum spanning tree from a slice of Edges
func MST() []*Edge {
	mst := make([]*Edge, 0, VertexAmount-1) // minimum size
	fringe := make(Fringe, 0)
	heap.Init(&fringe)

	for key := range Unvisited {
		delete(Unvisited, key)
		for _, e := range Vertices[key].Edges {
			item := &Item{Edge: e}
			heap.Push(&fringe, item)
			Vertices[e.To].FringeBy = item.Index
		}

		for fringe.Len() > 0 {
			mstEdge := heap.Pop(&fringe).(*Item).Edge
			delete(Unvisited, mstEdge.To)
			mst = append(mst, mstEdge)
			for _, edge := range Vertices[mstEdge.To].Edges {
				if _, ok := Unvisited[edge.To]; ok {
					if Vertices[edge.To].FringeBy == -1 {
						item := &Item{Edge: edge}
						heap.Push(&fringe, item)
						Vertices[edge.To].FringeBy = item.Index
					} else {
						if edge.Weight < fringe[Vertices[edge.To].FringeBy].Edge.Weight {
							fringe[Vertices[edge.To].FringeBy].Edge = edge
							heap.Fix(&fringe, Vertices[edge.To].FringeBy)
						}
					}
				}
			}
		}
	}
	return mst
}

// Max weights to consider to get correct output on given tests
var maxWeightComplete int32 = -80000

// Generate a fully connected graph with 'VertexAmount' vertices
func generateComplete() {
	// at least this big for starters
	for i := 0; i < VertexAmount-1; i++ {
		for j := i + 1; j < VertexAmount; j++ {
			weight := getEdgeWeight(i, j)
			if weight < maxWeightComplete {
				Vertices[i].AddEdge(&Edge{j, weight})
				Vertices[j].AddEdge(&Edge{i, weight})
			}
		}
	}
}

// Utility function
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Generate a 'numX' * 'numY' graph with connected rows and comlumns
func generateGrid(numX int, numY int) {
	before := makeTimestamp()

	// row edges
	for j := 0; j < numY; j++ {
		start := j * numX
		end := start + numX - 1
		for i := start; i < end; i++ {
			to := i + 1
			weight := getEdgeWeight(i, to)
			Vertices[i].AddEdge(&Edge{to, weight})
			Vertices[to].AddEdge(&Edge{i, weight})
		}
	}
	//column edges
	for i := 0; i < numX; i++ {
		for j := 0; j < numY-1; j++ {
			from := i + numX*j
			to := from + numX
			weight := getEdgeWeight(from, to)
			Vertices[from].AddEdge(&Edge{to, weight})
			Vertices[to].AddEdge(&Edge{from, weight})
		}
	}
	after := makeTimestamp()
	fmt.Printf("Grid generation time: %d\n", after-before)
}

// Read a graph from file with each lines being of format v1<tab>v2
// indicating an edge between the two vertices.
// Vertex number should always be less than 'VertexAmount'
// and there should be at most 'numOfEdges' edges.
func readGraph(filename string, numOfEdges int) []Edge {
	before := makeTimestamp()
	graph := make([]Edge, 0, numOfEdges) // known size
	if file, err := os.Open(filename); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "#") {
				continue
			}
			words := strings.Fields(scanner.Text())
			x, _ := strconv.Atoi(words[0])
			y, _ := strconv.Atoi(words[1])
			weight := getEdgeWeight(x, y)
			Vertices[x].AddEdge(&Edge{y, weight})
			Vertices[y].AddEdge(&Edge{x, weight})
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	after := makeTimestamp()
	fmt.Printf("File read time: %d\n", after-before)
	return graph
}

// Calculate the weight of a vertex based on its from, to vertices.
func getEdgeWeight(v1 int, v2 int) int32 {
	return xorshift32(Vertices[v1].Seed^Vertices[v2].Seed) % 100000
}

// Provided xorshift32 implementation.
func xorshift32(seed int32) int32 {
	ret := seed
	ret ^= ret << 13
	ret ^= ret >> 17
	ret ^= ret << 5
	return ret
}

// Fills 'Vertices' with Vertices with seeds generated from xorshift32.
func generateSeeds() {
	Vertices = make([]Vertex, VertexAmount, VertexAmount)
	Unvisited = make(map[int]bool)
	Vertices[0] = Vertex{xorshift32(Seed), nil, -1}
	for i := 1; i < VertexAmount; i++ {
		Vertices[i] = Vertex{xorshift32(Vertices[i-1].Seed ^ Seed), nil, -1}
	}
	for i := 0; i < VertexAmount; i++ {
		Unvisited[i] = false
	}
}

// Provided hash function.
func hashRand(inIndex int32) int32 {
	const b int32 = 0x5f375a86 //bunch of random bits
	for i := 0; i < 8; i++ {
		inIndex = (inIndex + 1) * ((inIndex >> 1) ^ b)
	}
	return inIndex
}

// Calculate the sum of all h(w) for all edges in MST.
func mstToInt(mst []*Edge) int32 {
	var total int32 = 0
	for i := 0; i < len(mst); i++ {
		total += hashRand(mst[i].Weight)
	}
	return total
}

// 2 arguments: <seed> <vertex amount>
//  Generate a fully connected graph with <vertex amount> vertices
// 3 arguments: <seed> <number of columns> <number of rows>
//  Generate a graph with connected rows and columns with dimensions
//  <number of columns> * <number of rows>
// 4 arguments: <seed> <filename> <vertex amount> <edge amount>
//  Generate a graph based on file -> see 'readGraph()'
func main() {
	debug.SetGCPercent(-1)
	args := os.Args[1:]
	switch len(args) {
	case 2:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		vertexAmount, _ := strconv.Atoi(args[1])
		VertexAmount = vertexAmount
		generateSeeds()
		generateComplete()
	case 3:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		X, _ := strconv.Atoi(args[1])
		Y, _ := strconv.Atoi(args[2])
		VertexAmount = X * Y
		generateSeeds()
		generateGrid(X, Y)
	case 4:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		filename := args[1]
		vertexAmount, _ := strconv.Atoi(args[2])
		VertexAmount = vertexAmount
		numOfEdges, _ := strconv.Atoi(args[3])
		generateSeeds()
		readGraph(filename, numOfEdges)
	}

	mst := MST()
	fmt.Println(mstToInt(mst))
}
