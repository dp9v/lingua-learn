package v2

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"learn_words/datasources/v2/models"
	"testing"
)

func TestPreferencesDataSource_Groups(t *testing.T) {
	app := test.NewApp()
	testDataSource := NewPreferencesDataSource(app)

	err := testDataSource.AddGroup(&models.Group{
		Id:    1,
		Name:  "testGroup",
		Words: []int64{1},
	}, false)
	if err != nil {
		return
	}
	assert.Empty(t, err, "AddGroup: Error should be empty")

	groups, err := testDataSource.ReadAllGroups()
	assert.Empty(t, err, "ReadAllGroups: Error should be empty")
	assert.Len(t, *groups, 1)
	assert.Equal(t, &models.Groups{1: models.Group{
		Id:    1,
		Name:  "testGroup",
		Words: []int64{1},
	}}, groups)

	err = testDataSource.DeleteGroup(1)
	assert.Empty(t, err, "DeleteGroup: Error should be empty")

	groups, err = testDataSource.ReadAllGroups()
	assert.Empty(t, err, "ReadAllGroups: Error should be empty")
	assert.Len(t, *groups, 0)

}

func TestPreferencesDataSource_Words(t *testing.T) {
	app := test.NewApp()
	testDataSource := NewPreferencesDataSource(app)

	err := testDataSource.AddWord(&models.Word{
		Id:          1,
		Original:    "TestOriginal",
		Translation: "TestTranslation",
	}, false)
	assert.Empty(t, err, "AddWord: Error should be empty")

	words, err := testDataSource.ReadWords([]int64{1})
	assert.Empty(t, err, "ReadWords: Error should be empty")
	assert.Len(t, *words, 1)
	assert.Equal(t, &models.Words{1: models.Word{
		Id:          1,
		Original:    "TestOriginal",
		Translation: "TestTranslation",
	}}, words)
}

func TestPreferencesDataSource_Stats(t *testing.T) {
	app := test.NewApp()
	testDataSource := NewPreferencesDataSource(app)
	stat := models.Stats{
		1: models.Stat{
			WordId:       1,
			ShowCount:    3,
			WrongAnswers: 2,
		},
	}
	err := testDataSource.UpdateStats(&stat)
	assert.Empty(t, err, "UpdateStats: Error should be empty")

	stats, err := testDataSource.LoadStats([]int64{1, 2})
	assert.Empty(t, err, "LoadStats: Error should be empty")
	assert.Len(t, *stats, 2)
	assert.Equal(t, stats, &models.Stats{
		1: models.Stat{
			WordId:       1,
			ShowCount:    3,
			WrongAnswers: 2,
		},
		2: models.Stat{
			WordId:       2,
			ShowCount:    0,
			WrongAnswers: 0,
		},
	})
}
