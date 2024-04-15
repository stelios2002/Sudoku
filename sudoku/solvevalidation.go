package sudoku

import (
	"math"
)

var (
	base        string
	finalsolved [][]rune
	og_solution [][]rune
)

func SolveValidation(args []string) (int, [][]rune) { // count of solution, solution grid
	og_solution = MakingTheSudoku(args)
	// PrintDrySudoku(solution)
	solution_count := 0
	base = BaseCreation(args)
	// fmt.Println("the base is: ", base)
	SolveRecur(0, 0, &og_solution, &solution_count)

	return solution_count, finalsolved
}

func BaseCreation(args []string) string {
	if len(args) == 4 {
		return "1234"
	}
	if len(args) == 9 {
		return "123456789"
	}
	return "123456789ABCDEFG"
}

func MakingTheSudoku(args []string) [][]rune {
	solution := [][]rune{}
	for _, row := range args {
		// fmt.Println(row)
		itsCol := []rune{}
		for _, col := range row {
			itsCol = append(itsCol, rune(col))
		}
		solution = append(solution, itsCol)
	}
	return solution
}

func isSafe(grid [][]rune, row, col int, num rune) bool {
	// Check if the number is not already present in the row and column
	for i := 0; i < len(base); i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check if the number is not already present in the 2x2,3x3,4x4 subgrid
	square := int(math.Sqrt(float64(len(base))))
	startRow := row - row%square
	startCol := col - col%square
	for i := 0; i < square; i++ {
		for j := 0; j < square; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

var counter int

func SolveRecur(row, col int, solution *[][]rune, solsFound *int) {
	// fmt.Println("in recur")

	if row == len(base)-1 && col == len(base) {
		*solsFound++
		// fmt.Println("exiting cause solution found!")
		// PrintFancySudoku(*solution)

		finalsolved = make([][]rune, len(og_solution))

		for i := range og_solution {
			finalsolved[i] = make([]rune, len(og_solution[i]))
			for j := range og_solution[i] {
				finalsolved[i][j] = og_solution[i][j]
			}
		}
		return
	}

	if col == len(base) {
		row++
		col = 0
	}

	tmp := *solution

	if tmp[row][col] != '.' {
		SolveRecur(row, col+1, solution, solsFound)
		// fmt.Println("exiting for some reason")
		return
	}

	for num := 0; num < len(base); num++ {
		if isSafe(*solution, row, col, rune(base[num])) {
			tmp = *solution
			tmp[row][col] = rune(base[num])
			solution = &tmp

			SolveRecur(row, col+1, solution, solsFound)

			// fmt.Println("solution")
		}

		if *solsFound > 1 {
			// fmt.Println("exiting cause more than 1 solution found")
			return
		}

		tmp = *solution
		tmp[row][col] = '.'
		solution = &tmp
	}
}
