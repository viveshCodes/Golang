package main
import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
func main(){
	var message = "Hello World , Stay Safe "
	wg.Add(1)
	go func(msg string){
		fmt.Println(msg)
		wg. Done()
	}(message)
	message = "Hello World , It's Lockdown"
	wg.Wait()
}
/*Hllo World , Stay Safe*/