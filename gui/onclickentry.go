package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type OnClickEntry struct {
	widget.Entry
	altReturnHandler func()
}

func (m *OnClickEntry) TypedShortcut(s fyne.Shortcut) {
	altReturn := &desktop.CustomShortcut{
		KeyName:  fyne.KeyReturn,
		Modifier: fyne.KeyModifierAlt,
	}
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s)
		return
	}
	if s.ShortcutName() == altReturn.ShortcutName() {
		m.altReturnHandler()
	}
}

func NewOnClickEntry(altEnterHandler func()) *OnClickEntry {
	res := &OnClickEntry{
		altReturnHandler: altEnterHandler,
	}
	res.ExtendBaseWidget(res)
	return res
}
