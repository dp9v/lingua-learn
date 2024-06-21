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

func TestPreferencesDataSource_ReadAllWords(t *testing.T) {
	app := test.NewApp()
	testDataSource := NewPreferencesDataSource(app)
	_ = testDataSource.AddWord(&models.Word{
		Id: 1,
	}, false)
	_ = testDataSource.AddWord(&models.Word{
		Id: 4,
	}, false)
	words, err := testDataSource.ReadAllWords()
	assert.NoError(t, err, "ReadAllWords: Error should be empty")
	assert.Len(t, *words, 2)
	assert.Equal(t, &models.Words{1: models.Word{Id: 1}, 4: models.Word{Id: 4}}, words)
}

func TestPreferencesDataSource_Stats(t *testing.T) {
	app := test.NewApp()
	testDataSource := NewPreferencesDataSource(app)
	_ = testDataSource.AddWord(&models.Word{Id: 1}, false)

	stat := models.Stats{
		1: models.Stat{
			WordId: 1,
			Statistic: map[string]int{
				models.SHOW:    3,
				models.WRONG:   2,
				models.CORRECT: 1,
			},
		},
	}
	err := testDataSource.UpdateStats(&stat)
	assert.Empty(t, err, "UpdateStats: Error should be empty")

	stats, err := testDataSource.ReadStats([]int64{1, 2})
	assert.Empty(t, err, "ReadStats: Error should be empty")
	assert.Len(t, *stats, 2)
	assert.Equal(t, stats, &models.Stats{
		1: models.Stat{
			WordId: 1,
			Statistic: map[string]int{
				models.SHOW:    3,
				models.WRONG:   2,
				models.CORRECT: 1,
			},
		},
		2: models.Stat{
			WordId:    2,
			Statistic: map[string]int{},
		},
	})

	allStats, err := testDataSource.ReadAllStats()
	assert.NoError(t, err, "ReadAllStats: Error should be empty")
	assert.Len(t, *allStats, 1)
}
