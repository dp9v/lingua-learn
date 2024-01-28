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

	groupSelector := widget.NewCheckGroup(groups.GetAllGroups(), func(strings []string) {})
	startBtn := widget.NewButton("Run check", func() {
		words := groups.GetWords(groupSelector.Selected)
		if len(words) == 0 {
			dialog.NewError(errors.New("список слов пуст"), myWindow).Show()
		} else {
			window := gui.NewShowWordsWindow(myApp, words.Shuffle(10))
			window.Show()
		}
	})
	myWindow.SetContent(container.NewVBox(groupSelector, startBtn))
	myWindow.ShowAndRun()
}
