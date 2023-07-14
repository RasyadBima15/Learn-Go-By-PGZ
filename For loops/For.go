package main

import (
	"fmt"
)

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	//for with statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	//can use for slice
	slice := []string{"Rasyad", "Bima", "Satya"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	//or using range
	for index, name := range slice {
		fmt.Println("Index", index, "=", name)
	}

	// can use for map
	person := make(map[string]string)
	person["name"] = "Rasyad"
	person["title"] = "Programmer"

	for key, values := range person {
		fmt.Println("Key:", key, ", Value:", values)
	}
}
