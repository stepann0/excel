package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

type FormulaInput struct {
	InputLine
	isInput bool
	table   *Table
}

func NewFormulaInput(name string, x, y, w int, frame bool, t *Table) *FormulaInput {
	return &FormulaInput{
		InputLine: InputLine{
			Widget: Widget{
				name:  name,
				x:     x,
				y:     y,
				w:     w,
				h:     2,
				frame: frame,
			},
		},
		table: t,
	}
}

func (i *FormulaInput) Layout(g *gocui.Gui) error {
	err := i.InputLine.Layout(g)
	if err != nil {
		return err
	}

	v, _ := g.View(i.name)
	if !i.isInput {
		v.Clear()
		fmt.Fprint(v, i.table.currCell().InputString())
	}
	return nil
}

type InputLine struct {
	Widget
}

func NewInputLine(name string, x, y, w int, frame bool) *InputLine {
	return &InputLine{
		Widget: Widget{
			name:  name,
			x:     x,
			y:     y,
			w:     w,
			h:     2,
			frame: frame,
		},
	}
}

func (i *InputLine) Layout(g *gocui.Gui) error {
	v, err := g.SetView(i.name, i.x, i.y, i.x+i.w, i.y+i.h, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Frame = i.frame
	v.BgColor = i.bg
	v.FgColor = i.fg
	v.Editable = true
	v.SetCursor(len(v.Buffer()), 0)
	return nil
}
