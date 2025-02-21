package main

import (
	"fmt"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 0
	var expectedResult2 int64 = 0
	day := "09"

	data, err := aoc.ReadFileToInt64Array("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(data)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(data)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

var invalidNumber int64

func PartOne(data []int64) int64 {
	scope := 25

	for idx, n := range data[scope:] {
		if !isValidCode(n, data[idx:idx+scope]) {
			invalidNumber = n
			break
		}
	}

	return invalidNumber
}

func PartTwo(data []int64) int64 {
	contigious := findContigiousRange(invalidNumber, data)
	var min int64 = invalidNumber
	var max int64 = 0
	for _, val := range contigious {
		min = aoc.Min64(min, val)
		max = aoc.Max64(max, val)
	}
	return min + max
}

func findContigiousRange(n int64, data []int64) []int64 {
	var lower int = 0
	var upper int = 0
	sum := data[upper]

	for true {
		if sum == n {
			break
		}
		if sum > n {
			sum -= data[lower]
			lower++
		}
		if sum < n {
			upper++
			sum += data[upper]
		}
	}

	return data[lower : upper+1]
}

func isValidCode(n int64, prev []int64) bool {
	for i := 0; i < len(prev)-1; i++ {
		for j := i + 1; j < len(prev); j++ {
			if prev[i]+prev[j] == n {
				return true
			}
		}
	}
	return false
}
