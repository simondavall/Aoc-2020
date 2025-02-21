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
	var expectedResult1 int64 = 1749
	var expectedResult2 int64 = 515
	day := "08"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(lines)
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

type Instruction struct {
	name  string
	value int
}

var instructions []Instruction

func PartOne(lines []string) int64 {
	for _, line := range lines {
		s := strings.Split(line, " ")
		value, _ := strconv.Atoi(s[1])
		instructions = append(instructions, Instruction{s[0], value})
	}

	ip := 0
	var acc int64 = 0
	var seen []int = []int{}

	for true {
		if ip < 0 || ip >= len(lines) || slices.Contains(seen, ip) {
			break
		}
		seen = append(seen, ip)
		ins := instructions[ip]
		ip, acc = ProcessInstruction(ins, ip, acc)
	}

	return acc
}

func PartTwo() int64 {
	var acc int64 = 0
	for idx, ins := range instructions {
		orig := ""
		if ins.name == "acc" {
			continue
		}
		if ins.name == "jmp" {
			orig = "jmp"
			instructions[idx].name = "nop"
		}
		if ins.name == "nop" {
			orig = "nop"
			instructions[idx].name = "jmp"
		}

		ip := 0
		acc = 0
		var seen []int = []int{}
		success := false

		for true {
			if ip < 0 {
				panic("Cannot have negative instruction pointer")
			}
			if ip >= len(instructions) {
				success = true
				break
			}
			if slices.Contains(seen, ip) {
				break
			}
			seen = append(seen, ip)
			ins := instructions[ip]
			ip, acc = ProcessInstruction(ins, ip, acc)
		}

		instructions[idx].name = orig

		if success {
			break
		}
	}

	return acc
}

func ProcessInstruction(ins Instruction, ip int, acc int64) (int, int64) {
	switch ins.name {
	case "acc":
		acc += int64(ins.value)
		ip++
	case "jmp":
		ip += ins.value
	case "nop":
		ip++
	default:
		panic("Unknown instruction")
	}
	return ip, acc
}
