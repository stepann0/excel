package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/stepann0/tercel/data"
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
}

func NewApp(
	g *gocui.Gui,
	mode Mode,
	table *Table,
	labels *AdressLabels,
	formulaInput *InputLine,
	cmdInput *InputLine,
	modeLabel *TextLabel) *App {
	app := &App{
		g:            g,
		mode:         mode,
		labels:       labels,
		table:        table,
		formulaInput: formulaInput,
		cmdInput:     cmdInput,
		modeLabel:    modeLabel,
	}
	app.SetModeLabel()
	return app
}

func (app *App) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("main", 0, 0, maxX-1, maxY-1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Visible = false
		v.Frame = false
		v.Editable = true
		v.Editor = app
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
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
	app.g.Cursor = false
	app.SetModeLabel()
	if _, err := app.g.SetCurrentView("main"); err != nil {
		panic(err)
	}
	switch {
	case ch == 'i':
		app.mode = INSERT
		app.InsertMode(v, key, ch, mod)
	case ch == ':':
		app.mode = COMMAND
		app.CommandMode(v, key, ch, mod)
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

func (app *App) InsertMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	app.g.Cursor = true
	app.SetModeLabel()
	line, err := app.g.SetCurrentView("formulaInput")
	if err != nil {
		panic(err)
	}
	line.Editor = gocui.EditorFunc(func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
		switch key {
		case gocui.KeySpace:
			line.EditWrite(' ')
		case gocui.KeyBackspace, gocui.KeyBackspace2:
			line.EditDelete(true)
		case gocui.KeyDelete:
			line.EditDelete(false)
		case gocui.KeyInsert:
			line.Overwrite = !line.Overwrite
		case gocui.KeyArrowDown:
			line.MoveCursor(0, 1)
		case gocui.KeyArrowUp:
			line.MoveCursor(0, -1)
		case gocui.KeyArrowLeft:
			line.MoveCursor(-1, 0)
		case gocui.KeyArrowRight:
			line.MoveCursor(1, 0)
		case gocui.KeyTab:
			line.EditWrite('\t')
		case gocui.KeyEnter:
			c := app.table.currCell()
			data, dtype := data.ConvertType(line.Buffer())
			app.table.DataTable.PutRef(c.adress, data, dtype)
		case gocui.KeyEsc:
			app.mode = NORMAL
			app.NormalMode(v, key, ch, mod)
		default:
			line.EditWrite(ch)
		}
	})
}

func (app *App) CommandMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	app.g.Cursor = true
	app.SetModeLabel()
	cmdInput, err := app.g.SetCurrentView("cmdInput")
	if err != nil {
		panic(err)
	}
	cmdInput.Editor = gocui.EditorFunc(func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
		switch key {
		case gocui.KeySpace:
			cmdInput.EditWrite(' ')
		case gocui.KeyBackspace, gocui.KeyBackspace2:
			cmdInput.EditDelete(true)
		case gocui.KeyDelete:
			cmdInput.EditDelete(false)
		case gocui.KeyInsert:
			cmdInput.Overwrite = !cmdInput.Overwrite
		case gocui.KeyArrowDown:
			cmdInput.MoveCursor(0, 1)
		case gocui.KeyArrowUp:
			cmdInput.MoveCursor(0, -1)
		case gocui.KeyArrowLeft:
			cmdInput.MoveCursor(-1, 0)
		case gocui.KeyArrowRight:
			cmdInput.MoveCursor(1, 0)
		case gocui.KeyTab:
			cmdInput.EditWrite('\t')
		case gocui.KeyEnter:
			// c := app.table.currCell()
			// c.data = cmdInput.Buffer()
		case gocui.KeyEsc:
			app.mode = NORMAL
			app.NormalMode(v, key, ch, mod)
		default:
			cmdInput.EditWrite(ch)
		}
	})
}

func (app *App) MoveCurrCell(dx, dy int) {
	app.table.SetCurrCell(dx, dy)
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
	v, err := app.g.View(name)
	if err != nil {
		panic(err)
	}
	v.FgColor = fg
	v.BgColor = bg
}

func (app *App) HighlightCurrIndexes(fg gocui.Attribute) {
	letter, num := app.table.currentCellAddr[0], app.table.currentCellAddr[1]
	app.labels.LetterLabels[letter].fg = fg
	app.labels.NumLabels[num].fg = fg
}

func (app *App) SetModeLabel() {
	app.modeLabel.text = app.mode.String()
	bg := gocui.ColorDefault
	fg := gocui.ColorDefault

	switch app.mode {
	case NORMAL:
		bg = gocui.ColorBlue
		fg = gocui.ColorBlack
	case INSERT:
		bg = gocui.ColorGreen
		fg = gocui.ColorBlack
	case COMMAND:
		bg = gocui.ColorMagenta
		fg = gocui.ColorWhite
	}
	app.modeLabel.bg = bg
	app.modeLabel.fg = fg | gocui.AttrBold
}
