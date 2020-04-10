package main

import (
	"fmt"
)

func main() {
	var i interface{} = 7
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float64")
	default:
		fmt.Println("Other data type")

	}
}
