package main

import "fmt"

func main(){
	var num1 int
	fmt.Println("Enter the number of your choice")
	fmt.Scanln(&num1)

	change_value(&num1)
	
	fmt.Println("The number after being passing to fucntion 'change_value' is ",num1)	
	
}
func change_value(receive_address *int){
	fmt.Println("Enter a new value for num1")
	var new_num int
	fmt.Scanln(&new_num)
	*receive_address = new_num
}