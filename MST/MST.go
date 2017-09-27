package main

import (
	"bufio"
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

type Vertex struct {
	Seed     int32
	Edges    []Edge
	Visisted bool
}

func (v Vertex) AddEdge(edge Edge) {
	v.Edges = append(v.Edges, edge)
}

type Edge struct {
	To     int32
	Weight int32
}

//https://golang.org/pkg/container/heap/#example__intHeap
type Fringe []Edge

func (f Fringe) Len() int           { return len(f) }
func (f Fringe) Less(i, j int) bool { return f[i].Weight < f[j].Weight }
func (f Fringe) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func (f *Fringe) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*f = append(*f, x.(Edge))
}

func (f *Fringe) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

func MST() []Edge {
	mst := make([]Edge, 0, VertexAmount-1) // minimum size
	var fringe Fringe
	Vertices[0].Visisted = true
	for len(fringe) > 0 {
		for _, edge := range Vertices[0].Edges {
			if !Vertices[edge.To].Visisted {
				//TODO only if no other cheaper edge in fringe is pointing to same
				fringe.Push(edge)
			}
		}
	}
	return mst
}

// generate a fully connected graph
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

/* Read a graph from file with each lines being of format v1<tab>v2
** indicating an edge between the two vertices.
** Vertex number should always be less than 'VertexAmount'
** and there should be at most 'numOfEdges' edges
 */
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

func getEdgeWeight(v1 int32, v2 int32) int32 {
	return xorshift32(Vertices[v1].Seed^Vertices[v2].Seed) % 100000
}

func xorshift32(seed int32) int32 {
	ret := seed
	ret ^= ret << uint32(13)
	ret ^= ret >> uint32(17)
	ret ^= ret << uint32(5)
	return ret
}

func generateSeeds() {
	Vertices = make([]Vertex, VertexAmount, VertexAmount)
	Vertices[0] = Vertex{xorshift32(Seed), nil, false}
	for i := int32(1); i < VertexAmount; i++ {
		Vertices[i] = Vertex{xorshift32(Vertices[i-1].Seed ^ Seed), nil, false}
	}
}

func hashRand(inIndex int32) int32 {
	const b = int32(0x5f375a86) //bunch of random bits
	for i := 0; i < 8; i++ {
		inIndex = (inIndex + 1) * ((inIndex >> 1) ^ b)
	}
	return inIndex
}

func mstToInt(mst []Edge) int32 {
	total := int32(0)
	for i := 0; i < len(mst); i++ {
		total += hashRand(mst[i].Weight)
	}
	return total
}

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

	mst := MST()
	fmt.Println(mstToInt(mst))
}
