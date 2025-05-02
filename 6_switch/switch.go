package main

import (
	"fmt"
	"time"
)

func main() {
	age := 4
	switch age {
	case 1:
		fmt.Println("Age is 1")
	case 2:
		fmt.Println("Age is 2")
	case 3:
		fmt.Println("Age is 3")
	case 4:
		fmt.Println("Age is 4")
		fallthrough // opposite of break
	case 5:
		fmt.Println("Age is 5")
	default:
		fmt.Println("Age is > 5")
	}

	// In Go, you don't need to add break statements at the end of 
	// each case in a switch statement. 
	// Go's switch statements automatically break after each case by default.

	// If you actually want to continue execution to the next case 
	// (similar to not having a break in C++), you need to 
	// explicitly use the fallthrough keyword


	//* Go explicitly supports comma-separated values in case statements
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Holiday!")
	case time.Friday:
		fmt.Println("Friday")
		fallthrough
	default:
		fmt.Println("Workday")
	}

	//* type switch
	getMyType := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other")
		}
	}

	getMyType("Hello")

	// func(i interface{})
	// func(i any)
	// interface{} is Go's way of representing a type that can hold 
	// any value - similar to a "void pointer" in C or Object in Java,
	//  but safer and with type checking at runtime. It's Go's 
	// implementation of a universal type.

	// any: A type alias for interface{} introduced in Go 1.18
}