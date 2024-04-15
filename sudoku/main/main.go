package main

import (
	"fmt"
	"os"
	"reflect"

	"sudoku"
)

// ERROR CODES:
const (
	OK               = -1
	WRONG_ARG_NUM    = 0
	WRONG_ROW_LENGTH = 1
	WRONG_CHARACTERS = 2
	TOO_MANY_FLAGS   = 3
	NOSOLUTION       = 4
	TOO_MANY_SOL     = 5
	SUDOKU_TOO_SMALL = 6
)

func main() {
	mainLoop()
}

var state int = 0

var (
	solved_sudoku [][]rune
	og_sudoku     [][]rune
)

// Finite state machine states:
const (
	LAUNCH       = 0
	WAITINGINPUT = 1
	HELPING      = 2
	SOLVING      = 3
	HINTING      = 4
	EXIT         = 5
	RESTART      = 6
)

// Flags used to determine how the program should behave
var (
	errorNum           int  = -1
	flaghelp           bool = false
	flaghint           bool = false
	flagsolve          bool = false
	flagexit           bool = false
	flagrestart        bool = false
	justfinishedsudoku bool = false
)

func mainLoop() {
	fmt.Println("")
	for {
		input := []string{}
		switch state {
		//
		//
		//
		//==============================================================================
		case LAUNCH:
			sudoku.PrintIntro()
			input = []string(os.Args[1:])
			if len(input) > 0 {
				if len(input) > 4 {
					if input[len(input)-3] == "[object Object]" && input[len(input)-2] == "cat" && input[len(input)-1] == "-e" {
						input = input[0 : len(input)-2]
					}
				}
				ParseInput(input)
				ParseFlags()

			} else {
				sudoku.PrintBasicInstructions()
				state = WAITINGINPUT
			}
		//
		//
		//
		//==============================================================================
		case WAITINGINPUT:

			if flags == 0 && len(solved_sudoku) > 0 {
				fmt.Println("\n\n\n\n\n\n\n")
				sudoku.PrintFancySudoku(og_sudoku)
				fmt.Println("     |[｡·  v  ·｡]| ?")
				fmt.Println("\nWhat would you like me to do with this sudoku? --solve or --hint?\n")

			} else {
				if len(solved_sudoku) == 0 && justfinishedsudoku == true {
					sudoku.GiveNextSudoku()
					justfinishedsudoku = false
				}
			}

			result := sudoku.ReadTerminalInput()
			// result := []string{"--solve", "--hint", ".96.4...1", "1...26...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"}
			// fmt.Println(" debug: ", result)
			if errorNum == TOO_MANY_SOL && (result[0] == "yes" || result[0] == "y") { // answer to the question on if there are more than one solutions, would you like to see one?
				sudoku.PrintFancySudoku(solved_sudoku)
				errorNum = OK
				state = WAITINGINPUT
			}
			if errorNum == TOO_MANY_SOL && (result[0] == "no" || result[0] == "n") {
				fmt.Println("\n\n\n\n\n\n\n     ¯\\_|[  ツ  ]|_/¯\n")
				errorNum = OK
				state = WAITINGINPUT
			}
			ParseInput(result)
			ParseFlags()
		//
		//
		//
		//==============================================================================
		case HELPING:
			sudoku.PrintFullHelp()
			state = WAITINGINPUT
		//
		//
		//
		//==============================================================================
		case SOLVING:

			if len(solved_sudoku) < 1 {
				fmt.Println("\n\n\n\n\n\n\n    |[   ˘ ³˘ ]|ノ°ﾟº❍｡\n")
				fmt.Println("\nYou need to provide me with a valid sudoku board first.\n")
				state = WAITINGINPUT
			} else {
				fmt.Println("\n\n\n\n\n\n\n    |[ ▀̿ ‿ ̿▀̿  ]|\n")
				fmt.Println("\n      Sudoku solved.\n         |\n         V")
				sudoku.PrintFancySudoku(solved_sudoku)
				solved_sudoku = [][]rune{}
				justfinishedsudoku = true

				state = WAITINGINPUT
			}
		//
		//
		//
		//==============================================================================
		case HINTING:
			if len(solved_sudoku) < 1 {
				fmt.Println("\n\n\n\n\n\n\n     |[ =^･ｪ･^=  ]|\n")
				fmt.Println("\nUnfortunately i can't provide you with life advice.")
				fmt.Println("If you'd like a hint for an unsolved sudoku board, you need to provide a valid one.\n")
			} else {
				fmt.Println("\n\n\n\n\n\n\n     |[  っ´ω`c ]|")
				fmt.Println("\nCertainly! Here is but a meager fragment of my power.\n")
				og_sudoku = sudoku.ShowHint(og_sudoku, solved_sudoku)

				if reflect.DeepEqual(og_sudoku, solved_sudoku) {
					fmt.Println("\n\n    |[  ･ิω･ิ#   ]|\n")
					fmt.Println("\nYour sudoku is now solved! ...i'm not sure why you didn't just use the solve function, but you do you, meatbag!\n")
					// sudoku.PrintFancySudoku(solved_sudoku)
					solved_sudoku = [][]rune{}
					justfinishedsudoku = true

				} else {
					fmt.Println("\n\n    |[૮  ˶•⤙•˶ ა  ]|\n")
					fmt.Println("\nUse --hint or -i to get the next hint!\n")
				}

			}
			state = WAITINGINPUT
		//
		//
		//
		//==============================================================================
		case EXIT:
			fmt.Println("\n\n\n\n\n\n\n   |[   ҂◡_◡ ]|\n")
			fmt.Println("\nShutting down...\n")
			return
		//
		//
		//
		//==============================================================================
		case RESTART:
			fmt.Println("\n\n\n\n\n\n\nᕕ|[   ᐛ  ]|ᕗ\n")
			fmt.Println("\nResetting...\n")
			solved_sudoku = [][]rune{}
			justfinishedsudoku = false
			og_sudoku = [][]rune{}
			sudoku.PrintIntro()
			state = WAITINGINPUT
		}
	}
}

var flags int

func ParseInput(input []string) {
	sudokustring := true
	sudokuinputs := make([]string, 0)
	flags = 0
	badInputReceived := false
	for _, arg := range input {
		sudokustring = true
		if len(arg) <= 9 {
			for _, chara := range arg {
				if !((chara >= '0' && chara <= '9') || chara == '.') {
					sudokustring = false
					break
				}
			}
		} else {
			for _, chara := range arg {
				if !((chara >= '0' && chara <= '9') || chara == '.' || (chara >= 'A' && chara <= 'G')) {
					sudokustring = false
					break
				}
			}
		}
		if sudokustring {
			sudokuinputs = append(sudokuinputs, arg)
		}
		if arg == "--help" || arg == "help" || arg == "-h" {
			flaghelp = true
			flags++
		} else if arg == "--hint" || arg == "hint" || arg == "-i" {
			flaghint = true
			flags++
		} else if arg == "--solve" || arg == "solve" || arg == "-s" {
			flagsolve = true
			flags++
		} else if arg == "--exit" || arg == "exit" || arg == "-e" {
			flagexit = true
			flags++
		} else if arg == "--restart" || arg == "restart" || arg == "-r" {
			flagrestart = true
			flags++
		} else if !sudokustring {
			badInputReceived = true
		}
	}

	if badInputReceived == true {
		fmt.Println("\n\n\n\n\n\n\n    |[ ´סּ︵סּ`  ]|\n")
		fmt.Println("\nWhat the heck's all this? I'm not cleverbot, you need to give me a valid input (command + board or just board or just command!).\n")
	}

	if flags > 1 {
		sudoku.PrintError(TOO_MANY_FLAGS)
		flaghelp = false
		flaghint = false
		flagsolve = false
		flagexit = false
		flagrestart = false
		flags = 0
	}

	if len(sudokuinputs) < 1 {
		if (len(solved_sudoku) < 1 && len(input) > 1) || (len(sudokuinputs) < 1 && (!flaghelp && !flaghint && !flagsolve && !flagexit && !flagrestart)) {
			// fmt.Println("\n\n\n\n\n\n\n      |[ ´סּ︵סּ`  ]|\n")
			// fmt.Println("\nWhat the heck's all this? I'm not cleverbot, you need to give me a valid input (command + board or just board).\n")
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			state = WAITINGINPUT
			return
		}
	}

	if len(sudokuinputs) != 0 {
		isValid, errorNum, numcount := sudoku.Validation(sudokuinputs)

		if isValid == true && numcount < 17 && len(sudokuinputs) == 9 {
			isValid = false
			errorNum = SUDOKU_TOO_SMALL
		}

		countSols := 0
		if !isValid {
			if badInputReceived == false {
				sudoku.PrintError(errorNum)
			}
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			return
		} else {
			og_sudoku = sudoku.MakingTheSudoku(sudokuinputs)
			countSols, solved_sudoku = sudoku.SolveValidation(sudokuinputs)
			if countSols == 0 {
				errorNum = NOSOLUTION
				sudoku.PrintError(NOSOLUTION)
				flaghelp = false
				flaghint = false
				flagsolve = false
				flagexit = false
				flagrestart = false
				solved_sudoku = [][]rune{}
				justfinishedsudoku = true
				state = WAITINGINPUT

				return
			} else if countSols > 1 {
				errorNum = TOO_MANY_SOL
				sudoku.PrintError(TOO_MANY_SOL)
				flaghelp = false
				flaghint = false
				flagsolve = false
				flagexit = false
				flagrestart = false
				state = WAITINGINPUT
				return
			}

		}
	}
	if len(solved_sudoku) > 1 && flags == 0 && state == LAUNCH {
		flagsolve = true
		flags++
	}
}

// Parses the flags that are currently on and sets the corresponding state
// ==============================================================================
func ParseFlags() {
	if flaghelp {
		if !flaghint && !flagsolve && !flagexit && !flagrestart {
			flaghelp = false
			state = HELPING
			return
		} else {
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			sudoku.PrintError(TOO_MANY_FLAGS)
			state = WAITINGINPUT
			return
		}
	}
	if flaghint {
		if !flaghelp && !flagsolve && !flagexit && !flagrestart {
			flaghint = false
			state = HINTING
			return
		} else {
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			sudoku.PrintError(TOO_MANY_FLAGS)
			state = WAITINGINPUT
			return
		}
	}
	if flagsolve {
		if !flaghelp && !flaghint && !flagexit && !flagrestart {
			flagsolve = false
			state = SOLVING
			return
		} else {
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			sudoku.PrintError(TOO_MANY_FLAGS)
			state = WAITINGINPUT
			return
		}
	}
	if flagexit {
		if !flaghelp && !flaghint && !flagsolve && !flagrestart {
			flagexit = false
			state = EXIT
			return
		} else {
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			sudoku.PrintError(TOO_MANY_FLAGS)
			state = WAITINGINPUT
			return
		}
	}
	if flagrestart {
		if !flaghelp && !flaghint && !flagexit && !flagsolve {
			flagrestart = false
			state = RESTART
			return
		} else {
			flaghelp = false
			flaghint = false
			flagsolve = false
			flagexit = false
			flagrestart = false
			sudoku.PrintError(TOO_MANY_FLAGS)
			state = WAITINGINPUT
			return
		}
	}
	state = WAITINGINPUT
}
