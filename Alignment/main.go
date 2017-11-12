package main

import (
	"fmt"
	"os"
	"strings"
)

const cost = 1
const takeA = "a"
const takeB = "b"
const takeBoth = "|"
const inf = 1<<63 - 1

// due to math's min being float
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//reverses a string
//https://github.com/golang/example/blob/master/stringutil/reverse.go
func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//reverses an intenger slice
func revInt(s []int) []int {
	r := []int(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

// returns the index of min value of x+y
func minSumIndex(x []int, y []int) int {
	minVal := inf
	minIndex := -1
	for i := 0; i < len(x); i++ {
		val := x[i] + y[i]
		if val < minVal {
			minVal = val
			minIndex = i
		}
	}
	return minIndex
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
	return matrix
}

// needleman-wunsch score two strings and return last line
// of the generated matrix
// -- based on wiki
// https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm
func nwScore(a string, b string) []int {
	score := generateMatrix(len(a)+1, len(b)+1)
	score[0][0] = 0

	needlemanWunsch(len(a), len(b), a, b, score)
	// fmt.Printf("A: %s, B: %s\n", a, b)
	fmt.Println("Matrix")
	printMatrix(score)

	//return the last line of score matrix
	return score[len(a)][1:]
}

func trace(i int, j int, a string, b string, m [][]int) string {
	if i == 1 && j == 1 {
		return takeBoth
	}
	if i == 1 {
		return strings.Repeat(takeA, j)
	}
	if j == 1 {
		return strings.Repeat(takeB, i)
	}
	if m[i][j]-m[i-1][j-1] != cost {
		prev := trace(i-1, j-1, a, b, m)
		return prev + takeBoth
	}
	if m[i][j]-m[i-1][j] == cost {
		prev := trace(i-1, j, a, b, m)
		return prev + takeA
	}
	if m[i][j]-m[i][j-1] == cost {
		prev := trace(i, j-1, a, b, m)
		return prev + takeB
	}
	prev := trace(i-1, j-1, a, b, m)
	return prev + takeBoth
}

func needlemanWunsch(i int, j int, a string, b string, m [][]int) int {
	if m[i][j] != -1 {
		return m[i][j]
	}
	res := -1
	if i == 0 {
		res = j * cost
	} else if j == 0 {
		res = i * cost
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
	if len(a) > len(b) {
		a, b = b, a
	}

	m := generateMatrix(len(a)+1, len(b)+1)
	m[0][0] = 0

	// fmt.Printf("A: %s, B: %s\n", a, b)
	distance := needlemanWunsch(len(a), len(b), a, b, m)

	printMatrix(m)

	output := trace(len(a), len(b), a, b, m)
	return distance, output
}

// i being index in a, j being index in b and m being the memoizer
func hirschberg(a string, b string) (int, string) {
	distance := -1
	output := ""
	if len(a) == 0 {
		distance = distance + cost*len(b)
		output = output + strings.Repeat(takeB, len(b))
	} else if len(b) == 0 {
		distance = distance + cost*len(a)
		output = output + strings.Repeat(takeA, len(a))
	} else if len(a) == 1 || len(b) == 1 {
		distance, output = NeedlemanWunsch(a, b)
	} else {
		alen := len(a)
		amid := alen / 2

		scoreL := nwScore(a[:amid], b)
		scoreR := nwScore(rev(a[amid+1:]), rev(b))
		bmid := minSumIndex(scoreL, revInt(scoreR))

		distanceUpper, outputUpper := hirschberg(a[:amid], b[:bmid])
		distanceLower, outputLower := hirschberg(b[amid+1:], b[bmid+1:])
		distance = distanceUpper + distanceLower
		output = outputUpper + outputLower
	}

	return distance, output
}

// Hirschberg outer call
func Hirschberg(a string, b string) (int, string) {
	if len(a) > len(b) {
		return hirschberg(b, a)
	}
	return hirschberg(a, b)
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	//distance, output := Hirschberg(a, b)
	distance, output := NeedlemanWunsch(a, b)
	fmt.Printf("Distance %d, Output: %s \n", distance, output)
}

func printMatrix(m [][]int) {
	for _, mi := range m {
		for _, val := range mi {
			fmt.Print(val)
		}
		fmt.Println()
	}
}
