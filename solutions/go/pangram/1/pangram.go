package pangram

import "strings"

func IsPangram(input string) bool {
	var letters = [26]rune{}
	var counter = 0

	for _, c := range strings.ToLower(input) {
		if c >= 'a' && c <= 'z' {
			if letters[c-'a'] == 0 {
				letters[c-'a'] = c
				counter++
			}

		}
	}

	return counter == 26
}
