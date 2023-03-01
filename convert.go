package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 128
		nilai64 int64 = int64(nilai32)
		nilai8 int8 = int8(nilai32)

	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Deril"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println(name+" "+eString)
}