package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return unicode.IsSpace(r) ||
			unicode.IsSymbol(r) ||
			(unicode.IsPunct(r) && r != '\'')
	})

	var r = make(map[string]int)
	for _, word := range words {
		word = normalizeWord(word)
		if len(word) > 0 {
			r[word]++
		}

	}
	return r
}

func normalizeWord(word string) string {
	word = strings.ToLower(word)
	word = strings.Trim(word, "'")
	return word
}
