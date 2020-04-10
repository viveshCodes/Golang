package main

import (
	"fmt"
	"time"
)

func main() {
	message := "Hello World , Stay Safe"
	go func() {
		fmt.Println(message)
	}()
	message = "Hello World , It's Lockdown"
	time.Sleep(100 * time.Millisecond)
}

/*output:
Hello World , It's Lockdown
*/
