package models

import "encoding/json"

const (
	SHOW    = iota
	CORRECT = iota
	WRONG   = iota
)

type Stat struct {
	WordId    int64       `json:"wordId"`
	Statistic map[int]int `json:"statistic"`
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
