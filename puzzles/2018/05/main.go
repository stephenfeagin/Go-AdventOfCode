package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	polymer := ReadInput(fname)
	fmt.Println(polymer)
}

// Part1 solves:
func Part1(polymer string) int {

}

// Part2 solves:
func Part2() {

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

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
