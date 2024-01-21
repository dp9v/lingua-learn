package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/gui"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Learn words")
	myWindow.Resize(fyne.Size{
		Width:  300,
		Height: 300,
	})

	startBtn := widget.NewButton("Run check", func() {
		if len(data) == 0 {
			dialog.NewError(errors.New("список слов пуст"), myWindow).Show()
		} else {
			window := gui.NewShowWordsWindow(myApp, data.Shuffle())
			window.Show()
			//showWindow(myApp, getRandomWords(5))
		}
	})
	myWindow.SetContent(container.NewCenter(startBtn))
	myWindow.ShowAndRun()
}
