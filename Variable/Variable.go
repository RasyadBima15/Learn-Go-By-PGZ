package main

import "fmt"

func main() {
	var name string
	name = "Rasyad Bimasatya"
	fmt.Println(name)
	//change the name
	name = "Rasyad"
	fmt.Println(name)

	//or can be declared like this

	var age = 18
	fmt.Println(age)
	age = 19
	fmt.Println(age)

	//or can be directly initialize the variable like this

	country := "Indonesia" //tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci :=
	fmt.Println(country)
	country = "Japan"
	fmt.Println(country)

	//can reclare multiple variables like this

	var (
		firstname = "Rasyad"    //directly use = can't use :=
		lastname  = "Bimasatya" //directly use = can't use :=
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

}
