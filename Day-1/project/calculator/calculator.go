// calculator/calculator.go

package calculator

// Square takes an integer number and calculates the square of the number.
func Square(num int) int {
	return num * num
}

// Sum takes two integers as input and calculates their sum.
func Sum(a, b int) int {
	return a + b
}

// Difference takes two integers as input and calculates their difference.
func Difference(a, b int) int {
	return a - b
}

// Product takes two integers as input and calculates their product.
func Product(a, b int) int {
	return a * b
}

// QuotientAndRemainder takes two integers as input and calculates their quotient and remainder.
func QuotientAndRemainder(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
