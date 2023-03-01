package main

import "fmt"

func main() {

	person := map[string]string{
		"name": "Deril",
		"address": "Ciamis",
	}

	person["title"] = "Programmer"


	fmt.Println(person)
	fmt.Println(person["name"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(len(book))
}