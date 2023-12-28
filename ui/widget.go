package ui

import (
	"github.com/awesome-gocui/gocui"
)

type Widget struct {
	name   string
	x, y   int
	w, h   int
	frame  bool
	bg, fg gocui.Attribute
}

func (w *Widget) SetBgColor(color gocui.Attribute) {
	w.bg = color
}

func (w *Widget) SetFgColor(color gocui.Attribute) {
	w.fg = color
}

func (w *Widget) SetFrame(b bool) {
	w.frame = b
}

func (w *Widget) BaseLayout(g *gocui.Gui) (*gocui.View, error) {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		v.Frame = w.frame
		v.BgColor = w.bg
		v.FgColor = w.fg
	}
	return v, nil
}
