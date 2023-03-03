package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	hostname, err := os.Hostname()
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])

	if err != nil {
		fmt.Println(nil, err.Error())
	} else {
		fmt.Println("Hostname: ", hostname)
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
