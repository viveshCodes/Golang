package main

import (
	"fmt"
)

func main() {
	var arrayOfNum = [3]int{7, 8, 1}
	//var arrayOfNum =[...]int{7,8,1}
	// arrayOfNum :=[3]int{7,8,1}
	// arrayOfNum :=[...]int{7,8,1}
	// var arrayOfNum [3]string  , just for declaration
	fmt.Printf("Values : %v", arrayOfNum)
	fmt.Printf("Size of array is : ", len(arrayOfNum))
}

/*We can find size of an array using len()*/
