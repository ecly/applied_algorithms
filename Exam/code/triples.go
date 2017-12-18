package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// Bucket stores a subset of the input
// that hashed to this specific bucket in a list of buckets.
type Bucket []int

// Returns true if there exists a triple a+b==c in the input
// Otherwise returns false
func findTriples(input []int, amount int) bool {
	numBuckets := amount / 16
	shift := uint(math.Log2(float64(numBuckets)))
	seed := rand.Int()
	buckets := make([]Bucket, numBuckets)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = Bucket{}
	}
	for _, x := range input {
		h := hash(x, seed, shift)
		buckets[h] = append(buckets[h], x)
	}
	for i := 0; i < len(buckets); i++ {
		for j := i; j < len(buckets); j++ {
			k1 := (i + j) % numBuckets
			if naiveCompare(buckets[i], buckets[j], buckets[k1]) {
				return true
			}
			k2 := (i + j + 1) % numBuckets
			if naiveCompare(buckets[i], buckets[j], buckets[k2]) {
				return true
			}
		}
	}

	return false
}

// Given hash function in Golang
func hash(input int, seed int, shift uint) int {
	return ((input * seed) >> (64 - shift)) & ((2 << (shift - 1)) - 1)
}

// Naively check all combinations of a+b==c.
// Runs in O(n^3)
func naiveCompare(A Bucket, B Bucket, C Bucket) bool {
	for i := 0; i < len(A); i++ {
		a := A[i]
		for j := 0; j < len(B); j++ {
			b := B[j]
			for k := 0; k < len(C); k++ {
				if a+b == C[k] {
					return true
				}
			}
		}
	}
	return false
}

// Reads input as seen on CodeJudge to a slice of ints
func readInput(filename string, amount int) []int {
	ints := make([]int, 0, amount)
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i, _ := strconv.Atoi(scanner.Text())
			ints = append(ints, i)
		}
	}
	return ints
}

func main() {
	filename := os.Args[1]
	amount, _ := strconv.Atoi(os.Args[2])
	input := readInput(filename, amount)
	if findTriples(input, amount) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
