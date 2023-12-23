package main

import (
	"log"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jroimartin/gocui"
)

type InputLine struct {
	name  string
	x, y  int
	w     int
	frame bool
}

func NewInputLine(name string, x, y, w int, frame bool) *InputLine {
	return &InputLine{
		name: name,
		x:    x, y: y, w: w,
		frame: frame,
	}
}

func (w *InputLine) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = w.frame
		v.Editable = true
		v.SetCursor(1, 0)
	}
	if _, err := g.SetCurrentView(w.name); err != nil {
		return err
	}
	return nil
}

type Cell string

type Table struct {
	name         string
	x, y         int
	cols, rows   int
	coloumnWidth int
	rowHeight    int
	data         [][]Cell
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

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.Cursor = true

	inp := NewInputLine("input", 1, 1, 70, true)
	table := &Table{
		name: "table",
		x:    1, y: 4,
		cols: 10, rows: 9,
		coloumnWidth: 4,
	}
	g.SetManager(inp, table)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
