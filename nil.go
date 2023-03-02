package main

import "fmt"

// nil just for interface, function, map, slice , pointer ,chanel

func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	}else {
		return map[string] string {
			"name":name,
		}
	}
}

func main() {

	var person map[string] string = NewMap("")

	if person == nil {
		fmt.Println("data kosong")
	}else{
		fmt.Println(person)
	}
	

	data := NewMap("Deril")

	fmt.Println(data)

}