package ui

import "github.com/jroimartin/gocui"

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
	v, err := i.Widget.BaseLayout(g)
	if err != nil {
		return err
	}
	v.Editable = true
	v.SetCursor(1, 0)
	return nil
}
