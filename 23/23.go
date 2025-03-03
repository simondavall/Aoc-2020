package main

import (
	"fmt"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 32897654
	var expectedResult2 int64 = 186715244496
	day := "23"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	first := Cup{int(lines[0][0] - '0'), nil, nil}
	prev := &first
	for _, b := range lines[0][1:] {
		current := Cup{int(b - '0'), prev, nil}
		prev.next = &current
		prev = &current
	}
	first.prev = prev
	prev.next = &first

	startPart1 := time.Now()
	resultPartOne := PartOne(&first, len(lines[0]))
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()

	first = Cup{int(lines[0][0] - '0'), nil, nil}
	prev = &first
	for _, b := range lines[0][1:] {
		current := Cup{int(b - '0'), prev, nil}
		prev.next = &current
		prev = &current
	}
	first.prev = prev
	prev.next = &first

	resultPartTwo := PartTwo(&first)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

type Cup struct {
	value int
	prev  *Cup
	next  *Cup
}

func PartOne(start *Cup, size int) int64 {
	curPtr := start

	ptr := start
	cups := make([]*Cup, size)
	for range size {
		cups[ptr.value-1] = ptr
		ptr = ptr.next
	}

	counter := 0
	for counter < 100 {
		counter++
		val := curPtr.value
		first := curPtr.next

		selectedValues, last := getSelectedValues(first)
		destVal := getDestinationValue(val, size, selectedValues)
		destination := cups[destVal-1]

		first.prev.next = last.next
		last.next.prev = first.prev
		destination.next.prev = last
		last.next = destination.next
		destination.next = first
		first.prev = destination

		curPtr = curPtr.next
	}

	oneCup := cups[0]
	curPtr = oneCup.next

	var tally int64 = 0
	for curPtr.value != 1 {
		tally *= 10
		tally += int64(curPtr.value)
		curPtr = curPtr.next
	}

	return tally
}

func PartTwo(start *Cup) int64 {
	curPtr := start
	last := curPtr.prev
	size := 1000000

	for i := range size - 9 {
		new_Cup := Cup{i + 10, last, nil}
		last.next = &new_Cup
		last = &new_Cup
	}
	curPtr.prev = last
	last.next = curPtr

	ptr := start
	cups := make([]*Cup, size)
	for range size {
		cups[ptr.value-1] = ptr
		ptr = ptr.next
	}

	counter := 0
	for counter < 10000000 {
		counter++
		val := curPtr.value
		first := curPtr.next

		selectedValues, last := getSelectedValues(first)
		destVal := getDestinationValue(val, size, selectedValues)
		destination := cups[destVal-1]

		first.prev.next = last.next
		last.next.prev = first.prev
		destination.next.prev = last
		last.next = destination.next
		destination.next = first
		first.prev = destination

		curPtr = curPtr.next
	}

	oneCup := cups[0]

	return int64(oneCup.next.value) * int64(oneCup.next.next.value)
}

func getDestinationValue(val int, size int, selectedValues map[int]bool) int {
	destVal := nextDestination(val, size)
	for selectedValues[destVal] {
		destVal = nextDestination(destVal, size)
	}
	return destVal
}

func nextDestination(cur int, size int) int {
	next := cur - 1
	if next < 1 {
		next += size
	}
	return next
}

func getSelectedValues(first *Cup) (map[int]bool, *Cup) {
	selectedValues := make(map[int]bool)
	selectedValues[first.value] = true
	last := first.next
	selectedValues[last.value] = true
	last = last.next
	selectedValues[last.value] = true
	return selectedValues, last
}
