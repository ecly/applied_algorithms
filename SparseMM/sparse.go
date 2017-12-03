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
	otherDimension int
	value          int
}

func prettyPrintSparseMatrix(matrix [][]sparseEntry) {
	for i, r := range matrix {
		fmt.Print("[")
		for _, v := range r {
			fmt.Printf("(%d, %d = %d)", i, v.otherDimension, v.value)
		}
		fmt.Println("]")
	}
}

// returns the result represented as [row][]sparseEntry.otherDimension=col
func sparseMultiply(a [][]sparseEntry, b [][]sparseEntry, n int) [][]sparseEntry {
	matrix := make([][]sparseEntry, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]sparseEntry, 0)
	}

	for k := 0; k < n; k++ {
		for i := 0; i < len(a[k]); i++ {
			for j := 0; j < len(b[k]); j++ {
				fst := a[k][i]
				snd := b[k][j]
				row := fst.otherDimension
				col := snd.otherDimension
				val := fst.value * snd.value
				matrix[row] = append(matrix[row], sparseEntry{col, val})
			}
		}
	}
	return matrix
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
	a := readMatrix(os.Args[3], n, true)
	b := readMatrix(os.Args[5], n, false)
	c := sparseMultiply(a, b, n)
	prettyPrintSparseMatrix(c)
}
