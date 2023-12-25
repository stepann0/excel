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
	formulaInput := ui.NewInputLine("formulaInput", 2, 1, 70, true)
	cmdInput := ui.NewInputLine("cmdInput", 7, maxY-2, maxX-10, false)
	cmdInput.SetBgColor(gocui.Attribute(tb.ColorLightGray))

	table := ui.NewTable("table", 2, 5, 13, 9)
	labelMode := ui.NewTextLabel("l1", "NORMAL", 0, maxY-2)
	labelMode.SetBgColor(gocui.Attribute(tb.ColorBlue))

	adressLabels := ui.NewAdressLabels(table)

	g.SetManager(formulaInput, cmdInput, table, labelMode, adressLabels)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, insertFormula); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func insertFormula(g *gocui.Gui, v *gocui.View) error {
	g.SetCurrentView("formulaInput")
	return nil
}
