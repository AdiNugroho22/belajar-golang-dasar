package main

import "fmt"

func main()  {
	var a = 17
	var b = 10

	var penjumlahan = a + b
	var pengurangan = a - b
	var perkalian = a * b
	var pembagian = a / b
	var sisabagi = a % b

	fmt.Println(penjumlahan)
	fmt.Println(pengurangan)
	fmt.Println(perkalian)
	fmt.Println(pembagian)
	fmt.Println(sisabagi)

	var i = 20
	i +=20	
	fmt.Println(i)

	i += 5
	fmt.Println(i)
}