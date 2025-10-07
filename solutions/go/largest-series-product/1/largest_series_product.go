package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	len := len(digits)

	if span > len || span < 0 {
		return 0, errors.New("span must be smaller than string length")
	}

	var max int64
	for i := range digits {
		j := i + span

		if j > len {
			break
		}

		res, err := calc(digits[i:j])

		if err != nil {
			return 0, err
		}

		if res > max {
			max = res
		}
	}
	return max, nil
}

func calc(digits string) (int64, error) {
	var res int64 = 1
	for _, v := range digits {
		if !unicode.IsDigit(v) {
			return 0, errors.New("must be digit")
		}
		res *= int64(v - '0')
	}
	return res, nil
}
