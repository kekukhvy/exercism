package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	var isLeap bool

	if year%4 == 0 {
		isLeap = true
	}

	if year%100 == 0 && year%400 != 0 {
		isLeap = false
	}

	return isLeap
}
