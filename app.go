package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"math/rand"
)

type Word struct {
	Original    string
	Translation string
}

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
			window := InitShowWordsWindow(myApp, getRandomWords(5))
			window.Show()
			//showWindow(myApp, getRandomWords(5))
		}
	})
	myWindow.SetContent(container.NewCenter(startBtn))
	myWindow.ShowAndRun()
}

func getRandomWords(count int) []Word {
	result := make([]Word, count)
	for i := 0; i < count; i++ {
		result[i] = data[rand.Intn(len(data))]
	}
	return result
}
