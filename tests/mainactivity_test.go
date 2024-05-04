package tests

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"learn_words/datasources/v2/models"
	"learn_words/gui"
	"sort"
	"testing"
)

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
	assert.IsType(t, &gui.ShowWordsActivity{}, guiApp.Content)

	words := guiApp.Content.(*gui.ShowWordsActivity).RoundWords
	sort.Slice(words, func(i, j int) bool {
		return words[i].Id < words[j].Id
	})
	assert.Equal(t, models.WordList{Words[2], Words[3], Words[4], Words[5]}, words)
}

func TestMainActivity_showGroupsClick(t *testing.T) {
	guiApp := gui.NewApplication(app)
	activity := gui.NewMainActivity(guiApp, "title", source)
	test.Tap(activity.ShowGroupsBtn)
	assert.IsType(t, &gui.GroupsActivity{}, guiApp.Content)
}
