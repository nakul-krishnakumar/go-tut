package main

import "fmt"

func main() {

	// while loop implementation
	i := 1
	for i <= 3  {
		fmt.Println(i)
		i += 1
	}

	// for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// for loop using range
	for i := range 3 {
		fmt.Println(i , " in range")
	}
}