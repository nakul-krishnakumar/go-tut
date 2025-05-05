package main

import "fmt"

//? variadic function
func add(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(add(10, 20))

	nums := []int {10, 20, 30, 40, 50}
	fmt.Println(add(nums...)) //spread operator, converts slice to int
	//* fmt.Println(add(nums)) This will raise error as we are passing slice instead of int
}