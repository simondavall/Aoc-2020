package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 13105044880745
	var expectedResult2 int64 = 3505392154485
	day := "14"

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
	var tally int64 = 0
	mask := ""
	memory := make(map[int]int64)
	binaryNumber := make([]byte, 36)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			msk := strings.Split(line, " = ")
			mask = msk[1]
			continue
		}

		memAddr, value, err := getMemory(line)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		setBinaryNumber(value, binaryNumber)

		var memStr strings.Builder
		for i := 0; i < len(mask); i++ {
			if mask[i] == 'X' {
				memStr.WriteByte(binaryNumber[i])
			} else {
				memStr.WriteByte(mask[i])
			}
		}

		decimalNumber, err := strconv.ParseInt(memStr.String(), 2, 64)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		memory[memAddr] = decimalNumber
	}

	for _, memoryValue := range memory {
		tally += memoryValue
	}

	return tally
}

func PartTwo(lines []string) int64 {
	var tally int64 = 0
	mask := ""
	memory := make(map[int64]int)
	binaryNumber := make([]byte, 36)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			msk := strings.Split(line, " = ")
			mask = msk[1]
			continue
		}

		memAddr, value, err := getMemory(line)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		setBinaryNumber(memAddr, binaryNumber)

		first := make([]byte, 36)
		memAddresses := [][]byte{first}
		for i := 0; i < len(mask); i++ {
			if mask[i] == 'X' {
				var newEntries [][]byte
				for _, current := range memAddresses {
					newEntry := slices.Clone(current)
					newEntry[i] = '1'
					current[i] = '0'
					newEntries = append(newEntries, newEntry)
				}
				for _, newEntry := range newEntries {
					memAddresses = append(memAddresses, newEntry)
				}
			} else if mask[i] == '1' {
				for _, address := range memAddresses {
					address[i] = mask[i]
				}
			} else {
				for _, address := range memAddresses {
					address[i] = binaryNumber[i]
				}
			}
		}

		for _, address := range memAddresses {
			memAddr, err := strconv.ParseInt(string(address[:]), 2, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			memory[memAddr] = value
		}
	}

	for _, memValue := range memory {
		tally += int64(memValue)
	}

	return tally
}

func getMemory(line string) (int, int, error) {
	memory := strings.Split(line, "] = ")
	value, err := strconv.Atoi(memory[1])
	if err != nil {
		return 0, 0, err
	}
	maddr := strings.Split(memory[0], "[")
	memAddr, err := strconv.Atoi(maddr[1])
	if err != nil {
		return 0, 0, err
	}

	return memAddr, value, nil
}

func setBinaryNumber(value int, binaryNumber []byte) {
	for i := range binaryNumber {
		binaryNumber[i] = '0'
	}

	valueAsBinary := strconv.FormatInt(int64(value), 2)
	offset := 36 - len(valueAsBinary)
	for i := 0; i < len(valueAsBinary); i++ {
		binaryNumber[i+offset] = valueAsBinary[i]
	}
}
