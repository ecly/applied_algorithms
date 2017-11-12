package main

import (
	"fmt"
	"os"
)

const cost = 1
const takeA = 'a'
const takeB = 'b'
const takeBoth = '|'
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

// needleman-wunsch score two strings and return last line
// of the generated matrix
// -- based on wiki
// https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm
func nwScore(a string, b string) []int {
	score := make([][]int, len(a)+1)
	for i := range score {
		mi := make([]int, len(b)+1)
		for j := range mi {
			mi[j] = -1
		}
		score[i] = mi
	}
	score[0][0] = 0
	for j := 1; j < len(b); j++ {
		score[0][j] = score[0][j-1] + cost
	}
	for i := 1; i < len(a); i++ {
		score[i][0] = score[i-1][0] + cost
		for j := 1; j < len(b); j++ {
			takeBothCost := 0
			if a[i-1] != b[j-1] {
				takeBothCost = cost
			}
			score[i][j] = min(takeBothCost+score[i-1][j-1],
				min(cost+score[i-1][j],
					cost+score[i][j-1]))
		}
	}
	//return the last line of score matrix
	return score[len(a)]
}

func trace(i int, j int, a string, b string, m [][]int) string {
	if i == 0 {
		return b[0:j]
	}
	if j == 0 {
		return a[0:i]
	}
	if m[i][j]-m[i][j-1] == cost {
		prev := trace(i, j-1, a, b, m)
		return prev + "b"
	}
	if m[i][j]-m[i-1][j] == cost {
		prev := trace(i-1, j, a, b, m)
		return prev + "a"
	}
	prev := trace(i-1, j-1, a, b, m)
	return prev + "|"
}

func needlemanWunsch(i int, j int, a string, b string, m [][]int) int {
	if m[i][j] != -1 {
		return m[i][j]
	}
	if len(a) < len(b) {
		a, b = b, a
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

	m := make([][]int, len(a)+1)
	for i := range m {
		mi := make([]int, len(b)+1)
		for j := range mi {
			mi[j] = -1
		}
		m[i] = mi
	}
	fmt.Printf("A: %s, B: %s\n", a, b)
	distance := needlemanWunsch(len(a), len(b), a, b, m)

	printMatrix(m)

	output := trace(len(a), len(b), a, b, m)
	return distance, output
}

// i being index in a, j being index in b and m being the memoizer
// recursive hirschberg call
func hirschberg(a string, b string) (int, string) {
	distance := -1
	output := ""
	if len(a) == 0 {
		for range b {
			distance = distance + cost
			output = output + "b"
		}
	} else if len(b) == 0 {
		for range a {
			distance = distance + cost
			output = output + "a"
		}
	} else if len(a) == 1 || len(b) == 1 {
		distance, output = NeedlemanWunsch(a, b)
	} else {
		alen := len(a)
		amid := alen / 2
		blen := len(b)

		scoreL := nwScore(a[:amid], b)
		scoreR := nwScore(rev(a[amid+1:alen]), rev(b))
		bmid := minSumIndex(scoreL, revInt(scoreR))

		distanceUpper, outputUpper := hirschberg(a[:amid], b[:bmid])
		distanceLower, outputLower := hirschberg(b[amid+1:alen], b[bmid+1:blen])
		distance = distanceUpper + distanceLower
		output = outputUpper + outputLower
	}

	return distance, output
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	//distance, output := hirschberg(a, b)
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
