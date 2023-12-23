package ui

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jroimartin/gocui"
)

type Cell string

type Table struct {
	Widget
	cols, rows   int
	coloumnWidth int
	data         [][]Cell
}

func NewTable(name string, x, y, cols, rows int) *Table {
	return &Table{
		Widget: Widget{
			name: name,
			x:    x, y: y,
		},
		cols: cols, rows: rows,
		coloumnWidth: 6,
	}
}

func (t *Table) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(t.name, t.x-1, t.y-1, maxX-3, maxY-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Frame = false
	v.Clear()

	t.drawGrid(v)
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
	T.SetStyle(my_style)
	T.Render()
}
