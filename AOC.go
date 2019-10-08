package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"regexp"
)

var (
	helpPattern = regexp.MustCompile(`-h`)
	yearPattern = regexp.MustCompile(`^201[5-8]$`)         // Puzzles only exist for 2015-2018
	dayPattern  = regexp.MustCompile(`^[012]?\d$|^3[01]$`) // Only 1-31 allowed
)

func main() {

	// Must be used with 2 arguments or with the first arg containing '-h'
	if len(os.Args) != 3 || helpPattern.MatchString(os.Args[1]) {
		exitWithUsage()
	}

	year, day := os.Args[1], os.Args[2]

	if !yearPattern.MatchString(year) {
		fmt.Println("Invalid <year>. Must be in [2015..2018]")
		os.Exit(1)
	}

	if !dayPattern.MatchString(day) {
		fmt.Println("Invalid <day>. Must be in [1..25].")
		os.Exit(1)
	}

	// Directories are zero-padded for days less than 10
	if len(day) == 1 {
		day = "0" + day
	}

	dir := filepath.Join("puzzles", year, day)

	// Try to open a plugin at puzzles/year/day
	pluginPath := filepath.Join(dir, year+day+".so")

	// Make sure that the file exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("No solution available for", year, day)
		os.Exit(1)
	}

	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Printf("No such file %s%s.so\n", year, day)
		fmt.Println("Try running ./build_plugin", year, day)
		os.Exit(1)
	}

	// My AOC plugins implement a Solve() function
	symbol, err := p.Lookup("Solve")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// We have to provide a type assertion to resolve the symbol. I use Solve(fname string) for the
	// puzzle input
	solve, ok := symbol.(func(string))
	if !ok {
		fmt.Println("Plugin has no `Solve` function")
		os.Exit(1)
	}

	// Get the puzzle input file
	inputFile := filepath.Join(dir, "input.txt")

	// Solve() prints the solutions and returns nothing
	solve(inputFile)
}

func exitWithUsage() {
	fmt.Println("Usage: ./AOC <year> <day>")
	os.Exit(1)
}
