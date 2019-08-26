package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	points := ReadInput(fname)
	canvas := DefineCanvas(points)
	solution1 := Part1(points, canvas)
	solution2 := Part2(points, canvas, 10000)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", solution1, solution2)
}

// Part1 solves:
func Part1(points, canvas []Point) int {
	// Make a map[Point]int to keep track of the areas of the points
	areas := make(map[Point]int)

	for _, pt := range canvas {
		nearest, err := pt.FindNearestPoint(points)
		if err == nil {
			areas[nearest]++
		}
	}

	areaVals := make([]int, len(areas))
	i := 0
	for _, val := range areas {
		areaVals[i] = val
		i++
	}
	sort.Ints(areaVals)
	return areaVals[len(areaVals)-1]
}

// Part2 solves:
func Part2(points, canvas []Point, limit int) int {
	// Initialize an accumulator variable
	area := 0

	// For each point in the canvas acting as origin, if the sum of the distances to all other
	// points is less than `limit`, increment `area`
	for _, origin := range canvas {
		sum := 0
		for _, pt := range points {
			sum += origin.GetDistance(pt)
		}
		if sum < limit {
			area++
		}
	}

	return area
}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) []Point {
	results := make([]Point, 0)
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		coords := strings.Split(scan.Text(), ", ")
		if len(coords) != 2 {
			log.Fatal("Failed to read input.txt")
		}
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		pt := Point{x, y}
		results = append(results, pt)
	}

	return results
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
