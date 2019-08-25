package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Claim represents a given claim about which area of fabric would be ideal
type Claim struct {
	id, left, top, width, height int
}

// Square defines the coordinates of a square inch of fabric
type Square struct {
	x, y int
}

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	claims := ReadInput(fname)
	squares := TallySquares(claims)
	fmt.Println("Part 1:", Part1(squares))
	fmt.Println("Part 2:", Part2(claims, squares))
}

// Part1 solves: How many square inches of fabric are within two or more claims?
// It takes the map provided by TallySquares as its input, and simply counts the number of times any
// square's count is greater than 1
func Part1(squares map[Square]int) int {
	result := 0

	for _, val := range squares {
		if val > 1 {
			result++
		}
	}
	return result
}

// Part2 solves: What is the ID of the only claim that doesn't overlap?
// It takes as inputs a slice of all the claims and the tally of squares of fabric
func Part2(claims []Claim, squares map[Square]int) int {
	for _, claim := range claims {
		claimSquares := TallySquares([]Claim{claim})
		tally := 0
		for sq := range claimSquares {
			tally += squares[sq]
		}
		if tally == len(claimSquares) {
			return claim.id
		}
	}
	return 0
}

// TallySquares calculates the number of times each square is represented in a claim. This function
// needs to be used for both Part1 and Part2, so it has been broken out into its own functoin
func TallySquares(claims []Claim) map[Square]int {
	results := make(map[Square]int)
	for _, claim := range claims {
		xMin := claim.left
		xMax := claim.left + claim.width
		yMin := claim.top
		yMax := claim.top + claim.height
		for x := xMin; x < xMax; x++ {
			for y := yMin; y < yMax; y++ {
				sq := Square{x, y}
				results[sq]++
			}
		}
	}

	return results
}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) []Claim {
	var results []Claim
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		claim := ParseClaim(scan.Text())
		results = append(results, claim)
	}
	return results
}

// ParseClaim creates a Claim object from the raw text of a claim from input.txt
func ParseClaim(str string) Claim {
	re := regexp.MustCompile("[0-9]+")
	numStr := re.FindAllString(str, -1)
	nums := make([]int, 5)
	for i, n := range numStr {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = num
	}
	return Claim{nums[0], nums[1], nums[2], nums[3], nums[4]}
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
