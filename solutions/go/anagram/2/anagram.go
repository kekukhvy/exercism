package anagram

import "strings"

func Detect(subject string, candidates []string) (anagrams []string) {

	for _, c := range candidates {
		if isAnagram(subject, c) {
			anagrams = append(anagrams, c)
		}
	}

	return anagrams
}

func isAnagram(subject, candidate string) bool {

	if len(candidate) != len(subject) {
		return false
	}

	if strings.EqualFold(subject, candidate) {
		return false
	}

	subMap := makeMap(subject)
	for _, c := range strings.ToUpper(candidate) {
		if subMap[c] == 0 {
			return false
		}
		subMap[c]--
	}

	return true
}

func makeMap(word string) (m map[rune]int) {
	m = make(map[rune]int)
	for _, c := range strings.ToUpper(word) {
		m[c]++
	}

	return m
}
