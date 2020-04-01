package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Value= %v  and Type= %T", num, num)
}

/*
	Signed Integer availabe :

	int (default)
	int8
	int16
	int32
	int64

	Unsigned Integer available :

	uint (default)
	uint8
	uint16
	uint32
	uint64 (Previously , we didn't have 64 bit uint)
___________________________________________________________________

	Arithmetic Operations available :
	+
	-
	*
	/
	%
___________________________________________________________________

	Bitwise Operations Available :
	&   (AND)
	|   (OR)
	^   (XOR)
	&^  (XNOR)

	Shift Available :
	right shift (>>)  (Number / 2^shift_amount)
	left shift (<<)   (Number * 2^shift_amount)
________________________________________________________________
*/
