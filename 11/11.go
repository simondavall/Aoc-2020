package main

import (
	"fmt"
	"strings"
	"time"

	aoc "aoc"
)

type direction struct {
	r int
	c int
}

var (
	_directions []direction = []direction{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	_height     int         = 0
	_width      int         = 0
)

func main() {
	var expectedResult1 int64 = 2483
	var expectedResult2 int64 = 2285
	day := "11"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_height = len(lines)
	_width = len(lines[0])

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

func PartOne(grid []string) int64 {
	return int64(getOccupiedSeats(grid, getNewSeat))
}

func PartTwo(grid []string) int64 {
	return int64(getOccupiedSeats(grid, getNewSeat2))
}

func getOccupiedSeats(grid []string, getNewSeat func(rune, int, int, []string) rune) int {
	hasChanged := true
	for hasChanged {
		hasChanged = false
		var newGrid []string

		for r, row := range grid {
			var sb strings.Builder
			for c, oldSeat := range row {
				newSeat := getNewSeat(oldSeat, r, c, grid)
				if newSeat != oldSeat {
					hasChanged = true
				}
				sb.WriteRune(newSeat)
			}
			newGrid = append(newGrid, sb.String())
		}
		grid = newGrid
	}

	occupiedSeats := 0
	for _, row := range grid {
		for _, ch := range row {
			if ch == '#' {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func getNewSeat(ch rune, r int, c int, grid []string) rune {
	switch ch {

	case '.':
		return '.'

	case 'L':
		for _, dir := range _directions {
			nr := r + dir.r
			nc := c + dir.c
			if !outOfBounds(nr, nc) && grid[nr][nc] == '#' {
				return 'L'
			}
		}
		return '#'

	case '#':
		occupied := 0
		for _, dir := range _directions {
			nr := r + dir.r
			nc := c + dir.c
			if !outOfBounds(nr, nc) && grid[nr][nc] == '#' {
				occupied++
			}
		}
		if occupied >= 4 {
			return 'L'
		}
		return '#'

	default:
		panic("Oops found unknown grid item")
	}
}

func getNewSeat2(ch rune, r int, c int, grid []string) rune {
	switch ch {

	case '.':
		return '.'

	case 'L':
		for _, dir := range _directions {
			nr, nc := r, c
			for {
				nr += dir.r
				nc += dir.c
				if outOfBounds(nr, nc) {
					break
				}
				if grid[nr][nc] == '.' {
					continue
				}
				if grid[nr][nc] == '#' {
					return 'L'
				}
				break
			}
		}
		return '#'

	case '#':
		occupied := 0
		for _, dir := range _directions {
			nr, nc := r, c
			for {
				nr += dir.r
				nc += dir.c
				if outOfBounds(nr, nc) {
					break
				}
				if grid[nr][nc] == '.' {
					continue
				}
				if grid[nr][nc] == '#' {
					occupied++
				}
				break
			}
		}
		if occupied >= 5 {
			return 'L'
		}
		return '#'

	default:
		panic("Oops found unknown grid item")
	}
}

func outOfBounds(r int, c int) bool {
	return r < 0 || r >= _height || c < 0 || c >= _width
}
