package main

import (
	"fmt"
)

func main() {
	var complexNum complex64 = 1 + 7i
	// var complexNum complex64 = complex(1,7)
	fmt.Printf("Value = %v and Type = %T", complex_num, complex_num)
	fmt.Printf("Real part = %v  and Imaginary part = %v", real(complex_num), imag(complex_num))
}

/*
Types of float available to us :

complex64
complex128
_________________________________
Arithmetic Operations available :
+
-
*
/
__________________________________
Bitwise Operations and Bit shift are not available for complex numbers
____________________________________________________________
*/
