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

type GuiApplication struct {
	app     fyne.App
	w       fyne.Window
	content Activity
	backBtn *widget.Button
}

func NewGuiApplication() *GuiApplication {
	myApp := app.New()
	myWindow := myApp.NewWindow("")
	res := GuiApplication{
		app: myApp,
		w:   myWindow,
	}
	res.backBtn = widget.NewButton("Switch", res.switchContent)
	res.switchContent()
	res.w.ShowAndRun()
	return &res
}

func (app *GuiApplication) switchContent() {
	if app.content == nil || app.content.GetTitle() == "test" {
		app.update(NewMainWindow("test2"))
	} else {
		app.update(NewMainWindow2("test"))
	}
}

func (app *GuiApplication) update(content Activity) {
	app.content = content
	app.w.SetContent(app.getMainContainer())
}

func (app *GuiApplication) getMainContainer() *fyne.Container {
	return container.NewBorder(
		container.NewHBox(layout.NewSpacer(), app.backBtn),
		nil,
		nil,
		nil,
		app.content.GetContent(),
	)
}
