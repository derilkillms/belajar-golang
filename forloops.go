package main

import "fmt"

func main() {

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// for i := 1; i < 10; i++ {
	// 	fmt.Println("Perulangan ke", i)
	// }

	slice := []string{"Muhammad","deril","akbar","pranata"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println("index",index,"=",name)
	}

	person := make(map[string]string)
	person["name"] = "Muhammad"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key,"=",value)
	}
}