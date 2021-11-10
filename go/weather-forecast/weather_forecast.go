// Package weather station manage the weather forecasting program.
package weather

// CurrentCondition is a string holds the value of current condition of the weather.
var CurrentCondition string

// CurrentLocation is the current location of the weather.
var CurrentLocation string

// Forecast is a function to receive the arguments city and condition, and returns a string of a current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
