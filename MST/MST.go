package main

import (
	"fmt"
	"os"
	"strconv"
)

// Globals
var Seed int32
var VertexAmount int32
var VertexSeed []int32
var Graph []Edge

type Edge struct {
	X      int32
	Y      int32
	Weight int32
}

func MST(graph []Edge) []Edge {
	return []Edge{Edge{1, 2, 3}}
}

// generate a fully connected graph
func generateComplete() []Edge {
	// at least this big for starters
	var graph []Edge
	for i := int32(0); i < VertexAmount; i++ {
		for j := i + 1; j < VertexAmount; j++ {
			graph = append(graph, Edge{i, j, getEdgeWeight(i, j)})
		}
	}
	return graph
}

func generateGrid(numX int32, numY int32) []Edge {
	// at least this big for starters
	size := numX * numY
	graph := make([]Edge, 0, size)

	// row connections
	for j := int32(0); j < numY; j++ {
		start := j * numX
		end := start + numX - 1
		for i := start; i < end; i++ {
			//fmt.Printf("ROW: from: %d, to: %d \n", i, i+1)
			graph = append(graph, Edge{i, i + 1, getEdgeWeight(i, i+1)})
		}
	}

	//column connections
	for i := int32(0); i < numX; i++ {
		for j := int32(0); j < numY-1; j++ {
			from := i + numX*j
			to := from + numX
			//fmt.Printf("COL: from: %d, to: %d \n", from, to)
			graph = append(graph, Edge{from, to, getEdgeWeight(from, to)})
		}
	}
	return graph
}

/* Read a graph from file with each lines being of format v1<tab>v2
** indicating an edge between the two vertices.
** Vertex number should always be less than 'VertexAmount'
** and there should be at most 'numOfEdges' edges
 */
func readGraph(filename string, numOfEdges int32) []Edge {
	complete := []Edge{Edge{1, 2, 3}}
	return complete
}

func getEdgeWeight(v1 int32, v2 int32) int32 {
	//fmt.Printf("v1: %d, v2: %d\n", v1, v2)
	return xorshift32(VertexSeed[v1]^VertexSeed[v2]) % 100000
}

func xorshift32(seed int32) int32 {
	ret := seed
	ret ^= ret << uint32(13)
	ret ^= ret >> uint32(17)
	ret ^= ret << uint32(5)
	return ret
}

func generateSeeds() {
	VertexSeed = make([]int32, VertexAmount, VertexAmount)
	VertexSeed[0] = xorshift32(Seed)
	for i := int32(1); i < VertexAmount; i++ {
		VertexSeed[i] = xorshift32(VertexSeed[i-1] ^ Seed)
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
	var graph []Edge

	switch len(args) {
	case 2:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)

		vertexAmount, _ := strconv.Atoi(args[1])
		VertexAmount = int32(vertexAmount)

		generateSeeds()
		graph = generateComplete()
	case 3:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)

		numX, _ := strconv.Atoi(args[1])
		numY, _ := strconv.Atoi(args[2])

		X := int32(numX)
		Y := int32(numY)
		VertexAmount = X * Y

		generateSeeds()
		for i := 0; i < len(VertexSeed); i++ {
			fmt.Printf("Seed: %d, Value: %d\n", i, VertexSeed[i])
		}
		graph = generateGrid(X, Y)
	case 4:
		seed, _ := strconv.Atoi(args[0])
		Seed = int32(seed)
		filename := args[1]
		vertexAmount, _ := strconv.Atoi(args[2])
		numOfEdges, _ := strconv.Atoi(args[3])
		VertexAmount = int32(vertexAmount)
		graph = readGraph(filename, int32(numOfEdges))
	}

	for i := 0; i < len(graph); i++ {
		x := graph[i].X
		y := graph[i].X
		w := graph[i].Weight
		fmt.Printf("X: %d, Y: %d, W: %d\n", x, y, w)
	}

	//mst := MST(graph)
	fmt.Println(mstToInt(graph))
}
