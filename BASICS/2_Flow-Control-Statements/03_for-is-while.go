package main 

import "fmt"

func main(){
	sum:=0
	for sum < 1000{
		sum += sum
	}

	fmt.Println("The final sum is ",sum)
}