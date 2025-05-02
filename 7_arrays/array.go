package main

import "fmt"

func main() {
	// Array Declaration
	var nums [4]int

	// Length of the Array
	fmt.Println("Length: ", len(nums))

	nums[1] = 4
	// Printing Array
	fmt.Println("Array: ", nums)

	// Declare and Initialize Array
	nums2 := [3]int{1, 2, 3}
	fmt.Println("Second Array: ", nums2)

	// 2D Array
	nums2D := [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println("2D Array: ", nums2D)

	// Compile-Time Size Determination
	nums3 := [...]int{1, 2, 3}
	fmt.Println(nums3)


	//* Arrays are only used when we know the size of the data structure
	//* For Dynamic Memory Allocation, we use SLICES

	//? advantages of arrays:
	//?	- fixed size, that is predictable
	//?	- memory optimization
	//?	- constant time access
}