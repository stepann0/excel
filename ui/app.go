package ui

import (
	"github.com/awesome-gocui/gocui"
)

type Mode int

func (m Mode) String() string {
	switch m {
	case NORMAL:
		return "NORMAL"
	case INSERT:
		return "INSERT"
	case COMMAND:
		return "COMMAND"
	default:
		return "UNKNOWN"
	}
}

const (
	NORMAL Mode = iota
	INSERT
	COMMAND
)

type App struct {
	g            *gocui.Gui
	mode         Mode
	labels       *AdressLabels
	table        *Table
	formulaInput *InputLine
	cmdInput     *InputLine
	modeLabel    *TextLabel
	i            int
}

func NewApp(
	g *gocui.Gui,
	mode Mode,
	table *Table,
	labels *AdressLabels,
	formulaInput *InputLine,
	cmdInput *InputLine,
	modeLabel *TextLabel) *App {
	return &App{
		g:            g,
		mode:         mode,
		labels:       labels,
		table:        table,
		formulaInput: formulaInput,
		cmdInput:     cmdInput,
		modeLabel:    modeLabel,
	}
}

func (app *App) Layout(g *gocui.Gui) error {
	v, err := g.SetView("main", 0, 0, 1, 1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Editable = true
		v.Editor = app
	}
	if _, err := g.SetCurrentView("main"); err != nil {
		return err
	}
	return nil
}

func (app *App) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch app.mode {
	case NORMAL:
		app.NormalMode(v, key, ch, mod)
	case INSERT:
		app.InsertMode(v, key, ch, mod)
	case COMMAND:
		app.CommandMode(v, key, ch, mod)
	}
}

func (app *App) NormalMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch == 'i':
		app.mode = INSERT
	case ch == ':':
		app.mode = COMMAND
	case ch == 'j' || key == gocui.KeyArrowDown:
		app.MoveDown()
	case ch == 'k' || key == gocui.KeyArrowUp:
		app.MoveUp()
	case ch == 'h' || key == gocui.KeyArrowLeft:
		app.MoveLeft()
	case ch == 'l' || key == gocui.KeyArrowRight:
		app.MoveRight()
	}
	// TODO: handle other keybindings...
}

func (app *App) InsertMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier)  {}
func (app *App) CommandMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {}

func (app *App) MoveCurrCell(dx, dy int) {
	// unhightlight previous
	app.Highlight(app.table.currCell().name, gocui.ColorYellow, gocui.ColorBlack)
	app.HighlightCurrIndexes(gocui.ColorDefault)

	// hightlight next
	app.table.SetCurrCell(dx, dy)
	app.Highlight(app.table.currCell().name, gocui.ColorGreen, gocui.ColorDefault)
	app.HighlightCurrIndexes(gocui.ColorYellow)

	table_view, _ := app.g.View("table")
	app.setCursorOnCurrCell(table_view)
}

func (app *App) setCursorOnCurrCell(table_view *gocui.View) {
	row, col := app.table.currentCellAddr[0], app.table.currentCellAddr[1]
	cell := app.table.data[row][col]
	table_view.SetCursor(cell.x-1, cell.y-4)
}

func (app *App) MoveDown() {
	app.MoveCurrCell(0, 1)

}

func (app *App) MoveUp() {
	app.MoveCurrCell(0, -1)
}

func (app *App) MoveLeft() {
	app.MoveCurrCell(-1, 0)
}

func (app *App) MoveRight() {
	app.MoveCurrCell(1, 0)
}

func (app *App) Highlight(name string, bg, fg gocui.Attribute) {
	app.g.Update(func(g *gocui.Gui) error {
		v, err := g.View(name)
		if err != nil {
			panic(err)
		}
		v.FgColor = fg
		v.BgColor = bg
		return nil
	})
}

func (app *App) HighlightCurrIndexes(fg gocui.Attribute) {
	l, n := app.table.currCell().decomposeAddress()
	app.Highlight("label_"+l, gocui.ColorDefault, fg)
	app.Highlight("label_"+n, gocui.ColorDefault, fg)
}

func (app *App) Update(func(*gocui.Gui) error) {

}
