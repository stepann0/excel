package main

import (
	"log"

	"github.com/jroimartin/gocui"
	tb "github.com/nsf/termbox-go"
	"github.com/stepann0/tercel/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.Cursor = true

	maxX, maxY := g.Size()
	formulaInput := ui.NewInputLine("formulaInput", 1, 1, 70, true)
	cmdInput := ui.NewInputLine("cmdInput", 8, maxY-2, maxX-10, true)
	table := ui.NewTable("table", 1, 4, 13, 15)
	label := ui.NewTextLabel("l1", "NORMAL", 0, maxY-2)
	label.SetBgColor(gocui.Attribute(tb.ColorBlue))

	g.SetManager(formulaInput, cmdInput, table, label)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
