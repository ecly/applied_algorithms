package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// a graph tracking only degree of vertices
type graph []int

func readVertices() graph {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(words[0])
	graph := make(graph, n)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		graph[x]++
		graph[y]++
	}
	return graph
}

func sumOfDegreesSquared(graph graph) int {
	sum := 0
	for _, degree := range graph {
		sum += degree * degree
	}
	return sum
}

func main() {
	graph := readVertices()
	fmt.Printf("%d", sumOfDegreesSquared(graph))
}
