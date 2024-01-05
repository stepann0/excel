package main

import (
	"fmt"
	"os"

	"github.com/awesome-gocui/gocui"
	"github.com/stepann0/tercel/data"
	"github.com/stepann0/tercel/formula"
	"github.com/stepann0/tercel/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	maxX, maxY := g.Size()
	formulaInput := ui.NewInputLine("formulaInput", 2, 1, 70, true)
	cmdInput := ui.NewInputLine("cmdInput", 7, maxY-2, maxX-10, false)
	cmdInput.SetBgColor(gocui.ColorBlack)

	data_table := data.NewTable(13, 10)
	data.LoadCSV(data_table, os.Args[1])
	table := ui.NewTable("table", 3, 5, data_table)

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

func _main() {
	exp := []string{
		"max((34+1), -(pow(2, 5)), 3.1415)",
		"max(pow(3, -(3)), 0)",
		"max(min(1, 10), 2/3*3.1415)",
		"sin(max(min(1, 10), 2/3*3.1415))",
		"-cos(max(min(1, 10), 2/3*3.1415))",
		"avg(-3, -2, -1, 0, 1, 0)",
		"max(4/24, 5/25/2*4, 6/26, 7/27, 8/28)",
		"-(-(-(-4)))",
		"-A1",
		"(B2:C7)",
		"B6:H5",
		"max((CC8), ((B2:B2)))",
	}
	for _, e := range exp {
		fmt.Println(formula.Parse(e))
	}
}
