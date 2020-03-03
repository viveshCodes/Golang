package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
/* Exported name starts with Capital letter like Pi is exported name. 
  Using small letter as starting of exported name will throw an error.
   For example : pi is not an exported name BUT Pi. */
