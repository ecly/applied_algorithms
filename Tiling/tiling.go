package main

import (
	"fmt"
	"os"
	"strconv"
    "time"
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

//TODO
/*func multiplyMatricesTiled(a []int, b []int, n int) []int {
}*/

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
	n_pow, _ := strconv.ParseUint(os.Args[1], 10, 64)
	s_max_pow, _ := strconv.ParseUint(os.Args[2], 10, 64)
    n := 1 << n_pow

    // time without tiling
	a := generateMatrix(n)
	start := time.Now()
    _ = transposeMatrix(a, n)
    fmt.Printf("Transpose with n=%d took %s\n", n, time.Since(start))

    for s := 1; s < 1 << s_max_pow; s = s << 1{
        // time with tiling
        start = time.Now()
        _ = transposeMatrixTiled(a, n, s)
        fmt.Printf("Tiled transpose with n=%d and s=%d, took %s\n", n, s, time.Since(start))
    }

    /*
    if sumMatrix(b,n) == sumMatrix(c,n){
        fmt.Print("Resulting matrices are identical.\n")
    } else {
        fmt.Print("Resulting matrices are different!\n")
    }
    */
}