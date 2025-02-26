package main

import (
	"fmt"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 240
	var expectedResult2 int64 = 1180
	day := "17"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var plainGrid [][]byte
	for _, line := range lines {
		plainGrid = append(plainGrid, []byte(line))
	}
	var grid [][][]byte
	grid = append(grid, plainGrid)

	startPart1 := time.Now()
	resultPartOne := PartOne(grid)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(grid)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(grid [][][]byte) int64 {
	var tally int64 = 0

	cycles := 0
	for cycles < 6 {
		cycles++
		grid = expandGrid(grid)

		var new_grid [][][]byte

		for z := 0; z < len(grid); z++ {
			var new_plain [][]byte
			for y := 0; y < len(grid[0]); y++ {
				var new_row []byte
				for x := 0; x < len(grid[0][0]); x++ {
					ch := grid[z][y][x]
					c := activeNeighbourCount(x, y, z, grid)
					new_row = updateNewRow(ch, c, new_row)
				}

				new_plain = append(new_plain, new_row)
			}
			new_grid = append(new_grid, new_plain)
		}
		grid = new_grid
	}

	for _, plain := range grid {
		for _, row := range plain {
			for _, ch := range row {
				if ch == '#' {
					tally++
				}
			}
		}
	}
	return tally
}

func PartTwo(small_grid [][][]byte) int64 {
	var hyper_grid [][][][]byte
	hyper_grid = append(hyper_grid, small_grid)

	cycles := 0
	for cycles < 6 {
		cycles++
		hyper_grid = expandHyperGrid(hyper_grid)
		var new_hyper [][][][]byte

		for w := 0; w < len(hyper_grid); w++ {
			var new_grid [][][]byte
			for z := 0; z < len(hyper_grid[0]); z++ {
				var new_plain [][]byte
				for y := 0; y < len(hyper_grid[0][0]); y++ {
					var new_row []byte
					for x := 0; x < len(hyper_grid[0][0][0]); x++ {
						ch := hyper_grid[w][z][y][x]
						c := activeHyperNeighbourCount(x, y, z, w, hyper_grid)
						new_row = updateNewRow(ch, c, new_row)
					}

					new_plain = append(new_plain, new_row)
				}
				new_grid = append(new_grid, new_plain)
			}
			new_hyper = append(new_hyper, new_grid)
		}
		hyper_grid = new_hyper
	}

	var tally int64 = 0
	for _, grid := range hyper_grid {
		for _, plain := range grid {
			for _, row := range plain {
				for _, ch := range row {
					if ch == '#' {
						tally++
					}
				}
			}
		}
	}
	return tally
}

func updateNewRow(ch byte, neighbourCount int, new_row []byte) []byte {
	if ch == '#' {
		if neighbourCount == 2 || neighbourCount == 3 {
			new_row = append(new_row, '#')
		} else {
			new_row = append(new_row, '.')
		}
	}
	if ch == '.' {
		if neighbourCount == 3 {
			new_row = append(new_row, '#')
		} else {
			new_row = append(new_row, '.')
		}
	}
	return new_row
}

func activeNeighbourCount(x int, y int, z int, grid [][][]byte) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				nx := x + dx
				ny := y + dy
				nz := z + dz
				if !outOfBounds(nx, ny, nz, grid) && grid[nz][ny][nx] == byte('#') {
					count++
				}
			}
		}
	}
	return count
}

func activeHyperNeighbourCount(x int, y int, z int, w int, hyper [][][][]byte) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					nx := x + dx
					ny := y + dy
					nz := z + dz
					nw := w + dw
					if !outOfBoundsHyper(nx, ny, nz, nw, hyper) && hyper[nw][nz][ny][nx] == byte('#') {
						count++
					}
				}
			}
		}
	}
	return count
}

func outOfBounds(x int, y int, z int, grid [][][]byte) bool {
	max_z := len(grid)
	max_y := len(grid[0])
	max_x := len(grid[0][0])
	return x < 0 || x >= max_x || y < 0 || y >= max_y || z < 0 || z >= max_z
}

func outOfBoundsHyper(x int, y int, z int, w int, hyper [][][][]byte) bool {
	max_w := len(hyper)
	max_z := len(hyper[0])
	max_y := len(hyper[0][0])
	max_x := len(hyper[0][0][0])
	return x < 0 || x >= max_x || y < 0 || y >= max_y || z < 0 || z >= max_z || w < 0 || w >= max_w
}
