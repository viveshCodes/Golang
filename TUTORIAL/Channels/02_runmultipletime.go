package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// Create channel  named chn
	chn := make(chan int) // Here , chan is keyword used to create channel
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-chn // get data from channnel
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			chn <- 42 // Pass data to channel
			wg.Done()
		}()
	}
	wg.Wait()
}
