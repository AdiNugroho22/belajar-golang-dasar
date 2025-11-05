package main

import (
	"fmt"

)

func main() {
	names := [...]string{
		"John",
		"Paul",
		"George",
		"Ringo",
		"adi",
		"Nugroho",
		"Rahmat",	
		"Rahmat",
		"Rahmat",
	}		
	slice := names[4:6]	
	slice2 := names[:3]
	slice3 := names[3:]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice2)
	fmt.Println(slice3)


	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
}
	daySlice1 := days[5:]	
	daySlice1[0] = "Sabtu baru"	
	daySlice1[1] = "Minggu baru"
	fmt.Println(days)// [senin selasa rabu kamis jumat sabtu minggu]

	daySlice2 := append(daySlice1, "Libur baru")

	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice2)
	fmt.Println(days)


	var newSlice []string = make([]string, 2, 5)
	// make([]T, length, capacity)
	newSlice[0] = "Adi"	
	newSlice[1] = "Adi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Adi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Nugroho"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice :=make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	} 