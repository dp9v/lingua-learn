package gui

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"learn_words/datasources"
)

type Activity interface {
	GetContent() fyne.CanvasObject
	GetTitle() string
}

type Application struct {
	app            fyne.App
	w              fyne.Window
	content        Activity
	backBtn        *widget.Button
	updateWordsBtn *widget.Button
	label          *widget.Label
}

func NewApplication(appId string) *Application {
	myApp := app.NewWithID(appId)
	myWindow := myApp.NewWindow("")
	res := Application{
		app:   myApp,
		w:     myWindow,
		label: widget.NewLabel(""),
	}
	res.label.TextStyle = fyne.TextStyle{Bold: true}
	res.backBtn = widget.NewButton(" << ", res.showMainActivity)
	res.updateWordsBtn = widget.NewButton("Update words", res.updateWords)
	res.showMainActivity()
	res.w.ShowAndRun()
	return &res
}

func (app *Application) showMainActivity() {
	app.update(NewMainActivity(app, "Main", datasources.NewPreferencesDataSource(app.app)))
}

func (app *Application) update(content Activity) {
	app.content = content
	app.w.SetTitle(content.GetTitle())
	app.label.SetText(app.content.GetTitle())
	app.w.SetContent(app.getMainContainer())
}

func (app *Application) getMainContainer() *fyne.Container {
	return container.NewBorder(
		container.NewHBox(app.label, layout.NewSpacer(), app.backBtn),
		app.updateWordsBtn,
		nil,
		nil,
		app.content.GetContent(),
	)
}

// Temp function to upload groups to Pref from dummyData
func (a *Application) updateWords() {
	pref := a.app.Preferences()
	allGroups, _ := datasources.NewGithubDataSource().ReadAllGroups()
	groups := allGroups.GetAllGroups()
	pref.SetStringList("groups", *groups)
	for groupName, words := range *allGroups {
		wordsJson, err := json.Marshal(words)
		if err != nil {
			dialog.NewError(err, a.w)
			return
		}
		pref.SetString(groupName+"__words", string(wordsJson))
	}
	a.showMainActivity()
}
