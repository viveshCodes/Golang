package main

import (
	"fmt"
)

func main() {
	result, error := divideNumbers(7, 0)
	if error != nil {
		fmt.Println("Error : ", error)
		return
	}
	fmt.Println("Result = ", result)
}
func divideNumbers(num1, num2 int) (int, error) { // error is predefined keyword to be used here
	if num2 == 0 {
		return 0, fmt.Errorf("Something went wrong")
	}
	return num1 / num2, nil
}

/*Output :
Error : Something went wrong
*/
