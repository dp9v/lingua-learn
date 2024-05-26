package v2

import "learn_words/datasources/v2/models"

type DataSourceV2 interface {
	ReadAllGroups() (*models.Groups, error)
	ReadWords([]int64) (*models.Words, error)
	LoadStat(id int64) (*models.Stat, error)
	LoadStats([]int64) (*models.Stats, error)
}

type RWDataSourceV2 interface {
	DataSourceV2
	AddGroup(group *models.Group, force bool) error
	DeleteGroup(groupId int64) error
	AddWord(word *models.Word, force bool) error
	UpdateStats(stats *models.Stats) error
	UpdateStat(stat *models.Stat) error
}
