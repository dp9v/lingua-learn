package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	v2 "learn_words/datasources/v2"
	"learn_words/datasources/v2/models"
)

type GroupsActivity struct {
	app       *Application
	source    v2.DataSourceV2
	groupList *widget.List
}

func (a *GroupsActivity) GetContent() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, a.groupList)
}

func (a *GroupsActivity) GetTitle() string {
	return "Group list"
}

func NewShowGroupsActivity(app *Application, source v2.DataSourceV2) *GroupsActivity {
	groups, err := source.ReadAllGroups()
	if err != nil {
		dialog.NewError(err, app.w).Show()
		*groups = make(models.Groups)
	}
	groupsBinding := binding.BindStringList(groups.AsList().Names())

	groupList := widget.NewListWithData(groupsBinding,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewButton("-", nil), widget.NewLabel("template"))
		},
		func(item binding.DataItem, object fyne.CanvasObject) {
			text := object.(*fyne.Container).Objects[0].(*widget.Label)
			text.Bind(item.(binding.String))
		},
	)
	return &GroupsActivity{
		app:       app,
		source:    source,
		groupList: groupList,
	}
}
