package main

import (
	"runtime/debug"
	"bufio"
    "fmt"
    "strconv"
    "os"
    "log"
    "math/bits"
    "strings"
)

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
    similarity += bits.OnesCount64(b.b & b1.b)
    similarity += bits.OnesCount64(b.c & b1.c)
    similarity += bits.OnesCount64(b.d & b1.d)
    return similarity
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
            vector := BitVector256{uint64(a),uint64(b),uint64(c),uint64(d)}
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

func correlatedPair(vectors []BitVector256) (int, int) {
    return 1, 1
}

func main(){
	debug.SetGCPercent(-1)
    filename := os.Args[1]
    // longAmount, _ := strconv.Atoi(os.Args[2])
    vectorAmount, _ := strconv.Atoi(os.Args[3])
    vectors := readVectors(filename, vectorAmount)

    // returns the indices of the correlated pair within vectors
    low, high := correlatedPair(vectors)
    fmt.Printf("%d %d\n", low, high)
}
