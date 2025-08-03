package raindrops

import (
	"strconv"
	"strings"
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
	var result = strings.Builder{}
	for _, pair := range sounds {
		if number%pair.divider == 0 {
			result.WriteString(pair.sound)
		}
	}

	if result.Len() == 0 {
		return strconv.Itoa(number)
	}

	return result.String()
}
