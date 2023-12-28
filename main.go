package main

import (
	"github.com/jroimartin/gocui"
	"github.com/stepann0/tercel/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.Cursor = true

	maxX, maxY := g.Size()
	formulaInput := ui.NewInputLine("formulaInput", 2, 1, 70, true)
	cmdInput := ui.NewInputLine("cmdInput", 7, maxY-2, maxX-10, false)
	cmdInput.SetBgColor(gocui.ColorBlack)
	table := ui.NewTable("table", 2, 5, 13, 9)

	adressLabels := ui.NewAdressLabels(table)
	modeLabel := ui.NewTextLabel("modeLabel", "NORMAL", 0, maxY-2)
	modeLabel.SetBgColor(gocui.ColorBlue)

	app := ui.NewApp(g, ui.NORMAL, table, adressLabels, formulaInput, cmdInput, modeLabel)
	g.SetManager(formulaInput, cmdInput, modeLabel, table, adressLabels, app)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
