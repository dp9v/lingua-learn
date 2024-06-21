package tests

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"learn_words/common"
	"learn_words/datasources/v2/models"
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

func TestShowWordsActivity_collectStatistics(t *testing.T) {
	CleanStat()
	guiApp := gui.NewApplication(app)
	dictService := common.NewDictionaryService(source)
	words := WordsList[:3]
	a := gui.NewShowWordsActivity(guiApp, words, dictService.IncrementStatValue)
	a.GetContent()

	a.Input.Text = WordsList[0].Original
	test.Tap(a.NextBtn)
	a.Input.Text = "Wrong1"
	test.Tap(a.NextBtn)
	a.Input.Text = "Wrong2"
	test.Tap(a.NextBtn)
	a.Input.Text = WordsList[1].Original
	test.Tap(a.NextBtn)
	a.Input.Text = WordsList[2].Original
	test.Tap(a.NextBtn)
	a.Input.Text = WordsList[1].Original
	test.Tap(a.NextBtn)

	stats, err := source.ReadStats([]int64{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, *stats, models.Stats{
		1: {
			WordId: 1,
			Statistic: map[string]int{
				models.SHOW:    1,
				models.CORRECT: 1,
			},
		},
		2: {
			WordId: 2,
			Statistic: map[string]int{
				models.SHOW:    2,
				models.WRONG:   1,
				models.CORRECT: 1,
			},
		},
		3: {
			WordId: 3,
			Statistic: map[string]int{
				models.SHOW:    1,
				models.CORRECT: 1,
			},
		},
	})
}
