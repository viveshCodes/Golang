package main

import (
	"fmt"
)

func main() {
	var s string = "This is a string"
	// Let's do byte slicing here
	byteSlice := []byte(s)

	fmt.Printf("Value for s[2] = %v  and Type = %T \n", s[2], s[2])
	fmt.Printf("Value for string(s[2]) = %v  and Type = %T \n", string(s[2]), string(s[2]))
	fmt.Printf("Value for b = %v  and Type = %T \n", byte_slice, byte_slice)
}

/*
____________________________________________
Output :

Value for s[2] = 105  and Type = uint8
Value for string(s[2]) = i  and Type = string
Value for b = [84 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103]  and Type = []uint8
___________________________________________
*/
