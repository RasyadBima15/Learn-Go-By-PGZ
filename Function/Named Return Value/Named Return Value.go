package main

import "fmt"

func Fullname() (firstname, lastname string) {
	firstname = "Rasyad"
	lastname = "Bimasatya"
	return //just type return because the variables have mentioned in return value function before
}

func main() {
	firstname, lastname := Fullname()
	fmt.Println(firstname, lastname)
}
