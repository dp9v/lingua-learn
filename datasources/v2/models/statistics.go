package models

import "encoding/json"

const (
	SHOW    = "SHOW"
	CORRECT = "CORRECT"
	WRONG   = "WRONG"
)

type Stat struct {
	WordId    int64          `json:"wordId"`
	Statistic map[string]int `json:"statistic"`
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
