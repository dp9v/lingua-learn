package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"learn_words/datasources"
)

type ShowGroupsActivity struct {
	app       *Application
	source    datasources.DataSource
	groupList *widget.List
}

func (a *ShowGroupsActivity) GetContent() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, a.groupList)
}

func (s *ShowGroupsActivity) GetTitle() string {
	return "Group list"
}

func NewShowGroupsActivity(app *Application, source datasources.DataSource) *ShowGroupsActivity {
	groups, err := source.ReadAllGroups()
	if err != nil {
		dialog.NewError(err, app.w).Show()
		return nil
	}
	groupsBinding := binding.BindStringList(groups.GetAllGroups())

	groupList := widget.NewListWithData(groupsBinding,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(item binding.DataItem, object fyne.CanvasObject) {
			object.(*widget.Label).Bind(item.(binding.String))
		},
	)
	return &ShowGroupsActivity{
		app:       app,
		source:    source,
		groupList: groupList,
	}
}
