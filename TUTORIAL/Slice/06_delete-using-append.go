package main

import (
	"fmt"
)

func main() {
	a := []int{7, 8, 9, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...) // delete 9 from slice a
	fmt.Println(b)
	fmt.Println(a)
}

/*
Output:
______
[7 8 9 4 5]
[7 8 4 5]
[7 8 4 5 5]  // Notice this . Why is 5 duplicated ?

*/
