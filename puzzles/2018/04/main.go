package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Solve prints out solutions for the puzzle
func Solve(fname string) {
	input := ReadInput(fname)
	fmt.Println(input)
}

// Part1 solves:
func Part1() {

}

// Part2 solves:
func Part2() {

}

// ReadInput parses the puzzle's input.txt file
func ReadInput(fname string) map[string]map[int]string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	timestampPattern := regexp.MustCompile(`\[(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2})\]`)
	var results = map[string]map[int]string{}

	for scanner.Scan() {
		var dateKey string
		var timeKey int

		line := scanner.Text()
		lineStr := timestampPattern.FindStringSubmatch(line)
		lineNum := make([]int, 5)
		for i := 1; i < len(lineStr); i++ {
			num, err := strconv.Atoi(lineStr[i])
			if err != nil {
				log.Fatal(err)
			}
			lineNum[i-1] = num
		}
		timestamp := time.Date(lineNum[0],
			time.Month(lineNum[1]),
			lineNum[2],
			lineNum[3],
			lineNum[4],
			0, 0, time.UTC)
		if timestamp.Hour() != 0 {
			timestamp = timestamp.AddDate(0, 0, 1)
			dateKey = fmt.Sprintf("%d-%d", timestamp.Month(), timestamp.Day())
			timeKey = -1
		} else {
			dateKey = fmt.Sprintf("%d-%d", timestamp.Month(), timestamp.Day())
			timeKey = timestamp.Minute()
		}
		if _, ok := results[dateKey]; !ok {
			results[dateKey] = make(map[int]string)
		}
		results[dateKey][timeKey] = line
	}
	return results
}

// Guard contains a guard's ID number and maps for asleep and awake minutes
type Guard struct {
	num           int
	asleep, awake map[int]int
}

/* func (g Guard) mostAsleepMinute() int {
	minute := 0
	countAsleep := 0
	for min, count := range g.asleep {
		if count > countAsleep {
			minute = min
			countAsleep = countAsleep
		}
	}
	return minute
} */

func (g Guard) sumAsleepMinutes() int {
	sum := 0
	for _, count := range g.asleep {
		sum += count
	}
	return sum
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
