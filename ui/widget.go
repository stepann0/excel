package ui

import "github.com/jroimartin/gocui"

type Widget struct {
	name   string
	x, y   int
	w, h   int
	bg, fg gocui.Attribute
}

func (w *Widget) SetBgColor(color gocui.Attribute) {
	w.bg = color
}

func (w *Widget) SetFgColor(color gocui.Attribute) {
	w.fg = color
}
