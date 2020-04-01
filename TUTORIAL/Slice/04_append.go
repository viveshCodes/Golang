// append pushes element in to  the stack
package main

import (
	"fmt"
)

func main() {
	sliceExample := []int{}
	// Using make,  sliceExample := make([]int,0)

	fmt.Println(sliceExample)
	fmt.Printf("Length : %v\n", len(sliceExample))
	fmt.Printf("Capacity : %v\n", cap(sliceExample))

	// Let's now use append here
	sliceExample = append(sliceExample, 1)
	fmt.Println(sliceExample)
	fmt.Printf("Length : %v\n", len(sliceExample))
	fmt.Printf("Capacity : %v\n", cap(sliceExample))

	// Let's append more values and concatenate with previous value
	sliceExample = append(sliceExample, 1, 7, 9, 2, 3, 5)
	fmt.Println(sliceExample)
	fmt.Printf("Length : %v\n", len(sliceExample))
	fmt.Printf("Capacity : %v\n", cap(sliceExample))
}

/*Output :
__________
[]
Length : 0
Capacity : 0
[1]
Length : 1
Capacity : 2
[1 1 7 9 2 3 5]
Length : 7
Capacity : 8


*/
