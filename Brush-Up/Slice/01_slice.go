// SLICE is refernece type
package main

import (
	"fmt"
)

func main() {
	var sliceExample = []int{7, 8, 1}
	//unlike Array
	pointerToSlice := sliceExample
	pointerToSlice[1] = 9
	fmt.Println(sliceExample)
	fmt.Println(pointerToSlice)
	fmt.Printf("Length : %v ", len(sliceExample))
	fmt.Println("Capacity : %v ", cap(sliceExample))

}

/* Functions available in Slice :
__________________________________
len()
cap()

*/

/*
	Output :
[7 9 1]
[7 9 1]
Length : 3 Capacity : 3
*/
