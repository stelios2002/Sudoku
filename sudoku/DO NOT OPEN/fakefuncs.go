package sudoku

import "fmt"

func ValidationFake(args []string) (bool, int) { // returns error code and grid
	return true, -1
}

func SolveValidationFake() (int, [][]rune) { // err code, count of solutoin, solution grid

	sudogrid := [][]rune{}
	for y := 0; y < 9; y++ {
		row := []rune{}
		for x := 0; x < 9; x++ {
			row = append(row, rune(x+'0'))
		}
		sudogrid = append(sudogrid, row)
	}
	// PrintFancySudoku(sudogrid)
	return 1, sudogrid
}

var (
	grid             [][]rune
	numscount        int
	xmax             int
	ymax             int
	solutions        int
	base2            string
	lastsolution     [][]rune
	disaster_end_all = false // in case there are more than 1 solution found
)

func TestFakeReal() {
	stringslice := []string{"E.B32.1G8D4..6.F", ".52ABC..E.7F9.3.", ".91.7AD.BC.5E4.8", "F7.DE45.1692.AG.", "...5CF.....83.12", "7..63285A.1.FC..", "2..B.7..D..3G5.6", ".3...6GA25BC.79.", "3..C91FB..5.627A", "5B6.G..E7F.148.3", "..G..3.7.AC.B.F.", "A27F6..C93E.5.DG", ".6.7AB92.ED.1F8C", "B.54.E.1C...A.67", ".DE.85.6FB..234.", "8.92FGC431...B.5"}
	PrintFancySudoku(MakingTheSudoku(stringslice))
	x, y := SolveValidation(stringslice)
	fmt.Println(y)
	a, b := Validation(stringslice)
	fmt.Println("validation: ", a, " ", b)
	if x == 1 {
		PrintFancySudoku(y)
	} else {
		fmt.Println("not only 1 solution: ", x)
		PrintDrySudoku(y)
	}
}

func SolveValidationFakeReal(args []string) (int, [][]rune) {
	lastsolution = [][]rune{}
	disaster_end_all = false

	base2 = BaseCreation(args)
	grid = MakingTheSudoku(args)
	numscount = len(grid)
	xmax = len(grid[0])
	ymax = len(grid)
	solutions = 0
	recur(0, 0)
	return solutions, lastsolution
}

func recur(prevx, prevy int) {
	fmt.Println("new recur")
	PrintFancySudoku(grid)
	if grid[prevy][prevx] == '.' {
		for tryi := 0; tryi < len(base2); tryi++ {
			tryr := rune(base2[tryi])
			if isSafe3(grid, prevx, prevy, tryr) {
				grid[prevy][prevx] = tryr
				if prevx+1 == xmax {
					if prevy+1 == ymax {
						fmt.Println("solution found!")
						solutions++
						if solutions > 1 {
							disaster_end_all = true
							return
						}
						saveSolution()

						grid[prevy][prevx] = '.'
						return // returning after finding solution, to check if there are more than 1

					} else {
						recur(0, prevy+1)
						if disaster_end_all == true {
							return
						}
					}
				} else {
					recur(prevx+1, prevy)
					if disaster_end_all == true {
						return
					}
				}
			}
		}
		grid[prevy][prevx] = '.'
		return // returning cause options have been exhausted
	} else {
		if prevx+1 == xmax {
			if prevy+1 == ymax {
				fmt.Println("solution found!")
				solutions++
				if solutions > 1 {
					disaster_end_all = true
					return
				}
				saveSolution()

				grid[prevy][prevx] = '.'
				return // returning after finding solution, to check if there are more than 1

			} else {
				recur(0, prevy+1)
				if disaster_end_all == true {
					return
				}
			}
		} else {
			recur(prevx+1, prevy)
			if disaster_end_all == true {
				return
			}
		}
	}
}

// func isSafe2(x, y int) bool {
// 	return true
// }

func saveSolution() {
	lastsolution := make([][]rune, len(grid))
	for i := range grid {
		lastsolution[i] = make([]rune, len(grid[i]))
		for j := range grid[i] {
			lastsolution[i][j] = grid[i][j]
		}
	}
}

func isSafe3(grid [][]rune, y, x int, num rune) bool {
	subboxlen := len(grid) / 3
	// Check if the number is not already present in the row and column
	for i := 0; i < len(base); i++ {
		if grid[y][i] == num || grid[i][x] == num {
			return false
		}
	}

	// Check if the number is not already present in the n x n subgrid

	startY := y - y%subboxlen
	startX := x - x%subboxlen

	for i := 0; i < subboxlen; i++ {
		for j := 0; j < subboxlen; j++ {
			if grid[i+startY][j+startX] == num {
				return false
			}
		}
	}

	return true
}
