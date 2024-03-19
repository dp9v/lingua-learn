package tests

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	v2 "learn_words/datasources/v2"
	"learn_words/gui"
	"testing"
)

var app = test.NewApp()
var source = v2.NewPreferencesDataSource(app)

func TestMainActivity_showGroups(t *testing.T) {
	guiApp := gui.NewApplication(app)
	activity := gui.NewMainActivity(guiApp, "title", source)
	assert.Len(t, activity.CheckGroup.Options, 3)
	assert.Equal(t, activity.CheckGroup.Options, []string{"Group1 (id: 1)", "Group2 (id: 2)", "Group3 (id: 3)"})
}

func TestMainActivity_startBtnClick(t *testing.T) {
	guiApp := gui.NewApplication(app)
	activity := gui.NewMainActivity(guiApp, "title", source)
	activity.CheckGroup.SetSelected(activity.CheckGroup.Options[1:3])
	test.Tap(activity.StartBtn)
	assert.Equal(t, guiApp.Content.GetTitle(), "Show words")
}

func init() {
	println("init test data")
	for _, word := range Words {
		_ = source.AddWord(&word, true)
	}
	for _, group := range Groups {
		_ = source.AddGroup(&group, true)
	}
}
