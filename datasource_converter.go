package main

import (
	"encoding/json"
	"learn_words/datasources/v2/models"
	"os"
	"path/filepath"
	"strings"
)

func Convert() {
	var groupCounter int64
	wordsCounter := getMaxId()
	groups := make(models.Groups)
	words := make(models.Words)
	files, err := os.ReadDir("./words")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var wordIds []int64
		groupWords := readV1Words(file)

		for i, word := range groupWords {
			if word.Id == 0 {
				groupWords[i].Id = wordsCounter
				word.Id = wordsCounter
				wordsCounter++
			}
			wordIds = append(wordIds, word.Id)
			words[word.Id] = word
		}
		groupName := strings.Replace(file.Name(), ".json", "", 1)
		groups[groupCounter] = models.Group{
			Id:    groupCounter,
			Name:  groupName,
			Words: wordIds,
		}
		updateV1Group(file, groupWords)
		groupCounter++
	}
	saveWords(words)
	saveGroups(groups)
}

func getMaxId() int64 {
	filePath, err := filepath.Abs(filepath.Join("words", "v2", "words.json"))
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var words models.Words
	err = json.Unmarshal(file, &words)
	if err != nil {
		panic(err)
	}
	maxId := int64(1)
	for _, word := range words {
		if word.Id > maxId {
			maxId = word.Id
		}
	}
	return maxId
}

func saveWords(words models.Words) {
	wordsJson, err := json.MarshalIndent(words, "", "  ")
	if err != nil {
		panic(err)
	}
	filePath, err := filepath.Abs(filepath.Join("words", "v2", "words.json"))
	err = os.WriteFile(filePath, wordsJson, 0777)
	if err != nil {
		panic(err)
	}
}

func saveGroups(groups models.Groups) {
	wordsJson, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		panic(err)
	}
	filePath, err := filepath.Abs(filepath.Join("words", "v2", "groups.json"))
	err = os.WriteFile(filePath, wordsJson, 0777)
	if err != nil {
		panic(err)
	}
}

func readV1Words(file os.DirEntry) models.WordList {
	absPath, err := filepath.Abs(filepath.Join("words", file.Name()))
	if err != nil {
		panic(err)
	}
	fileContent, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	groupWords := models.WordList{}
	err = json.Unmarshal(fileContent, &groupWords)
	if err != nil {
		panic(err)
	}
	return groupWords
}

func updateV1Group(file os.DirEntry, words models.WordList) {
	wordsJson, err := json.MarshalIndent(words, "", "  ")
	if err != nil {
		panic(err)
	}
	filePath, err := filepath.Abs(filepath.Join("words", file.Name()))
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filePath, wordsJson, 0777)
	if err != nil {
		panic(err)
	}
}
