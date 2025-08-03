package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	id = strings.ReplaceAll(id, "-", "")

	length := len(id)
	if length < 2 {
		return false
	}
	sum := 0

	for i := length - 1; i >= 0; i-- {

		if id[i] < '0' || id[i] > '9' {
			return false
		}

		num := int(id[i] - '0')

		if (length-i)%2 == 0 {
			num *= 2
			if num > 9 {
				num = num - 9
			}
		}
		sum += num
	}

	return sum%10 == 0
}
