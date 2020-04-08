package main

import "fmt"

func main(){
	
	get_address := get_number()
	
	fmt.Println("The number returned from fucntion 'get_number' is ",*get_address)	
	
}
func get_number() (*int){
	fmt.Println("Enter a number of your choice")
	var num int
	fmt.Scanln(&num)
	 return &num
}