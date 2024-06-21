package tests

import (
	"github.com/stretchr/testify/assert"
	"learn_words/common"
	"learn_words/gui"
	"testing"
)

func TestShowStatsActivity(t *testing.T) {
	guiApp := gui.NewApplication(app)
	testActivity := gui.NewShowStatsActivity(guiApp, common.NewDictionaryService(source))

	content := testActivity.GetContent()
	assert.NotEmpty(t, content, "content should not be empty")
	assert.NotEmpty(t, testActivity.StatText.Text)
}
