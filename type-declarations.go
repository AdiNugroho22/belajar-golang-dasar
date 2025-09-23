package main

import "fmt"

func main (){

	type NoKTP string	

	var ktpAdi NoKTP = "1111111"
	var contoh string = "222222222"
	var contohKTP NoKTP = NoKTP(contoh)
	fmt.Println(ktpAdi)
	fmt.Println(contohKTP)
}