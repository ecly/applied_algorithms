package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"math/rand"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// the min similarity we're looking for
const THRESHOLD = 70

// the cutoff for a single uint64's similarity to be considered
const CUTOFF = 15

// the number of bits each vector holds
const AMOUNT_OF_BITS = 256

// A BitVector represented as 4 uint64
type BitVector256 struct {
	a uint64
	b uint64
	c uint64
	d uint64
}

// Returns the similarity of two BitVectors
// meaning the number of bits set in both Vectors
func (b BitVector256) Compare(b1 BitVector256) int {
	similarity := 0
	similarity += bits.OnesCount64(b.a & b1.a)
	if similarity < CUTOFF {
		return 0
	}
	similarity += bits.OnesCount64(b.b & b1.b)
	similarity += bits.OnesCount64(b.c & b1.c)
	similarity += bits.OnesCount64(b.d & b1.d)
	//fmt.Printf("Similarity: %d\n", similarity)
	return similarity
}

func correlatedPair(vectors []BitVector256) (int, int) {
	for i, bv := range vectors {
		for j := i + 1; j < len(vectors); j++ {
			if bv.Compare(vectors[j]) > THRESHOLD {
				return i, j
			}
		}
	}
	return -1, -1
}

func compareInBuckets(buckets map[string][]BitVector256) (int, int) {
	for _, bucket := range buckets {
		i1, i2 := correlatedPair(bucket)
		if i1 != -1 {
			return i1, i2
		}
	}
	return -1, -1
}

/*
func findSetBit(permutation []int, vector BitVector256) int {
	for i, v := range permutation {

		if vecto
	}
}
*/

func groupInBuckets(vectors []BitVector256) map[string][]BitVector256 {
	rand.Seed(time.Now().UTC().UnixNano())
	buckets := make(map[string][]BitVector256)
	//b1 := rand.Perm(AMOUNT_OF_BITS)
	//b2 := rand.Perm(AMOUNT_OF_BITS)
	b3 := rand.Perm(AMOUNT_OF_BITS)
	for _, v := range b3 {
		fmt.Println(v)
	}

	return buckets
}

func minHash(vectors []BitVector256) (int, int) {
	buckets := groupInBuckets(vectors)
	i1, i2 := compareInBuckets(buckets)
	if i1 != -1 {
		return i1, i2
	}
	return minHash(vectors)
}

func readVectors(filename string, vectorAmount int) []BitVector256 {
	vectors := make([]BitVector256, 0, vectorAmount)
	if file, err := os.Open(filename); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			a, _ := strconv.ParseInt(words[0], 10, 64)
			b, _ := strconv.ParseInt(words[1], 10, 64)
			c, _ := strconv.ParseInt(words[2], 10, 64)
			d, _ := strconv.ParseInt(words[3], 10, 64)
			vector := BitVector256{uint64(a), uint64(b), uint64(c), uint64(d)}
			vectors = append(vectors, vector)
		}
		if scanErr := scanner.Err(); err != nil {
			log.Fatal(scanErr)
		}
	} else {
		log.Fatal(err)
	}
	return vectors
}

func main() {
	debug.SetGCPercent(-1)
	filename := os.Args[1]
	// longAmount, _ := strconv.Atoi(os.Args[2])
	vectorAmount, _ := strconv.Atoi(os.Args[3])
	vectors := readVectors(filename, vectorAmount)

	// returns the indices of the correlated pair within vectors
	b3 := rand.Perm(AMOUNT_OF_BITS)
	for _, v := range b3 {
		fmt.Println(v)
	}

	low, high := correlatedPair(vectors)
	fmt.Printf("%d %d\n", low, high)
}
