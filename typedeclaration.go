package main

import "fmt"

func main() {
	//alias type data
	type NoKTP string

	var NoKTPD NoKTP = "1231234123123"

	fmt.Println(NoKTPD)
}