package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

func ColorCode(color string) int {
	colors := Colors()

	for i, c := range colors {
		if color == c {
			return i
		}
	}

	return -1
}
