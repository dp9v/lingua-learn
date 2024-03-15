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
	app           *Application
	ds            v2.DataSourceV2
	groups        *models.Groups
	title         string
	checkGroup    *widget.CheckGroup
	startBtn      *widget.Button
	showGroupsBtn *widget.Button
}

func (a *MainActivity) GetContent() fyne.CanvasObject {
	return container.NewVBox(a.checkGroup, a.startBtn, a.showGroupsBtn)
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
	a.checkGroup.Options = *groups.AsList().Names()
	a.checkGroup.SetSelected([]string{})
}

func (a *MainActivity) startBtnClick() {
	var wordIds []int64
	r := regexp.MustCompile("\\(id: (\\d+)\\)")
	for _, s := range a.checkGroup.Selected {
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
	a.app.update(NewShowWordsActivity(a.app, wordsToShow))
}

func NewMainActivity(app *Application, title string, ds v2.DataSourceV2) *MainActivity {
	groupSelector := widget.NewCheckGroup([]string{}, nil)
	startBtn := widget.NewButton("Run check", nil)
	showGroupsBtn := widget.NewButton("ShowGroups", func() {
		app.update(NewShowGroupsActivity(app, v2.NewPreferencesDataSource(app.app)))
	})
	res := &MainActivity{
		app:           app,
		title:         title,
		checkGroup:    groupSelector,
		startBtn:      startBtn,
		showGroupsBtn: showGroupsBtn,
		ds:            ds,
	}
	res.startBtn.OnTapped = res.startBtnClick
	res.refresh()
	return res
}
