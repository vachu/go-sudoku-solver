package main

import (
	"fmt"

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
	for i := 0; i < 9; i++ {
		row, _ := board.GetRow(i)
		fmt.Println(row)
	}
}
