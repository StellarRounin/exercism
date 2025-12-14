/* 
Package weather provides the forecast for any city and condition.
*/
package weather

var (
    // CurrentCondition represents the actual weather condition.
	CurrentCondition string
    // CurrentLocation represents the actual weather location.
	CurrentLocation  string
)

// Forecast takes a city and condition and returns the the current weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
