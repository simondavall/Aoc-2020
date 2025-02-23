package main

import (
	"fmt"
	"strconv"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 759
	var expectedResult2 int64 = 45763
	day := "12"

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

var directions []byte = []byte{'N', 'E', 'S', 'W'}

func PartOne(lines []string) int64 {
	hor := 0  // east/west distance
	vert := 0 // north/south distance
	dir := 1

	for _, line := range lines {
		ins := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch ins {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			dh, dv := move(ins, value)
			hor += dh
			vert += dv
			break
		case 'F':
			dh, dv := move(directions[dir], value)
			hor += dh
			vert += dv
			break
		case 'L':
			fallthrough
		case 'R':
			dir = turn(ins, value, dir)
			break
		}
	}

	return int64(aoc.Abs(hor) + aoc.Abs(vert))
}

type waypoint struct {
	hor  int
	vert int
}

func PartTwo(lines []string) int64 {
	hor := 0  // east/west distance
	vert := 0 // north/south distance
	var w waypoint = waypoint{10, 1}

	for _, line := range lines {
		ins := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch ins {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			dh, dv := move(ins, value)
			w.hor += dh
			w.vert += dv
			break
		case 'F':
			hor += w.hor * value
			vert += w.vert * value
			break
		case 'L':
			fallthrough
		case 'R':
			rotate(ins, value, &w)
			break
		}
	}

	return int64(aoc.Abs(hor) + aoc.Abs(vert))
}

func move(ins byte, value int) (int, int) {
	hor, vert := 0, 0
	switch ins {
	case 'N':
		vert += value
		break
	case 'S':
		vert -= value
		break
	case 'E':
		hor += value
		break
	case 'W':
		hor -= value
		break
	default:
		panic("Invalid instruction passed to move()")
	}
	return hor, vert
}

func turn(ins byte, value int, dir int) int {
	dirLen := len(directions)
	var t int = value / 90
	if ins == 'R' {
		dir = (dir + t) % dirLen
	} else {
		dir = (dirLen + dir - t) % dirLen
	}
	return dir
}

func rotate(ins byte, value int, w *waypoint) {
	var t int = value / 90
	for range t {
		if ins == 'R' {
			w.hor, w.vert = w.vert, w.hor*-1
		} else {
			w.hor, w.vert = w.vert*-1, w.hor
		}
	}
}
