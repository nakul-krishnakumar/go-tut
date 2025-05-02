package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	} 

	// declaring variables on the if construct
	if id := 5; id >= 10 {
		fmt.Println("High ID", id)
	} else {
		fmt.Println("Low ID", id)	
	}
}

// GO DOES NOT HAVE TERNARY OPERATOR,
// INSTEAD WE USE NORMAL IF ELSE
