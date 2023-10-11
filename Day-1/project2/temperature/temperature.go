// temperature/temperature.go

package temperature

// FahrenheitToCelsius converts temperature from Fahrenheit to Celsius.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
