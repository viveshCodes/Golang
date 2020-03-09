package main

import (
	"fmt"
)

func main() {
	var complex_num complex64 = 1 + 7i
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
