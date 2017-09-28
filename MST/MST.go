package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Globals
var Seed int32
var VertexAmount int32
var Vertices []Vertex

// Seed: the seed of the vertex.
// Edges: edges connected to the vertex.
// Visited: whether is has already been visited by MST().
// FringeBy: a pointer to potential item in fringe pointing to this vertex.
type Vertex struct {
	Seed     int32
	Edges    []Edge
	Visisted bool
	FringeBy *Item
}

func (v *Vertex) AddEdge(edge Edge) {
	v.Edges = append(v.Edges, edge)
}

type Edge struct {
	To     int32
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
	heap.Fix(f, item.Index)
}

// Update modifies the weight and value of an Item in the heap
// and adjusts its position in the heap based on the new weight.
func (f *Fringe) update(item *Item, weight int32) {
	item.Edge.Weight = weight
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
func MST() []Edge {
	mst := make([]Edge, 0, VertexAmount-1) // minimum size
	fringe := &Fringe{}
	heap.Init(fringe)
	Vertices[0].Visisted = true
	// initial fringe
	for i, edge := range Vertices[0].Edges {
		fmt.Printf("Added E to: %d with Weight: %d to Fringe\n", edge.To, edge.Weight)
		item := &Item{i, &edge}
		heap.Push(fringe, item)
		heap.Fix(fringe, item.Index)
		Vertices[edge.To].FringeBy = item
	}
	for fringe.Len() > 0 {
		item := heap.Pop(fringe).(*Item)
		//fmt.Println("Pop1", heap.Pop(fringe).(Item).Edge.Weight)
		fmt.Println("Pop1", item.Edge.Weight)
	}

	for fringe.Len() > 0 {
		mstEdge := heap.Pop(fringe).(*Item).Edge
		//fmt.Println("Popped edge off with weight", mstEdge.Weight)
		Vertices[mstEdge.To].Visisted = true
		mst = append(mst, *mstEdge)
		for _, edge := range Vertices[mstEdge.To].Edges {
			if !Vertices[edge.To].Visisted {
				if Vertices[edge.To].FringeBy == nil {
					item := &Item{Edge: &edge}
					heap.Push(fringe, item)
					Vertices[edge.To].FringeBy = item
				} else {
					currentItem := Vertices[edge.To].FringeBy
					fmt.Println("hmm")
					if currentItem.Edge.Weight > edge.Weight {
						fringe.update(currentItem, edge.Weight)
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
			Vertices[i].AddEdge(Edge{j, weight})
			Vertices[j].AddEdge(Edge{i, weight})
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
			Vertices[i].AddEdge(Edge{to, weight})
			Vertices[to].AddEdge(Edge{i, weight})
		}
	}
	//column edges
	for i := int32(0); i < numX; i++ {
		for j := int32(0); j < numY-1; j++ {
			from := i + numX*j
			to := from + numX
			weight := getEdgeWeight(from, to)
			Vertices[from].AddEdge(Edge{to, weight})
			Vertices[to].AddEdge(Edge{from, weight})
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
			Vertices[x].AddEdge(Edge{y, weight})
			Vertices[y].AddEdge(Edge{x, weight})
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
	return xorshift32(Vertices[v1].Seed^Vertices[v2].Seed) % 100000
}

// Provided xorshift32 implementation.
func xorshift32(seed int32) int32 {
	ret := seed
	ret ^= ret << uint32(13)
	ret ^= ret >> uint32(17)
	ret ^= ret << uint32(5)
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
	const b = int32(0x5f375a86) //bunch of random bits
	for i := 0; i < 8; i++ {
		inIndex = (inIndex + 1) * ((inIndex >> 1) ^ b)
	}
	return inIndex
}

// Calculate the sum of all h(w) for all edges in MST.
func mstToInt(mst []Edge) int32 {
	total := int32(0)
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
	}
	/*for i, v := range Vertices {
		fmt.Printf("Vertex: %d, Seed:%d\n", i, v.Seed)
		for _, e := range v.Edges {
			fmt.Printf("To: %d, Weight: %d\n", e.To, e.Weight)
		}
		fmt.Println()
	}*/

	mst := MST()
	//mst := []Edge{Edge{0, -60078}, Edge{0, -78884}}
	fmt.Println("MST:")
	for _, e := range mst {
		fmt.Printf("To: %d, Weight: %d\n", e.To, e.Weight)
	}

	fmt.Println(mstToInt(mst))
}
