package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A BitVector represented as 4 uint64 with an additional
// index indicating it's index in the original slice of BitVectors.

type graph [][]int

func readVertices() graph {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(words[0])
	graph := make(graph, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	return graph
}

func sumOfDegreesSquared(graph graph) int {
	sum := 0
	for _, edges := range graph {
		sum += len(edges) * len(edges)
	}
	return sum
}

func main() {
	graph := readVertices()
	fmt.Printf("%d", sumOfDegreesSquared(graph))
}
