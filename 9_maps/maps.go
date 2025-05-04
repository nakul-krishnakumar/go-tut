package main

import (
	"fmt"
	"maps"
)

func main() {
	//? MAPS are similar to hashmap, objs, dicts

	// creating a map
	mpp := make(map[string]string)

			// or
	// mpp2 := map[string]int{}
	mpp3 := map[string]int{
		"price": 40,
		"quantity": 12,
	}
	fmt.Println(mpp3)

	// setting an element
	mpp["name"] = "nakul"
	mpp["place"] = "mangad"

	// getting an element
	fmt.Println(mpp["name"])
	//* NOTE: if key doesn't exists in the map, then it returns zero value

	// length of the map
	fmt.Println(len(mpp))

	// delete an element
	fmt.Println(mpp)
	delete(mpp, "name")
	fmt.Println(mpp)

	// empty the map
	clear(mpp)
	fmt.Println(mpp)

	// element exists or not
	val, flag := mpp3["price"]
	if flag {
		fmt.Println("exists with values", val)
	} else {
		fmt.Println("does not exists")
	}

	// equating two maps
	mpp4 := map[string]int{"price": 40, "quantity": 12}
	mpp5 := map[string]int{"price": 40, "quantity": 2}
	
	fmt.Println("Equal?", maps.Equal(mpp4, mpp5))
}