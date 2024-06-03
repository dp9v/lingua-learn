package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	v2 "learn_words/datasources/v2"
	"learn_words/datasources/v2/models"
	"regexp"
	"slices"
	"strconv"
)

type MainActivity struct {
	CheckGroup    *widget.CheckGroup
	StartBtn      *widget.Button
	ShowGroupsBtn *widget.Button
	app           *Application
	ds            v2.DataSourceV2
	groups        *models.Groups
	title         string
}

func (a *MainActivity) GetContent() fyne.CanvasObject {
	return container.NewVBox(a.CheckGroup, a.StartBtn, a.ShowGroupsBtn)
}

func (a *MainActivity) GetTitle() string {
	return a.title
}

func (a *MainActivity) refresh() {
	groups, err := a.ds.ReadAllGroups()
	if err != nil {
		dialog.ShowError(err, a.app.w)
		return
	}
	a.groups = groups
	a.CheckGroup.Options = *groups.AsList().Names()
	a.CheckGroup.SetSelected([]string{})
}

func (a *MainActivity) startBtnClick() {
	var wordIds []int64
	r := regexp.MustCompile("\\(id: (\\d+)\\)")
	for _, s := range a.CheckGroup.Selected {
		id, err := strconv.Atoi(r.FindStringSubmatch(s)[1])
		if err != nil {
			dialog.ShowError(err, a.app.w)
			return
		}
		wordIds = append(wordIds, (*a.groups)[int64(id)].Words...)
	}
	slices.Sort(wordIds)
	slices.Compact(wordIds)
	wordIds = slices.Clip(wordIds)
	words, err := a.ds.ReadWords(wordIds)
	if err != nil {
		dialog.ShowError(err, a.app.w)
		return
	}
	wordsToShow := words.Shuffle(13)
	a.app.Next(NewShowWordsActivity(a.app, wordsToShow))
}

func NewMainActivity(app *Application, title string, ds v2.DataSourceV2) *MainActivity {
	groupSelector := widget.NewCheckGroup([]string{}, nil)
	startBtn := widget.NewButton("Run check", nil)
	showGroupsBtn := widget.NewButton("ShowGroups", func() {
		app.Next(NewShowGroupsActivity(app, v2.NewPreferencesDataSource(app.app)))
	})
	res := &MainActivity{
		app:           app,
		title:         title,
		CheckGroup:    groupSelector,
		StartBtn:      startBtn,
		ShowGroupsBtn: showGroupsBtn,
		ds:            ds,
	}
	res.StartBtn.OnTapped = res.startBtnClick
	res.refresh()
	return res
}
