package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	//This is the number -12.3
	return "This is the number " + fmt.Sprintf("%.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
	num, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}

	n, err := strconv.Atoi(num.Value())
	if err != nil {
		return 0
	}

	return n
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num, ok := fnb.(FancyNumber)
	var val float64

	if !ok {
		val = 0.0
	}

	val = float64(ExtractFancyNumber(num))

	return fmt.Sprintf("This is a fancy box containing the number %.1f", val)
}

func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
