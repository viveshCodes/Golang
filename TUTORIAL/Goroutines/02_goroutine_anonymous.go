package main

import (
	"fmt"
	"time"
)

func main() {
	var message string
	message = "Hello World , Stay Safe"
	go func() {
		fmt.Println(message)
	}()
	time.Sleep(100 * time.Millisecond)
}
