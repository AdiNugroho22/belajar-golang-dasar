package main

import "fmt"

func main() {

	const firstname = "Adi"
	const lastname = "Nugroho"

	const (
		namadepan    = "Saya"
		namabelakang = "Ganteng"
	)
	//error
	//firstname = "Tidak Bisa Diubah"
	//lastname = "Tidak Bisa Diubah"
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(namadepan)
	fmt.Println(namabelakang)
}