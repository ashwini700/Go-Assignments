package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "Bou, Bou"
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow, Meow"
}

func main() {
	dog := Dog{}
	cat := Cat{}
	fmt.Printf("Dog says: %s\n", dog.MakeSound())

	fmt.Printf("Cat says: %s\n", cat.MakeSound())
}
