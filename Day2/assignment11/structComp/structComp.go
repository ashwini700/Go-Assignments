package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Person struct { //embedded addr
	Name string
	Address
}

func main() {
	person := Person{
		Name: "Ash",
		Address: Address{
			Street:  "Dare devil zone",
			City:    "kormangala",
			ZipCode: "560030",
		},
	}

	fmt.Printf("Person: %s\n", person.Name)
	fmt.Printf("Address: %s, %s, %s\n", person.Street, person.City, person.ZipCode)
}
