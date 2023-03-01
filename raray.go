package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Muhammad"
	name[1] = "Deril"
	name[2] = "Akbar"

	fmt.Println(name[1])

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

	fmt.Println(len(values))
}