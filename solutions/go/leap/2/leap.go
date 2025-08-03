package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	return year%100 == 0
}
