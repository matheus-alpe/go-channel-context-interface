package main

import "fmt"

/*
Unbuffered channel by default is blocking,
because only fit one data per time.
*/
func Example01() {
	dataChan := make(chan int)

	/*
	Because this line is on the main thread
	it will block and cause a deadlock error.
	*/
	dataChan <- 123

	n := <-dataChan 

	fmt.Printf("n = %d\n", n)
}
