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

var (
	currentWord      = 0
	roundWords       []Word
	showWordsWindow  fyne.Window
	nextBtn          = widget.NewButton("Start", onNextBtnClick)
	closeBtn         = widget.NewButton("Close", onCloseBtnClick)
	input            = widget.NewEntry()
	originalLabel    = widget.NewLabel("")
	correctWordLabel = widget.NewLabel("")
)

func showWindow(app fyne.App, words []Word) {
	showWordsWindow = app.NewWindow("Learn words")
	showWordsWindow.Resize(fyne.Size{
		Width:  300,
		Height: 200,
	})
	roundWords = words

	showWordsWindow.SetContent(container.NewBorder(
		originalLabel,
		container.NewHBox(nextBtn, layout.NewSpacer(), closeBtn),
		nil,
		nil,
		container.NewVBox(input, correctWordLabel),
	))

	originalLabel.SetText(roundWords[currentWord].Translation)
	closeBtn.Disable()
	showWordsWindow.Show()
}

func onNextBtnClick() {
	if strings.ToUpper(input.Text) != strings.ToUpper(roundWords[currentWord].Original) {
		correctWordLabel.SetText(roundWords[currentWord].Original)
		return
	}
	if currentWord+1 == len(roundWords) {
		dialog.ShowError(errors.New("Слова закончились"), showWordsWindow)
		return
	}

	currentWord++
	correctWordLabel.SetText("")
	input.SetText("")
	originalLabel.SetText(roundWords[currentWord].Translation)
}

func onCloseBtnClick() {

}
