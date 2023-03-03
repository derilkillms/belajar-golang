package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Deril", "Deril"))
	fmt.Println(strings.Split("Muhammad Deril", " "))
	fmt.Println(strings.ToLower("Deril"))
	fmt.Println(strings.ToUpper("Deril"))
	fmt.Println(strings.ToTitle("deril"))
	fmt.Println(strings.Trim("   Muhammad Deril   ", " "))
	fmt.Println(strings.ReplaceAll("Deril Deril Deril", "Deril", "Dede"))

}
