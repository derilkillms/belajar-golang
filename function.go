package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func sayHelloTo(firstname string, lastname string){
	fmt.Println("Hello",firstname,lastname)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullname() (string, string, string) {
	return "Muhammad", "-", "Deril"
}

func getFullnamed() (firstname string,middlename string,lastname string) {
	firstname = "Muhammad"
	middlename = "-"
	lastname = "Deril"
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodbye(name string) string {
	return "Good bye " + name
}

func main() {

	fmt.Println(getHello("Deril"))

	firstname, _, lastname := getFullnamed()

	fmt.Println(firstname, lastname)

	total := sumAll(10, 10, 11)

	fmt.Println(total)

	slice := []int{10,10,10,10}
	total = sumAll(slice...)
	fmt.Println(total)

	//make variable from func
	sayGoodBye := getGoodbye
	fmt.Println(sayGoodBye("Deril"))


}