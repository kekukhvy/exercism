package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
var colorsMap = map[string]int{

	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Value(colors []string) int {

	if len(colors) == 1 {
		return colorsMap[colors[0]]
	}

	return (colorsMap[colors[0]] * 10) + colorsMap[colors[1]]
}
