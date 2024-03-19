package main

import (
	"fyne.io/fyne/v2/app"
	"learn_words/gui"
)

func main() {
	gui.NewApplication(
		app.NewWithID("com.dp9v.lingua-learn"),
	)
}
