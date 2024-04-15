package sudoku

import "math/rand"

var HintsLeft = 0

type pos struct {
	x int
	y int
}

func GenerateHints(origin, solved [][]rune) {
	// whipe prev state

	// // get solution by calling solving call

	// hints ready
}

func ShowHint(old_grid, solved_grid [][]rune) [][]rune { // prints a hint and returns a new grid with that hint filled in
	// Generate 1 hint
	// print it with special way that highlights the hint

	// here's a hint

	positions := []pos{}
	// // create a slice of step by step solutions until the final which is the full solution
	for y := 0; y < len(old_grid); y++ {
		for x := 0; x < len(old_grid[0]); x++ {
			if old_grid[y][x] == '.' {
				positions = append(positions, pos{x, y})
			}
		}
	}

	HintsLeft = len(positions)

	// display

	// og_copy := [][]rune{} //TODO THIS TYPE OF COPY DOESNT WORK<
	// copy(og_copy[:], origin[:])

	og_copy := make([][]rune, len(old_grid))

	for i := range old_grid {
		og_copy[i] = make([]rune, len(old_grid[i]))
		for j := range old_grid[i] {
			og_copy[i][j] = old_grid[i][j]
		}
	}

	ranpos := positions[rand.Int()%len(positions)]

	og_copy[ranpos.y][ranpos.x] = solved_grid[ranpos.y][ranpos.x]

	PrintFancySudoku(og_copy, ranpos.x, ranpos.y)
	return og_copy
}
