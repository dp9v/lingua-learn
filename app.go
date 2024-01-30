package main

import "learn_words/gui"

func main() {
	app := gui.NewGuiApplication()
	println(app)
	//myApp := app.New()
	//myWindow := myApp.NewWindow("Learn words")
	//myWindow.Resize(fyne.Size{
	//	Width:  300,
	//	Height: 300,
	//})
	//
	//groupSelector := widget.NewCheckGroup(common.Groups.GetAllGroups(), func(strings []string) {})
	//startBtn := widget.NewButton("Run check", func() {
	//	words := common.Groups.GetWords(groupSelector.Selected)
	//	if len(words) == 0 {
	//		dialog.NewError(errors.New("список слов пуст"), myWindow).Show()
	//	} else {
	//		window := gui.NewShowWordsWindow(myApp, words.Shuffle(10))
	//		window.Show()
	//	}
	//})
	//myWindow.SetContent(container.NewVBox(groupSelector, startBtn))
	//myWindow.ShowAndRun()
}
