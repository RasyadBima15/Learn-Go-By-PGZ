package main

import (
	"fmt"
)

func main() {
	var array1 [2]string
	array1[0] = "Rasyad"
	array1[1] = "Bimasatya"
	fmt.Println(array1)
	fmt.Println(array1[0])
	fmt.Println(array1[1])
	fmt.Println(len(array1))

	var nim = [2]int{
		24,
		1,
	}
	fmt.Println(nim)
	fmt.Println(nim[0])
	fmt.Println(nim[1])
	fmt.Println(len(nim))

	var more [10]string
	fmt.Println(len(more)) //the len still count even though there is no data in array
}
