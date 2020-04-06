package main

import (
	"fmt"
)

func main() {
	num1, num2 := 1, 0

	fmt.Println("Let's find num1/num2")
	
	if num2==0{
		panic ("Something went wrong")
	} else{
		result := num1/num2
		fmt.Println("The result is ", result)
	}

}
/* Output :

Let's find num1/num2
panic: Something went wrong

goroutine 1 [running]:
main.main()
	/tmp/sandbox097738251/prog.go:13 +0xa0
	
*/