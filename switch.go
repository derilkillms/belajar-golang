package main

import "fmt"

func main() {
	name := "Akbar"

	switch name {
	case "Deril":
		fmt.Println("Hello Deril")
		fmt.Println("Hello Deril")
	case "Akbar":
		fmt.Println("Hello Akbar")
		fmt.Println("Hello Akbar")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default :
		fmt.Println("nama sudah benar")
	}

}