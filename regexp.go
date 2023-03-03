package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eka eko edo eto eyo eki", -1)
	fmt.Println(search)
}
