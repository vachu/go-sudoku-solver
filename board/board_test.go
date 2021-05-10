package board

import "testing"

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
		{8, 0, 9, false},
		{0, -1, 9, true},
		{0, 9, 9, true},
		{0, 10, 9, true},
		{0, 8, 9, false},
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
	expectedSquare := [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if row, e := b.GetSquare(7); e != nil {
		t.Errorf("Got error: %v", e)
	} else if row != expectedSquare {
		t.Errorf("wanted %v but got %v", expectedSquare, row)
	}
}
