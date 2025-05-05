package main

import "fmt"

func add(a, b int) int {
	return a + b
}

//? func taking func as param
// <fn_name> <fn_head> <return_type>
func processIt(fn func(a int) int, b int) int {
	result := fn(b)
	return result
}

//? func returning a func
//                  (<func_name> <return_type>) <- this is what 'returnIt' func returns
func returnIt(b int) func(a int) int {
	return func(a int) int {
		return a + b
	}
}

func main() {
	sum := add(5, 10)
	fmt.Println("Sum:", sum)

	fn := func(a int) int {
		fmt.Println("Entered function fn", a)
		return a + 100
	}
	fmt.Println(processIt(fn, 5))

	b := returnIt(5)
	fmt.Println("Sum:", b(40))
}