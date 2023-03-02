package main

import "fmt"

// menggabungkan data berbeda type

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var data Customer
	data.Name = "Deril"
	data.Address = "Ciamis"
	data.Age = 27

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}

	fmt.Println(joko)
	fmt.Println(data)

	budi := Customer{"budi", "jakarta", 32}
	fmt.Println(budi)

}
