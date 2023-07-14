package main

import "fmt"

func main() {
	var months = [...]string{ // ... mean that we don't know the len of array so we just type ...
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah" //this can be changed
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi" //this can be changed
	fmt.Println(months)

	//append slice
	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}
	daysSlice1 := days[5:]
	daysSlice1[0] = "sabtu baru"
	daysSlice1[1] = "minggu baru"
	fmt.Println(days)
	daysSlice2 := append(daysSlice1, "Libur Baru")
	daysSlice2[0] = "Ups"
	fmt.Println(daysSlice2) //make the new array because the capasity exceeding the maximum limit
	fmt.Println(days)       //daysSlice2 append doesn't affect to the days array because daysSlice2 make the new array

	//make slice

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rasyad Bimasatya"
	newSlice[1] = "Rayyan Adisatya"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice)) // have to adjust to length and cap of newSlice so the data is not truncated
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//Array vs Slice
	thisIsArray := [...]int{1, 2, 3, 4, 5} //this is array
	thisIsSlice := []int{1, 2, 3, 4, 5}    //this is slice
	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
