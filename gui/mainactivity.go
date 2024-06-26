package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/common"
	v2 "learn_words/datasources/v2"
	"learn_words/datasources/v2/models"
	"regexp"
	"strconv"
)

type MainActivity struct {
	CheckGroup    *widget.CheckGroup
	StartBtn      *widget.Button
	ShowGroupsBtn *widget.Button
	ShowStatsBtn  *widget.Button
	app           *Application
	ds            *common.DictionaryService
	groups        *models.Groups
	title         string
}

func (a *MainActivity) GetContent() fyne.CanvasObject {
	return container.NewVBox(a.CheckGroup, a.StartBtn, a.ShowStatsBtn)
}

func (a *MainActivity) GetTitle() string {
	return a.title
}

func (a *MainActivity) refresh() {
	groups, err := a.ds.GetGroupNames()
	if err != nil {
		dialog.ShowError(err, a.app.w)
		return
	}
	a.CheckGroup.Options = groups
	a.CheckGroup.SetSelected([]string{})
}

func (a *MainActivity) startBtnClick() {
	var groupIds []int64
	r := regexp.MustCompile("\\(id: (\\d+)\\)")
	for _, s := range a.CheckGroup.Selected {
		id, err := strconv.Atoi(r.FindStringSubmatch(s)[1])
		if err != nil {
			dialog.ShowError(err, a.app.w)
			return
		}

		groupIds = append(groupIds, int64(id))
	}
	words, err := a.ds.GetRandomWords(13, groupIds)
	if err != nil {
		dialog.ShowError(err, a.app.w)
		return
	}
	a.app.Next(NewShowWordsActivity(a.app, words, a.ds.IncrementStatValue))
}

func NewMainActivity(app *Application, title string, ds v2.RWDataSourceV2) *MainActivity {
	groupSelector := widget.NewCheckGroup([]string{}, nil)
	startBtn := widget.NewButton("Run check", nil)
	dictionaryService := common.NewDictionaryService(ds)
	showGroupsBtn := widget.NewButton("ShowGroups", func() {
		app.Next(NewShowGroupsActivity(app, ds)) //ToDo: remove this functionality
	})
	showStatsBtn := widget.NewButton("ShowStats", func() {
		app.Next(NewShowStatsActivity(app, dictionaryService))
	})
	res := &MainActivity{
		app:           app,
		title:         title,
		CheckGroup:    groupSelector,
		StartBtn:      startBtn,
		ShowGroupsBtn: showGroupsBtn,
		ShowStatsBtn:  showStatsBtn,
		ds:            dictionaryService,
	}
	res.StartBtn.OnTapped = res.startBtnClick
	res.refresh()
	return res
}
