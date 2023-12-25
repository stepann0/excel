package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type TextLabel struct {
	Widget
	text string
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
	}
}

func (l *TextLabel) Layout(g *gocui.Gui) error {
	v, err := l.Widget.BaseLayout(g)
	if err != nil {
		return err
	}

	fmt.Fprint(v, l.text)
	return nil
}

type AdressLabels struct {
	NumLabels    []*TextLabel
	LetterLabels []*TextLabel
	activeAdress string
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
	for _, l := range a.LetterLabels {
		err := l.Layout(g)
		if err != nil {
			return err
		}
	}
	for _, l := range a.NumLabels {
		err := l.Layout(g)
		if err != nil {
			return err
		}
	}
	return nil
}
