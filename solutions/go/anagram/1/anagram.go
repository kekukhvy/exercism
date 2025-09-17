package anagram

import "strings"

func Detect(subject string, candidates []string) (anagrams []string) {

	for _, c := range candidates {

		r := isAnagram(subject, c)
		if r {
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
	anagramMap := makeMap(candidate)

	if len(subMap) != len(anagramMap) {
		return false
	}

	for k, v := range anagramMap {
		value, exists := subMap[k]
		if !(exists && value == v) {
			return false
		}
	}
	return true
}

func makeMap(word string) (m map[rune]int) {
	m = make(map[rune]int)
	for _, c := range strings.ToUpper(word) {
		value, exists := m[c]
		if exists {
			m[c] = value + 1
		} else {
			m[c] = 1
		}
	}

	return m
}
