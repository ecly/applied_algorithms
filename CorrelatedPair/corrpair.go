package main

import (
    //"github.com/golang-collections/go-datastructures/bitarray"
    "github.com/willf/bitset"
	"bufio"
    "fmt"
    "strconv"
    "os"
    "log"
)

func read (filename string, longAmount int, vectorAmount int) []BitArray {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			x, _ := strconv.Atoi(words[0])
			y, _ := strconv.Atoi(words[1])
		}
		if scanErr = scanner.Err(); err != nil {
			log.Fatal(scanErr)
		}
	} else {
		log.Fatal(err)
	}
	return graph
}

func main(){
	debug.SetGCPercent(-1)
    filename := os.Args[1]
    longAmount, _ := strconv.Atoi(args[2])
    vectorAmount, _ := strconv.Atoi(args[3])
    vectors := readVectors(filename, longAmount, vectorAmount)
}
