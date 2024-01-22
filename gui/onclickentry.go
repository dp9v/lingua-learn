package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type OnKeyEntry struct {
	widget.Entry
	returnKeyHandler func()
}

func (m *OnKeyEntry) TypedKey(key *fyne.KeyEvent) {
	if key.Name == "Return" {
		m.returnKeyHandler()
	} else {
		m.Entry.TypedKey(key)
	}
}

func NewOnKeyEntry(returnKeyHandler func()) *OnKeyEntry {
	res := &OnKeyEntry{
		returnKeyHandler: returnKeyHandler,
	}
	res.ExtendBaseWidget(res)
	return res
}
