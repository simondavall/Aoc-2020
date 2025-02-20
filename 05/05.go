package main

import (
	"fmt"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 987
	var expectedResult2 int64 = 603
	day := "05"

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

var maxSeatId int

func PartOne(passes []string) int64 {
	for _, pass := range passes {
		if pass[0] != 'B' {
			continue
		}
		seatId := calcSeatId(pass)
		maxSeatId = aoc.Max(maxSeatId, seatId)
	}

	return int64(maxSeatId)
}

func PartTwo(passes []string) int64 {
	var mySeatId int = 0
	occupied := make([]bool, maxSeatId+1)

	for _, pass := range passes {
		current := calcSeatId(pass)
		occupied[current] = true
	}

	for i := maxSeatId; i >= 0; i-- {
		if !occupied[i] {
			mySeatId = i
			break
		}
	}

	return int64(mySeatId)
}

func calcSeatId(pass string) int {
	var rowCode string = pass[:len(pass)-3]
	var seatCode string = pass[len(pass)-3:]
	var row int
	for _, x := range rowCode {
		row <<= 1
		if x == 'B' {
			row++
		}
	}
	var seat int
	for _, y := range seatCode {
		seat <<= 1
		if y == 'R' {
			seat++
		}
	}
	return row*8 + seat
}
