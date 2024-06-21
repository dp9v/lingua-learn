package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"learn_words/common"
)

type ShowStatsActivity struct {
	app       *Application
	ds        *common.DictionaryService
	exportBtn *widget.Button
	importBtn *widget.Button
	StatText  *canvas.Text
}

func (s *ShowStatsActivity) GetContent() fyne.CanvasObject {
	defer s.Refresh()
	return container.NewBorder(
		nil,
		container.NewHBox(s.exportBtn, s.importBtn),
		nil,
		nil,
		container.NewScroll(s.StatText),
	)
}

func (s *ShowStatsActivity) Refresh() {
	stat, err := s.ds.GetFullStat()
	if err != nil {
		dialog.ShowError(err, s.app.w)
	}
	jsonStat, err := stat.Marshal()
	if err != nil {
		dialog.ShowError(err, s.app.w)
	}
	s.StatText.Text = string(jsonStat)
}

func (s *ShowStatsActivity) GetTitle() string {
	return "Show stats"
}

func NewShowStatsActivity(app *Application, ds *common.DictionaryService) *ShowStatsActivity {
	res := ShowStatsActivity{
		app:       app,
		ds:        ds,
		importBtn: widget.NewButton("Import..", func() {}),
		exportBtn: widget.NewButton("Export..", func() {}),
		StatText:  canvas.NewText("", color.White),
	}
	res.exportBtn.Disable()
	res.importBtn.Disable()
	return &res
}
