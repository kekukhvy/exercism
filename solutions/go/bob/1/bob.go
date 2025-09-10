// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	isYelling := true
	hasLetters := false

	for _, c := range remark {
		if unicode.IsLetter(c) {
			hasLetters = true
			if unicode.IsLower(c) {
				isYelling = false
				break
			}
		}
	}

	if isQuestion && isYelling && hasLetters {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if isYelling && hasLetters {
		return "Whoa, chill out!"
	}

	if !isYelling && hasLetters && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}

/*
"Sure." This is his response if you ask him a question, such as "How are you?" The convention used for questions is that it ends with a question mark.
"Whoa, chill out!" This is his answer if you YELL AT HIM. The convention used for yelling is ALL CAPITAL LETTERS.
"Calm down, I know what I'm doing!" This is what he says if you yell a question at him.
"Fine. Be that way!" This is how he responds to silence. The convention used for silence is nothing, or various combinations of whitespace characters.
"Whatever." This is what he answers to anything else.
*/
