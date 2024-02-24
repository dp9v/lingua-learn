package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/datasources"
)

type MainActivity struct {
	app           *Application
	ds            *datasources.DataSource
	title         string
	checkGroup    *widget.CheckGroup
	startBtn      *widget.Button
	showGroupsBtn *widget.Button
}

func (a *MainActivity) GetContent() fyne.CanvasObject {
	return container.NewVBox(a.checkGroup, a.startBtn, a.showGroupsBtn)
}

func (a *MainActivity) GetTitle() string {
	return a.title
}

func NewMainActivity(app *Application, title string, ds datasources.DataSource) *MainActivity {
	groups, err := ds.ReadAllGroups()
	if err != nil {
		dialog.ShowError(err, app.w)
		return nil
	}

	groupSelector := widget.NewCheckGroup(*groups.GetAllGroups(), func(strings []string) {})
	startBtn := widget.NewButton("Run check", func() {
		words := groups.GetWords(groupSelector.Selected)
		if len(words) == 0 {
			dialog.NewError(errors.New("список слов пуст"), app.w).Show()
		} else {
			app.update(NewShowWordsActivity(app, words.Shuffle(10)))
		}
	})
	showGroupsBtn := widget.NewButton("ShowGroups", func() {
		app.update(NewShowGroupsActivity(app, datasources.NewPreferencesDataSource(app.app)))
	})
	return &MainActivity{
		app:           app,
		title:         title,
		checkGroup:    groupSelector,
		startBtn:      startBtn,
		showGroupsBtn: showGroupsBtn,
	}
}
