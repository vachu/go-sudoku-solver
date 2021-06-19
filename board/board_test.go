package board

import (
	"bytes"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	if b == nil {
		t.Fatal("Error creating new Sudoku Board")
	}
	if len(b.cells) != 9 {
		t.Error("wrong no. of rows in the board")
	}
	for i := 0; i < 9; i++ {
		if len(b.cells[i]) != 9 {
			t.Error("wrong no. of columns in the board")
		}
	}
}

func TestGetRow(t *testing.T) {
	b := NewBoard()

	data := []struct {
		rowIndex  int
		wantError bool
	}{
		{-1, true},
		{0, false},
		{8, false},
		{9, true},
		{10, true},
	}
	for _, d := range data {
		if _, e := b.GetRow(d.rowIndex); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil && d.wantError {
			t.Error("expected an error")
		}
	}

	b.cells[4][5] = 9 // cell at 5th row and 6th column has the value '9'
	expectedRow := [9]byte{0, 0, 0, 0, 0, 9, 0, 0, 0}
	if row, e := b.GetRow(4); e != nil {
		t.Errorf("Got error: %v", e)
	} else if row != expectedRow {
		t.Errorf("wanted %v but got %v", expectedRow, row)
	}
}

func TestGetCol(t *testing.T) {
	b := NewBoard()
	data := []struct {
		colIndex  int
		wantError bool
	}{
		{-1, true},
		{0, false},
		{8, false},
		{9, true},
		{10, true},
	}
	for _, d := range data {
		if _, e := b.GetCol(d.colIndex); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil && d.wantError {
			t.Error("expected an error")
		}
	}

	b.cells[4][5] = 9 // cell at 5th row and 6th column has the value '9'
	expectedCol := [9]byte{0, 0, 0, 0, 9, 0, 0, 0, 0}
	if col, e := b.GetCol(5); e != nil {
		t.Errorf("Got error: %v", e)
	} else if col != expectedCol {
		t.Errorf("wanted %v but got %v", expectedCol, col)
	}
}

func TestGetCellValue(t *testing.T) {
	b := NewBoard()
	data := []struct {
		rowIndex  int
		colIndex  int
		wantError bool
	}{
		{-1, 0, true},
		{0, 0, false},
		{9, 0, true},
		{10, 0, true},
		{8, 0, false},
		{0, -1, true},
		{0, 9, true},
		{0, 10, true},
		{0, 8, false},
	}
	for _, d := range data {
		if _, e := b.GetCellValue(d.rowIndex, d.colIndex); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil && d.wantError {
			t.Error("expected an error")
		}
	}

	b.cells[4][5] = 9
	if value, e := b.GetCellValue(4, 5); e != nil {
		t.Errorf("got an unexpected error - %v", e)
	} else if value != 9 {
		t.Errorf("expected value=9 but got value=%v", value)
	}
}

func TestSetCellValue(t *testing.T) {
	b := NewBoard()
	data := []struct {
		rowIndex  int
		colIndex  int
		value     byte
		wantError bool
	}{
		{-1, 0, 9, true},
		{0, 0, 9, false},
		{9, 0, 9, true},
		{10, 0, 9, true},
		{8, 0, 9, true},
		{8, 1, 9, false},
		{0, -1, 9, true},
		{0, 9, 9, true},
		{0, 10, 9, true},
		{0, 8, 9, true},
		{7, 8, 9, false},
		{0, 8, 9, true},
		{0, 8, 8, false},
		{1, 6, 8, true},
		{0, 5, 9, true},
		{0, 5, 8, true},
		{0, 5, 7, false},
		{4, 1, 9, true},
		{4, 1, 8, false},
	}
	for _, d := range data {
		if e := b.SetCellValue(d.rowIndex, d.colIndex, d.value); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil && d.wantError {
			t.Error("expected an error")
		}
	}

	b = NewBoard()
	data2 := []struct {
		rowIndex  int
		colIndex  int
		value     byte
		wantError bool
	}{
		{4, 5, 0, true},
		{4, 5, 10, true},
		{4, 5, 9, false},
	}
	for _, d := range data2 {
		if e := b.SetCellValue(d.rowIndex, d.colIndex, d.value); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil {
			if d.wantError {
				t.Error("expected an error")
			} else if value := b.cells[d.rowIndex][d.colIndex]; value != d.value {
				t.Errorf("Expected board[%d][%d] set to %v but got %v", d.rowIndex, d.colIndex, d.value, value)
			}
		}
	}

	if e := b.SetCellValue(4, 5, 1); e == nil {
		t.Error("expected error")
	}

	b.SetOverwriteFlag(true)
	if e := b.SetCellValue(4, 5, 1); e != nil {
		t.Errorf("unexpected error - %v", e)
	}
}

func TestClearCellValue(t *testing.T) {
	b := NewBoard()
	data := []struct {
		rowIndex  int
		colIndex  int
		wantError bool
	}{
		{-1, 0, true},
		{0, 0, false},
		{9, 0, true},
		{10, 0, true},
		{8, 0, false},
		{0, -1, true},
		{0, 9, true},
		{0, 10, true},
		{0, 8, false},
	}
	for _, d := range data {
		if e := b.ClearCellValue(d.rowIndex, d.colIndex); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil && d.wantError {
			t.Error("expected an error")
		}
	}

	b.cells[4][5] = 9
	if e := b.ClearCellValue(4, 5); e == nil {
		t.Error("error expected")
	}

	b.SetOverwriteFlag(true)
	if e := b.ClearCellValue(4, 5); e != nil {
		t.Errorf("got an unexpected error - %v", e)
	} else if b.cells[4][5] != 0 {
		t.Errorf("Expected board[4][5] set to 0 (Zero) but got %v", b.cells[4][5])
	}
}

func TestGetSquare(t *testing.T) {
	b := NewBoard()

	data := []struct {
		sqIndex   int
		wantError bool
	}{
		{-1, true},
		{0, false},
		{8, false},
		{9, true},
		{10, true},
	}
	for _, d := range data {
		if _, e := b.GetSquare(d.sqIndex); e != nil && !d.wantError {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil && d.wantError {
			t.Error("expected an error")
		}
	}

	startRow := 6
	startCol := 3
	number := byte(1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.cells[startRow+i][startCol+j] = number
			number += 1
		}
	}
	expectedSquare := [3][3]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	if sq, e := b.GetSquare(7); e != nil {
		t.Errorf("Got error: %v", e)
	} else if sq != expectedSquare {
		t.Errorf("wanted %v but got %v", expectedSquare, sq)
	}
}

func Test_has(t *testing.T) {
	board := NewBoard()
	for i := 0; i < 9; i++ {
		board.cells[0][i] = byte(i) + 1
	}
	*board.rows[0][4] = 0

	data := []struct {
		value         byte
		expectedHas   bool
		expectedIndex int
	}{
		{0, true, 4},
		{1, true, 0},
		{8, true, 7},
		{5, false, -1},
		{10, false, -2},
		{9, true, 8},
	}
	for _, d := range data {
		if has, index := board.rows[0].has(d.value); has != d.expectedHas {
			t.Errorf("expected has to be '%v' but got '%v'", d.expectedHas, has)
		} else if index != d.expectedIndex {
			t.Errorf("expected index to be '%v' but got '%v'", d.expectedIndex, index)
		}
	}
}

func TestReadWrite(t *testing.T) {
	board := NewBoard()
	startCell := 3
	startCol := 3
	number := byte(1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.SetCellValue(startCell+i, startCol+j, number)
			number++
		}
	}

	var b bytes.Buffer
	if e := board.Write(&b); e != nil {
		t.Errorf("got an error while writing - %v", e)
	}

	board2 := NewBoard()
	if e := board2.Read(&b); e != nil {
		t.Errorf("got an error while reading - %v", e)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board.cells[i][j] != board2.cells[i][j] {
				t.Errorf("cells [%d][%d] are different - %d, %d",
					i, j, board.cells[i][j], board2.cells[i][j],
				)
			}
		}
	}
}

func TestClear(t *testing.T) {
	board := NewBoard()
	startRow, startCol := 3, 3
	number := byte(1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.SetCellValue(startRow+i, startCol+j, number)
			number++
		}
	}

	if e := board.Clear(); e == nil {
		t.Error("expected error") // canOverwrite flag is 'false' by default so error expected
	}

	board.canOverwrite = true
	board.Clear()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board.cells[i][j] != 0 {
				t.Errorf("cell[%d][%d] has non-zero value - %d", i, j, board.cells[i][j])
			}
		}
	}
}

func Test_getRowColIndices(t *testing.T) {
	data := []struct {
		sqIndex          int
		elemIndex        int
		wantError        bool
		expectedRowIndex int
		expectedColIndex int
	}{
		{-1, 0, true, 0, 0},
		{9, 0, true, 0, 0},
		{0, -1, true, 0, 0},
		{0, 9, true, 0, 0},
		{0, 0, false, 0, 0},
		{0, 4, false, 1, 1},
		{0, 8, false, 2, 2},
		{4, 0, false, 3, 3},
		{4, 4, false, 4, 4},
		{4, 8, false, 5, 5},
		{8, 0, false, 6, 6},
		{8, 4, false, 7, 7},
		{8, 8, false, 8, 8},
	}
	for _, d := range data {
		if r, c, e := GetRowColIndices(d.sqIndex, d.elemIndex); d.wantError && e == nil {
			t.Error("expected error")
		} else if !d.wantError && e != nil {
			t.Errorf("got unexpected error - %v", e)
		} else if e == nil {
			if r != d.expectedRowIndex || c != d.expectedColIndex {
				t.Errorf("expected (rowIndex, colIndex) to be (%d, %d) but got (%d, %d)",
					d.expectedRowIndex, d.expectedColIndex, r, c,
				)
			}
		}
	}
}

func TestCopyFrom(t *testing.T) {
	b1, b2 := NewBoard(), NewBoard()
	number := byte(1)
	for i := 6; i < 9; i++ {
		for j := 6; j < 9; j++ {
			b1.cells[i][j] = number
			number++
		}
	}

	b2.CopyFrom(b1)
	for i := 0; i < 9; i++ {
		r1, r2 := b1.rows[i], b2.rows[i]
		for j := 0; j < 9; j++ {
			if *r1[j] != *r2[j] {
				t.Errorf("expected value of cell[%d][%d] in boards b1 and b2 to be equal", i, j)
			}
		}
	}
}

func TestGetOverwriteFlag(t *testing.T) {
	b := NewBoard()
	if b.GetOverwriteFlag() {
		t.Error("expected overwrite Flag to be set to 'false'")
	}

	b.canOverwrite = true
	if !b.GetOverwriteFlag() {
		t.Error("expected overwrite Flag to be set to 'true'")
	}
}

func TestSetOverwriteFlag(t *testing.T) {
	b := NewBoard()
	if b.SetOverwriteFlag(true); !b.canOverwrite {
		t.Error("expected overwrite flag set to 'true'")
	}
	if b.SetOverwriteFlag(false); b.canOverwrite {
		t.Error("expected overwrite flag set to 'false'")
	}
}
