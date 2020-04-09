package main
import (
    "fmt"
)
func main(){
    g:= greeter{"Good afternoon ," ,"viveshCodes "}
    g.greet()
}
type greeter struct{
    wish string
    person string
}
func (g greeter) greet() {
    fmt.Println(g.wish , g.person)
}