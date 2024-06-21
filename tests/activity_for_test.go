package tests

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TestActivity struct {
	Name string
}

func (t TestActivity) GetContent() fyne.CanvasObject {
	return widget.NewLabel(t.Name)
}

func (t TestActivity) GetTitle() string {
	return t.Name
}

func NewTestActivity(name string) TestActivity {
	return TestActivity{Name: name}
}
