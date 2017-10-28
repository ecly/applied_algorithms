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

// A BitVector represented as 4 uint64 with an additional
// index indicating it's index in the original slice of BitVectors.
type BitVector256 struct {
	a     uint64
	b     uint64
	c     uint64
	d     uint64
	index int
}

// A Triple used as key in a map
type Triple struct {
	X int
	Y int
	Z int
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
				return bv.index, vectors[j].index
			}
		}
	}
	return -1, -1
}

func compareInBuckets(buckets map[Triple][]BitVector256) (int, int) {
	for _, bucket := range buckets {
		i1, i2 := correlatedPair(bucket)
		if i1 != -1 {
			return i1, i2
		}
	}
	return -1, -1
}

func findSetBit(permutation [AMOUNT_OF_BITS]uint, vector BitVector256) int {
	for i, v := range permutation {
		switch v / 4 {
		case 0:
			if vector.a&1<<(v%64) == 1 {
				return i
			}
		case 1:
			if vector.b&1<<(v%64) == 1 {
				return i
			}
		case 2:
			if vector.c&1<<(v%64) == 1 {
				return i
			}
		case 3:
			if vector.d&1<<(v%64) == 1 {
				return i
			}
		}
	}
	// this should never be the case with the input sake
	// will however happen for BitVector256 with all 0's
	return -1
}

func generatePermutation() [AMOUNT_OF_BITS]uint {
	signed_perm := rand.Perm(AMOUNT_OF_BITS)
	var permutation [AMOUNT_OF_BITS]uint
	for i, v := range signed_perm {
		permutation[i] = uint(v)
	}
	return permutation
}

func groupInBuckets(vectors []BitVector256) map[Triple][]BitVector256 {
	rand.Seed(time.Now().UTC().UnixNano())
	buckets := make(map[Triple][]BitVector256, AMOUNT_OF_BITS)
	b1 := generatePermutation()
	b2 := generatePermutation()
	b3 := generatePermutation()
	for _, v := range vectors {
		key := Triple{
			findSetBit(b1, v),
			findSetBit(b2, v),
			findSetBit(b3, v)}
		buckets[key] = append(buckets[key], v)
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
		index := 0
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			a, _ := strconv.ParseInt(words[0], 10, 64)
			b, _ := strconv.ParseInt(words[1], 10, 64)
			c, _ := strconv.ParseInt(words[2], 10, 64)
			d, _ := strconv.ParseInt(words[3], 10, 64)
			vector := BitVector256{uint64(a), uint64(b), uint64(c), uint64(d), index}
			index += 1
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

// Utility function
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	debug.SetGCPercent(-1)
	filename := os.Args[1]
	//longAmount, _ := strconv.Atoi(os.Args[2])
	vectorAmount, _ := strconv.Atoi(os.Args[3])

	before := makeTimestamp()
	vectors := readVectors(filename, vectorAmount)
	after := makeTimestamp()
	fmt.Printf("%d\n", after-before)

	// returns the indices of the correlated pair within vectors
	low, high := minHash(vectors)
	fmt.Printf("%d %d\n", low, high)
}
