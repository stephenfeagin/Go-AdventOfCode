package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {

	polymer := ReadInput(fname)
	solution1 := Part1(polymer)
	solution2 := Part2(polymer)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", solution1, solution2)
}

// Part1 solves: How many units remain after fully reacting the polymer you scanned?
func Part1(polymer string) int {

	return len(React(polymer))
}

// Part2 solves: What is the length of the shortest polymer you can produce by removing all units of
// exactly one type and fully reacting the result?
func Part2(polymer string) int {

	// get a string with all of the unique letters in `polymer` (lower-cased)
	letters := ""
	for _, char := range polymer {
		if c := unicode.ToLower(char); !strings.ContainsRune(letters, c) {
			letters += string(c)
		}
	}

	// make an int slice to hold the lengths of the reacted sub-polymers
	lengths := make([]int, len(letters))

	// For each letter, do a case-insensitive replace for that letter and save the length of the
	// reacted new sub-polymer in lengths
	for i, letter := range letters {
		re := regexp.MustCompile("(?i:" + string(letter) + ")")
		subPolymer := re.ReplaceAllLiteralString(polymer, "")
		lengths[i] = len(React(subPolymer))
	}

	// Sort the slice and return the first value (the min)
	sort.Ints(lengths)
	return lengths[0]
}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	polymer := strings.TrimSpace(string(content))
	return polymer
}

// React reacts
func React(polymer string) string {
	rd := strings.NewReader(polymer) // Create a new string reader from `polymer` (for efficiency)
	var result []rune                // Create a rune slice to hold the resulting characters

	for {
		char, _, err := rd.ReadRune() // Read a char
		if err != nil {
			if err == io.EOF {
				break // If err is io.EOF, it means we've reached the end of the buffer, so stop
			} else {
				log.Fatal(err) // If there's an error that isn't io.EOF, it's an actual error
			}
		}

		// If result is empty, append char to it and move on to the next character in rd
		if len(result) == 0 {
			result = append(result, char)
			continue
		}

		// Get the last char in result
		last := result[len(result)-1]

		// If last != char but they are equal with case changed, they react
		if last != char && (unicode.ToLower(last) == char || unicode.ToUpper(last) == char) {
			result = result[:len(result)-1] // So we pop the last item off of result
		} else { // If they don't react, we just add char to the end of result and continue on
			result = append(result, char)
		}
	}

	return string(result)
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
