package main

import (
	"fmt"
	"time"
)

//? Goroutines are used for concurrent processing (multithreading)

// func task(i int) {
// 	fmt.Println("doing task", i)
// }

func main() {
	for i:=0; i<=10; i++ {
		// go task(i)

		go func(i int) {
			fmt.Println(i) // here, i is accessible inside the func due to closure property
		}(i) // here, '()' means the func is being called when initialising itself
	}

	time.Sleep(time.Millisecond)
	// normally sleep is not used,
	//? instead waitgroups are used
}