package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
	"strings"
)

type ShowWordsWindow struct {
	currentWord      int
	roundWords       common.Words
	window           fyne.Window
	nextBtn          *widget.Button
	closeBtn         *widget.Button
	input            *OnKeyEntry
	translationLabel *widget.Label
	correctWordLabel *widget.Label
}

func NewShowWordsWindow(app fyne.App, words common.Words) ShowWordsWindow {
	window := ShowWordsWindow{
		currentWord:      0,
		roundWords:       words,
		window:           app.NewWindow("                                              "),
		translationLabel: widget.NewLabel(""),
		correctWordLabel: widget.NewLabel(""),
	}
	window.input = NewOnKeyEntry(window.onNextBtnClick)
	window.nextBtn = widget.NewButton("Next", window.onNextBtnClick)
	window.closeBtn = widget.NewButton("Close", window.onCloseBtnClick)
	return window
}

func (w *ShowWordsWindow) onNextBtnClick() {
	defer w.focusInput()
	if strings.ToUpper(w.input.Text) != strings.ToUpper(w.roundWords[w.currentWord].Original) {
		w.correctWordLabel.SetText(w.roundWords[w.currentWord].Original)
		return
	}
	if w.currentWord+1 == len(w.roundWords) {
		dialog.ShowError(errors.New("Слова закончились"), w.window)
		return
	}

	w.nextWord()
}

func (w *ShowWordsWindow) nextWord() {
	if len(w.correctWordLabel.Text) != 0 {
		w.roundWords = append(w.roundWords, w.roundWords[w.currentWord])
	}
	w.currentWord++
	w.correctWordLabel.SetText("")
	w.translationLabel.SetText(w.roundWords[w.currentWord].Translation)
	w.input.SetText("")
}

func (w *ShowWordsWindow) focusInput() {
	w.window.Canvas().Focus(w.input)
}

func (w *ShowWordsWindow) onCloseBtnClick() {
	w.window.Close()
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
	w.window.Show()
}
