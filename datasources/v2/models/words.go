package models

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"math/rand"
	"sort"
)

type Word struct {
	Id          int64  `json:"id"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
}

type Words map[int64]Word

type WordList []Word

func UnmarshalWord(wordsJson string) (*Word, error) {
	var res Word
	err := json.Unmarshal([]byte(wordsJson), &res)
	if err != nil {
		fyne.LogError("Unmarshal error", err)
		return nil, err
	}
	return &res, nil
}

func (w *Word) Marshal() (string, error) {
	jsonValue, err := json.Marshal(&w)
	if err != nil {
		return "", err
	}
	return string(jsonValue), nil
}

func (w *Words) AsList() WordList {
	res := make(WordList, len(*w))
	counter := 0
	for _, word := range *w {
		res[counter] = word
		counter++
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Id < res[j].Id
	})
	return res
}

func (w *Words) Shuffle(count int) WordList {
	words := w.AsList()

	if count <= 0 || count > len(words) {
		count = len(words)
	}
	shuffled := make(WordList, len(words))
	copy(shuffled, words)

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:count]
}
