package database

import "fmt"

var connection string

func init() {
	connection = "MySQL"
	fmt.Print("init di panggil")
}

func GetDatabase() string {
	return connection
}
