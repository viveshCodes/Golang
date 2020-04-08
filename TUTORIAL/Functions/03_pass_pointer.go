package main 
import (
	"fmt"
)
func modifyNumber(num *int){

	(*num) = (*num) + 7 

} 
func main(){
	var num int = 8
	modifyNumber(&num)
	fmt.Println(num)
}