package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var data = []Word{{
	Original:    "Dekuji",
	Translation: "Спасибо",
}, {
	Original:    "Ahoi",
	Translation: "Привет",
}, {
	Original:    "Chao",
	Translation: "Пока",
}}

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
			dialog.NewError(errors.New("Список ошибок пуст"), myWindow).Show()
		} else {
			showWindow(myApp, data)
		}
	})
	myWindow.SetContent(container.NewCenter(startBtn))
	myWindow.ShowAndRun()
}
