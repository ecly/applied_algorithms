package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"math/rand"
	"os"
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

// the depth of our hashing
const AMOUNT_OF_SUBBUCKETS = 5

// the size of a single permutation
// 64 instead of 256 as we only check first entry
// of the BitVector 256
const PERM_SIZE = 64

// A BitVector represented as 4 uint64 with an additional
// index indicating it's index in the original slice of BitVectors.
type BitVector256 struct {
	a     uint64
	b     uint64
	c     uint64
	d     uint64
	index int
}

// A Key used as key in a map
type Key [AMOUNT_OF_SUBBUCKETS]uint8
type BucketMap map[Key][]BitVector256

// Returns the similarity of two BitVectors
// meaning the number of bits set in both Vectors
func (b BitVector256) Compare(b1 BitVector256) int {
	similarity := 0
	similarity += bits.OnesCount64(b.a & b1.a)
	if similarity < CUTOFF {
		return 0
	}
	similarity += bits.OnesCount64(b.b & b1.b)
	if similarity < CUTOFF*2 {
		return 0
	}
	similarity += bits.OnesCount64(b.c & b1.c)
	similarity += bits.OnesCount64(b.d & b1.d)
	return similarity
}

// brute force in a cache efficient way
func correlatedPair(vectors []BitVector256) (int, int) {
	for i := 0; i < len(vectors); i++ {
		vec := vectors[i]
		for j := i + 1; j < len(vectors); j++ {
			if vec.Compare(vectors[j]) >= THRESHOLD {
				return vec.index, vectors[j].index
			}
		}
	}
	return -1, -1
}

func compareInBuckets(buckets BucketMap) (int, int) {
	//fmt.Printf("Amount of buckets %d\n", len(buckets))
	for _, bucket := range buckets {
		i1, i2 := correlatedPair(bucket)
		if i1 != -1 {
			return i1, i2
		}
	}
	return -1, -1
}

func findSetBit(permutation [PERM_SIZE]uint8, vector BitVector256) uint8 {
	for i := 0; i < AMOUNT_OF_BITS; i++ {
		v := permutation[i]
		if vector.a&(1<<v) != 0 {
			return v
		}
	}
	// this should never be the case with the input sake
	// will however happen for BitVector256 with all 0's
	return 0
}

// generate the numbers 0..PERM_SIZE as uint8
func defaultPermutation() [PERM_SIZE]uint8 {
	var arr [PERM_SIZE]uint8
	for i := 0; i < PERM_SIZE; i++ {
		arr[i] = uint8(i)
	}
	return arr
}

// permutates an array similarly to how golang's libraries do it
func generatePermutation(slice [PERM_SIZE]uint8) [PERM_SIZE]uint8 {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func groupInBuckets(vectors []BitVector256) BucketMap {
	rand.Seed(time.Now().UTC().UnixNano())
	buckets := make(BucketMap)
	var permutations [AMOUNT_OF_SUBBUCKETS][PERM_SIZE]uint8

	// generate all the permutations once
	defaultPermutation := defaultPermutation()
	for i := 0; i < AMOUNT_OF_SUBBUCKETS; i++ {
		permutations[i] = generatePermutation(defaultPermutation)
	}
	for _, v := range vectors {
		var key [AMOUNT_OF_SUBBUCKETS]uint8
		// create the key as an array of size AMOUNT_OF_SUBBUCKETS
		for j := 0; j < AMOUNT_OF_SUBBUCKETS; j++ {
			key[j] = findSetBit(permutations[j], v)
		}
		buckets[key] = append(buckets[key], v)
	}

	return buckets
}

// a single run of minhash, not guaranteed to find the pair
func minHash(vectors []BitVector256) (int, int) {
	buckets := groupInBuckets(vectors)
	return compareInBuckets(buckets)
}

// recursively try until result is found
func MinHash(vectors []BitVector256) (int, int) {
	i1, i2 := minHash(vectors)
	if i1 != -1 {
		return i1, i2
	}
	return MinHash(vectors)
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

func main() {
	filename := os.Args[1]
	//longAmount, _ := strconv.Atoi(os.Args[2])
	vectorAmount, _ := strconv.Atoi(os.Args[3])
	vectors := readVectors(filename, vectorAmount)

	// returns the indices of the correlated pair within vectors
	low, high := MinHash(vectors)
	fmt.Printf("%d %d\n", low, high)
}
