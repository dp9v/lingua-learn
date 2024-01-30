package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
)

type MainWindow2 struct {
	Title      string
	checkGroup *widget.CheckGroup
	startBtn   *widget.Button
}

func (w *MainWindow2) GetContent() fyne.CanvasObject {
	return container.NewVBox(w.checkGroup, w.startBtn)
}

func (w *MainWindow2) GetTitle() string {
	return w.Title
}

func NewMainWindow2(title string) *MainWindow2 {

	groupSelector := widget.NewCheckGroup(common.Groups.GetAllGroups(), func(strings []string) {})
	startBtn := widget.NewButton("Run check 2", func() {
		print("test btn click 2 ")
	})
	return &MainWindow2{
		Title:      title,
		checkGroup: groupSelector,
		startBtn:   startBtn,
	}
}
