// A Go program to explore variables in the Golang

package main

import "fmt"

func main() {
	var number1 int = 77
	var number2 int = 81

	var result int = number1 + number2
	fmt.Print("The sum of two numbers is ", result)
}

/*Exploring variable :
__________________________________________________________________
Method 1 : If we want to assign value to variable when we declare it :
------------------------------------------------------------------
   var number1 =77
   var number2 =81
   var result = number1 + number2
_________________________________________________________________________
Method 2: Below method is recommended
-------------------------------------------------------------------------
(1)
	//Declare variable
		var number1 int
		var number2 int
	//Then assign value whenever/wherever we want
(2)
	var number1 int =77
	var number2 int =81

__________________________________________________________________________
Shorthand :
--------------------------------------------------------------------------
(1)
	var number1 ,number2 int
	number1,number2 = 77,81
(2)
	var result int =7
	// Short-hand of above statement is :
	result :=7

*/

/******************* TYPES OF VARIABLE***********************/
/************************************************************/
/*
1.	bool
2. 	string
3.  int
	int8
	int16
	int32
	int64
	uint
	uint8
	uint16
	uint32
	uint64
	uintptr
	byte    // alias for uint8
	rune    // alias for int32 . It represents Unicode  Code point
4.  float32
	float64




*/
