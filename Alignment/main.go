package main

import (
    "fmt"
    "os"
)

const cost = 1
const takeA = 'a'
const takeB = 'b'
const takeBoth = '|'

// due to math's min being float
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// i being index in a, j being index in b and m being the memoizer
// recursive hirshberg call
func hirshberg(i int, j int, a string, b string, m [][]int) int {
    if m[i][j] != -1 {
        return m[i][j]
    }

    for ii, row := range m {
        for jj, col := range row {

        }
    }

    res := -1
    if i == 0 {
        res = j*cost
    } else if j == 0 {
        res = i*cost
    } else {
        takeBothCost := 0
        if a[i-1] != b[j-1] {
            takeBothCost = cost
        }
        res = min(takeBothCost + hirshberg(i-1, j-1, a, b, m),
              min(cost+hirshberg(i-1, j, a, b, m),
                  cost+hirshberg(i, j-1, a, b, m)))
        m[i][j] = res
    }

    return res
}

// Hirshberg outer call
func Hirshberg(a string, b string) (int, string) {
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
    res := hirshberg(len(a), len(b), a, b, m)

    return res, ""
}

func main(){
	a := os.Args[1]
	b := os.Args[2]
    fmt.Println(Hirshberg(a,b))
}
