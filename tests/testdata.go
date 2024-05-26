package tests

import (
	"fyne.io/fyne/v2/test"
	v2 "learn_words/datasources/v2"
	"learn_words/datasources/v2/models"
)

var app = test.NewApp()
var source = v2.NewPreferencesDataSource(app)

var Words = models.Words{
	1: {
		Id:          1,
		Original:    "Original1",
		Translation: "Translation1",
	},
	2: {
		Id:          2,
		Original:    "Original2",
		Translation: "Translation2",
	},
	3: {
		Id:          3,
		Original:    "Original3",
		Translation: "Translation3",
	},
	4: {
		Id:          4,
		Original:    "Original4",
		Translation: "Translation4",
	},
	5: {
		Id:          5,
		Original:    "Original5",
		Translation: "Translation5",
	},
}
var WordsList = Words.AsList()

var Groups = models.Groups{
	1: {
		Id:    1,
		Name:  "Group1",
		Words: []int64{1, 2},
	},
	2: {
		Id:    2,
		Name:  "Group2",
		Words: []int64{3, 4, 5},
	},
	3: {
		Id:    3,
		Name:  "Group3",
		Words: []int64{2, 3},
	},
}

func UpdateStatEmpty(id int64, k int) error {
	return nil
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
