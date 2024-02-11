package datasources

import "math/rand"

type Word struct {
	Original    string `json:"original"`
	Translation string `json:"translation"`
}

type WordGroups map[string]Words

type Words []Word

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

func (wg WordGroups) GetWords(nameGroups []string) Words {
	var res Words
	for _, groupName := range nameGroups {
		group, ok := wg[groupName]
		if ok {
			res = append(res, group...)
		}
	}
	return res
}

func (wg WordGroups) GetAllGroups() []string {
	res := make([]string, len(wg))
	i := 0
	for groupName := range wg {
		res[i] = groupName
		i++
	}
	return res
}
