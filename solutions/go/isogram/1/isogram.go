package isogram

func IsIsogram(word string) bool {
	var seen [26]bool

	for _, c := range word {
		var index int

		if c >= 'A' && c <= 'Z' {
			index = int(c - 'A')
		} else if c >= 'a' && c <= 'z' {
			index = int(c - 'a') // Convert lowercase to same index
		} else {
			continue // Skip non-letters
		}

		if seen[index] {
			return false
		}

		seen[index] = true
	}
	return true
}
