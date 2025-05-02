package main

import "fmt"

//? Slices are dynamic arrays.
//? They are the most used construct in GO.
//? Has many useful functions
func main() {
	// uninitialized slice is nil
	var nums []int
	// if nums == nil {
	// 	fmt.Println("Array is empty")
	// }

	fmt.Println("Nums:", nums)

	// dynamic memory allocation with initialized values
	num2 := make([]int, 2, 5)
	// here 5 is the capacity of the slice
	//? capacity -> max num of elements that can fit in it.
	// if you add elements more than capacity, 
	// the slice doubles its capacity on its own.
	fmt.Println("Nums 2:", num2)
	fmt.Println("Cap of Nums 2:", cap(num2))
	
	num2 = append(num2, 4)
	num2 = append(num2, 5)
	num2 = append(num2, 6)
	num2 = append(num2, 7)
	fmt.Println("Nums 2:", num2)
	fmt.Println("Cap of Nums 2:", cap(num2))
	//* notice that the cap doubles from 5 to 10

	//* -------- ANOTHER WAY TO DECLARE SLICES --------
	num3 := []int{}
	fmt.Println("Num 3:", num3)
	fmt.Printf("Type of num3: %T\n", num3)

	//? copy function
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	fmt.Println(nums, numsCopy)

	//? slice operator
	nums4 := []int{1, 2, 3 ,4 ,5}
	fmt.Println(nums4[2: 5])
	//! GO does NOT support negative indexing
}	