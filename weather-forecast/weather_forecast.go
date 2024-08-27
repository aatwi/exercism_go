// Package weather is a class that holds weather information.
package weather

// CurrentCondition is the current value of the weather condition.
var CurrentCondition string

// CurrentLocation is the current value of the location.
var CurrentLocation string

// Forecast is a function that returns a string based on the city and condition passed by the user.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
