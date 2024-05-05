package models

import "encoding/json"

type Stat struct {
	WordId       int64 `json:"wordId"`
	ShowCount    int   `json:"showCount"`
	WrongAnswers int64 `json:"wrongAnswers"`
}

type Stats map[int64]Stat

func UnmarshalStat(statJson string) (*Stat, error) {
	var res Stat
	err := json.Unmarshal([]byte(statJson), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *Stat) Marshal() (string, error) {
	jsonValue, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}
	return string(jsonValue), nil
}
