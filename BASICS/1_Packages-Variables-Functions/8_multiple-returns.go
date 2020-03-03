package main

import "fmt"

func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func main() {
	num1, num2 := swap(7, 3)
	fmt.Println("After swap ")
	fmt.Printf("num1= %d and num2=%d",num1,num2)
}
