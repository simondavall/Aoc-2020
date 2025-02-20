package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 6585
	var expectedResult2 int64 = 3276
	day := "06"

	blocks, err := aoc.ReadFileSplitBy("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(blocks)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(blocks)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(blocks []string) int64 {
	var tally int = 0

	for _, block := range blocks {
		var seen []rune
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			for _, ch := range line {
				if !slices.Contains(seen, ch) {
					seen = append(seen, ch)
				}
			}
		}
		tally += len(seen)
	}

	return int64(tally)
}

func PartTwo(blocks []string) int64 {
	var tally int = 0

	for _, block := range blocks {
		seen := make(map[rune]int)
		people := 0
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			people++
			for _, ch := range line {
				seen[ch]++
			}
		}
		for _, item := range seen {
			if item == people {
				tally++
			}
		}
	}

	return int64(tally)
}
