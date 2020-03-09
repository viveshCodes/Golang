package main

import (
	"fmt"
)

func main() {
	var num rune
	num = 'a' // Unlike string , we use single quote here

	fmt.Printf("Value = %v and Type = %T", num, num)
}

/*rune is type alias for int32*/

// Output : Value = 97 and Type = int32     Note that 97 is ASCII value for 'a'
