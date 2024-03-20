package main

import (
	"encoding/json"
	"learn_words/datasources/v2/models"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var groupCounter int64
	var wordsCounter int64
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
		groupWords := readWords(file)

		for _, word := range groupWords {
			word.Id = wordsCounter
			wordIds = append(wordIds, wordsCounter)
			words[wordsCounter] = word
			wordsCounter++
		}
		groupName := strings.Replace(file.Name(), ".json", "", 1)
		groups[groupCounter] = models.Group{
			Id:    groupCounter,
			Name:  groupName,
			Words: wordIds,
		}
		groupCounter++
	}
	saveWords(words)
	saveGroups(groups)
}

func readWords(file os.DirEntry) models.WordList {
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

func saveWords(words models.Words) {
	wordsJson, err := json.Marshal(words)
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
	wordsJson, err := json.Marshal(groups)
	if err != nil {
		panic(err)
	}
	filePath, err := filepath.Abs(filepath.Join("words", "v2", "groups.json"))
	err = os.WriteFile(filePath, wordsJson, 0777)
	if err != nil {
		panic(err)
	}
}
