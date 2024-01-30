package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type GuiApplication struct {
	app       fyne.App
	w         fyne.Window
	container fyne.CanvasObject
	state     string
	backBtn   *widget.Button
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
	if app.container == nil || app.state == "w2" {
		app.update(NewMainWindow("test").getContent())
		app.state = "w1"
	} else {
		app.update(NewMainWindow2("test2").getContent())
		app.state = "w2"
	}
}

func (app *GuiApplication) update(content fyne.CanvasObject) {
	app.container = container.NewBorder(
		container.NewHBox(layout.NewSpacer(), app.backBtn),
		nil, nil, nil, content,
	)
	app.w.SetContent(app.container)
}
