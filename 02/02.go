package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var expectedResult1 int64 = 620
	var expectedResult2 int64 = 727
	day := "02"

	lines, err := readLines("input.txt")
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
	var tally int64 = 0

	for _, line := range lines {
		min, max, ch, pwd := processLine(line)

		charCount := strings.Count(pwd, ch)
		if charCount >= min && charCount <= max {
			tally++
		}
	}

	return tally
}

func PartTwo(lines []string) int64 {
	var tally int64 = 0

	for _, line := range lines {
		min, max, ch, pwd := processLine(line)

		if (pwd[min-1] == ch[0] && pwd[max-1] == ch[0]) || (pwd[min-1] != ch[0] && pwd[max-1] != ch[0]) {
			continue
		}

		tally++
	}

	return tally
}

func processLine(line string) (int, int, string, string) {
	splits := strings.Split(line, " ")
	bounds := strings.Split(splits[0], "-")

	min, err1 := strconv.Atoi(bounds[0])
	if err1 != nil {
		println(err1)
		return 0, 0, "", ""
	}

	max, err2 := strconv.Atoi(bounds[1])
	if err2 != nil {
		println(err2)
		return 0, 0, "", ""
	}

	ch := splits[1][:len(splits[1])-1]
	pwd := splits[2]

	return min, max, ch, pwd
}
