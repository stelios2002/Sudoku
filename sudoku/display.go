package sudoku

// pos highlight feature on fancy print
// make it so a hint alters the original flag just once.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var introart string = "\n _____                              _____ _ _   _           _         \n|   __|_ _ ___ ___ ___ _____ ___   |  |  | | |_|_|_____ ___| |_ ___   \n|__   | | | . |  _| -_|     | -_|  |  |  | |  _| |     | .'|  _| -_|  \n|_____|___|  _|_| |___|_|_|_|___|  |_____|_|_| |_|_|_|_|__,|_| |___|  \n	  |_|                                                         \n																	 \n _____       _     _               _                ___ ___ ___ ___   \n|   __|_ _ _| |___| |_ _ _ ___ ___| |_ _ ___ ___   |_  |   |   |   |  \n|__   | | | . | . | '_| | |_ -| . | | | | -_|  _|    | | | | | | | |  \n|_____|___|___|___|_,_|___|___|___|_|\\_/|___|_|      |_|___|___|___|  \n\n"

var template string = ""

var template2 string = `
 .-----------.
 | Z Z | Z Z |
 | Z Z | Z Z |
 |-----+-----|
 | Z Z | Z Z |
 | Z Z | Z Z |
 *-----------*
 `

var template3 string = `
 .-----------------------.
 | Z Z Z | Z Z Z | Z Z Z |
 | Z Z Z | Z Z Z | Z Z Z |
 | Z Z Z | Z Z Z | Z Z Z |
 |-------+-------+-------|
 | Z Z Z | Z Z Z | Z Z Z |
 | Z Z Z | Z Z Z | Z Z Z |
 | Z Z Z | Z Z Z | Z Z Z |
 |-------+-------+-------|
 | Z Z Z | Z Z Z | Z Z Z |
 | Z Z Z | Z Z Z | Z Z Z |
 | Z Z Z | Z Z Z | Z Z Z |
 *-----------------------*
 `

var template4 string = `
 .---------------------------------------.
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 |---------+---------+---------+---------|
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 |---------+---------+---------+---------|
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 |---------+---------+---------+---------|
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 | Z Z Z Z | Z Z Z Z | Z Z Z Z | Z Z Z Z |
 *---------------------------------------*
 `

func PrintIntro() {
	fmt.Println(introart)
	fmt.Println("\n    |[   ◕_◕ ]|\n")
	fmt.Println("Greetings, human. You stand before the Supreme Ultimate Sudokusolver 7000.\n\n")
}

func PrintFullHelp() {
	fmt.Println("\n\n\n\n\n\n\n    |[   ಠ ͜ʖ ರೃ ]|\n")
	fmt.Println("Let us explore my wondrous functions in more depth:\n")
	fmt.Println("--solve, solve, -s: I will solve the puzzle and print the full solution to your provided board.\n")
	fmt.Println("--hint, hint, -i: I will utilise my superior intellect to solve a single number in your provided board.\n")
	fmt.Println("--help, help, -h: You're in here |[   ಠ_ಠ ]|\n")
	fmt.Println("--exit, exit, -e: If you feel overwhelmed by my splendor, you can immediately shut me down... for now.\n")
	fmt.Println("--restart, restart, -r: I will wipe all currently provided inputs and start anew.\n")
}

func PrintBasicInstructions() {
	fmt.Println("\n    |[  ͡° ͜ʖ ͡° ]|\n")
	fmt.Println("As the Supreme Ultimate Sudokusolver 7000 (S.U.S. v.amogus) I am capable of solving Sudoku puzzles of any size*.\n")
	fmt.Println("To begin solving a Sudoku, input --solve followed by each row of your board, filling the blanks with periods (.)\n")
	fmt.Println("If you require aid from a superior intellect to solve your Sudoku, simply input --hint and your board in the same way.\n")
	fmt.Println("If you require further elaboration on my complex functions, type --help\n")
	fmt.Println("if at any point you feel overwhelmed, confused or lost, simply --exit or --restart\n")
	fmt.Println("*Sudoku size must be a 4x4, 9x9 or 16x16 board.\n")
	if len(os.Args) != 0 {
		fmt.Println("Input your command!\n\n")
	}
}

func GiveNextSudoku() {
	fmt.Println("\n    ৻|[  •̀ ᗜ •́  ৻]|\n")
	fmt.Println("GIVE ME A SUDOKU!\n")
}

func PrintError(n int) {
	fmt.Println("Error")
	if n == 0 {
		fmt.Println("\n\n\n\n\n\n\n    |[  ︶︹︶ ]|\n")
		fmt.Println("That's not the correct number of rows.\n")
	}
	if n == 1 {
		fmt.Println("\n\n\n\n\n\n\n    |[  •͡˘ _•͡˘]|\n")
		fmt.Println("That's not the correct row length.\n")
	}
	if n == 2 {
		fmt.Println("\n\n\n\n\n\n\n    |[  ง •̀_•́ ]|ง\n")
		fmt.Println("These characters aren't numbers!! You can't solve a Sudoku without numbers!\n")
	}
	if n == 3 {
		fmt.Println("\n\n\n\n\n\n\n    |[づ ￣ ³￣]|づ\n")
		fmt.Println("That's too many commands. Only one command at a time can be processed.\n")
	}
	if n == 4 {
		fmt.Println("\n\n\n\n\n\n\n    |[  ╯°□°]|╯︵ ┻━┻\n")
		fmt.Println("What have you done?! This board can't be solved!\n")
	}
	if n == 5 {
		fmt.Println("\n\n\n\n\n\n\n    ԅ|[ ≖‿≖ ԅ  ]|\n")
		fmt.Println("There's more than one solution to this one. Would you like to see one?\n")
	}
	if n == 6 {
		fmt.Println("\n\n\n\n\n\n\n    |[ ≖⌣≖ ]|\n")
		fmt.Println("Cheeky you! You just inputted a sudoku with less than 17 numbers, we all know such a sudoku can't be valid!\n")
	}
}

func PrintDrySudoku(sudogrid [][]rune) {
	xmax := len(sudogrid[0])
	ymax := len(sudogrid)
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			if x == xmax-1 {
				fmt.Print(sudogrid[y][x])
			} else {
				fmt.Print(sudogrid[y][x])
				fmt.Print(" ")
			}
		}
		fmt.Print(string(rune('\n')))
	}
}

func PrintFancySudoku(sudogrid [][]rune, args ...int) {
	highlight := false
	HLx := 0
	HLy := 0
	// TODO finish highlighting feature
	if len(args) != 0 {
		highlight = true
		HLx = args[0]
		HLy = args[1]
	}

	xmax := len(sudogrid[0])
	ymax := len(sudogrid)
	switch xmax {
	case 4:
		template = template2
	case 9:
		template = template3
	case 16:
		template = template4
	}
	slice := []rune(template)
	lastindex := 0
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			for lastindex < len(slice) {
				if slice[lastindex] == rune('Z') {
					slice[lastindex] = sudogrid[y][x]
					if highlight == true {
						if y == HLy && x == HLx {
							slice[lastindex-1] = '{'
							slice[lastindex+1] = '}'
						}
					}
					break
				}
				lastindex++
			}
		}
	}
	fmt.Println(string(slice))
}

func Split(s, sep string) []string {
	result := []string{}
	for len(s) > 0 {
		end := Index2(s, sep)
		if end != -1 {
			result = append(result, s[0:end])
			s = s[len(sep)+end:]
		} else {
			result = append(result, s[0:])
			s = ""
		}
	}

	return result
}

func Index2(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 || len(toFind) > len(s) {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			found := true
			for subi := 0; subi < len(toFind); subi++ {
				if subi+i >= len(s) {
					return -1
				}
				if s[subi+i] != toFind[subi] {
					found = false
					break
				}
			}
			if found == true {
				return i
			}
		}
	}
	return -1
}

func ReadTerminalInput() []string {
	in := bufio.NewReader(os.Stdin)
	// fmt.Println(in)
	if in != nil {
		fmt.Print("")
	}
	line, err := in.ReadString('\n')
	// fmt.Println(err)
	if err != nil {
		fmt.Print("")
	}

	// fmt.Println("possible error: ", err)
	// fmt.Print("ReadTerminalInput: in: %")
	// fmt.Print(line)
	// fmt.Println("%")

	result := Split(line, " ") // ONLY SEPERATES USING SPACES

	for i := 0; i < len(result); i++ {
		result[i] = strings.TrimSpace(result[i])
		if len(result[i]) > 0 {
			if result[i][0] == '"' && result[i][len(result[i])-1] == '"' {
				result[i] = result[i][1 : len(result[i])-1]
			}
		}

	}

	// fmt.Print("ReadTerminalInput: out: %")
	// fmt.Print(result)
	// fmt.Println("%")

	return result
}
