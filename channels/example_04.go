package main

import (
	"fmt"
)

func Example04() {
	dataChan := make(chan int)

	go func() {

		for i := 0; i < 1e3; i++ {
			dataChan <- i
		}

		/*
		If we do not close the channel after our operation
		it will block and trigger a deadlock error.
		*/
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}
