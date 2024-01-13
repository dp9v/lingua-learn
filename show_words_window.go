package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	currentWord     = 0
	roundWords      []Word
	showWordsWindow fyne.Window
	nextBtn         = widget.NewButton("Next", onNextBtnClick)
	closeBtn        = widget.NewButton("Close", onCloseBtnClick)
	input           = widget.NewEntry()
	label           = widget.NewLabel("")
)

func showWindow(app fyne.App, words []Word) {
	showWordsWindow = app.NewWindow("Learn words")
	showWordsWindow.Resize(fyne.Size{
		Width:  300,
		Height: 200,
	})
	roundWords = words

	showWordsWindow.SetContent(container.NewBorder(
		label, container.NewHBox(nextBtn, layout.NewSpacer(), closeBtn), nil, nil, input,
	))
	label.SetText(roundWords[0].Original)

	showWordsWindow.Show()
}

func onNextBtnClick() {
	if currentWord == len(roundWords) {
		dialog.ShowError(errors.New("Слова закончились"), showWordsWindow)
		return
	}
	label.SetText(roundWords[currentWord].Original)
	currentWord++

}

func onCloseBtnClick() {

}
