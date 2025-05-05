package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()

	fmt.Println(increment()) // -> 1
	fmt.Println(increment()) // -> 2
	fmt.Println(increment()) // -> 3
	fmt.Println(increment()) // -> 4
	fmt.Println(increment()) // -> 5

	//? NOTE THAT THE FUNCTION DOES NOT FORGET THE PREVIOUS VALUE OF 'count'
	//? THIS IS POSSIBLE BECAUSE OF CLOSURE PROPERTY
	//? it is also possible in javascript
}