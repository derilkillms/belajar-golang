package main

import "fmt"

//secara default go mengirim variable dengan pass by value (menduplikat)

//pointer menggunakan pass by reference

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address)  {
	address.Country = "Indonesia"
	
}

func main(){

	var address1 Address = Address{"Ciamis", "Jawa Barat", "Indonesia"}

	var address2 *Address = &address1 //pointer with &
	var address3 *Address = &address1

	address2.City = "Bandung"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	*address2 = Address{"Jakarta", "Jakarta DKI", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address {
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}