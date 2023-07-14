package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	var NoKTP_Rasyad NoKTP = "27830720346081346801"
	var MarriedStatus Married = false
	fmt.Println(NoKTP_Rasyad)
	fmt.Println(NoKTP("aaaaaa"))
	fmt.Println(MarriedStatus)
}
