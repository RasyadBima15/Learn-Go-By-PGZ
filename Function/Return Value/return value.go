package main

import "fmt"

func sayHelloTo(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := sayHelloTo("Rasyad")
	fmt.Println(result)
	fmt.Println(sayHelloTo(""))
}
