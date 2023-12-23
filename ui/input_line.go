package ui

import "github.com/jroimartin/gocui"

type InputLine struct {
	Widget
	frame bool
}

func NewInputLine(name string, x, y, w int, frame bool) *InputLine {
	return &InputLine{
		Widget: Widget{
			name: name,
			x:    x,
			y:    y,
			w:    w,
		},
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
