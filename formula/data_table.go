package data

import (
	"fmt"
	"strconv"
)

type CellType int

const (
	Null CellType = iota
	Text
	Number__
	Formula
)

type FormulaData struct {
	Expr string
	Val  any
}

type DataCell struct {
	dtype CellType
	data  any
}

func (c *DataCell) Type() CellType {
	return c.dtype
}

func (c *DataCell) Data() any {
	return c.data
}

func (c *DataCell) Put(data any, dtype CellType) {
	c.data = data
	c.dtype = dtype
}

type DataTable struct {
	cols, rows int
	data       [][]*DataCell
}

func NewTable(cols, rows int) *DataTable {
	t := &DataTable{cols, rows, [][]*DataCell{}}
	for r := 0; r < rows; r++ {
		t.data = append(t.data, make([]*DataCell, cols))
		for c := 0; c < cols; c++ {
			t.data[r][c] = &DataCell{}
		}
	}
	return t
}

func (t *DataTable) Cols() int {
	return t.cols
}

func (t *DataTable) Rows() int {
	return t.rows
}

func (t *DataTable) At(x, y int) *DataCell {
	if !(x >= 0 && x < t.cols && y >= 0 && y < t.rows) {
		return nil
	}
	return t.data[y][x]
}

func (t *DataTable) AtRef(ref string) *DataCell {
	x, y := refToInd(ref)
	return t.At(x, y)
}

func (t *DataTable) PutRef(ref string, data any, dtype CellType) {
	x, y := refToInd(ref)
	t.Put(x, y, data, dtype)
}

func (t *DataTable) Put(x, y int, data any, dtype CellType) {
	if dtype == Formula {
		// expr := data.(string)
		// p := NewParser(expr, t)
		// data = FormulaData{expr, p.Eval()}
	}
	c := t.At(x, y)
	c.Put(data, dtype)
}

func (t *DataTable) GetCol(num int) []any {
	if !(num >= 0 && num < t.cols) {
		panic("colomn out of table")
	}
	res := []any{}
	for r := 0; r < t.rows; r++ {
		res = append(res, t.At(num, r).data)
	}
	return res
}

func (t *DataTable) GetRow(num int) []any {
	if !(num >= 0 && num < t.rows) {
		panic("row out of table")
	}
	res := []any{}
	for _, c := range t.data[num] {
		res = append(res, c.data)
	}
	return res
}

func (t *DataTable) GetRange(ref1, ref2 string) []any {
	col1, row1 := refToInd(ref1)
	col2, row2 := refToInd(ref2)
	if row1 != row2 && col1 != col2 {
		panic(fmt.Errorf("range dimentions error: %s:%s", ref1, ref2))
	}
	if row1 == row2 {
		// return a row
		return t.GetRow(row1)[col1 : col2+1]
	}
	if col1 == col2 {
		// return a coloumn
		return t.GetCol(col1)[row1 : row2+1]
	}
	panic(fmt.Errorf("range dimentions error: %s:%s", ref1, ref2))
}

func refToInd(ref string) (int, int) {
	x := int(ref[0] - 65) // 'A'-65 = 0
	y, _ := strconv.Atoi(ref[1:])
	return x, y - 1
}
