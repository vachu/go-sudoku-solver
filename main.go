package main

import (
	"fmt"
	"os"

	"vatsan.in/sudoku/board"
)

func main() {
	board := board.NewBoard()
	startCell := 3
	startCol := 3
	number := byte(1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.SetCellValue(startCell+i, startCol+j, number)
			number++
		}
	}

	if file, e := os.Create("C:\\temp\\sudoku.dat"); e != nil {
		fmt.Fprintf(os.Stderr, "Error - %v\n", e)
	} else {
		if e = board.Write(file); e != nil {
			fmt.Fprintf(os.Stderr, "Error - %v\n", e)
		}
		file.Close()
	}
	// board.Write(os.Stdout)

	if file, e := os.Open("C:\\temp\\sudoku.dat"); e != nil {
		fmt.Fprintln(os.Stderr, e)
	} else {
		if e = board.Read(file); e != nil {
			fmt.Fprintln(os.Stderr, e)
		} else {
			board.BeautyPrint(os.Stdout)
		}
		file.Close()
	}

	// if e := board.SetCellValue(0, 5, 9); e != nil {
	// 	fmt.Println("ERROR -", e)
	// }
	// if e := board.SetCellValue(3, 0, 3); e != nil {
	// 	fmt.Println("ERROR -", e)
	// }
	// if e := board.SetCellValue(4, 4, 1); e != nil {
	// 	fmt.Println("ERROR -", e)
	// }
}
