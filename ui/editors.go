package ui

import "github.com/jroimartin/gocui"

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

type TableEditor struct {
	mode Mode
}

func (te *TableEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch te.mode {
	case NORMAL:
		te.NormalMode(v, key, ch, mod)
	case INSERT:
		te.InsertMode(v, key, ch, mod)
	}
}

func (te *TableEditor) InsertMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyEsc:
		te.mode = NORMAL
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyEnter:
		v.EditNewLine()
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	}
	// TODO: handle other keybindings...
}

func (te *TableEditor) NormalMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch == 'i':
		te.mode = INSERT
	case ch == ':':
		te.mode = COMMAND
	case ch == 'j':
		v.MoveCursor(0, 1, false)
	case ch == 'k':
		v.MoveCursor(0, -1, false)
	case ch == 'h':
		v.MoveCursor(-1, 0, false)
	case ch == 'l':
		v.MoveCursor(1, 0, false)
	}
	// TODO: handle other keybindings...
}
