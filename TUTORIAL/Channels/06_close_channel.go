package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// Create channel  named chn
	chn := make(chan int) // Here , chan is keyword used to create channel
	wg.Add(2)
	go func() {
		for i := range chn {
			fmt.Println(i)
		}
		wg.Done()
	}()
	go func() {
		chn <- 42 // Pass data to channel
		chn <- 27
		close(chn)
		wg.Done()
	}()
	wg.Wait()
}

/*
Output:
42
27
*/
