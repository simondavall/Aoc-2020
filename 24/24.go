package main

import (
	"fmt"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 330
	var expectedResult2 int64 = 3711
	day := "24"

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

type delta struct {
	r int
	c int
}

type tile struct {
	isBlack bool
	r       int
	c       int
}

var (
	directions = map[string]delta{"e": {0, 2}, "se": {1, 1}, "sw": {1, -1}, "w": {0, -2}, "nw": {-1, -1}, "ne": {-1, 1}}
	tiles      = make(map[string]*tile)
)

func PartOne(lines []string) int64 {
	for _, line := range lines {
		r, c := 0, 0
		len := 0
		for idx, b := range line {
			if b == 's' || b == 'n' {
				len = 1
				continue
			}
			d := directions[line[idx-len:idx+1]]
			r += d.r
			c += d.c
			len = 0
		}

		key := fmt.Sprintf("%d,%d", r, c)
		t, exists := tiles[key]
		if !exists {
			tiles[key] = &tile{true, r, c}
		} else {
			t.isBlack = !t.isBlack
		}
	}

	var tally int64 = 0
	for _, t := range tiles {
		if t.isBlack {
			tally++
		}
	}

	return tally
}

func PartTwo() int64 {
	days := 0
	for days < 100 {
		days++
		copy := make(map[string]*tile)
		for k, v := range tiles {
			copy[k] = &tile{v.isBlack, v.r, v.c}
		}

		// add missing white tiles
		for _, t := range copy {
			if !t.isBlack {
				continue
			}
			for _, d := range directions {
				nr := t.r + d.r
				nc := t.c + d.c
				neighbour := fmt.Sprintf("%d,%d", nr, nc)
				_, exists := copy[neighbour]
				if !exists {
					copy[neighbour] = &tile{false, nr, nc}
				}
			}
		}

		for _, tile := range copy {
			c := countBlackNeighbours(*tile)
			if tile.isBlack {
				if c == 0 || c > 2 {
					tile.isBlack = false
				}
			} else {
				if c == 2 {
					tile.isBlack = true
				}
			}
		}

		tiles = copy
	}

	printTiles(tiles)

	var tally int64 = 0
	for _, t := range tiles {
		if t.isBlack {
			tally++
		}
	}

	return tally
}

func countBlackNeighbours(t tile) int {
	count := 0
	for _, d := range directions {
		nr := t.r + d.r
		nc := t.c + d.c
		key := fmt.Sprintf("%d,%d", nr, nc)
		nt, exists := tiles[key]
		if exists && nt.isBlack {
			count++
		}
	}
	return count
}

func printTiles(tiles map[string]*tile) {
	min_r, max_r := 0, 0
	min_c, max_c := 0, 0
	for _, t := range tiles {
		if t.r > max_r {
			max_r = t.r
		}
		if t.r < min_r {
			min_r = t.r
		}
		if t.c > max_c {
			max_c = t.c
		}
		if t.c < min_c {
			min_c = t.c
		}
	}

	for r := min_r; r <= max_r; r++ {
		for c := min_c; c <= max_c; c++ {
			key := fmt.Sprintf("%d,%d", r, c)
			t, exists := tiles[key]
			if !exists {
				fmt.Print(".")
				continue
			}
			if t.isBlack {
				fmt.Print("#")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
