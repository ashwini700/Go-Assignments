package main

import (
	"fmt"
)

func divide(num, denom int) int { //panic if its zero as it cant divide
	if denom == 0 {
		panic("Can't divide by 0")
	}
	return num / denom
}

func safeDivide(num, denom int) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r) //why error string should not be capitalised (doubt)
		}
	}()

	res = divide(num, denom)
	return res, nil
}

func main() {

	res1, err1 := safeDivide(8, 2) // Test cases with different inputs for safeDivide
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Res 1:", res1)
	}

	res2, err2 := safeDivide(5, 0) //panic situations denom 0
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Res 2:", res2)
	}

	res3, err3 := safeDivide(4, 4)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Res 3:", res3)
	}
}
