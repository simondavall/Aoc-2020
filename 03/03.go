package main

import (
	"fmt"
	"time"

	aoc "aoc"
)

var (
	_mapwidth  int
	_mapheight int
)

func main() {
	var expectedResult1 int64 = 200
	var expectedResult2 int64 = 3737923200
	day := "03"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_mapheight = len(lines)
	_mapwidth = len(lines[0])

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
	r, c := 0, 0
	dr, dc := 1, 3

	for r < _mapheight {
		if lines[r][c] == '#' {
			tally++
		}

		r = r + dr
		c = (c + dc) % _mapwidth
	}

	return tally
}

type Slope struct{ dr, dc int }

func PartTwo(lines []string) int64 {
	var tally int64 = 1

	slopes := []Slope{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	for _, slope := range slopes {
		r, c := 0, 0
		var trees int64 = 0
		for r < _mapheight {
			if lines[r][c] == '#' {
				trees++
			}

			r = r + slope.dr
			c = (c + slope.dc) % _mapwidth
		}
		tally *= trees
	}

	return tally
}
