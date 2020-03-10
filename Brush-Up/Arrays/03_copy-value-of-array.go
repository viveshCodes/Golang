package main

import (
	"fmt"
)

func main() {
	var arrayOfNum = [3]int{7, 8, 1}
	copyValues := arrayOfNum
	fmt.Println(arrayOfNum)
	fmt.Println(copyValues)

	copyValues[2] = 10
	fmt.Println(arrayOfNum)
	fmt.Println(copyValues)

	pointerToArray := &arrayOfNum
	pointerToArray[2] = 9
	fmt.Println(arrayOfNum)
	fmt.Println(pointerToArray)
	fmt.Println(*pointerToArray)
}

/*Output of this program :
[7 8 1]
[7 8 1]
[7 8 1]
[7 8 10]
[7 8 9]
&[7 8 9]
[7 8 9]
*/
