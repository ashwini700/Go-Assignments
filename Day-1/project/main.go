// main.go

package main

import (
	"fmt"
	"go-training/Day-1/project/calculator"
)

func main() {
	// Call functions
	num := 5
	squareResult := calculator.Square(num)
	fmt.Printf("Square of %d: %d\n", num, squareResult)

	sumResult := calculator.Sum(10, 5)
	fmt.Printf("Sum: %d\n", sumResult)

	diffResult := calculator.Difference(10, 5)
	fmt.Printf("Difference: %d\n", diffResult)

	productResult := calculator.Product(8, 3)
	fmt.Printf("Product: %d\n", productResult)

	quotient, remainder := calculator.QuotientAndRemainder(15, 4)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
