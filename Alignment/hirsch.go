package main

import (
	"fmt"
	"os"
	"strings"
)

const cost = 1
const takeBoth = "|"
const inf = 1<<63 - 1
const takeA = "a"
const takeB = "b"

// due to math's min being float
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rev(s string) string {
	size := len(s)
	r := make([]byte, size)
	for i := 0; i < size; i++ {
		r[i] = s[size-1-i]
	}
	return string(r)
}

//returns the reverse slice of given slice s
func revInt(s []int) []int {
	size := len(s)
	r := make([]int, size)
	for i := 0; i < size; i++ {
		r[i] = s[size-1-i]
	}
	return r
}

// returns the index of min value of x+y and the value
func minSum(x []int, y []int) (int, int) {
	minVal := inf
	minIndex := -1
	for i := 0; i < len(x); i++ {
		val := x[i] + y[i]
		if val < minVal {
			minVal = val
			minIndex = i
		}
	}
	return minIndex, minVal
}

func generateMatrix(n int, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		mi := make([]int, m)
		for j := 0; j < m; j++ {
			mi[j] = -1
		}
		matrix[i] = mi
	}
	matrix[0][0] = 0
	return matrix
}

// A trace function favoring a, to fit CodeJudge output
func traceRightLean(i int, j int, a string, b string, m [][]int, takeA string, takeB string) string {
	//fmt.Printf("m[%d][%d] = %d\n", i, j, m[i][j])
	if i == len(a) {
		return strings.Repeat(takeB, len(b)-j)
	}
	if j == len(b) {
		return strings.Repeat(takeA, len(a)-i)
	}

	aCost := m[i+1][j] - m[i][j]
	bCost := m[i][j+1] - m[i][j]
	bothCost := m[i+1][j+1] - m[i][j]
	if aCost <= bCost && aCost <= bothCost {
		return takeA + traceRightLean(i+1, j, a, b, m, takeA, takeB)
	}
	return trace(i, j, a, b, m, takeA, takeB)
}

func trace(i int, j int, a string, b string, m [][]int, takeA string, takeB string) string {
	//fmt.Printf("m[%d][%d] = %d\n", i, j, m[i][j])
	if i == len(a) {
		return strings.Repeat(takeB, len(b)-j)
	}
	if j == len(b) {
		return strings.Repeat(takeA, len(a)-i)
	}
	aCost := m[i+1][j] - m[i][j]
	bCost := m[i][j+1] - m[i][j]
	bothCost := m[i+1][j+1] - m[i][j]
	if bothCost <= bCost && bothCost <= aCost {
		return takeBoth + trace(i+1, j+1, a, b, m, takeA, takeB)
	} else if aCost <= bCost && aCost <= bothCost {
		return takeA + trace(i+1, j, a, b, m, takeA, takeB)
	} else if bCost <= bothCost && bCost <= aCost {
		return takeB + trace(i, j+1, a, b, m, takeA, takeB)
	}
	return "NOPE"
}

func needlemanWunsch(i int, j int, a string, b string, m [][]int) int {
	if m[i][j] != -1 {
		return m[i][j]
	}
	res := -1
	if i == 0 {
		res = j * cost
		m[i][j] = res
	} else if j == 0 {
		res = i * cost
		m[i][j] = res
	} else {
		takeBothCost := 0
		if a[i-1] != b[j-1] {
			takeBothCost = cost
		}
		res = min(takeBothCost+needlemanWunsch(i-1, j-1, a, b, m),
			min(cost+needlemanWunsch(i-1, j, a, b, m),
				cost+needlemanWunsch(i, j-1, a, b, m)))
		m[i][j] = res
	}
	return res
}

// NeedlemanWunsch outer call
func NeedlemanWunsch(a string, b string) (int, string) {
	m := generateMatrix(len(a)+1, len(b)+1)
	distance := needlemanWunsch(len(a), len(b), a, b, m)
	//printMatrix(m)
	output := trace(0, 0, a, b, m, takeA, takeB)
	return distance, output
}

// needleman-wunsch score two strings and return last line
// of the generated matrix
// -- based on wiki
// https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm
func nwScore(a string, b string) []int {
	//fmt.Printf("Comparing %s and %s\n", a, b)
	score := generateMatrix(len(a)+1, len(b)+1)
	needlemanWunsch(len(a), len(b), a, b, score)
	//printMatrix(score)
	//return the last line of score matrix without the first -1
	return score[len(score)-1]
}

// i being index in a, j being index in b and m being the memoizer
func hirschberg(a string, b string, takeA string, takeB string) (int, string) {
	distance := 0
	output := ""
	if len(a) == 0 {
		distance = distance + cost*len(b)
		output = output + strings.Repeat(takeB, len(b))
		//fmt.Printf("output for %s with %s = %s\n", a, b, output)
	} else if len(b) == 0 {
		distance = distance + cost*len(a)
		output = output + strings.Repeat(takeA, len(a))
		//fmt.Printf("output for %s with %s = %s\n", a, b, output)
	} else if len(a) == 1 || len(b) == 1 {
		distance, output = NeedlemanWunsch(a, b)
		//fmt.Printf("output for %s with %s = %s\n", a, b, output)
	} else {
		alen := len(a)
		amid := alen / 2

		scoreL := nwScore(a[:amid], b)
		scoreR := nwScore(rev(a[amid:]), rev(b))
		//fmt.Printf("scoreL %v\n", scoreL)
		//fmt.Printf("scoreR %v\n", revInt(scoreR))
		var bsplit int
		bsplit, distance = minSum(scoreL, revInt(scoreR))
		//fmt.Println("bsplit: ", bsplit)

		_, outputUpper := hirschberg(a[:amid], b[:bsplit], takeA, takeB)
		_, outputLower := hirschberg(b[amid:], b[bsplit:], takeA, takeB)
		//fmt.Printf("split1: '%s' with '%s'\n", a[:amid], b[:bsplit])
		//fmt.Printf("split2: '%s' with '%s'\n", a[amid:], b[bsplit:])
		//fmt.Printf("outputUpper: %s\n", outputUpper)
		//fmt.Printf("outputLower: %s\n", outputLower)
		output = outputUpper + rev(outputLower)
	}

	return distance, output
}

// Hirschberg outer call
func Hirschberg(a string, b string) (int, string) {
	if len(a) > len(b) {
		return hirschberg(b, a, takeB, takeA)
	}
	return hirschberg(a, b, takeA, takeB)
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	distance, output := Hirschberg(a, b)
	//fmt.Println(distance)
	//distance, output := NeedlemanWunsch(a, b)
	//fmt.Println(output)
	fmt.Printf("Distance %d, Output: %s \n", distance, output)
}

func printMatrix(m [][]int) {
	for _, mi := range m[:] {
		fmt.Printf("%v\n", mi[:])
	}
}
