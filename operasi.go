package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 10
		c = a * b
	)
	fmt.Println(c)
// augmented a += 10 
var i = 10
i += 10
fmt.Println(i)

//unary operation
i++
fmt.Println(i)

//operasi perbandingan > < >= <= == !=

var (
	nama1 = "Deril"
	nama2 = "Eril"
	result2 bool = nama1 >= nama2
)
fmt.Println(result2)

}