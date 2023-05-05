// Package weather provide some weather inforamtion.
package weather

// CurrentCondition is a string that store the current weather condition of CurrentLocation.
var CurrentCondition string

// CurrentLocation is a string that store the current city that we are trying to get the weather.
var CurrentLocation string

// Forecast function will return a sentence that describes the weather conditions at a certain location (city).
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
