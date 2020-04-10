package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(1)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc)
		fmt.Println(inc.Increment())
	}

}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (inc *IntCounter) Increment() int {
	fmt.Println(*inc)
	*inc++
	return int(*inc)
}

/*
0x40e020
1
2
0x40e020
2
3
0x40e020
3
4
0x40e020
4
5
0x40e020
5
6
0x40e020
6
7
0x40e020
7
8
0x40e020
8
9
0x40e020
9
10
0x40e020
10
11
*/
