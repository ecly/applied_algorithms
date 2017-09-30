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
)

// Globals
var Seed int32
var VertexAmount int32
var Vertices []Vertex

// Seed: the seed of the vertex.
// Edges: pointers to edges connected to the vertex.
// Visited: whether is has already been visited by MST().
// FringeBy: a pointer to potential item in fringe pointing to this vertex.
type Vertex struct {
	Seed     int32
	Edges    []*Edge
	Visited  bool
	FringeBy *Item
}

// Adds the pointer to Edge 'edge' to the Slice of Edges for Vertex 'v'.
func (v *Vertex) AddEdge(edge *Edge) {
	v.Edges = append(v.Edges, edge)
}

// An Edge has a pointer to the Vertex it points at
// and a int32 Weight.
type Edge struct {
	To     *Vertex
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
	f[j].Index = j
}

// Add an item to the fringe.
func (f *Fringe) Push(x interface{}) {
	n := len(*f)
	item := x.(*Item)
	item.Index = n
	*f = append(*f, item)
}

// Update modifies the Edge pointer to the given edge pointer
// and adjusts its position in the heap based on the new edge's weight
func (f *Fringe) update(item *Item, edge *Edge) {
	item.Edge = edge
	heap.Fix(f, item.Index)
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
	Vertices[0].Visited = true
	fringe := make(Fringe, len(Vertices[0].Edges))

	// initial fringe is first vertex' edges
	for i, e := range Vertices[0].Edges {
		item := &Item{Edge: e}
		fringe[i] = item
		e.To.FringeBy = item
	}

	heap.Init(&fringe)

	for fringe.Len() > 0 {
		mstEdge := heap.Pop(&fringe).(*Item).Edge

		// this check won't be needed if we fix the heap.Update()
		if mstEdge.To.Visited {
			continue
		}

		mstEdge.To.Visited = true
		mst = append(mst, mstEdge)
		for _, edge := range mstEdge.To.Edges {
			if !edge.To.Visited {
				if edge.To.FringeBy == nil {
					item := &Item{Edge: edge}
					edge.To.FringeBy = item
					heap.Push(&fringe, item)
				} else {
					if edge.Weight < edge.To.FringeBy.Edge.Weight {
						// this is terrible, we always push
						item := &Item{Edge: edge}
						edge.To.FringeBy = item
						heap.Push(&fringe, item)
					}
				}
			}
		}
	}
	return mst
}

// Generate a fully connected graph with 'VertexAmount' vertices
func generateComplete() {
	// at least this big for starters
	for i := int32(0); i < VertexAmount; i++ {
		for j := i + 1; j < VertexAmount; j++ {
			weight := getEdgeWeight(i, j)
			Vertices[i].AddEdge(&Edge{&Vertices[j], weight})
			Vertices[j].AddEdge(&Edge{&Vertices[i], weight})
		}
	}
}

// Generate a 'numX' * 'numY' graph with connected rows and comlumns
func generateGrid(numX int32, numY int32) {
	// row edges
	for j := int32(0); j < numY; j++ {
		start := j * numX
		end := start + numX - 1
		for i := start; i < end; i++ {
			to := i + 1
			weight := getEdgeWeight(i, to)
			Vertices[i].AddEdge(&Edge{&Vertices[to], weight})
			Vertices[to].AddEdge(&Edge{&Vertices[i], weight})
		}
	}
	//column edges
	for i := int32(0); i < numX; i++ {
		for j := int32(0); j < numY-1; j++ {
			from := i + numX*j
			to := from + numX
			weight := getEdgeWeight(from, to)
			Vertices[from].AddEdge(&Edge{&Vertices[to], weight})
			Vertices[to].AddEdge(&Edge{&Vertices[from], weight})
		}
	}
}

// Read a graph from file with each lines being of format v1<tab>v2
// indicating an edge between the two vertices.
// Vertex number should always be less than 'VertexAmount'
// and there should be at most 'numOfEdges' edges.
func readGraph(filename string, numOfEdges int32) []Edge {
	graph := make([]Edge, 0, numOfEdges) // known size
	if file, err := os.Open(filename); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			numY, _ := strconv.Atoi(words[0])
			numX, _ := strconv.Atoi(words[1])
			x := int32(numX)
			y := int32(numY)
			weight := getEdgeWeight(x, y)
			Vertices[x].AddEdge(&Edge{&Vertices[y], weight})
			Vertices[y].AddEdge(&Edge{&Vertices[x], weight})
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	return graph
}

// Calculate the weight of a vertex based on its from, to vertices.
func getEdgeWeight(v1 int32, v2 int32) int32 {
	//weight :=
	//fmt.Printf("Edge from %d, to %d, weight: %d\n", v1, v2, weight)
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
	Vertices[0] = Vertex{xorshift32(Seed), nil, false, nil}
	for i := int32(1); i < VertexAmount; i++ {
		Vertices[i] = Vertex{xorshift32(Vertices[i-1].Seed ^ Seed), nil, false, nil}
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
// 	Generate a fully connected graph with <vertex amount> vertices
// 3 arguments: <seed> <number of columns> <number of rows>
// 	Generate a graph with connected rows and columns with dimensions
// 	<number of columns> * <number of rows>
// 4 arguments: <seed> <filename> <vertex amount> <edge amount>
// 	Generate a graph based on file -> see 'readGraph()'
func main() {
	debug.SetGCPercent(-1)
	args := os.Args[1:]
	switch len(args) {
	case 2:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		vertexAmount, _ := strconv.Atoi(args[1])
		VertexAmount = int32(vertexAmount)
		generateSeeds()
		generateComplete()
	case 3:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		numX, _ := strconv.Atoi(args[1])
		numY, _ := strconv.Atoi(args[2])
		X := int32(numX)
		Y := int32(numY)
		VertexAmount = X * Y
		generateSeeds()
		generateGrid(X, Y)
	case 4:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		filename := args[1]
		vertexAmount, _ := strconv.Atoi(args[2])
		numOfEdges, _ := strconv.Atoi(args[3])
		VertexAmount = int32(vertexAmount)
		generateSeeds()
		readGraph(filename, int32(numOfEdges))
	default:
		// handcoded
		VertexAmount = 6
		generateSeeds()
		Vertices[0].AddEdge(&Edge{&Vertices[1], 3})
		Vertices[1].AddEdge(&Edge{&Vertices[0], 3})
		Vertices[0].AddEdge(&Edge{&Vertices[2], 7})
		Vertices[2].AddEdge(&Edge{&Vertices[0], 7})

		Vertices[1].AddEdge(&Edge{&Vertices[4], 9})
		Vertices[4].AddEdge(&Edge{&Vertices[1], 9})
		Vertices[1].AddEdge(&Edge{&Vertices[2], 10})
		Vertices[2].AddEdge(&Edge{&Vertices[1], 10})
		Vertices[1].AddEdge(&Edge{&Vertices[3], 4})
		Vertices[3].AddEdge(&Edge{&Vertices[1], 4})

		Vertices[2].AddEdge(&Edge{&Vertices[3], 5})
		Vertices[3].AddEdge(&Edge{&Vertices[2], 5})

		Vertices[3].AddEdge(&Edge{&Vertices[5], 8})
		Vertices[5].AddEdge(&Edge{&Vertices[3], 8})

		Vertices[4].AddEdge(&Edge{&Vertices[5], 1})
		Vertices[5].AddEdge(&Edge{&Vertices[4], 1})
	}

	mst := MST()
	//fmt.Println("MST:")
	/*for _, e := range mst {
		fmt.Printf("Weight: %d\n", e.Weight)
	}*/

	//fmt.Println("Len:", len(mst))
	fmt.Println(mstToInt(mst))
}
