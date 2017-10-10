package main

import "fmt"

var N int = 3

func generateMatrix(n int) []int {
	matrix := make([]int, n*n)
	for i := 0; i < n*n; i++ {
		matrix[i] = i
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

func main() {
	a := generateMatrix(N)
	mat := multiplyMatrices(a, a, N)
	for f := 0; f < N*N; f++ {
		fmt.Printf("%d ", mat[f])
	}
}
