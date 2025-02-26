package main

import "slices"

func expandGrid(grid [][][]byte) [][][]byte {
	min_z, max_z := 0, len(grid)-1
	expand_z := false
	for x := 0; x < len(grid[0][0]); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[min_z][y][x] == '#' || grid[max_z][y][x] == '#' {
				expand_z = true
				break
			}
		}
		if expand_z {
			break
		}
	}

	if expand_z {
		grid = expandGridInZ(grid)
	}

	min_y, max_y := 0, len(grid[0])-1
	expand_y := false
	for x := 0; x < len(grid[0][0]); x++ {
		for z := 0; z < len(grid); z++ {
			if grid[z][min_y][x] == '#' || grid[z][max_y][x] == '#' {
				expand_y = true
				break
			}
		}
		if expand_y {
			break
		}
	}

	if expand_y {
		grid = expandGridInY(grid)
	}

	min_x, max_x := 0, len(grid[0][0])-1
	expand_x := false
	for y := 0; y < len(grid[0]); y++ {
		for z := 0; z < len(grid); z++ {
			if grid[z][y][min_x] == '#' || grid[z][y][max_x] == '#' {
				expand_x = true
				break
			}
		}
		if expand_x {
			break
		}
	}

	if expand_x {
		grid = expandGridInX(grid)
	}
	return grid
}

func expandGridInZ(grid [][][]byte) [][][]byte {
	var new_grid [][][]byte
	var new_plain [][]byte
	for y := 0; y < len(grid[0]); y++ {
		var new_row []byte
		for x := 0; x < len(grid[0][0]); x++ {
			new_row = append(new_row, '.')
		}
		new_plain = append(new_plain, new_row)
	}
	top_plain := slices.Clone(new_plain)
	new_grid = append(new_grid, new_plain)
	for _, exisiting_plain := range grid {
		new_grid = append(new_grid, exisiting_plain)
	}
	new_grid = append(new_grid, top_plain)
	return new_grid
}

func expandGridInY(grid [][][]byte) [][][]byte {
	var new_grid [][][]byte
	for z := 0; z < len(grid); z++ {
		var new_plain [][]byte
		var new_row []byte
		for x := 0; x < len(grid[0][0]); x++ {
			new_row = append(new_row, '.')
		}
		top_row := slices.Clone(new_row)
		new_plain = append(new_plain, new_row)
		for _, exisiting_row := range grid[z] {
			new_plain = append(new_plain, exisiting_row)
		}
		new_plain = append(new_plain, top_row)
		new_grid = append(new_grid, new_plain)
	}
	return new_grid
}

func expandGridInX(grid [][][]byte) [][][]byte {
	var new_grid [][][]byte
	for z := 0; z < len(grid); z++ {
		var new_plain [][]byte
		for y := 0; y < len(grid[0]); y++ {
			var new_row []byte
			new_row = append(new_row, '.')
			for _, ch := range grid[z][y] {
				new_row = append(new_row, ch)
			}
			new_row = append(new_row, '.')
			new_plain = append(new_plain, new_row)
		}
		new_grid = append(new_grid, new_plain)
	}
	return new_grid
}
