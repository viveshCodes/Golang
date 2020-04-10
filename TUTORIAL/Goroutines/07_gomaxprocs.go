package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Threads : ", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	fmt.Println("Threads : ", runtime.GOMAXPROCS(-1))

}

/*
Threads : 4
Threads : 100  
*/