package main

import (
	"fmt"
)

func main() {
	sliceExample := make([]int, 0)
	oldSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(sliceExample)
	fmt.Printf("Length : %v\n", len(sliceExample))
	fmt.Printf("Capacity : %v\n", cap(sliceExample))

	// Let's now use append here
	sliceExample = append(sliceExample, 7)
	fmt.Println(sliceExample)
	fmt.Printf("Length : %v\n", len(sliceExample))
	fmt.Printf("Capacity : %v\n", cap(sliceExample))

	// Let's append more values and concatenate

	sliceExample = append(sliceExample, oldSlice...) // spread operators
	// sliceExample =append(sliceExample,[]int{1,2,3,4,5}...) this is equivalent to above statement if we don't declare oldSlice
	fmt.Println(sliceExample)
	fmt.Printf("Length : %v\n", len(sliceExample))
	fmt.Printf("Capacity : %v\n", cap(sliceExample))

}

/*
Output :
[]
Length : 0
Capacity : 0
[7]
Length : 1
Capacity : 2
[7 1 2 3 4 5]
Length : 6
Capacity : 8
*/
