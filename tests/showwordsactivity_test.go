package tests

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"learn_words/gui"
	"testing"
)

func TestShowWordsActivity_wrongAnswer(t *testing.T) {
	guiApp := gui.NewApplication(app)
	a := gui.NewShowWordsActivity(guiApp, WordsList, UpdateStatEmpty)
	a.GetContent()
	assert.Equal(t, WordsList[0].Translation, a.TranslationLabel.Text)
	test.Tap(a.NextBtn)
	assert.Equal(t, WordsList[0].Original, a.CorrectWordLabel.Text)
}

func TestShowWordsActivity_correctAnswer(t *testing.T) {
	guiApp := gui.NewApplication(app)
	a := gui.NewShowWordsActivity(guiApp, WordsList, UpdateStatEmpty)
	a.GetContent()
	a.Input.Text = WordsList[0].Original
	test.Tap(a.NextBtn)
	assert.Equal(t, WordsList[1].Translation, a.TranslationLabel.Text)
	assert.Empty(t, a.CorrectWordLabel.Text)
}
