package raindrops

import (
	"strconv"
)

var sounds = []struct {
	divider int
	sound   string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var result = ""
	for _, pair := range sounds {
		if number%pair.divider == 0 {
			result += pair.sound
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(number)
	}

	return result
}
