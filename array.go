package main

import "fmt"


func main()  {
	var names [4]string
	names[0] = "Adi"
	names[1] = "Nugroho"
	names[2] = "Abadi"
	names[3] = "Sejahtera"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [...]int {
		90,
		80,
		17,
		10,
		20,
		30,
		40,
		50,
		60,
		70,
		80,
		90,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
 