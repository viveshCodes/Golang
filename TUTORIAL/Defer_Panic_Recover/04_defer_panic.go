package main

import (
	"fmt"
)

func main() {
	num1, num2 := 1, 0

	fmt.Println("Let's find num1/num2")
	defer fmt.Println("Start")
	
	if num2==0{
		panic ("Something went wrong")
	} else{
		result := num1/num2
		fmt.Println("The result is ", result)
	}
	defer fmt.Println("end")
}
/*Output :
Let's find num1/num2
Start
panic: Something went wrong

goroutine 1 [running]:
main.main()
	/tmp/sandbox339133548/prog.go:14 +0x100


*/