package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type sparseEntry struct {
	otherDimension int
	value          int
}

func readMatrix(filename string, size int) [][]sparseEntry {
	matrix := make([][]sparseEntry, size)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]sparseEntry, 0)
	}
	if file, err := os.Open(filename); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			row, _ := strconv.Atoi(words[0])
			column, _ := strconv.Atoi(words[1])
			value, _ := strconv.Atoi(words[1])
			matrix[row] = append(matrix[row], sparseEntry{column, value})
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
	m, _ := strconv.Atoi(os.Args[2])
	a := readMatrix(os.Args[3], n)
	b := readMatrix(os.Args[4], m)
}
