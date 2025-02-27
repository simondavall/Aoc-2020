package main

import (
	"fmt"
	"slices"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 25190263477788
	var expectedResult2 int64 = 297139939002972
	day := "18"

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

type stackItem struct {
	value int64
	op    rune
}

func PartOne(lines []string) int64 {
	var tally int64 = 0
	var cur int64
	for _, str := range lines {
		var stack aoc.Stack[stackItem]
		cur = 0
		var op rune
		ops := []rune{'+', '*'}
		for _, ch := range str {
			if '0' <= ch && ch <= '9' {
				if op == 0 {
					cur = int64(ch - '0')
					continue
				} else {
					cur = applyOp(cur, op, int64(ch-'0'))
					op = 0
					continue
				}
			}
			if slices.Contains(ops, ch) {
				op = ch
				continue
			}
			if ch == '(' {
				stack.Push(stackItem{cur, op})
				cur = 0
				op = 0
			}
			if ch == ')' {
				p, err := stack.Pop()
				if err != nil {
					fmt.Println(err)
					return 0
				}
				if p.value > 0 && p.op > 0 {
					cur = applyOp(p.value, p.op, cur)
				}
			}
		}
		tally += cur
	}

	return tally
}

func PartTwo(lines []string) int64 {
	var tally int64 = 0
	var cur int64
	for _, str := range lines {
		var stack []aoc.Stack[stackItem] = make([]aoc.Stack[stackItem], 1)
		stackIndex := 0
		cur = 0
		var op rune
		ops := []rune{'+', '*'}
		for _, ch := range str {
			if '0' <= ch && ch <= '9' {
				if op == 0 {
					cur = int64(ch - '0')
					continue
				} else {
					cur = applyOp(cur, op, int64(ch-'0'))
					op = 0
					continue
				}
			}
			if slices.Contains(ops, ch) {
				if ch == '*' {
					stack[stackIndex].Push(stackItem{cur, ch})
					cur = 0
					op = 0
					continue
				}
				op = ch
				continue
			}
			if ch == '(' {

				stackIndex++
				stack = append(stack, aoc.Stack[stackItem]{})
				stack[stackIndex].Push(stackItem{cur, op})
				cur = 0
				op = 0
			}
			if ch == ')' {
				for !stack[stackIndex].IsEmpty() {
					p, _ := stack[stackIndex].Pop()
					if p.value > 0 && p.op > 0 {
						cur = applyOp(p.value, p.op, cur)
					}
				}
				stack = stack[:stackIndex]
				stackIndex--
			}
		}
		// flush stack
		for !stack[stackIndex].IsEmpty() {
			p, _ := stack[stackIndex].Pop()
			if p.value > 0 && p.op > 0 {
				cur = applyOp(p.value, p.op, cur)
			}
		}
		stack = stack[:stackIndex]
		stackIndex--

		tally += cur
	}

	return tally
}

func applyOp(cur int64, op rune, value int64) int64 {
	switch op {
	case '+':
		return cur + value
	case '*':
		return cur * value
	default:
		panic(fmt.Sprintf("Unknown operator:%c", op))
	}
}
