package isbn

import (
	"fmt"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	multiplyer := 10
	sum := 0

	for i := range isbn {

		if isbn[i] == 'X' {
			sum += 10
			continue
		}

		value := int(isbn[i] - '0')

		fmt.Println(value, "*", multiplyer, "=", value*multiplyer)
		fmt.Println(sum, value, multiplyer)
		sum = sum + (value * multiplyer)

		fmt.Println(sum)
		multiplyer = multiplyer - 1

	}

	return sum%11 == 0
}
