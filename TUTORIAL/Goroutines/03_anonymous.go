package main
import (
	"fmt"
	"time"
)
func main(){
	message := "Hello World , Stay Safe"
	go func(msg string){
		fmt.Println(msg)
	}(message)
	message = "Hello World , It's Lockdown"
	time.Sleep(100 * time.Millisecond)
}
/*Hello World , Stay Safe*/