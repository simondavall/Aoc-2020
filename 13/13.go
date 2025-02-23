package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aoc "aoc"
)

type Bus struct {
	id     int64
	offset int64
}

func main() {
	var expectedResult1 int64 = 174
	var expectedResult2 int64 = 780601154795940
	day := "13"

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
	time, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}

	buses := getBusData(lines)
	var timestamp int64 = int64(time)
	for true {
		for _, bus := range buses {
			if timestamp%bus.id == 0 {
				return int64((timestamp - int64(time)) * bus.id)
			}
		}
		timestamp++
	}

	return 0
}

func PartTwo(lines []string) int64 {
	buses := getBusData(lines)
	var timestamp int64 = 0
	found := false
	step := buses[0].id
	busCount := 2

	for !found {
		timestamp += step
		found = true
		for _, bus := range buses[:busCount] {
			if (timestamp+bus.offset)%bus.id != 0 {
				found = false
				break
			}
		}
		if found && busCount < len(buses) {
			found = false
			step *= buses[busCount-1].id
			busCount++
		}
	}

	return int64(timestamp)
}

func getBusData(lines []string) []Bus {
	var buses []Bus
	busData := strings.Split(lines[1], ",")
	for idx, bus := range busData {
		if bus != "x" {
			busId, err := strconv.Atoi(bus)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			buses = append(buses, Bus{int64(busId), int64(idx)})
		}
	}
	return buses
}
