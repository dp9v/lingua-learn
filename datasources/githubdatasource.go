package datasources

import (
	"encoding/json"
	"fmt"
	"learn_words/common"
)

const WORDS_URL = "https://api.github.com/repos/dp9v/lingua-learn/contents/words"

type GithubDataSource struct {
}

func (g *GithubDataSource) ReadAllGroups() (*WordGroups, error) {
	content, err := common.LoadRepositoryContent(WORDS_URL)
	if err != nil {
		return nil, fmt.Errorf("error while reading group list: %v", err)
	}

	result := make(WordGroups)
	for _, repositoryContent := range *content {
		content, err := common.LoadPageContent(repositoryContent.DownloadUrl)
		if err != nil {
			return nil, fmt.Errorf("error while reading group content: %v", err)
		}
		var words Words
		err = json.Unmarshal([]byte(content), &words)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %v", err)
		}
		result[repositoryContent.GetShortName()] = words
	}
	return &result, nil
}
