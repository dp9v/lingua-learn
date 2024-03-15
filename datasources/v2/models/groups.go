package models

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Group struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Words []int64 `json:"words"`
}

type Groups map[int64]Group

type GroupList []Group

func UnmarshalGroups(groupsJson string) (*Groups, error) {
	var res Groups
	err := json.Unmarshal([]byte(groupsJson), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (g *Group) Str() string {
	return fmt.Sprintf("%s (id: %d)", g.Name, g.Id)
}

func (g *Groups) Marshal() (string, error) {
	jsonValue, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(jsonValue), nil
}

func (g *Groups) AsList() *GroupList {
	res := make(GroupList, len(*g))
	counter := 0
	for _, group := range *g {
		res[counter] = group
		counter++
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Id < res[j].Id
	})
	return &res
}

func (gl *GroupList) Names() *[]string {
	res := make([]string, len(*gl))
	for i, group := range *gl {
		res[i] = group.Str()
	}
	return &res
}
