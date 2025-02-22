package main

import (
	"fmt"
	"sort"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 2414
	var expectedResult2 int64 = 21156911906816
	day := "10"

	joltages, err := aoc.ReadFileToIntArray("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	joltages = append(joltages, 0)

	startPart1 := time.Now()
	resultPartOne := PartOne(joltages)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(joltages)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(joltages []int) int64 {
	sort.Ints(joltages[:])
	joltages = append(joltages, joltages[len(joltages)-1]+3)
	diff := make([]int, 4)
	prev := 0
	for _, n := range joltages {
		diff[n-prev]++
		prev = n
	}

	return int64(diff[1] * diff[3])
}

func PartTwo(joltages []int) int64 {
	dp := make([]int64, len(joltages))
	dp[0] = 1

	for i := 1; i < len(joltages); i++ {
		dp[i] = 0
		for j := i - 1; j >= 0 && joltages[i]-joltages[j] <= 3; j-- {
			dp[i] += dp[j]
		}
	}

	return dp[len(dp)-1]
}
