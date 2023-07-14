package main

import "fmt"

func main() {
	var (
		number1 int32 = 100000
		number2 int64 = int64(number1)
		number3 int8  = int8(number2) //can't convert to int8 because exceeding the maximum limit of int8
	)
	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3) //can't convert to int8 because exceeding the maximum limit of int8

	var (
		name    = "Rasyad"
		e       = name[0]   //a byte code (Data Type uint8 )
		estring = string(e) //converter
	)

	var (
		nameBro = "Rayyan"
		a       = nameBro[2]
		astring = string(a)
	)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(estring)

	fmt.Println(nameBro)
	fmt.Println(a)
	fmt.Println(astring)
}
