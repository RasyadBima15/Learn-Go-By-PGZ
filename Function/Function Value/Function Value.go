package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name + "!"
}
func main() {
	hello := getHello
	fmt.Println(hello("Rasyad")) // the variable changed into function
}
