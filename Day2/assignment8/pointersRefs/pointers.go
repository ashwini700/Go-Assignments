package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func IncrAge(s *Student) {
	s.Age++
}

func main() {
	student := Student{
		Name: "Ash",
		Age:  23,
	}

	fmt.Printf("beginning age of student %s: %d\n", student.Name, student.Age)

	IncrAge(&student)

	fmt.Printf("Updated age of student %s: %d\n", student.Name, student.Age) //updated  age
}
