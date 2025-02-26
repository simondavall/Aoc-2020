package main

import "slices"

func expandHyperGrid(hyper [][][][]byte) [][][][]byte {
	min_w, max_w := 0, len(hyper)-1
	expand_w := false

	for x := 0; x < len(hyper[0][0][0]); x++ {
		for y := 0; y < len(hyper[0][0]); y++ {
			for z := 0; z < len(hyper[0]); z++ {
				if hyper[min_w][z][y][x] == '#' || hyper[max_w][z][y][x] == '#' {
					expand_w = true
					break
				}
			}
			if expand_w {
				break
			}
		}
		if expand_w {
			break
		}
	}

	if expand_w {
		hyper = expandHyperInW(hyper)
	}

	min_z, max_z := 0, len(hyper[0])-1
	expand_z := false
	for x := 0; x < len(hyper[0][0][0]); x++ {
		for y := 0; y < len(hyper[0][0]); y++ {
			for w := 0; w < len(hyper); w++ {
				if hyper[w][min_z][y][x] == '#' || hyper[w][max_z][y][x] == '#' {
					expand_z = true
					break
				}
			}
			if expand_z {
				break
			}
		}
		if expand_z {
			break
		}
	}

	if expand_z {
		hyper = expandHyperInZ(hyper)
	}

	min_y, max_y := 0, len(hyper[0][0])-1
	expand_y := false
	for x := 0; x < len(hyper[0][0][0]); x++ {
		for z := 0; z < len(hyper[0]); z++ {
			for w := 0; w < len(hyper); w++ {
				if hyper[w][z][min_y][x] == '#' || hyper[w][z][max_y][x] == '#' {
					expand_y = true
					break
				}
			}
			if expand_y {
				break
			}
		}
		if expand_y {
			break
		}
	}

	if expand_y {
		hyper = expandHyperInY(hyper)
	}

	min_x, max_x := 0, len(hyper[0][0][0])-1
	expand_x := false
	for y := 0; y < len(hyper[0][0]); y++ {
		for z := 0; z < len(hyper[0]); z++ {
			for w := 0; w < len(hyper); w++ {
				if hyper[w][z][y][min_x] == '#' || hyper[w][z][y][max_x] == '#' {
					expand_x = true
					break
				}
			}
			if expand_x {
				break
			}
		}
		if expand_x {
			break
		}
	}

	if expand_x {
		hyper = expandHyperInX(hyper)
	}
	return hyper
}

func expandHyperInW(hyper [][][][]byte) [][][][]byte {
	var new_hyper [][][][]byte
	var new_grid [][][]byte
	for z := 0; z < len(hyper[0]); z++ {
		var new_plain [][]byte
		for y := 0; y < len(hyper[0][0]); y++ {
			var new_row []byte
			for x := 0; x < len(hyper[0][0][0]); x++ {
				new_row = append(new_row, '.')
			}
			new_plain = append(new_plain, new_row)
		}
		new_grid = append(new_grid, new_plain)
	}
	top_grid := slices.Clone(new_grid)
	new_hyper = append(new_hyper, new_grid)
	for _, exisiting_grid := range hyper {
		new_hyper = append(new_hyper, exisiting_grid)
	}
	new_hyper = append(new_hyper, top_grid)
	return new_hyper
}

func expandHyperInZ(hyper [][][][]byte) [][][][]byte {
	var new_hyper [][][][]byte
	for w := 0; w < len(hyper); w++ {
		var new_grid [][][]byte
		var new_plain [][]byte
		for y := 0; y < len(hyper[0][0]); y++ {
			var new_row []byte
			for x := 0; x < len(hyper[0][0][0]); x++ {
				new_row = append(new_row, '.')
			}
			new_plain = append(new_plain, new_row)
		}
		top_plain := slices.Clone(new_plain)
		new_grid = append(new_grid, new_plain)
		for _, exisiting_plain := range hyper[w] {
			new_grid = append(new_grid, exisiting_plain)
		}
		new_grid = append(new_grid, top_plain)
		new_hyper = append(new_hyper, new_grid)
	}
	return new_hyper
}

func expandHyperInY(hyper [][][][]byte) [][][][]byte {
	var new_hyper [][][][]byte
	for w := 0; w < len(hyper); w++ {
		var new_grid [][][]byte
		for z := 0; z < len(hyper[0]); z++ {
			var new_plain [][]byte
			var new_row []byte
			for x := 0; x < len(hyper[0][0][0]); x++ {
				new_row = append(new_row, '.')
			}
			top_row := slices.Clone(new_row)
			new_plain = append(new_plain, new_row)
			for _, exisiting_row := range hyper[w][z] {
				new_plain = append(new_plain, exisiting_row)
			}
			new_plain = append(new_plain, top_row)
			new_grid = append(new_grid, new_plain)
		}
		new_hyper = append(new_hyper, new_grid)
	}
	return new_hyper
}

func expandHyperInX(hyper [][][][]byte) [][][][]byte {
	var new_hyper [][][][]byte
	for w := 0; w < len(hyper); w++ {
		var new_grid [][][]byte
		for z := 0; z < len(hyper[0]); z++ {
			var new_plain [][]byte
			for y := 0; y < len(hyper[0][0]); y++ {
				var new_row []byte
				new_row = append(new_row, '.')
				for _, ch := range hyper[w][z][y] {
					new_row = append(new_row, ch)
				}
				new_row = append(new_row, '.')
				new_plain = append(new_plain, new_row)
			}
			new_grid = append(new_grid, new_plain)
		}
		new_hyper = append(new_hyper, new_grid)
	}
	return new_hyper
}
