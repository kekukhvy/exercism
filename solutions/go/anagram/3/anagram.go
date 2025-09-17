package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (anagrams []string) {

	subjSignature := getSignature(subject)

	for _, c := range candidates {
		if strings.EqualFold(subject, c) {
			continue
		}

		if subjSignature == getSignature(c) {
			anagrams = append(anagrams, c)
		}
	}

	return anagrams
}

func getSignature(w string) string {
	runes := []rune(strings.ToUpper(w))
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
