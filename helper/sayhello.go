package helper

import "fmt"

var version = 1                    // tidak bisa diakses package luar
var Application = "Belajar golang" //bisa diakses package luar

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
