package value

import (
	"fmt"
	"strconv"
)

type CellType int

const (
	Empty      CellType = iota // Just empty cell; nil
	ConstValue                 // Contains inserted number, text, bool or other
	Formula                    // Starts with '='
)

type FormulaData struct {
	Expr string
	Val  any
}

type DataCell struct {
	Type CellType
	Data Value
	Text string // Inserted text (e.g "=sum(A1:A10)", "2.1828", "TRUE")
}

func (c *DataCell) Put(text string) {
	if len(text) == 0 {
		return
	}
	c.Text = text
	if text[0] == '=' {
		c.Type = Formula
		return
	}

	c.Type = ConstValue
	var int_n int64
	var float_n float64
	var err error

	int_n, err = strconv.ParseInt(text, 10, 64)
	if err == nil {
		c.Data = Number[int]{int(int_n)}
		return
	}

	float_n, err = strconv.ParseFloat(text, 64)
	if err == nil {
		c.Data = Number[float64]{float_n}
		return
	}
	c.Data = String{text}
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

func (t *DataTable) GetRange(ref1, ref2 string) []any {
	col1, row1 := refToInd(ref1)
	col2, row2 := refToInd(ref2)
	if row1 != row2 && col1 != col2 {
		panic(fmt.Errorf("range dimentions error: %s:%s", ref1, ref2))
	}
	if row1 == row2 {
		// return a row
		// return t.GetRow(row1)[col1 : col2+1]
	}
	if col1 == col2 {
		// return a coloumn
		// return t.GetCol(col1)[row1 : row2+1]
	}
	panic(fmt.Errorf("range dimentions error: %s:%s", ref1, ref2))
}

func refToInd(ref string) (int, int) {
	x := int(ref[0] - 65) // 'A'-65 = 0
	y, _ := strconv.Atoi(ref[1:])
	return x, y - 1
}