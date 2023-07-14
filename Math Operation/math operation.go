package main

import "fmt"

func main() {
	var addition = 10 + 10
	fmt.Println(addition)

	var a = 10
	var b = 20
	var result = a + b
	fmt.Println(result)

	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var positive = +10
	var negative = -10
	fmt.Println(positive)
	fmt.Println(negative)

	// var x float32 = 10
	// var y float32 = 5.5

	var (
		x float32 = 10
		y float32 = 5.5
	)
	var divide = x / y
	fmt.Println(divide)
}
