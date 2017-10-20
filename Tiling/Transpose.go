package main

import (
	"fmt"
	"os"
	"strconv"
)

var MAXVAL = 1 << 16

func generateMatrix(n int) []int {
	matrix := make([]int, n*n)
	for i := 0; i < n*n; i++ {
		matrix[i] = xorshift64(i) % MAXVAL
	}
	return matrix
}

func xorshift64(seed int) int {
	ret := seed
	ret ^= ret >> 12
	ret ^= ret << 25
	ret ^= ret >> 27
	return ret
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


func transposeMatrix(a []int, n int) []int {
	mat := make([]int, n*n)
    for i := 0; i < n; i ++{
        for j := 0; j < n; j++{
            mat[i+j*n] = a[j+i*n]
        }
    }
    return mat
}

func transposeMatrixTiled(a []int, n int, s int) []int {
	mat := make([]int, n*n)
	for ii := 0; ii < n; ii++ {
		l := ii + s
		if l > n {
			l = n
		}
		for j := 0; j < n; j++ {
            for i := ii; i < l; i++{
                mat[i+j*n] = a[j+i*n]
            }
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
	//b := transposeMatrix(a, n)
    b := transposeMatrixTiled(a, n, s)
	fmt.Println("\n\nafter")
	for f := 0; f < n*n; f++ {
		if f%n == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", b[f])
	}
}
