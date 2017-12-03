package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sparseEntry struct {
	otherDim int
	val      int
}

type sparseOutputCell struct {
	row int
	col int
	val int
}

func prettyPrintSparseMatrix(matrix [][]sparseEntry) {
	for i, r := range matrix {
		fmt.Print("[")
		for _, v := range r {
			fmt.Printf("(%d, %d = %d)", i, v.otherDim, v.val)
		}
		fmt.Println("]")
	}
}

// returns the result as a list of cells - to calculate the complete matrix
// these would have to be summed up for each row/col set
func sparseMultiply(a [][]sparseEntry, b [][]sparseEntry, n int) []sparseOutputCell {
	result := make([]sparseOutputCell, 0)

	for k := 0; k < n; k++ {
		for i := 0; i < len(a[k]); i++ {
			for j := 0; j < len(b[k]); j++ {
				fst := a[k][i]
				snd := b[k][j]
				val := fst.val * snd.val
				result = append(result, sparseOutputCell{fst.otherDim, snd.otherDim, val})
			}
		}
	}
	return result
}

// assumes square matrix of size*size
func readMatrix(filename string, n int, rowFirst bool) [][]sparseEntry {
	matrix := make([][]sparseEntry, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]sparseEntry, 0)
	}
	if file, err := os.Open(filename); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			if len(words) != 3 { // test output contains empty lines
				continue
			}
			row, _ := strconv.Atoi(words[0])
			col, _ := strconv.Atoi(words[1])
			val, _ := strconv.Atoi(words[2])
			if rowFirst {
				matrix[row] = append(matrix[row], sparseEntry{col, val})
			} else {
				matrix[col] = append(matrix[col], sparseEntry{row, val})
			}
		}
		if scanErr := scanner.Err(); err != nil {
			log.Fatal(scanErr)
		}
	} else {
		log.Fatal(err)
	}
	return matrix
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	a := readMatrix(os.Args[3], n, false)
	b := readMatrix(os.Args[5], n, true)
	c := sparseMultiply(a, b, n)
	for i := 0; i < len(c); i++ {
		entry := c[i]
		fmt.Printf("%d %d %d\n", entry.row, entry.col, entry.val)
	}
}
