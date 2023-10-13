// Function with Variadic Parameters
// Description: Create a program that defines a function with variadic parameters and uses it.
// Instructions:
// Create a function called sum that takes an arbitrary number of integers and returns their sum.
// Use this function to calculate and print the sum of different sets of numbers.

package main

import "fmt"

// arbitrary nums
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {

	res1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of set 1: %d\n", res1) //diff set of sum of nums

	res2 := sum(10, 20, 30)
	fmt.Printf("Sum of set 2: %d\n", res2)

	res3 := sum(3, 5, 7, 9, 10)
	fmt.Printf("Sum of set 3: %d\n", res3)
}
