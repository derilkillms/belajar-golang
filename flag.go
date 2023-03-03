package main

import (
	"flag"
	"fmt"
)

// app basis terminal
func main() {
	host := flag.String("host", "localhost", "Put your database host")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 100, "put your number")

	flag.Parse()

	fmt.Println("Host	:", *host)
	fmt.Println("User	:", *user)
	fmt.Println("Password :", *password)
	fmt.Println("Numbaer", *number)
}
