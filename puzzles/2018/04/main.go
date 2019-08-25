package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

// Guard contains a guard's ID number and maps for asleep and awake minutes
type Guard struct {
	num           int
	asleep, awake map[int]int
}

// g.mostAsleepMinute() calculates which minute number guard g was asleep during the most
func (g Guard) mostAsleepMinute() int {
	minute := 0
	maxAsleep := 0
	for min, count := range g.asleep {
		if count > maxAsleep {
			minute = min
			maxAsleep = count
		}
	}
	return minute
}

// g.sumAsleepMinutes() calculates how much total time guard g was asleep
func (g Guard) sumAsleepMinutes() int {
	sum := 0
	for _, count := range g.asleep {
		sum += count
	}
	return sum
}

// TrackGuards creates a map from minute to guard on duty, and calculates whether that guard was
// asleep or awake
func TrackGuards(dateEntries map[string]map[int]string) map[int]Guard {
	guards := make(map[int]Guard)
	guardPattern := regexp.MustCompile(`#(\d+)`)

	// For each date in dateEntries, there will be a map of entries
	// For each date, sort the entries by their keys, which are the minute number that the block in
	// question starts. Determine the guard ID, and if that guard is not already in the guards
	// map, add it as a Guard object with their ID. Then assign the Guard instance at guards[id] to
	// a variable
	for date, entries := range dateEntries {
		sortedEntries := make([]int, len(entries)) // create a new int slice to hold the entry keys
		i := 0                                     // start a counter
		for en := range entries {                  // iterate over the keys of entries
			sortedEntries[i] = en // assign the key to sortedEntries at the counter
			i++                   // increment the counter
		}
		sort.Ints(sortedEntries)

		stringID := guardPattern.FindAllStringSubmatch(entries[sortedEntries[0]], -1)

		// We expect the results to be a [][]string, where the first entry is len 2 (with the
		// original string followed by the number extracted from it)
		// If this is not the case, exit
		if len(stringID[0]) != 2 {
			log.Fatalf("Failed to extract guard ID for date %s\n", date)
		}
		guardID, err := strconv.Atoi(stringID[0][1])
		if err != nil {
			log.Fatal(err)
		}

		// check if this guard ID is already in the guards map
		// if it isn't, add it and create empty maps for the asleep and awake fields
		guard, ok := guards[guardID]
		if !ok {
			guard = Guard{num: guardID}
			guard.asleep = make(map[int]int)
			guard.awake = make(map[int]int)
			guards[guardID] = guard
		}

		// For each entry (through len-1 to allow for using i+1), identify the current entry and the
		// next entry
		for i := 0; i < len(sortedEntries)-1; i++ {
			thisEntry := sortedEntries[i]
			nextEntry := sortedEntries[i+1]

			// We start with i=0, i+1=1
			// In the first block, the guard is always awake
			// So in the blocks where i+1 is even, the guard is asleep
			if (i+1)%2 == 0 {
				for j := thisEntry; j < nextEntry; j++ {
					if _, ok := guard.asleep[j]; !ok {
						guard.asleep[j] = 0
					}
					guard.asleep[j]++
				}
			} else {
				for j := thisEntry; j < nextEntry; j++ {
					if _, ok := guard.awake[j]; !ok {
						guard.awake[j] = 0
					}
					guard.awake[j]++
				}
			}
		}
	}

	return guards
}

// I include main() so that go doesn't yell about a main package having no main() function
func main() {
}
