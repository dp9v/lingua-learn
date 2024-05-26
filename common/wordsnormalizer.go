package common

var czechMap = map[rune]rune{
	'á': 'a',
	'č': 'c',
	'ď': 'd',
	'é': 'e',
	'ě': 'e',
	'í': 'i',
	'ň': 'n',
	'ó': 'o',
	'ř': 'r',
	'š': 's',
	'ť': 't',
	'ú': 'u',
	'ů': 'u',
	'ý': 'y',
	'ž': 'z',
}

func Normalize(word string) string {
	translated := ""
	for _, char := range word {
		if val, ok := czechMap[char]; ok {
			translated += string(val)
		} else {
			translated += string(char)
		}
	}
	return translated
}
