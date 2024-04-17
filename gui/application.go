package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"learn_words/datasources"
	v2 "learn_words/datasources/v2"
	"learn_words/datasources/v2/models"
)

type Activity interface {
	GetContent() fyne.CanvasObject
	GetTitle() string
}

type Application struct {
	Content        Activity
	app            fyne.App
	w              fyne.Window
	backBtn        *widget.Button
	updateWordsBtn *widget.Button
	label          *widget.Label
}

type TestApplication struct {
	*Application
}

func (t TestApplication) update(content Activity) {

}

func NewApplication(myApp fyne.App) *Application {
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
	activity := NewMainActivity(app, "Main", v2.NewPreferencesDataSource(app.app))
	if activity != nil {
		app.update(activity)
	}
}

func (app *Application) update(content Activity) {
	app.Content = content
	app.w.SetTitle(content.GetTitle())
	app.label.SetText(app.Content.GetTitle())
	app.w.SetContent(app.getMainContainer())
}

func (app *Application) getMainContainer() *fyne.Container {
	return container.NewBorder(
		container.NewHBox(app.label, layout.NewSpacer(), app.backBtn),
		app.updateWordsBtn,
		nil,
		nil,
		app.Content.GetContent(),
	)
}

// Temp function to upload groups to Pref from dummyData
func (app *Application) updateWords() {
	writeDatasource := v2.NewPreferencesDataSource(app.app)
	allGroups, err := datasources.NewGithubDataSource().ReadAllGroups()
	if err != nil {
		dialog.ShowError(err, app.w)
	}
	wordIdCounter := int64(0)
	groupIdCounter := int64(0)

	for groupName, words := range *allGroups {
		wordIds := make([]int64, len(words))
		for i, word := range words {
			err := writeDatasource.AddWord(&models.Word{
				Id:          wordIdCounter,
				Original:    word.Original,
				Translation: word.Translation,
			}, true)
			if err != nil {
				dialog.ShowError(err, app.w)
				return
			}
			wordIds[i] = wordIdCounter
			wordIdCounter++
		}
		err := writeDatasource.AddGroup(&models.Group{
			Id:    groupIdCounter,
			Name:  groupName,
			Words: wordIds,
		}, true)
		if err != nil {
			dialog.ShowError(err, app.w)
			return
		}
		groupIdCounter++
	}
	app.showMainActivity()
}
