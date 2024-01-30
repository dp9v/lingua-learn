package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
)

type MainActivity struct {
	app        *Application
	title      string
	checkGroup *widget.CheckGroup
	startBtn   *widget.Button
}

func (a *MainActivity) GetContent() fyne.CanvasObject {
	return container.NewVBox(a.checkGroup, a.startBtn)
}

func (a *MainActivity) GetTitle() string {
	return a.title
}

func NewMainActivity(app *Application, title string) *MainActivity {

	groupSelector := widget.NewCheckGroup(common.Groups.GetAllGroups(), func(strings []string) {})
	startBtn := widget.NewButton("Run check", func() {
		words := common.Groups.GetWords(groupSelector.Selected)
		if len(words) == 0 {
			dialog.NewError(errors.New("список слов пуст"), app.w).Show()
		} else {
			app.update(NewShowWordsActivity(app, words.Shuffle(10)))
		}
	})
	return &MainActivity{
		app:        app,
		title:      title,
		checkGroup: groupSelector,
		startBtn:   startBtn,
	}
}
