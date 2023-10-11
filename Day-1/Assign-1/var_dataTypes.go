package main

import "fmt"

func main() {
	//Declare a variable of type int and assign it a value of 42. Print the variable.
	// i := 42
	// fmt.Println(i)

	//Declare a variable of type float64 and assign it a value of 3.14. Print the variable
	// var a float64 = 3.14
	// fmt.Println(a)

	//Create a constant called pi and assign it the value of Pi (math.Pi). Print the constant.
	// const pi = math.Pi
	// fmt.Println(pi)

	//create a group of variables using var () variable declaration syntax, assign it some values & print it.
	var (
		a = 20
		b = true
		c = "Hello"
	)

	fmt.Println(a, b, c)

	//Create a group of constants using var() group constant declaration syntax, assign it some values & print it.
	const (
		c1 = 10
		c2 = "me"
		c3 = false
		// c1 = 11
	)
	fmt.Println(c1, c2, c3)

}
