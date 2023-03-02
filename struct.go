package main

import "fmt"

// menggabungkan data berbeda type

//struct
type Customer struct {
	Name, Address string
	Age           int
}

//struct method
func (customer Customer) sayHi( name string){
	fmt.Println("Hello", name, "my name is",customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuu from",a.Name)
}

func main() {
	var data Customer
	data.Name = "Deril"
	data.Address = "Ciamis"
	data.Age = 27

	data.sayHi("Muhammad")
	data.sayHuuu()

	// joko := Customer{
		// 	Name:    "Joko",
		// 	Address: "Cirebon",
		// 	Age:     35,
		// }

		// fmt.Println(joko)
		// fmt.Println(data)

		// budi := Customer{"budi", "jakarta", 32}
		// fmt.Println(budi)

	}
