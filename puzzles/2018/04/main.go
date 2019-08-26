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
	dateEntries := ReadInput(fname)
	guards := TrackGuards(dateEntries)
	solution1 := Part1(guards)
	solution2 := Part2(guards)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", solution1, solution2)
}

// Part1 solves: Find the guard that has the most minutes asleep. What minute does that guard spend
// asleep the most? What is the ID of the guard you chose multiplied by the minute you chose?
func Part1(guards map[int]Guard) int {
	guardID := 0
	maxMinutesAsleep := 0
	mostAsleepMinute := 0

	for _, g := range guards {
		if g.sumAsleepMinutes() > maxMinutesAsleep {
			guardID = g.num
			maxMinutesAsleep = g.sumAsleepMinutes()
			mostAsleepMinute = g.mostAsleepMinute()
		}
	}

	return guardID * mostAsleepMinute
}

// Part2 solves: Of all guards, which guard is most frequently asleep on the same minute? What is
// the ID of the guard you chose multiplied by the minute you chose?
func Part2(guards map[int]Guard) int {
	guardID := 0
	mostAsleepMinute := 0
	minutesSleptThen := 0

	for _, g := range guards {
		if g.asleep[g.mostAsleepMinute()] > minutesSleptThen {
			guardID = g.num
			mostAsleepMinute = g.mostAsleepMinute()
			minutesSleptThen = g.asleep[g.mostAsleepMinute()]
		}
	}

	return guardID * mostAsleepMinute
}

// ReadInput parses the puzzle's input.txt file
// Returns a map with date strings as keys, and maps from minute to line text as its values
func ReadInput(fname string) map[string]map[int]string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	timestampPattern := regexp.MustCompile(`\[(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2})\]`)
	var results = make(map[string]map[int]string)

	for scanner.Scan() {
		var dateKey string
		var timeKey int

		line := scanner.Text()
		lineStr := timestampPattern.FindStringSubmatch(line)
		var lineNum [5]int
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

		// if Hour is not 0, it's because it's started before midnight. So add one day, and set the
		// timeKey to -1 to signify that it started before midnight
		if timestamp.Hour() != 0 {
			timestamp = timestamp.AddDate(0, 0, 1)
			dateKey = fmt.Sprintf("%d-%d", timestamp.Month(), timestamp.Day())
			timeKey = -1
		} else {
			dateKey = fmt.Sprintf("%d-%d", timestamp.Month(), timestamp.Day())
			timeKey = timestamp.Minute()
		}

		// if dateKey isn't already in the results map, add it as an empty map
		if _, ok := results[dateKey]; !ok {
			results[dateKey] = make(map[int]string)
		}
		results[dateKey][timeKey] = line
	}
	return results
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
