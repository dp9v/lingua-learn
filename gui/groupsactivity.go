package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/datasources"
)

type GroupsActivity struct {
	app       *Application
	source    datasources.DataSource
	groupList *widget.List
}

func (a *GroupsActivity) GetContent() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, a.groupList)
}

func (s *GroupsActivity) GetTitle() string {
	return "Group list"
}

func NewShowGroupsActivity(app *Application, source datasources.DataSource) *GroupsActivity {
	groups, err := source.ReadAllGroups()
	if err != nil {
		dialog.NewError(err, app.w).Show()
		*groups = make(datasources.WordGroups)
	}
	groupsBinding := binding.BindStringList(groups.GetAllGroups())

	groupList := widget.NewListWithData(groupsBinding,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewButton("-", nil), widget.NewLabel("template"))
		},
		func(item binding.DataItem, object fyne.CanvasObject) {
			object.(*widget.Label).Bind(item.(binding.String))
		},
	)
	return &GroupsActivity{
		app:       app,
		source:    source,
		groupList: groupList,
	}
}
