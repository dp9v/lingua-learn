package v2

import (
	"fmt"
	"fyne.io/fyne/v2"
	"learn_words/datasources/v2/models"
)

const GROUPS_PATTERN = "GroupsV2"
const WORD_ID_PATTERN = "WordsV2_%d"
const STAT_ID_PATTERN = "StatsV2_%d"

type PreferencesDataSource struct {
	fyne.App
}

func (p *PreferencesDataSource) ReadAllGroups() (*models.Groups, error) {
	result := p.Preferences().String(GROUPS_PATTERN)
	if len(result) == 0 {
		return &models.Groups{}, nil
	}
	return models.UnmarshalGroups(result)
}

func (p *PreferencesDataSource) ReadWords(ids []int64) (*models.Words, error) {
	res := make(models.Words)
	for _, id := range ids {
		wordJson := p.Preferences().String(fmt.Sprintf(WORD_ID_PATTERN, id))
		if len(wordJson) == 0 {
			continue
		}
		word, err := models.UnmarshalWord(wordJson)
		if err != nil {
			fyne.LogError("Word can not be unmarshalled", err)
			return nil, err
		}
		res[word.Id] = *word
	}
	return &res, nil
}

func (p *PreferencesDataSource) AddGroup(group *models.Group, force bool) error {
	groups, err := p.ReadAllGroups()
	if err != nil {
		return err
	}

	_, exists := (*groups)[group.Id]
	if exists && !force {
		return fmt.Errorf("group with id: %d already saved", group.Id)
	}
	(*groups)[group.Id] = *group
	jsonGroups, err := groups.Marshal()
	if err != nil {
		return err
	}
	p.Preferences().SetString(GROUPS_PATTERN, jsonGroups)
	return nil
}

func (p *PreferencesDataSource) DeleteGroup(groupId int64) error {
	groups, err := p.ReadAllGroups()
	if err != nil {
		return err
	}
	delete(*groups, groupId)
	jsonGroups, err := groups.Marshal()
	if err != nil {
		return err
	}
	p.Preferences().SetString(GROUPS_PATTERN, jsonGroups)
	return nil
}

func (p *PreferencesDataSource) AddWord(word *models.Word, force bool) error {
	words, err := p.ReadWords([]int64{word.Id})
	if err != nil {
		return err
	}
	if len(*words) == 1 && !force {
		return fmt.Errorf("word with id: %d already saved", word.Id)
	}

	jsonWord, err := word.Marshal()
	if err != nil {
		return err
	}

	p.Preferences().SetString(fmt.Sprintf(WORD_ID_PATTERN, word.Id), jsonWord)
	return nil
}

func (p *PreferencesDataSource) LoadStat(id int64) (*models.Stat, error) {
	wordJson := p.Preferences().StringWithFallback(fmt.Sprintf(STAT_ID_PATTERN, id), "{}")
	if wordJson == "{}" {
		return &models.Stat{WordId: id}, nil
	}
	return models.UnmarshalStat(wordJson)
}

func (p *PreferencesDataSource) LoadStats(ids []int64) (*models.Stats, error) {
	res := make(models.Stats)
	for _, id := range ids {
		stat, err := p.LoadStat(id)
		if err != nil {
			fyne.LogError("Stats can not be unmarshalled", err)
			return nil, fmt.Errorf("stat if:%d can not be unmarshalled: %s", id, err)
		}
		res[stat.WordId] = *stat
	}
	return &res, nil
}

func (p *PreferencesDataSource) UpdateStats(stats *models.Stats) error {
	errorCount := 0
	for _, stat := range *stats {
		err := p.UpdateStat(&stat)
		if err != nil {
			fyne.LogError("stats can not be updated", err)
			errorCount++
		}
	}
	if errorCount > 0 {
		return fmt.Errorf("%d error(s) occurred in updating stats", errorCount)
	}
	return nil
}

func (p *PreferencesDataSource) UpdateStat(stat *models.Stat) error {
	statJson, err := stat.Marshal()
	if err != nil {
		return fmt.Errorf("stat id=%d can not be marshalled: %s", stat.WordId, err)
	}
	p.Preferences().SetString(fmt.Sprintf(STAT_ID_PATTERN, stat.WordId), statJson)
	return nil
}

func NewPreferencesDataSource(app fyne.App) RWDataSourceV2 {
	return &PreferencesDataSource{app}
}
