package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	input := ReadInput(fname)
	solution1 := Part1(input)
	solution2 := Part2(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", solution1, solution2)
}

// Part1 solves: What is the solution to your captcha?
func Part1(numstr string) int {
	// Convert the string into an array of runes
	numArr := []rune(numstr)

	// Calclulate the length of the array once, for reference
	numLen := len(numArr)

	// Initialize an int variable to store results
	var sum int

	// Loop over the items in numArr
	for i, n := range numArr {
		// Assign an index, j, to be one greater than i, mod numLength (to account for
		// the list being circular)
		j := (i + 1) % numLen

		if n == numArr[j] {
			// If the current item is equal to the item at index j, convert the current
			// item to an int and add it to sum
			nInt, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatal(err)
			}
			sum += nInt
		}
	}
	return sum
}

// Part2 solves: What is the solution to your new captcha?
func Part2(numstr string) int {
	// As above, convert numstr into a rune slice; save the slice length; and initialize
	// a variable for results. Additionally, save half of the slice length to reference,
	// because Part 2 requires checking the number at halfway around the list.
	numArr := []rune(numstr)
	numLen := len(numArr)
	halfLen := numLen / 2
	var sum int

	for i, n := range numArr {
		// As in Part1, get an index j that is the current index plus halfLen (instead of
		// 1), take mod numLen to get the index of the item that is halfway around the
		// list. Again, if the two are equal, convert to int and add to sum
		j := (i + halfLen) % numLen
		if n == numArr[j] {
			nInt, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatal(err)
			}
			sum += nInt
		}
	}

	return sum
}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) string {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	strContents := strings.TrimRight(string(contents), "\r\n")
	return strContents
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
