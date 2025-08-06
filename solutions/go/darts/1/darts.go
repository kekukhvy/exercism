package darts

import "math"

func Score(x, y float64) int {
	c := math.Sqrt(x*x + y*y)

	switch {
	case c > 10:
		return 0
	case c > 5:
		return 1
	case c > 1:
		return 5
	case c >= 0:
		return 10
	}

	return -1
}
