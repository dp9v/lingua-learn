package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
)

type ShowStatsActivity struct {
	app       *Application
	ds        *common.DictionaryService
	exportBtn *widget.Button
	importBtn *widget.Button
	StatText  *widget.TextGrid
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
	println("test")
	stat, err := s.ds.GetFullStat()
	if err != nil {
		dialog.ShowError(err, s.app.w)
	}
	jsonStat, err := stat.Marshal()
	if err != nil {
		dialog.ShowError(err, s.app.w)
	}
	s.StatText.SetText(jsonStat)
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
		StatText:  widget.NewTextGrid(),
	}
	res.exportBtn.Disable()
	res.importBtn.Disable()
	return &res
}
