package board

import "fmt"

type block [9]*byte

type Board struct { // the whole 9x9 sudoku square
	cells   [9][9]byte // individual cells containing numbers 1 to 9
	rows    [9]block   // each row has pointers to 9 cells horizontally across the board
	cols    [9]block   // each col has pointers to 9 cells vertically across the board
	squares [9]block   // 3x3 squares containing 9 cells
}

// NewBoard returns an zero-initialized new Sudoku Board
func NewBoard() *Board {
	board := Board{}
	for i := 0; i < 9; i++ {
		// fmt.Println("i =", i)
		kBase := (i % 3) * 3
		for j := 0; j < 9; j++ {
			board.rows[i][j] = &board.cells[i][j]
			board.cols[i][j] = &board.cells[j][i]

			sqIdx := (i/3)*3 + j/3
			k := j%3 + kBase
			board.squares[sqIdx][k] = &board.cells[i][j]
			// fmt.Printf("\tsqIdx = %d; k = %d\n", sqIdx, k)
		}
		// fmt.Println()
	}
	return &board
}

// GetRow returns an array containing the cell values for the specified row index (0 - 8).
// The returned array always contains 9 elements.  The error object is non-nil if there is an error
func (b *Board) GetRow(rowIndex int) (row [9]byte, e error) {
	if rowIndex < 0 || rowIndex > 8 {
		e = fmt.Errorf("rowIndex out of range")
	} else {
		for i := 0; i < 9; i++ {
			row[i] = *b.rows[rowIndex][i]
		}
	}
	return
}

// GetCol returns an array containing the cell values for the specified column index (0 - 8).
// The returned array always contains 9 elements.  The error object is non-nil if there is an error
func (b *Board) GetCol(colIndex int) (col [9]byte, e error) {
	if colIndex < 0 || colIndex > 8 {
		e = fmt.Errorf("colIndex out of range")
	} else {
		for i := 0; i < 9; i++ {
			col[i] = *b.cols[colIndex][i]
		}
	}
	return
}

// GetCellValue returns the value contained in the specified cell (by row and column Indices).
// The error object is non-nil if there is an error
func (b *Board) GetCellValue(rowIndex, colIndex int) (value byte, e error) {
	if rowIndex < 0 || rowIndex > 8 {
		e = fmt.Errorf("rowIndex out of range")
	} else if colIndex < 0 || colIndex > 8 {
		e = fmt.Errorf("colIndex out of range")
	} else {
		value = b.cells[rowIndex][colIndex]
	}
	return
}

// SetCellValue sets the specified value to the cell identified by the specified row and column
// indices.  The error object is non-nil if there is an error
func (b *Board) SetCellValue(rowIndex, colIndex int, value byte) (e error) {
	if rowIndex < 0 || rowIndex > 8 {
		e = fmt.Errorf("rowIndex out of range")
	} else if colIndex < 0 || colIndex > 8 {
		e = fmt.Errorf("colIndex out of range")
	} else if value < 1 || value > 9 {
		e = fmt.Errorf("value out of range")
	} else {
		b.cells[rowIndex][colIndex] = value
	}
	return
}

// ClearCellValue clears / zero-es the value in the cell identified by the specified
// row and column indices.  The error object is non-nil if there is an error
func (b *Board) ClearCellValue(rowIndex, colIndex int) (e error) {
	if rowIndex < 0 || rowIndex > 8 {
		e = fmt.Errorf("rowIndex out of range")
	} else if colIndex < 0 || colIndex > 8 {
		e = fmt.Errorf("colIndex out of range")
	} else {
		b.cells[rowIndex][colIndex] = 0
	}
	return
}

// GetSquare returns the cell-values of the specified 3x3 square in a uni-dimensional
// array of 9 elements.  The error object is non-nil if there is an error
func (b *Board) GetSquare(sqIndex int) (square [9]byte, e error) {
	if sqIndex < 0 || sqIndex > 8 {
		e = fmt.Errorf("sqIndex out of range")
	} else {
		startRow := (sqIndex / 3) * 3
		startCol := (sqIndex % 3) * 3
		si := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				square[si] = b.cells[startRow+i][startCol+j]
				si++
			}
		}
	}
	return
}
