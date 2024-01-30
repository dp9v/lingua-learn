package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Activity interface {
	GetContent() fyne.CanvasObject
	GetTitle() string
}

type Application struct {
	app     fyne.App
	w       fyne.Window
	content Activity
	backBtn *widget.Button
	label   *widget.Label
}

func NewApplication() *Application {
	myApp := app.New()
	myWindow := myApp.NewWindow("")
	res := Application{
		app:   myApp,
		w:     myWindow,
		label: widget.NewLabel(""),
	}
	res.label.TextStyle = fyne.TextStyle{Bold: true}
	res.backBtn = widget.NewButton(" << ", res.showMainActivity)
	res.showMainActivity()
	res.w.ShowAndRun()
	return &res
}

func (app *Application) showMainActivity() {
	app.update(NewMainActivity(app, "Main"))
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
		nil,
		nil,
		nil,
		app.content.GetContent(),
	)
}
