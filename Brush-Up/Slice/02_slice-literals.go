package main

import (
	"fmt"
)

func main() {
	sliceExample := []int{7, 8, 9, 1, 2, 3, 4, 10}
	a := sliceExample[:]
	b := sliceExample[2:]
	c := sliceExample[:6]
	d := sliceExample[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

/*
	(I)Slicing operations wew have :
	_____________________________
	[:]
	[i:j] // i : inclusive , j : exclusive
	[i:]
	[:j]

	(II) Slicing operations are also available for Array:
	______________________________________________________
*/
