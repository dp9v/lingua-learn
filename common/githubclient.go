package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RepositoryContent struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
}

func (c *RepositoryContent) GetShortName() string {
	return strings.Split(c.Name, ".")[1]
}

func LoadRepositoryContent(url string) (*[]RepositoryContent, error) {
	body, err := LoadPageContent(url)
	if err != nil {
		return nil, fmt.Errorf("error reading url body: %v", err)
	}

	result := &[]RepositoryContent{}
	err = json.Unmarshal([]byte(body), result)
	if err != nil {
		return nil, fmt.Errorf("error parsing body json: %v", err)
	}
	return result, nil
}

func LoadPageContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("URL is not reachable: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading url body: %v", err)
	}
	return string(body), nil
}
