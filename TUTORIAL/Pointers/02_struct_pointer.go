package main
import (
	"fmt"
)
type mystruct  struct {
	foo int
}
func main(){
	var ptr *mystruct 
	fmt.Println(ptr)        
	ptr = new(mystruct)
	fmt.Println(ptr)        
	fmt.Println(ptr.foo)
	/*
	(*ptr).foo is same as ptr.foo in Golang
	ptr = &mystruct can be written as ptr = new(mystruct) in Golang
	*/
}
/*Output:
<nil>
&{0}
0
*/