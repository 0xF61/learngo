// Package weather provides a summary of conditions in a given location.
package weather

// CurrentCondition stores the weather report at a location.
var CurrentCondition string

// CurrentLocation stores the name of the city.
var CurrentLocation string

// Forecast returns a summary of current weather condition at a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
