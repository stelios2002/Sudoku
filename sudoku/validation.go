package sudoku

func Validation(args []string) (bool, int, int) {
	numcounter := 0
	if len(args) != 16 {
		if len(args) != 9 {
			if len(args) != 4 {
				return false, 0, numcounter
			}
		}
	}
	if len(args) == 9 {
		for _, arg := range args {
			if len(arg) != 9 {
				return false, 1, numcounter
			}
			for _, char := range arg {
				if (char < '1' || char > '9') && char != '.' {
					return false, 2, numcounter
				} else {
					if char != '.' {
						numcounter++
					}
				}
			}
		}
	}
	if len(args) == 4 {
		for _, arg := range args {
			if len(arg) != 4 {
				return false, 1, numcounter
			}
			for _, char := range arg {
				if (char < '1' || char > '4') && char != '.' {
					return false, 2, numcounter
				} else {
					if char != '.' {
						numcounter++
					}
				}
			}
		}
	}
	if len(args) == 16 {
		for _, arg := range args {
			if len(arg) != 16 {
				return false, 1, numcounter
			}
			for _, char := range arg {
				if (char < '1' || char > '9') && (char < 'A' || char > 'G') && char != '.' {
					return false, 2, numcounter
				} else {
					if char != '.' {
						numcounter++
					}
				}
			}
		}
	}
	return true, -1, numcounter
}
