package main

import (
	"fmt"
)

func main() {
	fmt.Print("Go runs on ")
	switch os:=7;os {
	case 1:
		fmt.Println("OS X.")
	case 2:
		fmt.Println("Linux.")
	default:

		fmt.Printf("This is default option.\n")
	}
}

/*
Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

*/