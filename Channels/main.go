package main

import (
	"fmt"
)

/*
Channels: Go channels are typed communication pipes that allow concurrent goroutines to safely send and receive data. They eliminate the need for explicit locks or mutexes. Data flows in the direction of the arrow operator (<-), and operations block until both sides are ready

- In Go, channels are like a conveyor belt or a pipe that connects different tasks (goroutines) so they can safely pass data back and forth.
- Channels are (sync) means blocking in nature.
- buffered channels are non-blocking in nature.
- Non buffered channels are blocking in nature.



*/

// receiving
func sum(result chan int, num1 int, num2 int) {
	ans := num1 + num2

	result <- ans
}

// goroutine synchronizer
func task(done chan bool) {
	// defer func() { // defer just like a finally block in js, inside this always execute.
	// 	done <- true
	// }()

	done <- true // send

	fmt.Println("Processing...")

}

// emailChan only received krega means(<-) or for done send only.
func emailSender(emailChan <-chan string, done chan<- bool) {
	// defer func() {
	// 	done <- true
	// }()

	// emailChan <- "hello" // show error

	for email := range emailChan {
		fmt.Println("Sending email to", email)
		// time.Sleep(time.Millisecond * 200)
	}

	done <- true
}

func main() {
	// // 1. channel creation using make method - unbuffered channel
	// messageChan := make(chan string)

	// go func() { // sending goroutine
	// 	// time.Sleep(time.Second * 5)

	// 	// Send string into the channel
	// 	messageChan <- "hello go channel!"

	// }()

	// fmt.Println("hii!!")

	// // Receive a string from the channel
	// msg := <-messageChan

	// fmt.Println(msg)

	// // 2. send and received after processing
	// result := make(chan int)

	// go sum(result, 4, 6)

	// res := <-result

	// fmt.Println("Sum is:", res)

	// done := make(chan bool)
	// go task(done)

	// // receiving
	// res := <-done // this is blocking

	// fmt.Println(res)

	// 3. Buffered channel - this is non-blocking
	// implement queue system

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// close(emailChan)

	// fmt.Println("Done sending...")

	// res := <-done
	// fmt.Println("Processing done...", res)

	// multi-channel handle
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

}
