package main

import (
	"fmt"
	"strconv"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 545789
	var expectedResult2 int64 = 0
	day := "25"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	cardPublicKey, err := strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	doorPublicKey, err := strconv.ParseInt(lines[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(cardPublicKey, doorPublicKey)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo()
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(cardPublicKey int64, doorPublicKey int64) int64 {
	subjectNumber := int64(7)

	cardLoopSize := 0
	publicKey := int64(1)
	for publicKey != cardPublicKey {
		cardLoopSize++
		publicKey = transform(publicKey, subjectNumber)
	}

	doorLoopSize := 0
	publicKey = 1
	for publicKey != doorPublicKey {
		doorLoopSize++
		publicKey = transform(publicKey, subjectNumber)
	}

	subjectNumber = cardPublicKey
	encryptionKey := int64(1)
	for range doorLoopSize {
		encryptionKey = transform(encryptionKey, subjectNumber)
	}

	return encryptionKey
}

func transform(key int64, n int64) int64 {
	key *= n
	return key % 20201227
}

func PartTwo() int64 {
	return 0
}
