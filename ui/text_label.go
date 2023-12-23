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
			x:    x, y: y,
		},
		text,
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
		v.BgColor = l.bg
		v.FgColor = l.fg
		fmt.Fprint(v, l.text)
	}
	return nil
}
