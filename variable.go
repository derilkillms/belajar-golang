package main

import "fmt"

func main() {
	var name string

	name = "Deril"
	fmt.Println(name)

	name = "Muhammad Deril"
	fmt.Println(name)


	var lastname = "Pranata"
	fmt.Println(lastname)

	var age int8 = 27
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		kata1 = "Kata1"
		kata2 = "Kata2"
	)
	fmt.Println(kata1+" "+kata2, 3)
}