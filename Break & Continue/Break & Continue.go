package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println("Iterasi ke-", i)
	}
	fmt.Println("==============================================")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Iterasi ke-", i)
	}
}
