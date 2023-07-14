package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Rasyad",
		"adress": "Makassar",
	}
	person["Title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["adress"])

	//make map
	book := make(map[string]string) //make a new map
	book["Title"] = "Buku Golang"
	book["Author"] = "Rasyad Bimasatya"
	book["wrong"] = "ups"
	fmt.Println(book)
	delete(book, "wrong") //delete one of the key in the map
	fmt.Println(book)
	fmt.Println(len(book))

}
