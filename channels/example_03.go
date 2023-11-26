package main

import "fmt"

/*
Buffered channel is like a queue with limited space
because only fit N data per time. e.g.: N = 2
*/
func Example03() {
	dataChan := make(chan int, 2)

	dataChan <- 123
	dataChan <- 692

	n := <-dataChan 
	fmt.Printf("n = %d\n", n)


	n = <-dataChan 
	fmt.Printf("n = %d\n", n)

	/*
	if our input is bigger then the limit
	then it will block and trigger a deadlock error.
	*/
	dataChan <- 782
	dataChan <- 912
	// dataChan <- 323 // uncomment to trigger deadlock

	n = <-dataChan 
	fmt.Printf("n = %d\n", n)
	n = <-dataChan 
	fmt.Printf("n = %d\n", n)
}
