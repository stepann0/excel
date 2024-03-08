package main

import (
	"github.com/awesome-gocui/gocui"
	"github.com/stepann0/excel/formula"
	"github.com/stepann0/excel/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	maxX, maxY := g.Size()
	data_table := formula.NewTable(13, 10)
	table := ui.NewTable("table", 3, 5, data_table)
	table.DataTable.LoadCSV("/home/stepaFedora/Документы/small.csv")

	formulaInput := ui.NewFormulaInput("formulaInput", 3, 1, 104, true, table)
	cmdInput := ui.NewInputLine("cmdInput", 7, maxY-2, maxX-10, false)
	cmdInput.SetBgColor(gocui.ColorBlack)
	adressLabels := ui.NewAdressLabels(table)
	modeLabel := ui.NewTextLabel("modeLabel", "      ", 0, maxY-2)

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
