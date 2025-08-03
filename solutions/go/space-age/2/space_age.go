package space

import (
	"math"
)

type Planet string

var planetsMap = map[Planet]float64{
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
	"Mercury": 0.2408467,
}

func Age(seconds float64, planet Planet) float64 {

	factor, ok := planetsMap[planet]
	if !ok {
		return -1
	}
	age := seconds / 31557600 / factor
	return math.Round(age*100) / 100
}
