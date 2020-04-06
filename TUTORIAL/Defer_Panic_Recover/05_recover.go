package main

import (
	"fmt"
	"log"
)

func main() {
	num1, num2 := 1, 0

	fmt.Println("Let's find num1/num2")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error : ", err)
		}
	}()

	if num2 == 0 {
		panic("Something went wrong")
	} else {
		result := num1 / num2
		fmt.Println("The result is ", result)
	}
	defer fmt.Println("end")
}

/*Output:
Let's find num1/num2
2009/11/10 23:00:00 Error :  Something went wrong
*/