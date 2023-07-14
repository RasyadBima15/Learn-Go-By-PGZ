package main

import "fmt"

func main() {
	const firstname = "Rasyad"
	const lastname = "Bimasatya"
	const value = 1000

	// firstname = "Budi" //Can't assign to firstname
	// lastname = "laksana" //Can't assign to firstname
	// value = 500 //Can't assign to firstname

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

	//declare multiple constant

	const (
		country = "Indonesia"
		city    = "Makassar"
		age     = 18
	)
	fmt.Println(country)
	fmt.Println(city)
	fmt.Println(age)
}
