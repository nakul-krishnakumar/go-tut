package main

import (
	"fmt"
	"sync"
)

//? MUTEX IS USED TO AVOID RACE CONDITIONS

type post struct {
	likes int
	views int
	mu sync.Mutex
}

func (p *post) incLikes(wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
		p.mu.Unlock()
	} ()

	p.mu.Lock()
	p.likes += 1
}

func (p *post) incViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	} ()

	p.views += 1
}

func main() {
	post1 := post{
		views: 0,
		likes: 0,
	}

	var wg sync.WaitGroup
	for i:=0; i<100; i++ {
		wg.Add(1)
		go post1.incLikes(&wg)
	}

	for i:=0; i<100; i++ {
		wg.Add(1)
		go post1.incViews(&wg)
	}

	wg.Wait()
	fmt.Println("likes:", post1.likes)
	fmt.Println("views:", post1.views)

	//! TRY RUNNING THIS PROGRAM QUICKLY MULTIPLE TIMES
	//? you will notice that view count may not always be outputted as 100
	//? this is due to RACE CONDITIONS
	//? like count will always be 100 as it uses MUTEX LOCKS which prevents RACE conditions
	//? But the DRAWBACK of MUTEX is that when one goroutine acquires a mutex lock (with mu.Lock()), any other goroutines that try to lock it are blocked (i.e., paused) until the first one unlocks it.
	//? But this drawback can be ignored as it avoiding RACE CONDITIONS is the priority
}