package sudoku





func TestPrint() {
	sudogrid := [][]int{}
	for y := 0; y < 9; y++ {
		row := []int{}
		for x := 0; x < 9; x++ {
			row = append(row, x)
		}
		sudogrid = append(sudogrid, row)
	}
	// fmt.Println("dry: ")
	// PrintDrySudoku(sudogrid)
	// fmt.Println(" ")
	// fmt.Println("fancy: ")
	// PrintFancySudoku(sudogrid)
	// fmt.Println(" ")
}

// func ReadTerminalInput() []string {
// 	in := bufio.NewReader(os.Stdin)
// 	line, err := in.ReadString('\n')

// 	fmt.Println("possible error: ", err)
// 	// fmt.Print("ReadTerminalInput: in: %")
// 	// fmt.Print(line)
// 	// fmt.Println("%")
// 	result := Split(line, " ") // ONLY SEPERATES USING SPACES

// 	for i := 0; i < len(result); i++ {
// 		if result[i][0] == '"' && result[i][len(result[i])-1] == '"' {
// 			result[i] = result[i][1 : len(result[i])-1]
// 		}
// 		result[i] = strings.TrimSpace(result[i])
// 	}

// 	// fmt.Print("ReadTerminalInput: out: %")
// 	// fmt.Print(result)
// 	// fmt.Println("%")

// 	return result
// }
