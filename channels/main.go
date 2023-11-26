package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}

		for i := 0; i < 1e3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}

		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}
