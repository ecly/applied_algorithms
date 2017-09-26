package main

import (
	"fmt"
)

// Globals
var Seed int32
var VertexAmount int32
var VertexSeed []int32

type Edge struct {
	Y      int32
	X      int32
	Weight int32
}

func MST(graph []int32) {
	fmt.Println("Solved")
}

func generateComplete(seed int32, numOfVertices int32) []int32 {
	complete := []int32{1}
	return complete
}

func generateGrid(seed int32, numX int32, numY int32) []int32 {
	complete := []int32{1}
	return complete
}

func readGraph(seed int32, filename string, numOfVertices int32, numOfEdges int32) []int32 {
	complete := []int32{1}
	return complete
}

const EDGE_MOD = 100000

func getEdgeWeight(v1 int32, v2 int32) int32 {
	return xorshift32(VertexSeed[v1]^VertexSeed[v2]) % EDGE_MOD
}

func xorshift32(seed int32) int32 {
	ret := seed
	ret ^= ret << 13
	ret ^= ret >> 17
	ret ^= ret << 5
	return ret
}

func generateSeeds(seed int32) {
	VertexSeed[0] = xorshift32(seed)
	for i := int32(1); i < VertexAmount; i++ {
		VertexSeed[i] = xorshift32(VertexSeed[i-1] ^ seed)
	}
}

func hashRand(inIndex int32) int32 {
	const b = 0x5f375a86 //bunch of random bits
	for i := 0; i < 8; i++ {
		inIndex = (inIndex + 1) * ((inIndex >> 1) ^ b)
	}
	return inIndex
}

func mstToint32(mst []Edge, mstsize int32) int32 {
	total := int32(0)
	for i := int32(0); i < mstsize; i++ {
		total += hashRand(mst[i].Weight)
	}
	return total
}

func createEdge(v1 int32, v2 int32) Edge {
	return Edge{v1, v2, getEdgeWeight(v1, v2)}
}

func main() {
	//args := os.Args[1:]
	//var graph []int32
	/*
		switch len(args) {
		case 2:
			Seed, _ := strconv.Atoi(args[0])
			VertexAmount, _ := strconv.Atoi(args[1])
			graph = generateComplete(seed, numOfVertices)
		case 3:
			Seed, _ := strconv.Atoi(args[0])
			numX, _ := strconv.Atoi(args[1])
			numY, _ := strconv.Atoi(args[2])
			NumOfVertices = numX * numY
			graph = generateGrid(seed, numX, numY)
		case 4:
			Seed, _ := strconv.Atoi(args[0])
			filename := args[1]
			VertexAmount, _ := strconv.Atoi(args[2])
			numOfEdges, _ := strconv.Atoi(args[3])
			graph = readGraph(seed, filename, numOfVertices, numOfEdges)
		}
	*/

	VertexSeed = []int32{899619065, 857390176, -1009232974}
	Seed = 1597463007
	Graph := []Edge{createEdge(0, 1), createEdge(0, 2), createEdge(1, 2)}
	fmt.Println(mstToint32(Graph, 3))
	//MST(graph)
}
