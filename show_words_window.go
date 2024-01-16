package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type ShowWordsWindow struct {
	currentWord      int
	roundWords       []Word
	window           fyne.Window
	nextBtn          *widget.Button
	closeBtn         *widget.Button
	input            *widget.Entry
	translationLabel *widget.Label
	correctWordLabel *widget.Label
}

func InitShowWordsWindow(app fyne.App, words []Word) ShowWordsWindow {
	window := ShowWordsWindow{
		currentWord:      0,
		roundWords:       words,
		window:           app.NewWindow("Learn words"),
		input:            widget.NewEntry(),
		translationLabel: widget.NewLabel(""),
		correctWordLabel: widget.NewLabel(""),
	}
	window.nextBtn = widget.NewButton("Next", window.onNextBtnClick)
	window.closeBtn = widget.NewButton("Close", window.onCloseBtnClick)
	return window
}

func (w *ShowWordsWindow) onNextBtnClick() {
	if strings.ToUpper(w.input.Text) != strings.ToUpper(w.roundWords[w.currentWord].Original) {
		w.correctWordLabel.SetText(w.roundWords[w.currentWord].Original)
		return
	}
	if w.currentWord+1 == len(w.roundWords) {
		dialog.ShowError(errors.New("Слова закончились"), w.window)
		return
	}

	w.currentWord++
	w.correctWordLabel.SetText("")
	w.input.SetText("")
	w.translationLabel.SetText(w.roundWords[w.currentWord].Translation)
}

func (w *ShowWordsWindow) onCloseBtnClick() {

}

func (w *ShowWordsWindow) Show() {
	w.window.Resize(fyne.Size{
		Width:  300,
		Height: 200,
	})

	w.window.SetContent(container.NewBorder(
		w.translationLabel,
		container.NewHBox(w.nextBtn, layout.NewSpacer(), w.closeBtn),
		nil,
		nil,
		container.NewVBox(w.input, w.correctWordLabel),
	))

	w.translationLabel.SetText(w.roundWords[w.currentWord].Translation)
	w.closeBtn.Disable()
	w.window.Show()
}
