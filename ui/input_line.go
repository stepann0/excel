package ui

import (
	"github.com/awesome-gocui/gocui"
)

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
	return nil
}
