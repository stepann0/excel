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
