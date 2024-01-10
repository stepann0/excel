package ui

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/stepann0/tercel/data"
)

type Cell struct {
	Widget
	cell   *data.DataCell
	adress string
}

func NewCell(adress string, x, y, w, h int, cell *data.DataCell) *Cell {
	return &Cell{
		Widget: Widget{
			name: "cell_" + adress,
			x:    x,
			y:    y,
			w:    w,
			h:    h,
			bg:   gocui.ColorDefault,
		},
		cell:   cell,
		adress: adress,
	}
}

func (c *Cell) Layout(g *gocui.Gui) error {
	v, err := g.SetView(c.name, c.x, c.y, c.x+c.w, c.y+c.h, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}
	switch c.Type() {
	case data.Number:
		c.fg = gocui.ColorBlue
	case data.Text:
		c.fg = gocui.ColorGreen
	case data.Formula:
		c.fg = gocui.ColorMagenta
	}

	v.BgColor = c.bg
	v.FgColor = c.fg
	v.Clear()
	fmt.Fprint(v, c)
	return nil
}

func (c *Cell) decomposeAddress() (string, string) {
	i := 0
	for i < len(c.adress) && c.adress[i] >= 'A' && c.adress[i] <= 'Z' {
		i++
	}
	return c.adress[:i], c.adress[i:]
}

func (c Cell) String() string {
	if c.Data() == nil {
		return ""
	}
	text := fmt.Sprint(c.Data())
	if c.Type() == data.Formula {
		result := c.Data().(data.FormulaData).Val
		switch result.Type() {
		case data.ResError:
			text = fmt.Sprintf("%5s", "###")
		case data.ResNil:
			text = ""
		case data.ResNumber:
			text = fmt.Sprint(result.Value().(float64))
		case data.ResRange:
			text = "#rng"
		}
	}
	if len(text) >= c.w {
		text = text[:c.w-2] + "…"
	}
	return text
}

func (c Cell) InputString() string {
	if c.Data() == nil {
		return ""
	}
	text := fmt.Sprint(c.Data())
	if c.Type() == data.Formula {
		text = c.Data().(data.FormulaData).Expr
	}
	return text
}

func (c *Cell) Type() data.CellType {
	return c.cell.Type()
}

func (c *Cell) Data() any {
	return c.cell.Data()
}

type Table struct {
	Widget
	DataTable       *data.DataTable
	cols, rows      int
	cells           [][]*Cell
	coloumnWidth    int
	currentCellAddr []int // address is []int{x, y}
}

func NewTable(name string, x, y int, data *data.DataTable) *Table {
	t := &Table{
		Widget: Widget{
			name: name,
			x:    x, y: y,
		},
		DataTable: data,
		cols:      data.Cols(), rows: data.Rows(),
		coloumnWidth: 7,
	}
	t.createCells()
	t.currentCellAddr = []int{0, 0}
	return t
}

func (t *Table) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	table_view, err := g.SetView(t.name, t.x-1, t.y-1, maxX-3, maxY-2, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		table_view.Frame = false
	}
	table_view.Clear()
	t.drawGrid(table_view)

	for i, r := range t.cells {
		for j, c := range r {
			if c == nil {
				continue
			}
			c.bg = gocui.ColorDefault
			c.fg = gocui.ColorDefault
			if t.isCurrCell(j, i) {
				c.bg = gocui.GetColor("gray")
			}
			if err := c.Layout(g); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Table) drawGrid(view *gocui.View) {
	T := table.NewWriter()
	T.SetOutputMirror(view)
	row := make([]any, t.cols)
	for c := 0; c < t.cols; c++ {
		row[c] = strings.Repeat("*", t.coloumnWidth)
	}

	for r := 0; r < t.rows; r++ {
		T.AppendRow(row)
	}
	my_style := table.StyleLight
	my_style.Options.SeparateRows = true
	my_style.Box.PaddingLeft = ""
	my_style.Box.PaddingRight = ""
	T.SetStyle(my_style)
	T.Render()
}

func (t *Table) createCells() {
	for r := 0; r < t.rows; r++ {
		t.cells = append(t.cells, make([]*Cell, t.cols))
		for c := 0; c < t.cols; c++ {
			x, y := c*(t.coloumnWidth+1)+t.x, r*2+t.y
			adress := fmt.Sprintf("%c%d", c%26+65, r+1)
			t.cells[r][c] = NewCell(adress, x, y, t.coloumnWidth+1, 2, t.DataTable.At(c, r))
		}
	}
}

func (t *Table) isCurrCell(x, y int) bool {
	return x == t.currentCellAddr[0] && y == t.currentCellAddr[1]
}

func (t *Table) currCell() *Cell {
	x, y := t.currentCellAddr[0], t.currentCellAddr[1]
	return t.cells[y][x]
}

func (t *Table) SetCurrCell(dx, dy int) {
	x, y := t.currentCellAddr[0], t.currentCellAddr[1]
	if x+dx >= 0 && x+dx < t.cols && y+dy >= 0 && y+dy < t.rows {
		t.currentCellAddr[0] += dx
		t.currentCellAddr[1] += dy
	}
}
