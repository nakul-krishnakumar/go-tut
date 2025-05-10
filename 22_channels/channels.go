package main

import (
	"fmt"
	// "math/rand"
	"time"
)

//? CHANNELS ARE USED FOR COMMUNICATION BETWEEN CONCURRENT GOROUTINES
//? THEY ALLOW TO SAFELY PASS DATA WITHOUT USING SHARED MEMORY

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

func readNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("reading number", num)
		time.Sleep(time.Second)
	}
}

//! ---------------------------------
// WAITGROUPS USING CHANNELS
func task(done chan bool) {
	defer func(){ done <- true }()
	fmt.Println("doing concurrent process...")
}

//! ---------------------------------

//* IMPORTANT CONCEPTS
//? emailChan <-chan string => means that the channel is recieve-only
//? done chan<- bool        => means that the channel is send-only
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func(){ done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	//! UNBUFFERED CHANNELS
	//! ------------------------------
	//* WAITGROUPS USING CHANNELS
	//? When to use channel or waitgroups for synchronous tasks?
	//? -- Use channels when there is only a single goroutine
	//? -- Use waitgroups when there are multiple goroutines

	// done := make(chan bool)
	
	// go task(done)
	
	// <-done
	//! ------------------------------

	// numChan := make(chan int) //* unbuffered channel

	// go processNum(numChan) // goroutine with a channel setup
	// go readNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100) // this is a blocking process (synchronous)
	// 	// inifinetly sending data through the numChan channel
	// }
	//! ------------------------------
	
	// emailChanUN := make(chan string) //* unbuffered channel

	// emailChanUN <- "1@example.com"
	// emailChanUN <- "2@example.com"

	// fmt.Println(<-emailChanUN)
	// fmt.Println(<-emailChanUN)
	//? THIS WILL CAUSE DEADLOCK AS FOR UNBUFFERED CHANNELS, BOTH SENDER AND RECIEVER SHOULD BE RUNNING PARALLELY (AT THE SAME TIME), THAT IS, GOROUTINES SHOULD BE USED

	//! BUFFERED CHANNELS
	// emailChan := make(chan string, 100) //* buffered channel

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	//? THIS WONT CAUSE DEADLOCK AS FOR BUFFERED CHANNELS, WE CAN KEEP DATA IN BUFFER AND RECIEVER CAN RECIEVE IT LATER

	emailChan := make(chan string, 5) // buffered channel
	done := make(chan bool)	// unbuffered channel (used as synchronizer)

	go emailSender(emailChan, done)

	for i:=0; i<=5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending!") //* this will be shown as soon as everything is send into the emailChan , it will be very very quick as it is a buffered channel

	close(emailChan) //? Here, if this buffered channel is not closed, it will lead to deadlock as the 'done' synchronizer channel will be waiting infinitely if not closed, even though everything is send and recieved
	<- done
	

	//! ------------------------------
	//? Both accessing and uploading to an UNBUFFERED CHANNEL are blocking processes
	//? -- newChan := make(chan <data_type>)
	//? -- As buffered channels are blocking, they tend to be slow while sending multiple data

	//* So for multiple data handling, BUFFERED CHANNELS are used
	//* -- In BUFFERED CHANNELS, we can send a limited amount of data without blocking
}