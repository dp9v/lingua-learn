package common

import "math/rand"

type WordGroups map[string]Words

type Words []Word

type Word struct {
	Original    string
	Translation string
}

func (w Words) GetRandomWords(count int) Words {
	result := make(Words, count)
	for i := 0; i < count; i++ {
		result[i] = (w)[rand.Intn(len(w))]
	}
	return result
}

func (w Words) Shuffle(count int) Words {
	if count <= 0 || count > len(w) {
		count = len(w)
	}
	shuffled := make(Words, len(w))
	copy(shuffled, w)

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:count]
}

func (wg WordGroups) GetWords(groupName string) Words {
	if groupName == "All" {
		return wg.getAll()
	}
	group, ok := wg[groupName]
	if ok {
		return group
	}
	return make(Words, 0)
}

func (wg WordGroups) GetAllGroups() []string {
	res := make([]string, len(wg)+1)
	i := 0
	res[i] = "All"
	for groupName := range wg {
		i++
		res[i] = groupName
	}
	return res
}

func (wg WordGroups) getAll() Words {
	var res Words
	for _, words := range wg {
		res = append(res, words...)
	}
	return res
}
