package main

import (
	"flag"
	"fyne.io/fyne/v2/app"
	"learn_words/gui"
)

func main() {
	convertFlag := flag.Bool("convert", false, "Set this flag to perform conversion")
	flag.Parse()

	if *convertFlag {
		Convert()
	} else {
		gui.NewApplication(
			app.NewWithID("com.dp9v.lingua-learn"),
		)
	}

}
