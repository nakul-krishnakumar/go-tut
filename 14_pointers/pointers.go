package main

import "fmt"

func changeNum(num *int) {
	*num = 4
	fmt.Println("Inside func: num =", *num)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println("Outside func: num =", num)
}