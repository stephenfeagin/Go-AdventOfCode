package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	freqs := readInput(fname)
	solution1 := part1(freqs)
	solution2 := part2(freqs)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", solution1, solution2)
}

// readInput parses the puzzle's input.txt file
func readInput(fname string) []int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var freqs []int

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		i, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		freqs = append(freqs, i)
	}

	return freqs
}

// Starting with a frequency of zero, what is the resulting frequency after all of the changes in
// frequency have been applied?
func part1(freqs []int) int {
	var currentFreq int
	for _, i := range freqs {
		currentFreq += i
	}
	return currentFreq
}

// What is the first frequency your device reaches twice?
func part2(freqs []int) int {
	results := make([]int, len(freqs))
	results[0] = 0
	nextFreq := 0
	for {
		for _, val := range freqs {
			nextFreq += val
			index := sort.SearchInts(results, nextFreq)
			if index < len(results) && results[index] == nextFreq {
				return nextFreq
			}

			results = append(results, 0)
			copy(results[index+1:], results[index:])
			results[index] = nextFreq

		}
	}
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
