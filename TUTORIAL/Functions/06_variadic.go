package main
import (
	"fmt"
)
func getSum(values ...int)(int){
	var sum int
	for _,value := range values {
		sum += value
	}
	return sum
}
func main(){
	sum := getSum(1,2,3,4,5)
	fmt.Println("The sum is ",sum)
}