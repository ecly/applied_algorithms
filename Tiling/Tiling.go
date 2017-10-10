package main

import (
	"fmt"
	"os"
	"strconv"
)

var MAXVAL = 1 << 8

func generateMatrix(n int) []int {
	matrix := make([]int, n*n)
	for i := 0; i < n*n; i++ {
		matrix[i] = xorshift64(i) % MAXVAL
	}
	return matrix
}

func multiplyMatrices(a []int, b []int, n int) []int {
	mat := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := a[i*n+j]
			for k := 0; k < n; k++ {
				mat[k+i*n] += x * b[j*n+k]
			}
		}
	}
	return mat
}

func xorshift64(seed int) int {
	ret := seed
	ret ^= ret >> 12
	ret ^= ret << 25
	ret ^= ret >> 27
	return ret
}

func transposMatrix(a []int, n int, s int) []int {
	N = n * n
	mat := make([]int, N)
	for ii := 0; ii < N; ii++ {
		l := i + s
		if l > N {
			l = N
		}
		for j := 1; j < N; j++ {
			for i = ii; 
			mat[j*n+i] = a[i*n+j]

		}
	}
	return mat
}

func sumMatrix(a []int, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += a[i]
	}
	return total
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	s, _ := strconv.Atoi(os.Args[2])
	a := generateMatrix(n)
	fmt.Println("before")
	for f := 0; f < n*n; f++ {
		if f%n == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", a[f])
	}
	b := transposMatrix(a, n)
	fmt.Println("\n\nafter")
	for f := 0; f < n*n; f++ {
		if f%n == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", b[f])
	}
}
