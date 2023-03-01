package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string
func sayHellowWithFilter(name string, filter Filter) {
	namefiltered := filter(name)
	fmt.Println("Hello", namefiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}else{
		return name
	}
}

func upperFilter(name string) string {
	strconv := strings.ToUpper(name)
	return strconv
}

func main() {
	sayHellowWithFilter("Deril", spamFilter)
	sayHellowWithFilter("Anjing", spamFilter)
	sayHellowWithFilter("Deril", upperFilter)
}