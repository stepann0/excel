package data

type CellType int

const (
	Null CellType = iota
	Text
	Number
	Formula
)

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
	return t.data[y][x]
}

func (t *DataTable) Put(x, y int, data any, dtype CellType) {
	c := t.At(x, y)
	c.Put(data, dtype)
}
