package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
)

type MainWindow struct {
	Title      string
	checkGroup *widget.CheckGroup
	startBtn   *widget.Button
}

func NewMainWindow(title string) *MainWindow {

	groupSelector := widget.NewCheckGroup(common.Groups.GetAllGroups(), func(strings []string) {})
	startBtn := widget.NewButton("Run check 1", func() {
		print("test btn click 1 ")
	})
	return &MainWindow{
		Title:      title,
		checkGroup: groupSelector,
		startBtn:   startBtn,
	}
}

func (w MainWindow) getContent() fyne.CanvasObject {
	return container.NewVBox(w.checkGroup, w.startBtn)
}
