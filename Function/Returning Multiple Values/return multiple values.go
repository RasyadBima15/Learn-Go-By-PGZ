package main

import "fmt"

func getFullName() (string, string) {
	return "Rasyad", "Bimasatya"
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)
	fmt.Println(firstname)
	fmt.Println(lastname)
	//or can be just like this
	fmt.Println(getFullName())

	_, lstname := getFullName() // _ meeans we don't care the value (we just declared value as a variable as we want (We don't declare all the returning values in function))
	fmt.Println(lstname)
}
