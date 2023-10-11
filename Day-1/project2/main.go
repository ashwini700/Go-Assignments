// main.go

package main

import (
	"fmt"
	"go-training/Day-1/project2/temperature"
)

func main() {
	var fahrenheitInput float64

	// Prompt the user for input
	fmt.Print("Enter temperature in Fahrenheit: ")
	_, err := fmt.Scanf("%f", &fahrenheitInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert Fahrenheit to Celsius using the temperature package
	celsiusResult := temperature.FahrenheitToCelsius(fahrenheitInput)

	// Print the converted temperature
	fmt.Printf("%.2f°F is %.2f°C\n", fahrenheitInput, celsiusResult)
}
