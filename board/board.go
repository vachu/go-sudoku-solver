package board

import (
	"fmt"
	"io"
	"sync"
)

type block [9]*byte

// Block defines the type for a 9-element array of bytes
type Block [9]byte

// Type Board defines the main Sudoku Boarod of 9x9 cells.  None of the fields
// are exported
type Board struct {
	cells   [9][9]byte  // individual cells containing numbers 1 to 9
	rows    [9]block    // each row has pointers to 9 cells horizontally across the board
	cols    [9]block    // each col has pointers to 9 cells vertically across the board
	squares [9]block    // linear array of 9 cells having pointers to 3x3 squares
	m       *sync.Mutex // mutex to synchronize read and writes to cells
}

// NewBoard returns an zero-initialized new Sudoku Board
func NewBoard() *Board {
	board := Board{}
	board.m = &sync.Mutex{}
	for i := 0; i < 9; i++ {
		kBase := (i % 3) * 3
		for j := 0; j < 9; j++ {
			board.rows[i][j] = &board.cells[i][j]
			board.cols[i][j] = &board.cells[j][i]

			sqIdx := (i/3)*3 + j/3
			k := j%3 + kBase
			board.squares[sqIdx][k] = &board.cells[i][j]
		}
	}
	return &board
}

// GetRow returns an array containing the cell values for the specified row index (0 - 8).
// The returned array always contains 9 elements.  The error object is non-nil if there is an error
func (b *Board) GetRow(rowIndex int) (row [9]byte, e error) {
	if rowIndex < 0 || rowIndex > 8 {
		e = fmt.Errorf("rowIndex out of range")
	} else {
		b.m.Lock()
		for i := 0; i < 9; i++ {
			row[i] = *b.rows[rowIndex][i]
		}
		b.m.Unlock()
	}
	return
}

// GetCol returns an array containing the cell values for the specified column index (0 - 8).
// The returned array always contains 9 elements.  The error object is non-nil if there is an error
func (b *Board) GetCol(colIndex int) (col [9]byte, e error) {
	if colIndex < 0 || colIndex > 8 {
		e = fmt.Errorf("colIndex out of range")
	} else {
		b.m.Lock()
		for i := 0; i < 9; i++ {
			col[i] = *b.cols[colIndex][i]
		}
		b.m.Unlock()
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
		b.m.Lock()
		value = b.cells[rowIndex][colIndex]
		b.m.Unlock()
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
		e = fmt.Errorf("'%d' - value out of range", value)
	} else {
		sqIndex := (rowIndex/3)*3 + colIndex/3
		b.m.Lock()
		prevValue := b.cells[rowIndex][colIndex]
		b.cells[rowIndex][colIndex] = 0 // clear the cell for which new value has to be set
		if hasValue, index := b.rows[rowIndex].has(value); hasValue {
			// if the value to be set is present in any other cell in the containing row
			e = fmt.Errorf("'%v' already present in cell[%d][%d]", value, rowIndex, index)
			b.cells[rowIndex][colIndex] = prevValue // revert to previous value
		} else if hasValue, index = b.cols[colIndex].has(value); index != rowIndex && hasValue {
			// if the value to be set is present in any other cell in the containing column
			e = fmt.Errorf("'%v' already present in cell[%d][%d]", value, index, colIndex)
			b.cells[rowIndex][colIndex] = prevValue // revert to previous value
		} else if hasValue, _ = b.squares[sqIndex].has(value); hasValue {
			// if the value to be set is present in any other cell in the containing 3x3 square
			e = fmt.Errorf("'%v' already present in square[%d]", value, sqIndex)
			b.cells[rowIndex][colIndex] = prevValue // revert to previous value
		} else {
			b.cells[rowIndex][colIndex] = value // set the new value
		}
		b.m.Unlock()
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
		b.m.Lock()
		b.cells[rowIndex][colIndex] = 0
		b.m.Unlock()
	}
	return
}

// GetSquare returns the cell-values of the specified 3x3 square in a uni-dimensional
// array of 9 elements.  The error object is non-nil if there is an error
func (b *Board) GetSquare(sqIndex int) (square [3][3]byte, e error) {
	if sqIndex < 0 || sqIndex > 8 {
		e = fmt.Errorf("sqIndex out of range")
	} else {
		startRow := (sqIndex / 3) * 3
		startCol := (sqIndex % 3) * 3
		b.m.Lock()
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				square[i][j] = b.cells[startRow+i][startCol+j]
			}
		}
		b.m.Unlock()
	}
	return
}

func (b *block) has(value byte) (hasValue bool, index int) {
	if value > 9 {
		index = -2
	} else {
		index = -1
		for i := 0; i < 9; i++ {
			if hasValue = (*b[i] == value); hasValue {
				index = i
				break
			}
		}
	}
	return
}

func (b *Board) Clear() {
	b.m.Lock()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.cells[i][j] = 0
		}
	}
	b.m.Unlock()
}

func (b *Board) BeautyPrint(w io.Writer) {
	const (
		BORDER               = "====================="
		HORIZONTAL_SEPERATOR = "------+-------+------"
	)

	fmt.Fprintln(w, BORDER)
	b.m.Lock()
	for i := 0; i < 9; i++ {
		if i == 3 || i == 6 {
			fmt.Fprintln(w, HORIZONTAL_SEPERATOR)
		}
		r := b.rows[i]
		fmt.Fprintln(w, *r[0], *r[1], *r[2], "|", *r[3], *r[4], *r[5], "|", *r[6], *r[7], *r[8])
	}
	b.m.Unlock()
	fmt.Fprintln(w, BORDER)
}

func (b *Board) Write(w io.Writer) error {
	b.m.Lock()
	for i := 0; i < 9; i++ {
		r := b.rows[i]
		if _, e := fmt.Fprintln(w, *r[0], *r[1], *r[2], *r[3], *r[4], *r[5], *r[6], *r[7], *r[8]); e != nil {
			return e
		}
	}
	b.m.Unlock()
	return nil
}

func (b *Board) Read(r io.Reader) error {
	newBoard := NewBoard()
	for i := 0; i < 9; i++ {
		var c [9]byte
		n, e := fmt.Fscanln(r, &c[0], &c[1], &c[2], &c[3], &c[4], &c[5], &c[6], &c[7], &c[8])
		if e != nil {
			return fmt.Errorf("error reading line from input - %v", e)
		}
		if n != 9 {
			return fmt.Errorf("only %d instead of item(s) read", n)
		}
		for j := 0; j < 9; j++ {
			if c[j] > 0 {
				if e = newBoard.SetCellValue(i, j, c[j]); e != nil {
					return e
				}
			}
		}
	}

	b.m.Lock()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.cells[i][j] = newBoard.cells[i][j]
		}
	}
	b.m.Unlock()
	return nil
}
