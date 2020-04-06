package main
import (
    
	"fmt"
)
func main(){
	collection := []int{ 7,7,8,1}
	fmt.Println("Index	Value  ")
	for key , value := range collection{
		fmt.Printf("%v       %v\n",key,value)
	}
}