package datasources

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
)

type PreferencesDataSource struct {
	fyne.App
}

func (i *PreferencesDataSource) ReadAllGroups() (*WordGroups, error) {
	groups := i.Preferences().StringList("groups")
	result := make(WordGroups)
	for _, group := range groups {
		wordsJson := i.Preferences().String(group + "__words")
		var words Words
		err := json.Unmarshal([]byte(wordsJson), &words)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %v", err)
		}
		result[group] = words
	}
	return &result, nil
}

func NewPreferencesDataSource(app fyne.App) DataSource {
	return &PreferencesDataSource{app}
}
