package main

import "fmt"

func main() {
	type File []bool

	type Chessboard map[string]File

	var cb Chessboard = map[string]File{
		"A": {true, false, true, false, false, false, false, true},
		"B": {false, false, false, false, true, false, false, false},
		"C": {false, false, true, false, false, false, false, false},
		"D": {false, false, false, false, false, false, false, false},
		"E": {false, false, false, false, false, true, false, true},
		"F": {false, false, false, false, false, false, false, false},
		"G": {false, false, false, true, false, false, false, false},
		"H": {true, true, true, true, true, true, false, true},
	}

	// file := "A"
	rank := 2

	count := 0

	// for k, v := range cb {
	// 	if k == file {
	// 		for _, val := range v {
	// 			if val == true {
	// 				count += 1
	// 			}
	// 		}
	// 	}
	// }

	l := len(cb)

	if rank >= 1 && rank <= l {
		for _, f := range cb {
			for i, v := range f {
				if i+1 == rank && v {
					count += 1
				}
			}
		}
	}

	fmt.Println(cb)

	fmt.Println(count)
}
