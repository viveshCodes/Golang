package main

import (
	"fmt"
	"time"
)

func greeting() {
	fmt.Println("Hello Go !")
}

func main() {

	go greeting()
	time.Sleep(100 * time.Millisecond)

}
