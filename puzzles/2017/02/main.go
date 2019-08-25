package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	input := ReadInput(fname)
	solution1 := Part1(input)
	solution2 := Part2(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", solution1, solution2)
}

// Part1 solves: What is the checksum for the spreadsheet in your puzzle input?
func Part1(input [][]int) int {
	// Initialize results variable
	sum := 0

	// For each row, sort the row and then add to sum the difference between the last
	// item and the first item
	// (We'll see if this is viable for the actual puzzle input)
	for _, row := range input {
		sort.Ints(row)
		sum += row[len(row)-1] - row[0]
	}
	return sum
}

// Part2 solves: What is the sum of each row's result in your puzzle input?
func Part2(input [][]int) int {
	// Initialize results variable
	sum := 0

	// For each row, sort the row
	for _, row := range input {
		sort.Ints(row)

		// Working in descending order, search for the two values that divide evenly
		// i is the larger number
		for i := len(row) - 1; i > 0; i-- {
			// j is the smaller number
			for j := i - 1; j >= 0; j-- {
				// if row[j] divides evenly into row[i], add the quotient of the two to
				// the sum variable, and break out of the loop
				if row[i]%row[j] == 0 {
					sum += row[i] / row[j]
					break
				}
			}
		}
	}

	return sum
}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) [][]int {
	// Read the provided file
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader, changing the "Comma" field to \t
	r := csv.NewReader(file)
	r.Comma = '\t'

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new slice of slices to hold the output
	output := make([][]int, len(rows))

	// For each row, for each value in that row, convert it to int and append to the
	// corresponding row in output var
	for i, row := range rows {
		for _, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			output[i] = append(output[i], num)
		}
	}
	return output
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
