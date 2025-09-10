package romannumerals

import "errors"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {

	if input > 3999 || input <= 0 {
		return "", errors.New("wrong value")
	}

	var result = ""
	for input > 0 {
		for _, v := range romanNumerals {
			mod := input % v.Value

			if mod == input {
				continue
			} else {
				result += v.Symbol
				input -= v.Value
				break
			}
		}
	}
	return result, nil
}
