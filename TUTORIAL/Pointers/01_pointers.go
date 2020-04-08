package main 
import (
	"fmt"
)
func main(){
	var number int = 7
	var ptr *int 
	ptr = &number
	fmt.Printf("number = %v  &number = %v \n",number , &number)
	fmt.Printf("*ptr = %v  ptr= %v " , (*ptr) ,ptr)
}