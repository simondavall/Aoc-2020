package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 870
	var expectedResult2 int64 = 9136
	day := "15"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(lines)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(lines)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(lines []string) int64 {
	line := lines[0]
	numMap := make(map[int]int)
	var diff int = 0
	initial := strings.Split(line, ",")
	for idx, n := range initial {
		n, _ := strconv.Atoi(n)
		diff = addToNumMap(n, idx, numMap)
	}

	counter := len(initial)
	for counter < 2019 {
		diff = addToNumMap(diff, counter, numMap)
		counter++
	}

	return int64(diff)
}

func PartTwo(lines []string) int64 {
	line := lines[0]
	numMap := make(map[int]int)
	var diff int = 0
	initial := strings.Split(line, ",")
	for idx, n := range initial {
		n, _ := strconv.Atoi(n)
		diff = addToNumMap(n, idx, numMap)
	}

	counter := len(initial)
	for counter < 30000000-1 {
		diff = addToNumMap(diff, counter, numMap)
		counter++
	}

	return int64(diff)
}

func addToNumMap(n int, idx int, numMap map[int]int) int {
	diff := 0
	if cur, ok := numMap[n]; ok {
		diff = idx - cur
	}
	numMap[n] = idx
	return diff
}
