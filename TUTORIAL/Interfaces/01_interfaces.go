//Unlike struct , interface doesn't describe data but behavior.
package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	fmt.Println(w.Write([]byte("Hello Gopher")))
}

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*
Hello Gopher
13 <nil>
*/
