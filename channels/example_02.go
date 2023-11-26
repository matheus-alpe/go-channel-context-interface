package main

import "fmt"

/*
Unbuffered channel by default is blocking,
because only fit one data per time.
*/
func Example02() {
	dataChan := make(chan int)

	/*
	Because this line is on the main thread
	it will block and cause a deadlock error.
	To resolve this problem, we need to execute
	this code inside of a goroutine.
	*/
	go func() {
		dataChan <- 123
	}()

	n := <-dataChan 

	fmt.Printf("n = %d\n", n)
}
