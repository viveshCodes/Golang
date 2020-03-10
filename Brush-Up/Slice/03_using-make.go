package main

import (
	"fmt"
)

func main() {
	var sliceExample = make([]int, 3, 100) // 3 is the length of slice , 100 is capacity
	fmt.Println(sliceExample)
	fmt.Printf("Length : %v \n", len(sliceExample))
	fmt.Println("Capacity : %v", cap(sliceExample))
}

/*
Output:
[0 0 0]
Length : 3
Capacity : 100

*/
