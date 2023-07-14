package main

import "fmt"

func SumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sum := SumAll(10, 10, 10, 10, 10, 10, 10) //this can be empty
	fmt.Println(sum)

	slice := []int{10, 20, 30, 40, 50} //variable arguments as a slice
	sum = SumAll(slice...)             //don't forger to type ... to put the slice in variadic function
	fmt.Println(sum)
}
