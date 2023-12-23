package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jroimartin/gocui"
	tb "github.com/nsf/termbox-go"
)

type TextLabel struct {
	name             string
	text             string
	x, y             int
	w                int
	BgColor, FgColor gocui.Attribute
}

func NewTextLabel(name, text string, x, y int) *TextLabel {
	return &TextLabel{
		name: name,
		text: text,
		x:    x, y: y,
	}
}

func (l *TextLabel) Layout(g *gocui.Gui) error {
	v, err := g.SetView(l.name, l.x, l.y, l.x+len(l.text)+1, l.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Editable = false
		v.BgColor = l.BgColor
		v.FgColor = l.FgColor
		fmt.Fprint(v, l.text)
	}
	return nil
}

func (l *TextLabel) SetBgColor(color gocui.Attribute) {
	l.BgColor = color
}

func (l *TextLabel) SetFgColor(color gocui.Attribute) {
	l.FgColor = color
}

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
	data         [][]Cell
}

func NewTable(name string, x, y, cols, rows int) *Table {
	return &Table{
		name: name,
		x:    x, y: y,
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

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.Cursor = true

	maxX, maxY := g.Size()
	formulaInput := NewInputLine("formulaInput", 1, 1, 70, true)
	cmdInput := NewInputLine("cmdInput", 8, maxY-2, maxX-10, true)
	table := NewTable("table", 1, 4, 13, 15)
	label := NewTextLabel("l1", "NORMAL", 0, maxY-2)
	label.SetBgColor(gocui.Attribute(tb.ColorBlue))

	g.SetManager(formulaInput, cmdInput, table, label)

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
