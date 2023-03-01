package main

import "fmt"

func main() {
	name := "Muhammad"
	counter := 0
	increment := func() {
		name := "Deril"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	fmt.Println(counter)
	increment()
	increment()
	fmt.Println(name)
	fmt.Println(counter)
}
