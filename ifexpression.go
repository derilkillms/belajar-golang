package main

import "fmt"

func main() {
	name := "Pranata"

	if name == "Deril" {
		fmt.Println("Hallo "+name)
	} else if name == "Pranata" {
		fmt.Println("Hallo "+name)
	} else {
		fmt.Println("Hi, Boleh Kenalan ?")
	}

	
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}