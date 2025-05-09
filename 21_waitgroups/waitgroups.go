package main

import (
	"fmt"
	"sync"
)

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done() // defer is used to run it just before the function ends
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i:=0; i<=10; i++{
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}

//? wg.Add(1) -> it adds one to the waitgroup for each goroutine
//? wg.Done() -> it decreases one from the waitgroup when each goroutine is finished
//? wg.Wait() -> it waits for all workers to finish