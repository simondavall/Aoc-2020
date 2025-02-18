package main

import (
	"fmt"
	"time"
)

func main() {
	var expectedResult1 int64 = 605364
	var expectedResult2 int64 = 128397680
	day := "01"

	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := toInt32Array(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(expenses)
	fmt.Printf("Day_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(expenses)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(expenses []int32) int64 {
	for i := 0; i < len(expenses)-1; i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return int64(expenses[i] * expenses[j])
			}
		}
	}

	return 0
}

func PartTwo(expenses []int32) int64 {
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return int64(expenses[i] * expenses[j] * expenses[k])
				}
			}
		}
	}

	return 0
}
