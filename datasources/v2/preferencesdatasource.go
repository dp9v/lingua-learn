package v2

import (
	"fmt"
	"fyne.io/fyne/v2"
	"learn_words/datasources/v2/models"
	"slices"
	"sort"
)

const GROUPS_PATTERN = "GroupsV2"
const SAVED_WORDS = "SavedWords"
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

func (p *PreferencesDataSource) ReadAllWords() (*models.Words, error) {
	wordIds := p.Preferences().IntList(SAVED_WORDS)
	wordIds64 := make([]int64, len(wordIds))
	for i, wordId := range wordIds {
		wordIds64[i] = int64(wordId)
	}
	return p.ReadWords(wordIds64)
}

func (p *PreferencesDataSource) ReadAllStats() (*models.Stats, error) {
	wordIds := p.Preferences().IntList(SAVED_WORDS)
	wordIds64 := make([]int64, len(wordIds))
	for i, wordId := range wordIds {
		wordIds64[i] = int64(wordId)
	}
	return p.ReadStats(wordIds64)
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
	p.addWordId(int(word.Id))

	return nil
}

func (p *PreferencesDataSource) addWordId(id int) {
	list := p.Preferences().IntList(SAVED_WORDS)
	savedIds := append(list, id)
	sort.Ints(savedIds)
	slices.Compact(savedIds)
	slices.Clip(savedIds)
	p.Preferences().SetIntList(SAVED_WORDS, savedIds)
}

func (p *PreferencesDataSource) ReadStat(id int64) (*models.Stat, error) {
	wordJson := p.Preferences().StringWithFallback(fmt.Sprintf(STAT_ID_PATTERN, id), "{}")
	if wordJson == "{}" {
		return &models.Stat{WordId: id, Statistic: map[string]int{}}, nil
	}
	return models.UnmarshalStat(wordJson)
}

func (p *PreferencesDataSource) ReadStats(ids []int64) (*models.Stats, error) {
	res := make(models.Stats)
	for _, id := range ids {
		stat, err := p.ReadStat(id)
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
