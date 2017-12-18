package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Interval struct representing the beginning and end of and interval.
type Interval struct {
	From int
	To   int
}

// Utiliy function to remove an Interval from a slice of Intervals
func remove(intervals *[]Interval, elem Interval) {
	for i := 0; i < len(*intervals); i++ {
		if (*intervals)[i] == elem {
			*intervals = append((*intervals)[:i], (*intervals)[i+1:]...)
		}
	}
}

// Finds the maximum independent set in a list of Intervals
func maxIndependentSet(intervals []Interval) []Interval {
	// Make a copy of input and sort in increasing order of Interval.To
	endsFirst := make([]Interval, len(intervals))
	copy(endsFirst, intervals)
	sort.SliceStable(endsFirst, func(i, j int) bool {
		return endsFirst[i].To < endsFirst[j].To
	})

	// Make a copy of input and sort in increasing order of Interval.From
	startsFirst := make([]Interval, len(intervals))
	copy(startsFirst, intervals)
	sort.SliceStable(startsFirst, func(i, j int) bool {
		return startsFirst[i].From < startsFirst[j].From
	})

	maxSet := make([]Interval, 0)
	for len(endsFirst) > 0 {
		fst := endsFirst[0]
		maxSet = append(maxSet, fst)
		endsFirst = endsFirst[1:]
		remove(&startsFirst, fst)

		snd := startsFirst[0]
		for snd.From <= fst.To {
			remove(&endsFirst, snd)
			startsFirst = startsFirst[1:]
			if len(startsFirst) == 0 { // avoid index out of range
				break
			}
			snd = startsFirst[0]
		}
	}

	return maxSet
}

// Reads input as seen on CodeJudge to a slice of Intervals
func readIntervals(filename string, amount int) []Interval {
	intervals := make([]Interval, 0, amount)
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			from, _ := strconv.Atoi(words[0])
			to, _ := strconv.Atoi(words[1])
			intervals = append(intervals, Interval{from, to})
		}
	}
	return intervals
}

func main() {
	filename := os.Args[1]
	amount, _ := strconv.Atoi(os.Args[2])
	intervals := readIntervals(filename, amount)
	independentSet := maxIndependentSet(intervals)
	for _, interval := range independentSet {
		fmt.Printf("%d %d\n", interval.From, interval.To)
	}
}
