package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {

	err := validate(n)
	if err != nil {
		return 0, err
	}

	i := 0

	for {
		if n == 1 {
			return i, nil
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		i++
	}
}

func validate(n int) error {
	if n <= 0 {
		return errors.New("invalid input number")
	}
	return nil
}
