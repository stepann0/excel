package ui

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jroimartin/gocui"
)

type CellType int

const (
	Text CellType = iota
	Number
	Formula
)

type Cell struct {
	Widget
	dtype  CellType
	data   any
	adress string
}

func NewCell(adress string, x, y, w, h int) *Cell {
	return &Cell{
		Widget: Widget{
			name: "cell_" + adress,
			x:    x,
			y:    y,
			w:    w,
			h:    h,
			bg:   gocui.ColorYellow,
		},
		adress: adress,
	}
}

func (c *Cell) Layout(g *gocui.Gui) error {
	v, err := c.Widget.BaseLayout(g)
	if err != nil {
		return err
	}
	v.Frame = false
	v.Clear()
	fmt.Fprint(v, c.name)
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
	return "==="
}

type Table struct {
	Widget
	cols, rows      int
	coloumnWidth    int
	data            [][]*Cell
	currentCellAddr []int // address is []int{x, y}
}

func NewTable(name string, x, y, cols, rows int) *Table {
	t := &Table{
		Widget: Widget{
			name: name,
			x:    x, y: y,
		},
		cols: cols, rows: rows,
		coloumnWidth: 7,
	}
	t.createCells()
	t.currentCellAddr = []int{0, 0}
	return t
}

func (t *Table) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	table_view, err := g.SetView(t.name, t.x-1, t.y-1, maxX-3, maxY-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		table_view.Frame = false
	}
	table_view.Clear()
	t.drawGrid(table_view)

	for _, r := range t.data {
		for _, c := range r {
			if c == nil {
				continue
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
		t.data = append(t.data, make([]*Cell, t.cols))
		for c := 0; c < t.cols; c++ {
			x, y := c*(t.coloumnWidth+1)+t.x, r*2+t.y
			adress := fmt.Sprintf("%c%d", c%26+65, r+1)
			t.data[r][c] = NewCell(adress, x, y, t.coloumnWidth+1, 2)
		}
	}
}

func (t *Table) isCurrCell(row, col int) bool {
	return row == t.currentCellAddr[0] && col == t.currentCellAddr[1]
}

func (t *Table) currCell() *Cell {
	x, y := t.currentCellAddr[0], t.currentCellAddr[1]
	return t.data[y][x]
}

func (t *Table) SetCurrCell(dx, dy int) {
	x, y := t.currentCellAddr[0], t.currentCellAddr[1]
	if x+dx >= 0 && x+dx < t.cols && y+dy >= 0 && y+dy < t.rows {
		t.currentCellAddr[0] += dx
		t.currentCellAddr[1] += dy
	}
}
