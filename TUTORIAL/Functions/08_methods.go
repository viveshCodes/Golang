package main

import (
	"fmt"
)

func main() {
	g := greeter{"Good afternoon ,", "viveshCodes "}
	g.greet()
}

type greeter struct {
	wish   string
	person string
}

func (g greeter) greet() { // here greeter is called receiver
	fmt.Println(g.wish, g.person)
}

/*We can have pointer receiver as following:
func(g *greeter) (greet){

}
*/
