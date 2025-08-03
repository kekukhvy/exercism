// Package weather provides a simple weather forecast functionality.
package weather

// CurrentCondition holds the current weather condition for a given city.
var CurrentCondition string

// CurrentLocation holds the name of the city for which the weather is forecasted.
var CurrentLocation string

// Forecast returns a string with the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
