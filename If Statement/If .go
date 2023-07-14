package main

import "fmt"

func main() {
	person := "Rasyad"
	if person == "Rasyad" {
		fmt.Println("Hii Rasyad")
	} else if person == "Joko" {
		fmt.Println("Hii Joko")
	} else {
		fmt.Println("Nama tidak cocok")
	}

	//if short statement
	if length := len(person); length > 5 {
		fmt.Println("Ini Adalah nama saya")
	} else {
		fmt.Println("Ini bukan nama saya")
	}
}
