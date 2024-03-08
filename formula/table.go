package formula

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	V "github.com/stepann0/excel/value"
)

type CellType int

const (
	Empty      CellType = iota // Just empty cell; nil
	ConstValue                 // Contains inserted number, text, bool or other
	Formula                    // Starts with '='
)

type DataCell struct {
	Type CellType
	Data V.Value
	Text string // Inserted text (e.g "=sum(A1:A10)", "2.1828", "TRUE")
}

func (c *DataCell) Put(text string, t *DataTable) {
	if len(text) == 0 {
		return
	}
	c.Text = text
	if text[0] == '=' {
		c.Type = Formula
		c.CalcFormula(t)
		return
	}

	c.Type = ConstValue
	if text == TRUE_LITERAL {
		c.Data = V.Boolean(true)
		return
	}
	if text == FALSE_LITERAL {
		c.Data = V.Boolean(false)
		return
	}
	var int_n int64
	var float_n float64
	var err error

	int_n, err = strconv.ParseInt(text, 10, 64)
	if err == nil {
		c.Data = V.Int(int_n)
		return
	}

	float_n, err = strconv.ParseFloat(text, 64)
	if err == nil {
		c.Data = V.Float(float_n)
		return
	}
	c.Data = V.String(text)
}

func (c *DataCell) CalcFormula(t *DataTable) {
	var expr V.Value
	defer func() {
		if r := recover(); r != nil {
			expr = V.Error{Msg: fmt.Errorf("%v", r)}
		}
	}()
	expr = NewParser(c.Text[1:], t).Parse().Eval()
	c.Data = expr
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

func refToInd(ref string) (int, int) {
	x := int(ref[0] - 65) // 'A'-65 = 0
	y, _ := strconv.Atoi(ref[1:])
	return x, y - 1
}

func (t *DataTable) GetRange(ref1, ref2 string) []V.Value {
	col1, row1 := refToInd(ref1)
	col2, row2 := refToInd(ref2)
	if row1 != row2 && col1 != col2 {
		panic(fmt.Errorf("range dimentions error: %s:%s", ref1, ref2))
	}
	if row1 == row2 {
		return t.GetRow(row1)[col1 : col2+1]
	}
	if col1 == col2 {
		return t.GetCol(col1)[row1 : row2+1]
	}
	return nil
}

func (t *DataTable) GetRow(i int) []V.Value {
	if !(i >= 0 && i < t.rows) {
		panic("row index out of table")
	}
	area := make([]V.Value, len(t.data[i]))
	for j := range t.data[i] {
		area[j] = t.data[i][j].Data
	}
	return area
}

func (t *DataTable) GetCol(i int) []V.Value {
	if !(i >= 0 && i < t.cols) {
		panic("col index out of table")
	}
	area := make([]V.Value, len(t.data[i]))
	for j := range t.data {
		area[j] = t.data[j][i].Data
	}
	return area
}

func (t *DataTable) LoadCSV(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'
	records, err := csvReader.ReadAll()
	if err != nil {
		return
	}
	if len(records) == 0 {
		return
	}
	rows, cols := len(records), len(records[0])
	if x, y := t.cols, t.rows; rows > y || cols > x {
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			t.At(j, i).Put(records[i][j], t)
		}
	}
}
