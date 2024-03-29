package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

type TextLabel struct {
	Widget
	text      string
	writeFunc func(*gocui.View)
}

func NewTextLabel(name, text string, x, y int) *TextLabel {
	return &TextLabel{
		Widget{
			name: name,
			x:    x,
			y:    y,
			w:    len(text) + 1,
			h:    2,
		},
		text,
		func(v *gocui.View) {
			v.Clear()
			fmt.Fprint(v, text)
		},
	}
}

func (l *TextLabel) Layout(g *gocui.Gui) error {
	v, err := g.SetView(l.name, l.x, l.y, l.x+len(l.text)+1, l.y+l.h, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Frame = l.frame
	v.BgColor = l.bg
	v.FgColor = l.fg
	v.Clear()
	fmt.Fprint(v, l.text)
	return nil
}

type AdressLabels struct {
	NumLabels    []*TextLabel
	LetterLabels []*TextLabel
	table        *Table
}

func NewAdressLabels(t *Table) *AdressLabels {
	AdLab := AdressLabels{
		NumLabels:    make([]*TextLabel, t.rows),
		LetterLabels: make([]*TextLabel, t.cols),
		table:        t,
	}

	for r := 0; r < t.rows; r++ {
		text := fmt.Sprint(r + 1)
		AdLab.NumLabels[r] = NewTextLabel("label_"+text, text, 0, r*2+t.y)
	}
	for c := 0; c < t.cols; c++ {
		text := fmt.Sprintf("%c", c%26+65)
		AdLab.LetterLabels[c] = NewTextLabel("label_"+text, text, c*(t.coloumnWidth+1)+t.x+3, t.y-2)
	}
	return &AdLab
}

func (a *AdressLabels) Layout(g *gocui.Gui) error {
	x, y := a.table.currentCellAddr[0], a.table.currentCellAddr[1]
	for i, l := range a.LetterLabels {
		l.fg = gocui.ColorDefault
		if i == x {
			l.fg = gocui.ColorYellow | gocui.AttrBold
		}

		err := l.Layout(g)
		if err != nil {
			return err
		}
	}
	for i, l := range a.NumLabels {
		l.fg = gocui.ColorDefault
		if i == y {
			l.fg = gocui.ColorYellow | gocui.AttrBold
		}

		err := l.Layout(g)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AdressLabels) GetLabelsByIndex(row, col int) (*TextLabel, *TextLabel) {
	if row < a.table.rows && col < a.table.cols {
		return a.LetterLabels[col], a.NumLabels[row]
	}
	return nil, nil
}

func (a *AdressLabels) GetLabelsByAddress(address string) (*TextLabel, *TextLabel) {
	col, row := address[0]-65, address[1]-49 // because 'A'-65 = 0, '1'-49 = 0
	return a.GetLabelsByIndex(int(row), int(col))
}
