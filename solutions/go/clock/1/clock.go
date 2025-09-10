package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

const minInDay = 60 * 24

func New(h, m int) Clock {

	clock := Clock{}
	minutes := 60*h + m

	if minutes < 0 {
		return clock.Subtract(minutes * -1)
	}

	return clock.Add(minutes)
}

func (c Clock) Add(m int) Clock {
	m = m % minInDay
	currentMinutes := (c.hour * 60) + c.minute
	res := currentMinutes + m
	res = res % minInDay
	c.hour = res / 60
	c.minute = res % 60
	return c
}

func (c Clock) Subtract(m int) Clock {
	m = m % minInDay
	currentMinutes := (c.hour * 60) + c.minute
	res := currentMinutes - m

	if res < 0 {
		res = minInDay - (res * -1)
	}

	c.hour = res / 60
	c.minute = res % 60
	return c

}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
