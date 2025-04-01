package main

import "fmt"

const age = 30
// x := 10 //! not allowed
var s string = "nakul"


func main() {
	const name string = "golang"
	// name = "nakul" //! not allowed

	// MULTIPLE CONSTANT GROUPING
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
} 