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
	matrix[0][0] = 0
	return matrix
}

func trace(i int, j int, a string, b string, m [][]int, takeA string, takeB string) string {
	//fmt.Printf("m[%d][%d] == %d - ", i, j, m[i][j])
	if i == 0 {
		//fmt.Printf("Taking: %s*%d\n", takeB, j)
		return strings.Repeat(takeB, j)
	}
	if j == 0 {
		// fmt.Printf("Taking: %s*%d\n", takeA, i)
		return strings.Repeat(takeA, i)
	}
	if m[i][j]-m[i][j-1] == cost {
		// fmt.Printf("Taking: %s\n", takeB)
		prev := trace(i, j-1, a, b, m, takeA, takeB)
		return prev + takeB
	}
	if m[i][j]-m[i-1][j-1] == cost {
		// fmt.Printf("Taking: %s\n", takeBoth)
		prev := trace(i-1, j-1, a, b, m, takeA, takeB)
		return prev + takeBoth
	}
	if m[i][j]-m[i-1][j] == cost {
		// fmt.Printf("Taking: %s\n", takeA)
		prev := trace(i-1, j, a, b, m, takeA, takeB)
		return prev + takeA
	}
	if m[i][j]-m[i-1][j-1] == 0 {
		// fmt.Printf("Taking: %s\n", takeBoth)
		prev := trace(i-1, j-1, a, b, m, takeA, takeB)
		return prev + takeBoth
	}
	if m[i][j]-m[i][j-1] == 0 {
		// fmt.Printf("Taking: %s\n", takeB)
		prev := trace(i, j-1, a, b, m, takeA, takeB)
		return prev + takeB
	}
	if m[i][j]-m[i-1][j] == 0 {
		// fmt.Printf("Taking: %s\n", takeA)
		prev := trace(i-1, j, a, b, m, takeA, takeB)
		return prev + takeA
	}
	// fmt.Printf("Taking: %s\n", takeBoth)
	prev := trace(i-1, j-1, a, b, m, takeA, takeB)
	return prev + takeBoth
}

func needlemanWunsch(i int, j int, a string, b string, m [][]int) int {
	if m[i][j] != -1 {
		return m[i][j]
	}
	res := -1
	if i == 0 {
		if strings.Contains(b[:j], a[:1]) {
			res = j*cost - 1
		} else {
			res = j * cost
		}
		m[i][j] = res
	} else if j == 0 {
		if strings.Contains(a[:i], b[:1]) {
			res = i*cost - 1
		} else {
			res = i * cost
		}
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
	traceTakeA, traceTakeB := takeA, takeB
	if len(a) > len(b) {
		a, b = b, a
		//fmt.Println("Swapping")
		traceTakeA, traceTakeB = traceTakeB, traceTakeA
	}
	m := generateMatrix(len(a)+1, len(b)+1)
	distance := needlemanWunsch(len(a), len(b), a, b, m)
	printMatrix(m)
	output := trace(len(a), len(b), a, b, m, traceTakeA, traceTakeB)
	return distance, output
}

// needleman-wunsch score two strings and return last line
// of the generated matrix
// -- based on wiki
// https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm
func nwScore(a string, b string) []int {
	score := generateMatrix(len(a)+1, len(b)+1)
	needlemanWunsch(len(a), len(b), a, b, score)
	//return the last line of score matrix without the first -1
	return score[len(a)][1:]
}

// i being index in a, j being index in b and m being the memoizer
func hirschberg(a string, b string, takeA string, takeB string) (int, string) {
	distance := 0
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

		distanceUpper, outputUpper := hirschberg(a[:amid], b[:bmid], takeA, takeB)
		fmt.Printf("Distance upper: %d\n", distanceUpper)
		distanceLower, outputLower := hirschberg(b[amid+1:], b[bmid+1:], takeA, takeB)
		fmt.Printf("Distance lower: %d\n", distanceLower)
		distance = distanceUpper + distanceLower
		output = outputUpper + outputLower
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
	//distance, output := Hirschberg(a, b)
	_, output := NeedlemanWunsch(a, b)
	fmt.Println(output)
	//fmt.Printf("Distance %d, Output: %s \n", distance, output)
}

func printMatrix(m [][]int) {
	for _, mi := range m[:] {
		fmt.Printf("%v\n", mi[:])
	}
}
