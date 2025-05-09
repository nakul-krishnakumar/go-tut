package main

import "fmt"

/*
//* - Generics in Go allow you to write functions, types, and data structures that work with any data type, while maintaining type safety.
//* - With generics, you define a type parameter — like T — that acts as a placeholder for actual type
*/

// GENERICS IN FUNCTIONS
//? Here, T is the type parameter (generic type)
func printSlice[T comparable](items []T) {
	for _, item := range items{
		fmt.Println(item)
	}
}
//* COMPARABLE -> interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). 

//* func printSlice[T string | int](items []T) {

//! Multiple generics are also possible:
//? func printSlice[T comparable | V int](items []T, nums []V) {

// GENERICS IN STRUCTS
type stack[T string | int] struct {
	elements []T
}

func main() {
	// GENERICS IN FUNCTIONS
	nums := []int{1, 2, 3, 4, 5}
	langs := []string{"a", "b", "c", "d"}
	printSlice(nums)
	printSlice(langs)
	// -----------------

	// GENERICS IN STRUCTS
	st1 := stack[int]{
		elements: nums,
	}
	fmt.Println("Integer stack:", st1)

	st2 := stack[string]{
		elements: langs,
	}
	fmt.Println("Integer stack:", st2)
}