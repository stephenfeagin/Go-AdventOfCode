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
	yearPattern = regexp.MustCompile(`^201[5-8]$`)
	dayPattern  = regexp.MustCompile(`^[012]?\d$|^3[01]$`)
)

func main() {
	if len(os.Args) != 3 || helpPattern.MatchString(os.Args[1]) {
		exitWithUsage()
	}

	year, day := os.Args[1], os.Args[2]

	if !yearPattern.MatchString(year) {
		fmt.Println("Invalid <year>. Must be in [2015..2018]")
		os.Exit(1)
	}

	if !dayPattern.MatchString(day) {
		fmt.Println("Invalid <day>. Must be in [1..31].")
		os.Exit(1)
	}

	if len(day) == 1 {
		day = "0" + day
	}

	dir := filepath.Join("puzzles", year, day)

	pluginPath := filepath.Join(dir, year+day+".so")
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Printf("No solution available for %s %s\n", year, day)
		os.Exit(1)
	}

	symbol, err := p.Lookup("Solve")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	solve, ok := symbol.(func(string))
	if !ok {
		fmt.Println("Plugin has no `Solve` function")
		os.Exit(1)
	}

	inputFile := filepath.Join(dir, "input.txt")
	solve(inputFile)
}

func exitWithUsage() {
	fmt.Println("Usage: ./AOC <year> <day>")
	os.Exit(1)
}
