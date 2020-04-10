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
	go func(chn <-chan int) {
		i := <-chn // get data from channnel
		fmt.Println(i)
		wg.Done()
	}(chn)
	go func(chn chan<- int) {
		chn <- 42 // Pass data to channel
		wg.Done()
	}(chn)
	wg.Wait()
}
