// Package weather shows us what the climate condition of the place is like.
package weather

// CurrentCondition represents the current climate condition of where are we.
var CurrentCondition string

// CurrentLocation represents the current location of where are we.
var CurrentLocation string

// Forecast returns the current location and weather conditions of the city you sent via the parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
