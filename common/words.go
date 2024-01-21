package common

import "math/rand"

type WordGroups map[string]Words

type Words []Word

type Word struct {
	Original    string
	Translation string
}

func (w Words) getRandomWords(count int) Words {
	result := make(Words, count)
	for i := 0; i < count; i++ {
		result[i] = (w)[rand.Intn(len(w))]
	}
	return result
}

func (w Words) Shuffle() Words {
	shuffled := make(Words, len(w))
	copy(shuffled, w)

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
