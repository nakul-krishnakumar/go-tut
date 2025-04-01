package main

import "fmt"

func main() {
	var name string = "golang"
	// var name = "golang" -> also works but explicit type assigning is prefered, here automatic type infering is done

	//shorthand syntax
	age := 18

	var place string
	place = "Kannur"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(place)
}