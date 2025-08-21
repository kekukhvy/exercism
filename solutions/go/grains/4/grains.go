package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}
	return 1 << (number - 1), nil
}
func Total() uint64 {
	return math.MaxUint64
}
