package tests

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"learn_words/gui"
	"testing"
)

func TestApplication_navigation(t *testing.T) {
	app := gui.NewApplication(app)
	app.Next(NewTestActivity("test1"))
	app.Next(NewTestActivity("test2"))
	assert.Equal(t, app.Content.GetTitle(), "test2", "Activity is not expected")

	test.Tap(app.BackBtn)
	assert.Equal(t, app.Content.GetTitle(), "test1", "Activity is not expected")

	app.Next(NewTestActivity("test3"))
	assert.Equal(t, app.Content.GetTitle(), "test3", "Activity is not expected")
	test.Tap(app.BackBtn)
	test.Tap(app.BackBtn)
	assert.Equal(t, app.Content.GetTitle(), "Main", "Activity is not expected")
	assert.True(t, app.BackBtn.Disabled(), "Back button should be disabled")
}
