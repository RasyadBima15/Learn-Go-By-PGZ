package main

import "fmt"

func main() {
	name := "Rasyyyyyyyyyyyyyyyyyyyy"
	switch name {
	case "Rasyad":
		fmt.Println("Hello Rasyad")
	case "Rasyy":
		fmt.Println("Hello Rasyy")
	default:
		fmt.Println("Bukan nama saya")
	}

	//switch short hand
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama lumayan panjang")
	// case false:
	// 	fmt.Println("Nama pendek")
	// }

	//switch without a condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama pendek")
	}

}
