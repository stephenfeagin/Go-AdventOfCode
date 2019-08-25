package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	boxes := ReadInput(fname)
	solution1 := Part1(boxes)
	solution2 := Part2(boxes)

	fmt.Printf("Part 1: %d\nPart 2: %s\n", solution1, solution2)
}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ids []string

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		ids = append(ids, scan.Text())
	}
	return ids
}

// Part1 solves: What is the checksum for your list of box IDs?
func Part1(boxes []string) int {
	twos := 0
	threes := 0
	for _, box := range boxes {
		counts := make(map[rune]int, len(box))
		for _, r := range box {
			counts[r]++
		}
		for _, val := range counts {
			if val == 2 {
				twos++
				break
			}
		}
		for _, val := range counts {
			if val == 3 {
				threes++
				break
			}
		}
	}
	return twos * threes
}

// Part2 solves: What letters are common between the two correct box IDs?
func Part2(boxes []string) string {
	// Make an empty string slice for results
	results := make([]string, 0)

	// Initialize a string variable for the solution
	var solution string

	// Iterate over each box ID
	for _, box := range boxes {
		// Iterate over each character index
		for j := range box {

			// Make a new string by creating a string slice from the box ID up to the
			// given index, "*", then the remaining characters in the ID
			replaced := strings.Join([]string{box[:j], "*", box[j+1:]}, "")

			// Iterate over the results slice and check for a match
			for _, val := range results {
				if val == replaced {
					ind := strings.Index(val, "*")
					solution = strings.Join([]string{val[:ind], val[ind+1:]}, "")
				}
			}
			results = append(results, replaced)
		}
	}

	return solution
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
