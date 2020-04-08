package main
import (
	"fmt"
)
func printSum(msg string ,values ...int){
	var sum int
	for _,value := range values {
		sum += value
	}
	fmt.Println(msg , sum)
}
func main(){
	printSum( "The sum is ",1,2,3,4,5)
}